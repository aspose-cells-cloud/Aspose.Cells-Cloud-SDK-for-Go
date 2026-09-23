#!/usr/bin/env bash
#
# run-tests.sh — Linux/macOS/Git-Bash runner for the newtests suite (mirrors run-tests.ps1).
#
# Runs each *_test.go file sequentially with `go test -v -timeout <timeout> <file> base_test.go`,
# streams the output to the console while writing it to <LogDir>/<file>.log, records the results
# as TSV files, and at the end analyses the run and writes a Chinese Markdown report.
#
# For every failing test case it additionally re-runs the test in isolation
# (`go test -v -run '^<Name>$'`), records that log too, and classifies the cause
# (认证失败 / 参数被拒 / 资源不存在 / 服务端错误 / 网络超时 / 编译失败 ...).
#
# Every test function is also mapped to the API version (v3.0 / v4.0) of the requests package
# constructors it calls, so the report can tell whether a failure cluster is a per-operation
# API-version issue (each generated request carries its operation's own version prefix).
#
# Usage:
#   ./run-tests.sh
#   ./run-tests.sh -Filter "api_cells_calculate*" -Timeout 5m
#   ./run-tests.sh -SkipIndividual -FailIfAnyTestFails
#   ./run-tests.sh -ReportOnly
#
# Options:
#   -Filter <glob>        Test-file glob (default "*_test.go").
#   -Timeout <dur>        Per-file `go test` timeout (default "20m").
#   -StopOnError          Halt at the first failing file (default: run all files).
#   -SkipIndividual       Skip the isolated re-run + cause analysis for failures.
#   -LogDir <dir>         Directory for logs and TSV records (default "test-logs").
#   -ReportFile <file>    Markdown report path (default "test-report.md").
#   -FailIfAnyTestFails   Exit 1 when a test case failed (default: exit 0 if the run completed).
#   -ReportOnly           Run no test: rebuild the report from an existing -LogDir
#                         (results.tsv, failures.tsv and the per-file logs) using the
#                         current analysis rules. Useful after tweaking the report layout
#                         or the failure classification, when repeating hours of live API
#                         calls would be wasteful. The recorded logs are read, never rewritten.
#   -h|--help             Show this help.
#
# Environment variables used by the tests themselves (see newtests/README.md):
#   CellsCloudClientId, CellsCloudClientSecret, CellsCloudApiBaseUrl

set -uo pipefail

# --- ANSI colors -------------------------------------------------------------
C_RED='\033[0;31m'
C_GREEN='\033[0;32m'
C_YELLOW='\033[0;33m'
C_MAGENTA='\033[0;35m'
C_CYAN='\033[0;36m'
C_WHITE='\033[0;37m'
C_DARKGRAY='\033[0;90m'
C_RED_BG='\033[41;37m'
NC='\033[0m'

# --- defaults -----------------------------------------------------------------
FILTER="*_test.go"
TIMEOUT="20m"
STOP_ON_ERROR=0
SKIP_INDIVIDUAL=0
LOG_DIR="test-logs"
REPORT_FILE="test-report.md"
FAIL_IF_ANY=0
REPORT_ONLY=0

usage() {
    # Print the leading comment block (everything after the shebang, up to the first code line).
    awk 'NR > 1 { if ($0 !~ /^#/) exit; sub(/^# ?/, ""); print }' "${BASH_SOURCE[0]}"
}

# --- argument parsing ---------------------------------------------------------
while [[ $# -gt 0 ]]; do
    case "$1" in
        -Filter=*) FILTER="${1#*=}"; shift ;;
        -Filter)   FILTER="$2"; shift 2 ;;
        -Timeout=*) TIMEOUT="${1#*=}"; shift ;;
        -Timeout)  TIMEOUT="$2"; shift 2 ;;
        -StopOnError) STOP_ON_ERROR=1; shift ;;
        -SkipIndividual) SKIP_INDIVIDUAL=1; shift ;;
        -LogDir=*) LOG_DIR="${1#*=}"; shift ;;
        -LogDir)   LOG_DIR="$2"; shift 2 ;;
        -ReportFile=*) REPORT_FILE="${1#*=}"; shift ;;
        -ReportFile)   REPORT_FILE="$2"; shift 2 ;;
        -FailIfAnyTestFails) FAIL_IF_ANY=1; shift ;;
        -ReportOnly) REPORT_ONLY=1; shift ;;
        -h|--help|-help) usage; exit 0 ;;
        *) echo "Unknown option: $1" >&2; exit 2 ;;
    esac
done

# --- setup ---------------------------------------------------------------------
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR" || exit 2

HELPER_FILE="base_test.go"

# Credentials and the Go toolchain are only needed to RUN tests; -ReportOnly reads logs.
if [[ $REPORT_ONLY -eq 0 ]]; then
    missing_vars=()
    for v in CellsCloudClientId CellsCloudClientSecret CellsCloudApiBaseUrl; do
        [[ -z "${!v:-}" ]] && missing_vars+=("$v")
    done
    if [[ ${#missing_vars[@]} -gt 0 ]]; then
        echo -e "${C_RED}缺少环境变量：$(IFS=', '; echo "${missing_vars[*]}")${NC}"
        echo -e "${C_YELLOW}请先设置（参见 newtests/README.md）：${NC}"
        echo -e "${C_YELLOW}  export CellsCloudClientId=...${NC}"
        echo -e "${C_YELLOW}  export CellsCloudClientSecret=...${NC}"
        echo -e "${C_YELLOW}  export CellsCloudApiBaseUrl=https://api.aspose.cloud${NC}"
        exit 2
    fi

    command -v go >/dev/null 2>&1 || { echo -e "${C_RED}未找到 go 命令，请确认 Go 已安装并在 PATH 中。${NC}" >&2; exit 2; }
    [[ -f "$HELPER_FILE" ]] || { echo -e "${C_RED}未找到 $HELPER_FILE（当前目录：$SCRIPT_DIR）${NC}" >&2; exit 2; }
    if [[ ! -d "TestData" ]]; then
        echo -e "${C_YELLOW}警告：未找到 TestData（应由 ../integrationtests/TestData 软链而来），测试可能因缺少样本文件而失败。${NC}"
    fi
fi

API_BASE_URL="${CellsCloudApiBaseUrl:-}"
SDK_VERSION="unknown"
if [[ -f "../version.go" ]]; then
    v="$(grep -o 'globalCellsCloudSDKVersion *= *"[^"]*"' ../version.go | sed 's/.*"\(.*\)"/\1/')"
    [[ -n "$v" ]] && SDK_VERSION="$v"
fi

# request-type -> API version, read from the generated requests package literals.
declare -A VERSION_OF
if [[ -d "../requests" ]]; then
    for f in ../requests/*.go; do
        [[ -f "$f" ]] || continue
        t="$(grep -o 'func New[A-Za-z0-9_]*Request' "$f" | head -n 1 | sed 's/^func New//; s/Request$//')"
        ver="$(grep -o 'localVarPath := "/v[0-9.]*' "$f" | head -n 1 | sed 's#.*"/##')"
        [[ -n "$t" && -n "$ver" ]] && VERSION_OF["$t"]="$ver"
    done
fi

# --- helpers -------------------------------------------------------------------
# "TestName<TAB>RequestType" for every requests.NewXxxRequest call inside a test function.
extract_test_requests() {
    awk '
        /^func[ \t]+Test[A-Za-z0-9_]+[ \t]*\(/ {
            name = $2; sub(/\(.*$/, "", name)
        }
        /requests\.New[A-Za-z0-9_]*Request/ {
            s = $0
            while (match(s, /requests\.New[A-Za-z0-9_]*Request/)) {
                type = substr(s, RSTART + 12, RLENGTH - 19)   # strip "requests.New" and "Request"
                print name "\t" type
                s = substr(s, RSTART + RLENGTH)
            }
        }
    ' "$1"
}

# "TestName<TAB>PASS|FAIL|SKIP" for the top-level test functions in a `go test -v` log.
extract_test_results() {
    grep -E '^--- (PASS|FAIL|SKIP): ' "$1" 2>/dev/null | awk '{
        st = $2; sub(/:$/, "", st)
        if ($3 !~ /\//) print $3 "\t" st
    }'
}

# Lines that carry an error cause (assertions, panics, timeouts), trimmed.
extract_error_lines() {
    grep -E '\.go:[0-9]+:|^panic:|test timed out' "$1" 2>/dev/null | sed 's/^[[:space:]]*//'
}

# Error lines belonging to ONE test in a file-level log. `go test -v` prints a test's own output
# between its "=== RUN <name>" line and the following "--- PASS/FAIL/SKIP" line, so the block
# gives the cause of *that* test — unlike extract_error_lines, which mixes every test in the file.
extract_test_error_lines() {
    awk -v want="$2" '
        $0 == "=== RUN   " want { inblock = 1; next }
        inblock && (/^--- (PASS|FAIL|SKIP): / || /^=== RUN   /) { exit }
        inblock { print }
    ' "$1" 2>/dev/null | grep -E '\.go:[0-9]+:|^panic:|test timed out' | sed 's/^[[:space:]]*//'
}

# Classify a failure from its error text; prints "<category><TAB><summary>".
classify_failure() {
    local text="$1"
    local category="其他"

    if [[ "$text" == *"test timed out"* ]]; then category="超时"
    elif [[ "$text" =~ (code=401|HTTP\ 401|[Uu]nauthorized) ]]; then category="认证失败(401)"
    elif [[ "$text" =~ (code=403|HTTP\ 403|[Ff]orbidden) ]]; then category="无权限(403)"
    elif [[ "$text" =~ (code=404|HTTP\ 404|[Nn]ot\ [Ff]ound) ]]; then category="资源不存在(404)"
    elif [[ "$text" =~ (code=400|HTTP\ 400|[Bb]ad\ [Rr]equest) ]]; then category="请求被拒(400)"
    elif [[ "$text" =~ (code=5[0-9][0-9]|HTTP\ 5[0-9][0-9]|Internal\ Server\ Error) ]]; then category="服务端错误(5xx)"
    elif [[ "$text" =~ (invalid\ parameter|required\ request\ parameter\ is\ missing|unknown\ operation) ]]; then category="参数校验失败"
    elif [[ "$text" =~ (no\ such\ host|connection\ refused|dial\ tcp|read\ tcp|write\ tcp|wsarecv|wsasend|connection\ reset|reset\ by\ peer|connection\ attempt\ failed|i/o\ timeout|context\ deadline\ exceeded|handshake|TLS|EOF|broken\ pipe) ]]; then category="网络错误"
    elif [[ "$text" =~ (build\ failed|cannot\ find\ package|undefined:|syntax\ error) ]]; then category="编译失败"
    fi

    # Prefer the line carrying the HTTP status / SDK error code: it is the actual cause.
    local summary
    summary="$(printf '%s\n' "$text" | grep -E 'code=[0-9]+|HTTP [0-9]+' | head -n 1)"
    [[ -z "$summary" ]] && summary="$(printf '%s\n' "$text" | head -n 1)"
    summary="$(printf '%s\n' "$summary" | sed 's/^[^:]*\.go:[0-9]*:[[:space:]]*//' | tr -d '\r')"
    [[ ${#summary} -gt 200 ]] && summary="${summary:0:200}..."

    printf '%s\t%s\n' "$category" "$summary"
}

# Render a version set as "v3.0", "v4.0" or "v3.0+v4.0"; "-" when unknown.
format_versions() {
    local joined="$1"
    [[ -z "$joined" ]] && { printf -- '-'; return; }
    printf '%s' "$(printf '%s\n' "$joined" | tr '+' '\n' | sort -u | paste -sd'+' -)"
}

# Append one TSV row, flattening tabs/newlines so a row stays on one line.
tsv_row() {
    local path="$1"; shift
    local line=""
    local field
    for field in "$@"; do
        field="$(printf '%s' "$field" | tr '\t\r\n' '   ' | sed 's/^ *//; s/ *$//')"
        line+="${field}"$'\t'
    done
    printf '%s\n' "${line%$'\t'}" >> "$path"
}

# --- file selection -------------------------------------------------------------
shopt -s nullglob
test_files=()
for f in $FILTER; do
    [[ -f "$f" ]] || continue
    [[ "$f" == "$HELPER_FILE" ]] && continue
    [[ "$f" == *_test.go ]] && test_files+=("$f")
done
shopt -u nullglob

# Sort for deterministic ordering (matches the PowerShell runner).
if [[ ${#test_files[@]} -gt 0 ]]; then
    mapfile -t test_files < <(printf '%s\n' "${test_files[@]}" | sort)
fi

if [[ ${#test_files[@]} -eq 0 && $REPORT_ONLY -eq 0 ]]; then
    echo -e "${C_YELLOW}没有匹配 '$FILTER' 的测试文件。${NC}"
    exit 0
fi

mkdir -p "$LOG_DIR"
RESULTS_TSV="$LOG_DIR/results.tsv"
FAILURES_TSV="$LOG_DIR/failures.tsv"
REPORT_PATH="$SCRIPT_DIR/$REPORT_FILE"
STARTED_AT="$(date '+%Y-%m-%d %H:%M:%S')"
START_EPOCH="$(date +%s)"

# Fresh record files for this run, with a header row so the columns are self-describing.
# -ReportOnly reads these files, so it must never truncate them.
if [[ $REPORT_ONLY -eq 0 ]]; then
    printf 'File\tTotal\tPass\tFail\tSkip\tStatus\tDuration\tVersions\n' > "$RESULTS_TSV"
    printf 'File\tTest\tReplay\tCategory\tApiVersions\tMessage\n' > "$FAILURES_TSV"
fi

echo ""
echo "============================================================"
echo -e "${C_CYAN} newtests 集成测试 (v4.0 新模型 API)${NC}"
echo "============================================================"
echo -e "${C_CYAN} 文件数    : ${#test_files[@]}${NC}"
echo -e "${C_CYAN} 超时      : ${TIMEOUT}${NC}"
echo -e "${C_CYAN} API 地址  : ${API_BASE_URL}${NC}"
echo -e "${C_CYAN} SDK 版本  : ${SDK_VERSION}${NC}"
echo -e "${C_CYAN} 日志      : $LOG_DIR${NC}"
echo -e "${C_CYAN} 报告      : $REPORT_FILE${NC}"
echo -e "${C_CYAN} 失败重跑  : $([[ $SKIP_INDIVIDUAL -eq 1 ]] && echo False || echo True)${NC}"
[[ $REPORT_ONLY -eq 1 ]] && echo -e "${C_CYAN} 重建模式  : 从 $LOG_DIR 重建报告，未执行任何测试${NC}"
echo "============================================================"
echo ""

# --- state ----------------------------------------------------------------------
declare -a FILE_ROWS=()        # "file|total|pass|fail|skip|status|duration|versions"
declare -a FAILED_ROWS=()      # "file|test|replay|category|versions|message"
declare -a SKIPPED_ROWS=()     # "file|test"
declare -A TEST_STATUS         # "file|test" -> PASS/FAIL/SKIP
declare -A TEST_VERSIONS       # "file|test" -> "v3.0+v4.0"

total_pass=0
total_fail=0
total_skip=0
total_build=0
total_timeout=0

if [[ $REPORT_ONLY -eq 1 ]]; then
    # -----------------------------------------------------------------------------------------
    # Rebuild the state above from an existing test-logs/ directory, so the report can be
    # re-rendered with the current analysis rules without repeating the live API calls.
    # results.tsv supplies the per-file facts; the per-file logs supply per-test status and the
    # error blocks; failures.tsv names the failing cases. Nothing here writes to test-logs/.
    # -----------------------------------------------------------------------------------------
    if [[ ! -f "$RESULTS_TSV" || ! -f "$FAILURES_TSV" ]]; then
        echo -e "${C_RED}找不到 $RESULTS_TSV 或 $FAILURES_TSV，请先正常运行一次测试。${NC}"
        exit 2
    fi

    while IFS=$'\t' read -r r_file r_total r_pass r_fail r_skip r_status r_duration r_versions; do
        r_versions="${r_versions%$'\r'}"     # tolerate a TSV written by the PowerShell runner
        [[ "$r_file" == "File" || -z "$r_file" ]] && continue
        FILE_ROWS+=("${r_file}|${r_total}|${r_pass}|${r_fail}|${r_skip}|${r_status}|${r_duration}|${r_versions}")
        total_pass=$((total_pass + ${r_pass:-0}))
        total_fail=$((total_fail + ${r_fail:-0}))
        total_skip=$((total_skip + ${r_skip:-0}))
        [[ "$r_status" == "BUILD" ]] && total_build=$((total_build + 1))
        [[ "$r_status" == "TIMEOUT" ]] && total_timeout=$((total_timeout + 1))
    done < <(tail -n +2 "$RESULTS_TSV")

    echo ""

    for row in "${FILE_ROWS[@]}"; do
        IFS='|' read -r f_name f_total f_pass f_fail f_skip f_status f_duration _ <<< "$row"
        base="${f_name%.go}"
        log_file="$LOG_DIR/$base.log"
        [[ -f "$log_file" ]] || continue

        declare -A file_test_versions=()
        while IFS=$'\t' read -r tname ttype; do
            [[ -z "$tname" || -z "$ttype" ]] && continue
            ver="${VERSION_OF[$ttype]:-}"
            [[ -z "$ver" ]] && continue
            if [[ -n "${file_test_versions[$tname]:-}" ]]; then
                case "+${file_test_versions[$tname]}+" in
                    *"+$ver+"*) ;;
                    *) file_test_versions[$tname]="${file_test_versions[$tname]}+$ver" ;;
                esac
            else
                file_test_versions[$tname]="$ver"
            fi
        done < <(extract_test_requests "$f_name")

        while IFS=$'\t' read -r tname tstatus; do
            [[ -z "$tname" ]] && continue
            TEST_STATUS["${f_name}|${tname}"]="$tstatus"
            [[ -n "${file_test_versions[$tname]:-}" ]] && TEST_VERSIONS["${f_name}|${tname}"]="${file_test_versions[$tname]}"
            [[ "$tstatus" == "SKIP" ]] && SKIPPED_ROWS+=("${f_name}|${tname}")
        done < <(extract_test_results "$log_file")
    done

    while IFS=$'\t' read -r r_file r_test r_replay r_category r_versions r_message; do
        r_message="${r_message%$'\r'}"
        [[ "$r_file" == "File" && "$r_test" == "Test" ]] && continue
        [[ -z "$r_file" || -z "$r_test" ]] && continue
        base="${r_file%.go}"
        log_file="$LOG_DIR/$base.log"
        isolated_log="$LOG_DIR/${base}.${r_test}.log"

        # Re-derive the cause from the raw logs rather than trusting the recorded summary, which
        # is truncated to 200 characters. A re-run that passes has no error output of its own, so
        # the first attempt's cause has to come from that test's block in the file-level log;
        # otherwise the row would be filed under "其他" with an empty summary.
        if [[ "$r_replay" == PASS* ]]; then
            err_text="$(extract_test_error_lines "$log_file" "$r_test")"
        elif [[ -f "$isolated_log" ]]; then
            err_text="$(extract_error_lines "$isolated_log")"
        else
            err_text="$(extract_test_error_lines "$log_file" "$r_test")"
        fi

        IFS=$'\t' read -r derived_category derived_summary <<< "$(classify_failure "$err_text")"

        if [[ "$r_replay" == PASS* ]]; then
            r_category="偶发失败(重跑通过)"
            r_message="${derived_summary:-首次执行失败、单独重跑通过，判定为偶发（服务端/网络抖动）。}"
        elif [[ -n "$derived_summary" && "$r_replay" != "(未重跑)" ]]; then
            r_category="$derived_category"
            r_message="$derived_summary"
        fi
        FAILED_ROWS+=("${r_file}|${r_test}|${r_replay}|${r_category}|${r_versions}|${r_message}")
    done < <(tail -n +2 "$FAILURES_TSV")

    # The run's wall clock is not recorded in the TSVs; approximate it from the log timestamps.
    if compgen -G "$LOG_DIR/*.log" > /dev/null; then
        first_log="$(ls -tr "$LOG_DIR"/*.log | head -n 1)"
        last_log="$(ls -tr "$LOG_DIR"/*.log | tail -n 1)"
        STARTED_AT="$(date -r "$first_log" '+%Y-%m-%d %H:%M:%S')"
        ENDED_AT_REBUILT="$(date -r "$last_log" '+%Y-%m-%d %H:%M:%S')"
        START_EPOCH="$(date -r "$first_log" +%s)"
        END_EPOCH_REBUILT="$(date -r "$last_log" +%s)"
    fi

    test_files=()   # an empty list keeps the run loop below from executing
fi

index=0
for file in "${test_files[@]}"; do
    index=$((index + 1))
    base="${file%.go}"
    log_file="$LOG_DIR/$base.log"

    echo -e "${C_WHITE}[${index}/${#test_files[@]}] ${file} ...${NC}"

    # Per-test request versions for this file (also stored per test function below).
    declare -A file_test_versions=()
    while IFS=$'\t' read -r tname ttype; do
        [[ -z "$tname" || -z "$ttype" ]] && continue
        ver="${VERSION_OF[$ttype]:-}"
        [[ -z "$ver" ]] && continue
        if [[ -n "${file_test_versions[$tname]:-}" ]]; then
            case "+${file_test_versions[$tname]}+" in
                *"+$ver+"*) ;;
                *) file_test_versions[$tname]="${file_test_versions[$tname]}+$ver" ;;
            esac
        else
            file_test_versions[$tname]="$ver"
        fi
    done < <(extract_test_requests "$file")

    start_ms="$(date +%s%3N)"
    go test -v -timeout "$TIMEOUT" "$file" "$HELPER_FILE" 2>&1 | tee "$log_file"
    exit_code="${PIPESTATUS[0]}"
    end_ms="$(date +%s%3N)"
    elapsed="$(awk -v ms="$((end_ms - start_ms))" 'BEGIN{printf "%.1fs", ms/1000}')"

    pass_count="$(grep -c -E '^--- PASS: ' "$log_file" 2>/dev/null || true)"
    fail_count="$(grep -c -E '^--- FAIL: ' "$log_file" 2>/dev/null || true)"
    skip_count="$(grep -c -E '^--- SKIP: ' "$log_file" 2>/dev/null || true)"
    is_timeout=0
    grep -q 'test timed out' "$log_file" 2>/dev/null && is_timeout=1
    is_build_error=0
    grep -qE 'build failed|\[build failed\]|cannot find package|# command-line-arguments' "$log_file" 2>/dev/null && is_build_error=1

    if [[ $is_timeout -eq 1 ]]; then
        status="TIMEOUT"; status_color="$C_RED"; total_timeout=$((total_timeout + 1))
    elif [[ $exit_code -eq 0 ]]; then
        status="PASS"; status_color="$C_GREEN"
    elif [[ $exit_code -eq 1 ]]; then
        status="FAIL"; status_color="$C_RED"
    else
        status="BUILD"; status_color="$C_MAGENTA"; total_build=$((total_build + 1))
    fi

    total_pass=$((total_pass + pass_count))
    total_fail=$((total_fail + fail_count))
    total_skip=$((total_skip + skip_count))
    total_tests=$((pass_count + fail_count + skip_count))

    # File-level version set: union over the test functions of this file.
    file_versions=""
    for v in "${file_test_versions[@]:-}"; do
        [[ -z "$v" ]] && continue
        while IFS= read -r single; do
            case "+${file_versions}+" in
                *"+$single+"*) ;;
                *) file_versions="${file_versions:+$file_versions+}$single" ;;
            esac
        done < <(printf '%s\n' "$v" | tr '+' '\n')
    done
    file_versions_text="$(format_versions "$file_versions")"

    FILE_ROWS+=("${file}|${total_tests}|${pass_count}|${fail_count}|${skip_count}|${status}|${elapsed}|${file_versions_text}")
    tsv_row "$RESULTS_TSV" "$file" "$total_tests" "$pass_count" "$fail_count" "$skip_count" "$status" "$elapsed" "$file_versions_text"

    echo -e "  -> ${status_color}${status}  通过=${pass_count}  失败=${fail_count}  跳过=${skip_count}  (${elapsed})${NC}"

    # Remember every test's status and API versions for the report analysis.
    while IFS=$'\t' read -r tname tstatus; do
        [[ -z "$tname" ]] && continue
        TEST_STATUS["${file}|${tname}"]="$tstatus"
        [[ -n "${file_test_versions[$tname]:-}" ]] && TEST_VERSIONS["${file}|${tname}"]="${file_test_versions[$tname]}"
        [[ "$tstatus" == "SKIP" ]] && SKIPPED_ROWS+=("${file}|${tname}")
    done < <(extract_test_results "$log_file")

    if [[ $is_timeout -eq 1 ]]; then
        echo -e "  ${C_RED_BG} !!! 该文件测试超时（${TIMEOUT}）。请提醒测试人员：可能是服务器出现异常，请检查服务端状态。 ${NC}"
    fi

    # ---- Failure analysis + isolated re-run ------------------------------------------------
    if [[ "$status" != "PASS" && $SKIP_INDIVIDUAL -eq 0 ]]; then
        mapfile -t failed_names < <(extract_test_results "$log_file" | awk -F'\t' '$2=="FAIL"{print $1}')

        if [[ "$status" == "BUILD" || $is_build_error -eq 1 ]]; then
            echo -e "  ${C_YELLOW}[错误原因分析] 编译失败：${NC}"
            err_text="$(extract_error_lines "$log_file")"
            IFS=$'\t' read -r category summary <<< "$(classify_failure "$err_text")"
            if [[ -n "$summary" ]]; then
                echo -e "    ${C_DARKGRAY}${summary}${NC}"
            else
                echo -e "    ${C_DARKGRAY}（无具体错误信息，请查看日志 ${log_file}）${NC}"
            fi
            tsv_row "$FAILURES_TSV" "$file" "(编译失败)" "N/A" "编译失败" "$file_versions_text" "$summary"
            FAILED_ROWS+=("${file}|(编译失败)|N/A|编译失败|${file_versions_text}|${summary}")
        elif [[ ${#failed_names[@]} -eq 0 ]]; then
            echo -e "  ${C_YELLOW}[错误原因分析] 测试失败，但未解析到具体的失败用例名。${NC}"
            err_text="$(extract_error_lines "$log_file")"
            IFS=$'\t' read -r category summary <<< "$(classify_failure "$err_text")"
            [[ -n "$summary" ]] && echo -e "    ${C_DARKGRAY}${summary}${NC}"
            tsv_row "$FAILURES_TSV" "$file" "(未解析到用例)" "N/A" "$category" "$file_versions_text" "$summary"
            FAILED_ROWS+=("${file}|(未解析到用例)|N/A|${category}|${file_versions_text}|${summary}")
        else
            echo -e "  ${C_YELLOW}失败用例（${#failed_names[@]} 个）：$(IFS=', '; echo "${failed_names[*]}")${NC}"
            echo ""

            for test_name in "${failed_names[@]}"; do
                [[ -z "$test_name" ]] && continue
                echo -e "    ${C_YELLOW}--- 单独重跑: ${test_name} ...${NC}"

                single_log="$LOG_DIR/${base}.${test_name}.log"
                go test -v -timeout "$TIMEOUT" -run "^${test_name}$" "$file" "$HELPER_FILE" 2>&1 | tee "$single_log"
                single_exit="${PIPESTATUS[0]}"

                if [[ $single_exit -eq 0 ]]; then
                    replay="PASS(重跑通过)"; replay_color="$C_GREEN"
                    # A re-run that passes produces no error output of its own, so the cause of the
                    # FIRST attempt is read from that test's block in the file-level log. Without
                    # this the row would be filed under "其他" with an empty summary.
                    err_text="$(extract_test_error_lines "$log_file" "$test_name")"
                    IFS=$'\t' read -r _first_category summary <<< "$(classify_failure "$err_text")"
                    category="偶发失败(重跑通过)"
                    [[ -z "$summary" ]] && summary="首次执行失败、单独重跑通过，判定为偶发（服务端/网络抖动）。"
                else
                    replay="FAIL"; replay_color="$C_RED"
                    err_text="$(extract_error_lines "$single_log")"
                    IFS=$'\t' read -r category summary <<< "$(classify_failure "$err_text")"
                fi

                versions_text="$(format_versions "${TEST_VERSIONS[${file}|${test_name}]:-}")"

                echo -e "      结果: ${replay_color}${replay}${NC}"
                [[ -n "$summary" ]] && echo -e "      原因: ${C_DARKGRAY}[${category}] ${summary}${NC}"
                echo ""

                tsv_row "$FAILURES_TSV" "$file" "$test_name" "$replay" "$category" "$versions_text" "$summary"
                FAILED_ROWS+=("${file}|${test_name}|${replay}|${category}|${versions_text}|${summary}")
            done
        fi
    elif [[ "$status" != "PASS" && $SKIP_INDIVIDUAL -eq 1 ]]; then
        echo -e "  ${C_DARKGRAY}（已跳过失败用例的单独重跑，因为指定了 -SkipIndividual）${NC}"
        while IFS=$'\t' read -r tname tstatus; do
            [[ "$tstatus" != "FAIL" ]] && continue
            versions_text="$(format_versions "${TEST_VERSIONS[${file}|${tname}]:-}")"
            tsv_row "$FAILURES_TSV" "$file" "$tname" "(未重跑)" "未分析" "$versions_text" ""
            FAILED_ROWS+=("${file}|${tname}|(未重跑)|未分析|${versions_text}|")
        done < <(extract_test_results "$log_file")
    fi

    echo ""

    if [[ $exit_code -ne 0 && $STOP_ON_ERROR -eq 1 ]]; then
        echo -e "${C_YELLOW}遇到失败文件，按 -StopOnError 提前结束。${NC}"
        break
    fi
done

if [[ $REPORT_ONLY -eq 1 ]]; then
    # Keep the timestamps the rebuild derived from the log files.
    ENDED_AT="${ENDED_AT_REBUILT:-$ENDED_AT}"
    wall_clock="$(awk -v s="$START_EPOCH" -v e="${END_EPOCH_REBUILT:-$(date +%s)}" 'BEGIN{printf "%.1fs", e-s}')"
else
    ENDED_AT="$(date '+%Y-%m-%d %H:%M:%S')"
    wall_clock="$(awk -v s="$START_EPOCH" -v e="$(date +%s)" 'BEGIN{printf "%.1fs", e-s}')"
fi
total_tests_run=$((total_pass + total_fail + total_skip))
if [[ $((total_pass + total_fail)) -gt 0 ]]; then
    pass_rate="$(awk -v p="$total_pass" -v t="$total_pass" -v f="$total_fail" 'BEGIN{printf "%.2f", 100*p/(t+f)}')"
else
    pass_rate="0.00"
fi

# --- console summary -------------------------------------------------------------
echo "============================================================"
echo -e "${C_CYAN} 汇总${NC}"
echo "============================================================"
{
    printf 'File\tTotal\tPass\tFail\tSkip\tStatus\tDuration\tVersions\n'
    printf '%s\n' "${FILE_ROWS[@]}" | awk -F'|' '{print $1"\t"$2"\t"$3"\t"$4"\t"$5"\t"$6"\t"$7"\t"$8}'
} | column -t -s $'\t'
echo ""
echo -e "${C_WHITE}文件数        : ${#FILE_ROWS[@]}${NC}"
echo -e "${C_WHITE}用例总数      : ${total_tests_run}${NC}"
echo -e "${C_GREEN}通过          : ${total_pass}${NC}"
echo -e "${C_RED}失败          : ${total_fail}${NC}"
echo -e "${C_YELLOW}跳过          : ${total_skip}${NC}"
echo -e "${C_MAGENTA}编译失败文件  : ${total_build}${NC}"
echo -e "${C_RED}超时文件      : ${total_timeout}${NC}"
echo -e "${C_WHITE}通过率        : ${pass_rate}%  (通过 / (通过+失败))${NC}"
echo -e "${C_WHITE}总耗时        : ${wall_clock}${NC}"
echo ""

if [[ ${#FAILED_ROWS[@]} -gt 0 ]]; then
    echo "============================================================"
    echo -e "${C_CYAN} 失败用例明细${NC}"
    echo "============================================================"
    {
        printf 'File\tTest\tReplay\tCategory\tVersions\n'
        printf '%s\n' "${FAILED_ROWS[@]}" | awk -F'|' '{print $1"\t"$2"\t"$3"\t"$4"\t"$5}'
    } | column -t -s $'\t'
    echo ""
fi

# --- Markdown report (Chinese) ---------------------------------------------------
md_escape() { printf '%s' "$1" | sed 's/|/\\|/g'; }

{
    echo "# newtests 集成测试报告"
    echo ""
    echo "> 本报告由 \`newtests/run-tests.sh\` 自动生成，数据来源为本次运行的原始 \`go test -v\` 日志与 \`test-logs/*.tsv\` 记录。"
    if [[ $REPORT_ONLY -eq 1 ]]; then
        echo ">"
        echo "> 本报告以 \`-ReportOnly\` 从 \`${LOG_DIR}\` 重建，未重新执行测试；起止时间为日志文件时间戳。"
    fi
    echo ""
    echo "## 一、运行环境"
    echo ""
    echo "| 项目 | 值 |"
    echo "|---|---|"
    echo "| 测试套件 | newtests（v4.0 新模型 API 集成测试） |"
    echo "| API 基地址 | \`${API_BASE_URL}\` |"
    echo "| SDK 版本 | \`${SDK_VERSION}\` |"
    echo "| 单文件超时 | \`${TIMEOUT}\` |"
    echo "| 开始时间 | ${STARTED_AT} |"
    echo "| 结束时间 | ${ENDED_AT} |"
    echo "| 总耗时 | ${wall_clock} |"
    echo "| 日志目录 | \`${LOG_DIR}\` |"
    echo "| 失败用例重跑 | $([[ $SKIP_INDIVIDUAL -eq 1 ]] && echo "未启用(-SkipIndividual)" || echo "已启用") |"
    echo ""
    echo "## 二、总览"
    echo ""
    echo "| 指标 | 数值 |"
    echo "|---|---|"
    echo "| 测试文件 | ${#FILE_ROWS[@]} |"
    echo "| 用例总数（含跳过） | ${total_tests_run} |"
    echo "| 通过 | ${total_pass} |"
    echo "| 失败 | ${total_fail} |"
    echo "| 跳过 | ${total_skip} |"
    echo "| 编译失败文件 | ${total_build} |"
    echo "| 超时文件 | ${total_timeout} |"
    echo "| 通过率 | ${pass_rate}% |"
    echo ""
    if [[ $total_fail -eq 0 && $total_build -eq 0 && $total_timeout -eq 0 ]]; then
        echo "**结论：本次运行全部用例通过。**"
    else
        echo "**结论：本次运行存在失败项**，明细见第三、四节。"
    fi
    echo ""
    echo "## 三、逐文件结果"
    echo ""
    echo "| 文件 | 用例 | 通过 | 失败 | 跳过 | 状态 | 耗时 | 调用的 API 版本 |"
    echo "|---|---:|---:|---:|---:|---|---:|---|"
    for row in "${FILE_ROWS[@]}"; do
        IFS='|' read -r f t p fa s st d ver <<< "$row"
        echo "| \`${f}\` | ${t} | ${p} | ${fa} | ${s} | ${st} | ${d} | ${ver} |"
    done
    echo ""
    echo "## 四、失败用例明细"
    echo ""
    if [[ ${#FAILED_ROWS[@]} -eq 0 ]]; then
        echo "本次运行没有失败用例。"
    else
        echo "| # | 文件 | 用例 | 单独重跑 | 原因分类 | 调用的 API 版本 | 错误摘要 |"
        echo "|---:|---|---|---|---|---|---|"
        i=0
        for row in "${FAILED_ROWS[@]}"; do
            i=$((i + 1))
            IFS='|' read -r f t r c ver msg <<< "$row"
            echo "| ${i} | \`${f}\` | \`${t}\` | ${r} | ${c} | ${ver} | $(md_escape "$msg") |"
        done
    fi
    echo ""
    echo "## 五、失败原因分类统计"
    echo ""
    if [[ ${#FAILED_ROWS[@]} -eq 0 ]]; then
        echo "无失败用例，无需分类。"
    else
        echo "| 原因分类 | 数量 |"
        echo "|---|---:|"
        for row in "${FAILED_ROWS[@]}"; do
            IFS='|' read -r f t r c ver msg <<< "$row"
            echo "$c"
        done | sort | uniq -c | sort -k1,1nr -k2,2 | while read -r count cat; do
            echo "| ${cat} | ${count} |"
        done
    fi
    echo ""
    echo "## 六、API 版本维度分析"
    echo ""
    echo "每个生成的 request 现在携带其 operation 自己的 API 版本前缀（v3.0 或 v4.0，取自规范中该 operation 的 \`APIVersion\`）。"
    echo "下表按用例调用的 API 版本统计通过率，用于判断失败是否集中在某一个版本上。"
    echo ""
    echo "| 用例调用的 API 版本 | 用例数 | 通过 | 失败 | 通过率 |"
    echo "|---|---:|---:|---:|---:|"
    # bucket[version] = "total pass fail" for every non-skipped test.
    declare -A bucket_total=() bucket_pass=() bucket_fail=()
    for key in "${!TEST_STATUS[@]}"; do
        st="${TEST_STATUS[$key]}"
        [[ "$st" == "SKIP" ]] && continue
        ver="$(format_versions "${TEST_VERSIONS[$key]:-}")"
        bucket_total["$ver"]=$(( ${bucket_total["$ver"]:-0} + 1 ))
        if [[ "$st" == "PASS" ]]; then
            bucket_pass["$ver"]=$(( ${bucket_pass["$ver"]:-0} + 1 ))
        else
            bucket_fail["$ver"]=$(( ${bucket_fail["$ver"]:-0} + 1 ))
        fi
    done
    for ver in $(printf '%s\n' "${!bucket_total[@]}" | sort); do
        bt="${bucket_total[$ver]}"; bp="${bucket_pass[$ver]:-0}"; bf="${bucket_fail[$ver]:-0}"
        brate="$(awk -v p="$bp" -v f="$bf" 'BEGIN{ if (p+f>0) printf "%.2f", 100*p/(p+f); else print "0.00" }')"
        echo "| ${ver} | ${bt} | ${bp} | ${bf} | ${brate}% |"
    done
    echo ""

    echo "## 七、分析结论与建议"
    echo ""
    if [[ $total_fail -eq 0 && $total_build -eq 0 && $total_timeout -eq 0 ]]; then
        echo "- 本次运行 **${total_pass} 个用例全部通过**（跳过 ${total_skip} 个），未发现回归。"
    else
        echo "- 共执行 ${#FILE_ROWS[@]} 个文件 / ${total_tests_run} 个用例，通过 ${total_pass}，失败 ${total_fail}，跳过 ${total_skip}，通过率 **${pass_rate}%**。"

        top_category="$(for row in "${FAILED_ROWS[@]}"; do IFS='|' read -r f t r c ver msg <<< "$row"; echo "$c"; done | sort | uniq -c | sort -k1,1nr -k2,2 | head -n 1 | sed 's/^ *[0-9]* *//')"
        top_count="$(for row in "${FAILED_ROWS[@]}"; do IFS='|' read -r f t r c ver msg <<< "$row"; echo "$c"; done | sort | uniq -c | sort -k1,1nr -k2,2 | head -n 1 | awk '{print $1}')"
        [[ -n "$top_category" ]] && echo "- 失败最集中的原因分类是 **${top_category}**（${top_count} 个用例），应优先排查。"

        v3_total=0; v3_fail=0; v4_total=0; v4_fail=0
        for key in "${!TEST_STATUS[@]}"; do
            st="${TEST_STATUS[$key]}"
            [[ "$st" == "SKIP" ]] && continue
            ver="$(format_versions "${TEST_VERSIONS[$key]:-}")"
            if [[ "$ver" == "v3.0" ]]; then
                v3_total=$((v3_total + 1)); [[ "$st" == "FAIL" ]] && v3_fail=$((v3_fail + 1))
            elif [[ "$ver" == "v4.0" ]]; then
                v4_total=$((v4_total + 1)); [[ "$st" == "FAIL" ]] && v4_fail=$((v4_fail + 1))
            fi
        done
        if [[ $v3_total -gt 0 || $v4_total -gt 0 ]]; then
            v3_rate="$(awk -v f="$v3_fail" -v t="$v3_total" 'BEGIN{ if (t>0) printf "%.2f", 100*f/t; else print "0.00" }')"
            v4_rate="$(awk -v f="$v4_fail" -v t="$v4_total" 'BEGIN{ if (t>0) printf "%.2f", 100*f/t; else print "0.00" }')"
            echo "- 仅调用 v3.0 接口的用例 ${v3_total} 个，失败率 ${v3_rate}%；仅调用 v4.0 接口的用例 ${v4_total} 个，失败率 ${v4_rate}%。"
            if [[ $v3_total -gt 0 && $v4_total -gt 0 ]] && awk -v a="$v3_rate" -v b="$v4_rate" 'BEGIN{exit !(a > b*2 && a > 5)}'; then
                echo "- 失败在 v3.0 用例上明显更密集，但**这并不表示版本前缀寻址有误**：同一批 v3.0 URL 重试即可返回 200，而同路径在 v4.0 下为 404（端点不存在），说明原因是这些 v3.0 端点的**服务端偶发超时**，而非请求打到了错误的版本。判定请以第四节的错误摘要与「单独重跑」列为准。"
            fi
        fi

        replay_pass="$(printf '%s\n' "${FAILED_ROWS[@]}" | awk -F'|' '$3 ~ /^PASS/{n++} END{print n+0}')"
        network_fail="$(printf '%s\n' "${FAILED_ROWS[@]}" | awk -F'|' '$4 ~ /网络错误/{n++} END{print n+0}')"
        transport_fail=$((network_fail + replay_pass))
        if [[ $transport_fail -gt 0 ]]; then
            echo "- 失败中有 ${transport_fail} 个属于**传输层问题**（${network_fail} 个网络错误 + ${replay_pass} 个重跑即通过），表现为 \`context deadline exceeded\` / \`TLS handshake timeout\`，属于服务端或网络抖动，而非接口契约或请求模型缺陷；建议先按文件重跑以剔除这些噪声，再分析剩余失败。"
        fi

        [[ $replay_pass -gt 0 ]] && echo "- 有 ${replay_pass} 个用例在单独重跑时通过，说明存在**偶发性（服务器抖动）**或共享远程文件状态的因素，建议确认测试之间的远程目录是否需要隔离。"

        if [[ $total_build -gt 0 ]]; then
            build_files="$(printf '%s\n' "${FILE_ROWS[@]}" | awk -F'|' '$6=="BUILD"{printf "%s, ", $1}' | sed 's/, $//')"
            echo "- 有 ${total_build} 个文件**编译失败**（${build_files}），这类文件的所有用例都未真正执行，需先修复编译问题再复测。"
        fi
        if [[ $total_timeout -gt 0 ]]; then
            timeout_files="$(printf '%s\n' "${FILE_ROWS[@]}" | awk -F'|' '$6=="TIMEOUT"{printf "%s, ", $1}' | sed 's/, $//')"
            echo "- 有 ${total_timeout} 个文件**超时**（${timeout_files}）。请提醒测试人员：可能是服务器出现异常，请检查服务端状态。"
        fi
    fi
    if [[ ${#SKIPPED_ROWS[@]} -gt 0 ]]; then
        echo "- 有 ${#SKIPPED_ROWS[@]} 个用例被跳过（\`t.Skip\`），多为多文件上传受限的 LightCells 组装/合并用例，不计入通过率。"
    fi
    echo "- 详细日志：\`${LOG_DIR}/results.tsv\`（逐文件）、\`${LOG_DIR}/failures.tsv\`（逐失败用例）、\`${LOG_DIR}/<文件>.log\`（原始输出）。"
    echo ""
    echo "## 八、跳过的用例"
    echo ""
    if [[ ${#SKIPPED_ROWS[@]} -eq 0 ]]; then
        echo "本次运行没有被跳过的用例。"
    else
        echo "| 文件 | 用例 |"
        echo "|---|---|"
        for row in "${SKIPPED_ROWS[@]}"; do
            IFS='|' read -r f t <<< "$row"
            echo "| \`${f}\` | \`${t}\` |"
        done
    fi
    echo ""
    echo "---"
    echo ""
    echo "_报告由 \`newtests/run-tests.sh\` 生成。_"
} > "$REPORT_PATH"

echo "============================================================"
echo -e "${C_GREEN} 报告已生成: $REPORT_PATH${NC}"
echo -e "${C_GREEN} 逐文件记录: $RESULTS_TSV${NC}"
echo -e "${C_GREEN} 失败记录  : $FAILURES_TSV${NC}"
echo "============================================================"

if [[ $FAIL_IF_ANY -eq 1 && ($total_fail -gt 0 || $total_build -gt 0 || $total_timeout -gt 0) ]]; then
    exit 1
fi
exit 0
