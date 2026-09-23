<#
.SYNOPSIS
    Runs the newtests suite (v4.0 new-model integration tests), records every run, analyses the
    results and generates a Chinese Markdown report.

.DESCRIPTION
    Iterates over every *_test.go file in this directory (excluding base_test.go), runs it with
    `go test -v -timeout <timeout> <file> base_test.go`, streams the output to the console while
    writing it to <LogDir>/<file>.log, and aggregates the per-file results.

    All files are executed regardless of individual failures (use -StopOnError to halt at the
    first failing file).

    For each failing test case it additionally:
      1. Extracts the failed test function names (--- FAIL: <Name>).
      2. Re-runs each one in isolation with `go test -v -run '^<Name>$'` and logs it separately.
      3. Classifies the cause (认证失败 / 参数被拒 / 资源不存在 / 服务端错误 / 网络超时 / 编译失败 ...).

    Every test function is also mapped to the API version (v3.0 / v4.0) of the requests package
    constructors it calls, so the report can tell whether a failure cluster is a per-operation
    API-version issue (each generated request now carries its operation's own version prefix).

    Artefacts (all under <LogDir> unless noted):
      results.tsv          per-file summary, written as the run progresses
      failures.tsv         one row per failed test case, written as the run progresses
      <file>.log           raw `go test -v` output of the file run
      <file>.<Test>.log    raw output of the isolated re-run of a failed test
      <ReportFile>         the Markdown report (default: test-report.md, next to this script)

    The script does NOT modify the test sources: newtests/ is generated output
    (see tools/convert_integrationtests.py in the repository root).

.PARAMETER Filter
    Glob pattern for the test files to run (defaults to all *_test.go files).

.PARAMETER Timeout
    Timeout passed to `go test` for a single file (defaults to 20 minutes).

.PARAMETER StopOnError
    Stop after the first failing file (default: run ALL files regardless of failures).

.PARAMETER SkipIndividual
    Skip the isolated re-run + cause analysis for failed test cases.

.PARAMETER LogDir
    Directory (relative to this script) where logs and TSV records are written.

.PARAMETER ReportFile
    Markdown report path (relative to this script).

.PARAMETER FailIfAnyTestFails
    Exit with code 1 when at least one test case failed (default: exit 0 whenever the run
    itself completed, so a CI job can archive the report instead of stopping at the first run).

.PARAMETER ReportOnly
    Do not run any test. Rebuild the report from an existing -LogDir (results.tsv, failures.tsv
    and the per-file `go test -v` logs) using the current analysis rules. Useful after tweaking
    the report layout or the failure classification, when repeating hours of live API calls
    would be wasteful. The recorded logs are read, never rewritten.

.EXAMPLE
    .\run-tests.ps1
    .\run-tests.ps1 -Filter "api_cells_calculate*" -Timeout 5m
    .\run-tests.ps1 -SkipIndividual -FailIfAnyTestFails
    .\run-tests.ps1 -ReportOnly

.NOTES
    Credentials are read from the environment (same names as ../integrationtests):
      CellsCloudClientId, CellsCloudClientSecret, CellsCloudApiBaseUrl
    Every test performs real HTTP calls against the Aspose.Cells Cloud service.
#>

[CmdletBinding()]
param(
    [string]$Filter = "*_test.go",
    [string]$Timeout = "20m",
    [switch]$StopOnError,
    [switch]$SkipIndividual,
    [string]$LogDir = "test-logs",
    [string]$ReportFile = "test-report.md",
    [switch]$FailIfAnyTestFails,
    [switch]$ReportOnly
)

$ErrorActionPreference = "Stop"

$helperFile = "base_test.go"
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $scriptDir

# ---------------------------------------------------------------------------------------------
# Helper functions
# ---------------------------------------------------------------------------------------------

# Console line in a fixed colour (keeps the output readable in both PS 5.1 and PS 7).
function Write-Step {
    param([string]$Text, [string]$Color = "Gray")
    Write-Host $Text -ForegroundColor $Color
}

# Read a table of results and write it as TSV, replacing tabs/newlines so a row stays one line.
function Add-TsvRow {
    param([string]$Path, [string[]]$Fields)
    $clean = $Fields | ForEach-Object {
        if ($null -eq $_) { "" } else { ($_ -replace "[\t\r\n]+", " ").Trim() }
    }
    $line = $clean -join "`t"
    Add-Content -Path $Path -Value $line -Encoding UTF8
}

# Classify a failure from its error text. Returns a category and a one-line summary.
function Get-FailureCategory {
    param([string[]]$Output)

    $errorLines = @($Output | Where-Object {
            $_ -match '\.go:\d+:' -or $_ -match '^panic:' -or $_ -match 'test timed out'
        })
    $text = ($errorLines -join "`n")

    $summary = ""
    if ($errorLines.Count -gt 0) {
        $summary = $errorLines[0].Trim()
        # Prefer the line carrying the HTTP status / SDK error code: it is the actual cause.
        $preferred = $errorLines | Where-Object { $_ -match 'code=\d+|HTTP \d+' } | Select-Object -First 1
        if ($preferred) { $summary = $preferred.Trim() }
        $summary = $summary -replace '^\S+\.go:\d+:\s*', ''
    }

    if ($text -match 'test timed out') { $category = "超时" }
    elseif ($text -match 'code=401|HTTP 401|unauthorized|Unauthorized') { $category = "认证失败(401)" }
    elseif ($text -match 'code=403|HTTP 403|forbidden|Forbidden') { $category = "无权限(403)" }
    elseif ($text -match 'code=404|HTTP 404|Not Found|not found') { $category = "资源不存在(404)" }
    elseif ($text -match 'code=400|HTTP 400|Bad Request|bad request') { $category = "请求被拒(400)" }
    elseif ($text -match 'code=5\d\d|HTTP 5\d\d|Internal Server Error') { $category = "服务端错误(5xx)" }
    elseif ($text -match 'invalid parameter|required request parameter is missing|unknown operation') { $category = "参数校验失败" }
    elseif ($text -match 'no such host|connection refused|dial tcp|read tcp|write tcp|wsarecv|wsasend|connection reset|reset by peer|connection attempt failed|i/o timeout|context deadline exceeded|handshake|TLS|EOF|broken pipe') { $category = "网络错误" }
    elseif ($text -match 'build failed|cannot find package|undefined:|syntax error') { $category = "编译失败" }
    else { $category = "其他" }

    if ($summary.Length -gt 200) { $summary = $summary.Substring(0, 200) + "..." }
    return [pscustomobject]@{ Category = $category; Summary = $summary }
}

# Extract top-level test results (name -> PASS/FAIL/SKIP) from `go test -v` output.
function Get-TestResults {
    param([string[]]$Output)
    $map = [ordered]@{}
    foreach ($line in $Output) {
        # Sub-test lines are indented ("    --- FAIL: Test/sub"); only top-level lines start at col 0.
        if ($line -match '^--- (PASS|FAIL|SKIP):\s+(\S+)') {
            $status = $Matches[1]
            $name = $Matches[2]
            if ($name -notmatch '/') { $map[$name] = $status }
        }
    }
    return $map
}

# Extract the error lines a single test produced during a file-level run.
# `go test -v` groups a test's own output between its `=== RUN` line and the following
# `--- PASS/FAIL/SKIP` line, so the block gives the cause of *that* test — unlike scraping the
# whole file log, which mixes every test in the file together.
function Get-TestErrorLines {
    param([string[]]$Output, [string]$TestName)

    $lines = [System.Collections.Generic.List[string]]::new()
    $inBlock = $false
    foreach ($line in $Output) {
        if ($line -eq "=== RUN   $TestName") { $inBlock = $true; continue }
        if ($inBlock -and ($line -match '^--- (PASS|FAIL|SKIP): ' -or $line -match '^=== RUN   ')) { break }
        if ($inBlock) { $lines.Add($line) }
    }
    return @($lines | Where-Object { $_ -match '\.go:\d+:' -or $_ -match '^panic:' -or $_ -match 'test timed out' })
}

# Map every test function in a file to the API versions (v3.0/v4.0) of the requests it builds.
function Get-TestRequestVersions {
    param([string]$Path, [hashtable]$VersionMap)

    $map = @{}
    $current = $null
    foreach ($line in [System.IO.File]::ReadAllLines($Path)) {
        if ($line -match '^func\s+(Test[A-Za-z0-9_]+)\s*\(') { $current = $Matches[1] }
        if ($current -and $line -match 'requests\.New[A-Za-z0-9_]*Request') {
            foreach ($m in [regex]::Matches($line, 'requests\.New([A-Za-z0-9_]+)Request')) {
                $type = $m.Groups[1].Value
                if ($VersionMap.ContainsKey($type)) {
                    if (-not $map.ContainsKey($current)) {
                        $map[$current] = New-Object System.Collections.Generic.HashSet[string]
                    }
                    [void]$map[$current].Add($VersionMap[$type])
                }
            }
        }
    }
    return $map
}

# Render a HashSet of versions as "v3.0", "v4.0" or "v3.0+v4.0"; "-" when unknown.
function Format-Versions {
    param($Versions)
    if (-not $Versions -or $Versions.Count -eq 0) { return "-" }
    return (($Versions | Sort-Object) -join "+")
}

# ---------------------------------------------------------------------------------------------
# Pre-flight checks
# ---------------------------------------------------------------------------------------------

# Credentials and the Go toolchain are only needed to RUN tests; -ReportOnly reads logs.
$credentialVars = @("CellsCloudClientId", "CellsCloudClientSecret", "CellsCloudApiBaseUrl")
if (-not $ReportOnly) {
    $missingVars = @($credentialVars | Where-Object { [string]::IsNullOrWhiteSpace([Environment]::GetEnvironmentVariable($_)) })
    if ($missingVars.Count -gt 0) {
        Write-Step "缺少环境变量：$($missingVars -join ', ')" "Red"
        Write-Step "请先设置（参见 newtests/README.md）：" "Yellow"
        Write-Step '  export CellsCloudClientId=...' "Yellow"
        Write-Step '  export CellsCloudClientSecret=...' "Yellow"
        Write-Step '  export CellsCloudApiBaseUrl=https://api.aspose.cloud' "Yellow"
        exit 2
    }
}

if (-not $ReportOnly) {
    if (-not (Test-Path $helperFile)) {
        Write-Step "未找到 $helperFile（当前目录：$scriptDir）" "Red"
        exit 2
    }
    if (-not (Test-Path "TestData")) {
        Write-Step "警告：未找到 TestData（应由 ../integrationtests/TestData 软链而来），测试可能因缺少样本文件而失败。" "Yellow"
    }

    $goCmd = Get-Command go -ErrorAction SilentlyContinue
    if (-not $goCmd) {
        Write-Step "未找到 go 命令，请确认 Go 已安装并在 PATH 中。" "Red"
        exit 2
    }
}

# API base URL, SDK version, and the request-type -> API-version map (used by the analysis).
$apiBaseUrl = [Environment]::GetEnvironmentVariable("CellsCloudApiBaseUrl")
$sdkVersion = "unknown"
$versionGoPath = Join-Path (Split-Path -Parent $scriptDir) "version.go"
if (Test-Path $versionGoPath) {
    $versionGo = Get-Content $versionGoPath -Raw
    if ($versionGo -match 'globalCellsCloudSDKVersion\s*=\s*"([^"]+)"') { $sdkVersion = $Matches[1] }
}

$versionMap = @{}
$requestsDir = Join-Path (Split-Path -Parent $scriptDir) "requests"
if (Test-Path $requestsDir) {
    foreach ($reqFile in Get-ChildItem $requestsDir -Filter "*.go") {
        $content = Get-Content $reqFile.FullName -Raw
        $typeName = $null
        $apiVersion = $null
        if ($content -match 'func\s+New([A-Za-z0-9_]+)Request\s*\(') { $typeName = $Matches[1] }
        if ($content -match 'localVarPath\s*:=\s*"/(v[0-9]+\.[0-9]+)') { $apiVersion = $Matches[1] }
        if ($typeName -and $apiVersion) { $versionMap[$typeName] = $apiVersion }
    }
}

# ---------------------------------------------------------------------------------------------
# Run
# ---------------------------------------------------------------------------------------------

$testFiles = Get-ChildItem -Path $scriptDir -Filter $Filter -File |
    Where-Object { $_.Name -ne $helperFile -and $_.Name -like "*_test.go" } |
    Sort-Object Name

if (-not $testFiles -and -not $ReportOnly) {
    Write-Step "没有匹配 '$Filter' 的测试文件。" "Yellow"
    exit 0
}

$logPath = Join-Path $scriptDir $LogDir
New-Item -ItemType Directory -Force -Path $logPath | Out-Null

$resultsTsv = Join-Path $logPath "results.tsv"
$failuresTsv = Join-Path $logPath "failures.tsv"
$reportPath = Join-Path $scriptDir $ReportFile
$startedAt = Get-Date

# Fresh record files for this run, with a header row so the columns are self-describing.
# -ReportOnly reads these files, so it must never truncate them.
if (-not $ReportOnly) {
    Set-Content -Path $resultsTsv -Value ("File`tTotal`tPass`tFail`tSkip`tStatus`tDuration`tVersions") -Encoding UTF8
    Set-Content -Path $failuresTsv -Value ("File`tTest`tReplay`tCategory`tApiVersions`tMessage") -Encoding UTF8
}

Write-Step ""
Write-Step "============================================================" "Cyan"
Write-Step " newtests 集成测试 (v4.0 新模型 API)" "Cyan"
Write-Step "============================================================" "Cyan"
Write-Step " 文件数    : $($testFiles.Count)" "Cyan"
Write-Step " 超时      : $Timeout" "Cyan"
Write-Step " API 地址  : $apiBaseUrl" "Cyan"
Write-Step " SDK 版本  : $sdkVersion" "Cyan"
Write-Step " 日志      : $logPath" "Cyan"
Write-Step " 报告      : $ReportFile" "Cyan"
Write-Step " 失败重跑  : $(-not $SkipIndividual)" "Cyan"
Write-Step "============================================================" "Cyan"
Write-Step ""

$fileResults = [System.Collections.Generic.List[object]]::new()
$failedCases = [System.Collections.Generic.List[object]]::new()
$skippedTests = [System.Collections.Generic.List[object]]::new()
$testVersionInfo = @{}     # "file|test" -> HashSet of versions
$allTestResults = @{}      # "file|test" -> PASS/FAIL/SKIP

$totalPass = 0
$totalFail = 0
$totalSkip = 0
$totalBuild = 0
$totalTimeout = 0

if ($ReportOnly) {
    # -----------------------------------------------------------------------------------------
    # Rebuild the state above from an existing test-logs/ directory, so the report can be
    # re-rendered with the current analysis rules without repeating the live API calls.
    # results.tsv supplies the per-file facts; the per-file logs supply per-test status and the
    # error blocks; failures.tsv names the failing cases. Nothing here writes to test-logs/.
    # -----------------------------------------------------------------------------------------
    if (-not (Test-Path $resultsTsv) -or -not (Test-Path $failuresTsv)) {
        Write-Step "找不到 $resultsTsv 或 $failuresTsv，请先正常运行一次测试。" "Red"
        exit 2
    }

    foreach ($row in (Import-Csv -Path $resultsTsv -Delimiter "`t")) {
        $fileResults.Add([pscustomobject]@{
                File     = $row.File
                Total    = [int]$row.Total
                Pass     = [int]$row.Pass
                Fail     = [int]$row.Fail
                Skip     = [int]$row.Skip
                Status   = $row.Status
                Duration = $row.Duration
                Versions = $row.Versions
            })
        $totalPass += [int]$row.Pass
        $totalFail += [int]$row.Fail
        $totalSkip += [int]$row.Skip
        if ($row.Status -eq "BUILD") { $totalBuild++ }
        if ($row.Status -eq "TIMEOUT") { $totalTimeout++ }

        $logFile = Join-Path $logPath (([System.IO.Path]::GetFileNameWithoutExtension($row.File)) + ".log")
        if (-not (Test-Path $logFile)) { continue }

        $testResultsMap = Get-TestResults -Output ([System.IO.File]::ReadAllLines($logFile))
        $fileVersionMap = Get-TestRequestVersions -Path (Join-Path $scriptDir $row.File) -VersionMap $versionMap
        foreach ($name in $testResultsMap.Keys) {
            $key = "$($row.File)|$name"
            $allTestResults[$key] = $testResultsMap[$name]
            if ($fileVersionMap.ContainsKey($name)) { $testVersionInfo[$key] = $fileVersionMap[$name] }
            if ($testResultsMap[$name] -eq "SKIP") {
                $skippedTests.Add([pscustomobject]@{ File = $row.File; Test = $name })
            }
        }
    }

    foreach ($row in (Import-Csv -Path $failuresTsv -Delimiter "`t")) {
        $category = $row.Category
        $message = $row.Message

        $base = [System.IO.Path]::GetFileNameWithoutExtension($row.File)
        $logFile = Join-Path $logPath ($base + ".log")
        $isolatedLog = Join-Path $logPath ($base + "." + $row.Test + ".log")

        # Re-derive the cause from the raw logs rather than trusting the recorded summary, which
        # is truncated to 200 characters. A re-run that passes has no error output of its own, so
        # the first attempt's cause has to come from that test's block in the file-level log;
        # otherwise the row would be filed under "其他" with an empty summary.
        $errText = @()
        if ($row.Replay -like "PASS*") {
            if (Test-Path $logFile) {
                $errText = Get-TestErrorLines -Output ([System.IO.File]::ReadAllLines($logFile)) -TestName $row.Test
            }
        }
        elseif (Test-Path $isolatedLog) {
            $errText = [System.IO.File]::ReadAllLines($isolatedLog)
        }
        elseif (Test-Path $logFile) {
            $errText = Get-TestErrorLines -Output ([System.IO.File]::ReadAllLines($logFile)) -TestName $row.Test
        }

        $derived = Get-FailureCategory -Output $errText
        if ($row.Replay -like "PASS*") {
            $category = "偶发失败(重跑通过)"
            $message = if ($derived.Summary) { $derived.Summary } else { "首次执行失败、单独重跑通过，判定为偶发（服务端/网络抖动）。" }
        }
        elseif ($derived.Summary -and $row.Replay -ne "(未重跑)") {
            $category = $derived.Category
            $message = $derived.Summary
        }

        $failedCases.Add([pscustomobject]@{
                File     = $row.File
                Test     = $row.Test
                Replay   = $row.Replay
                Category = $category
                Versions = $row.ApiVersions
                Message  = $message
            })
    }

    # The run's wall clock is not recorded in the TSVs; approximate it from every log timestamp,
    # counting the isolated re-runs too (they are the last thing the run writes).
    $logStamps = @(Get-ChildItem -Path $logPath -Filter "*.log" -ErrorAction SilentlyContinue |
        ForEach-Object { $_.LastWriteTime })
    if ($logStamps.Count -gt 0) {
        $startedAt = ($logStamps | Sort-Object | Select-Object -First 1)
        $finishedAt = ($logStamps | Sort-Object | Select-Object -Last 1)
        $wallClock = "{0:F1}s" -f ($finishedAt - $startedAt).TotalSeconds
    }

    Write-Step " 重建模式  : 从 $LogDir 重建报告，未执行任何测试" "Cyan"
    Write-Step ""

    $testFiles = @()   # an empty list keeps the run loop below from executing
}

$index = 0
foreach ($file in $testFiles) {
    $index++
    $logFile = Join-Path $logPath ($file.BaseName + ".log")
    Write-Step ("[{0}/{1}] {2} ..." -f $index, $testFiles.Count, $file.Name) "White"

    # Request-type -> version map for this file (also stored per test function).
    $fileVersionMap = Get-TestRequestVersions -Path $file.FullName -VersionMap $versionMap

    $sw = [System.Diagnostics.Stopwatch]::StartNew()
    # Stream the output to the console and capture it (Tee-Object returns what it writes).
    $output = & go test -v -timeout $Timeout $file.Name $helperFile 2>&1 | Tee-Object -FilePath $logFile
    $exitCode = $LASTEXITCODE
    $sw.Stop()

    $testResults = Get-TestResults -Output $output
    $passCount = @($testResults.Values | Where-Object { $_ -eq "PASS" }).Count
    $failCount = @($testResults.Values | Where-Object { $_ -eq "FAIL" }).Count
    $skipCount = @($testResults.Values | Where-Object { $_ -eq "SKIP" }).Count
    $isTimeout = @($output | Select-String -Pattern 'test timed out').Count -gt 0
    $isBuildError = @($output | Select-String -Pattern 'build failed|\[build failed\]|cannot find package|# command-line-arguments').Count -gt 0

    if ($isTimeout) {
        $status = "TIMEOUT"; $statusColor = "Red"; $totalTimeout++
    }
    elseif ($exitCode -eq 0) {
        $status = "PASS"; $statusColor = "Green"
    }
    elseif ($exitCode -eq 1) {
        $status = "FAIL"; $statusColor = "Red"
    }
    else {
        $status = "BUILD"; $statusColor = "Magenta"; $totalBuild++
    }

    $totalPass += $passCount
    $totalFail += $failCount
    $totalSkip += $skipCount
    $totalTests = $passCount + $failCount + $skipCount
    $elapsed = "{0:F1}s" -f $sw.Elapsed.TotalSeconds

    # File-level version set: union over the test functions of this file.
    $fileVersions = New-Object System.Collections.Generic.HashSet[string]
    foreach ($versions in $fileVersionMap.Values) { foreach ($v in $versions) { [void]$fileVersions.Add($v) } }
    $fileVersionsText = Format-Versions $fileVersions

    $fileResults.Add([pscustomobject]@{
            File      = $file.Name
            Total     = $totalTests
            Pass      = $passCount
            Fail      = $failCount
            Skip      = $skipCount
            Status    = $status
            Duration  = $elapsed
            Versions  = $fileVersionsText
        })
    Add-TsvRow -Path $resultsTsv -Fields @($file.Name, $totalTests, $passCount, $failCount, $skipCount, $status, $elapsed, $fileVersionsText)

    Write-Step ("  -> {0}  通过={1}  失败={2}  跳过={3}  ({4})" -f $status, $passCount, $failCount, $skipCount, $elapsed) $statusColor

    # Remember every test's status and API versions for the report analysis.
    foreach ($kv in $testResults.GetEnumerator()) {
        $key = "$($file.Name)|$($kv.Key)"
        $allTestResults[$key] = $kv.Value
        $json = $null
        if ($fileVersionMap.ContainsKey($kv.Key)) { $testVersionInfo[$key] = $fileVersionMap[$kv.Key] }
        if ($kv.Value -eq "SKIP") {
            $skippedTests.Add([pscustomobject]@{ File = $file.Name; Test = $kv.Key })
        }
    }

    if ($isTimeout) {
        Write-Step "  !!! 该文件测试超时（$Timeout）。请提醒测试人员：可能是服务器出现异常，请检查服务端状态。" "White"
    }

    # ---- Failure analysis + isolated re-run ------------------------------------------------
    if ($status -ne "PASS" -and -not $SkipIndividual) {
        $failedNames = @($testResults.GetEnumerator() | Where-Object { $_.Value -eq "FAIL" } | ForEach-Object { $_.Key })

        if ($status -eq "BUILD" -or $isBuildError) {
            Write-Step "  [错误原因分析] 编译失败：" "Yellow"
            $cause = Get-FailureCategory -Output $output
            if ($cause.Summary) { Write-Step ("    {0}" -f $cause.Summary) "DarkYellow" }
            else { Write-Step ("    （无具体错误信息，请查看日志 {0}）" -f $logFile) "DarkYellow" }

            Add-TsvRow -Path $failuresTsv -Fields @($file.Name, "(编译失败)", "N/A", "编译失败", $fileVersionsText, $cause.Summary)
            $failedCases.Add([pscustomobject]@{
                    File      = $file.Name
                    Test      = "(编译失败)"
                    Replay    = "N/A"
                    Category  = "编译失败"
                    Versions  = $fileVersionsText
                    Message   = $cause.Summary
                })
        }
        elseif ($failedNames.Count -eq 0) {
            Write-Step "  [错误原因分析] 测试失败，但未解析到具体的失败用例名。" "Yellow"
            $cause = Get-FailureCategory -Output $output
            if ($cause.Summary) { Write-Step ("    {0}" -f $cause.Summary) "DarkYellow" }

            Add-TsvRow -Path $failuresTsv -Fields @($file.Name, "(未解析到用例)", "N/A", $cause.Category, $fileVersionsText, $cause.Summary)
            $failedCases.Add([pscustomobject]@{
                    File      = $file.Name
                    Test      = "(未解析到用例)"
                    Replay    = "N/A"
                    Category  = $cause.Category
                    Versions  = $fileVersionsText
                    Message   = $cause.Summary
                })
        }
        else {
            Write-Step ("  失败用例（{0} 个）：{1}" -f $failedNames.Count, ($failedNames -join ', ')) "Yellow"
            Write-Step ""

            foreach ($testName in $failedNames) {
                Write-Step ("    --- 单独重跑: {0} ..." -f $testName) "Yellow"

                $singleOut = & go test -v -timeout $Timeout -run "^$testName$" $file.Name $helperFile 2>&1
                $singleExit = $LASTEXITCODE

                $singleLog = Join-Path $logPath ($file.BaseName + "." + $testName + ".log")
                $singleOut | Out-File -FilePath $singleLog -Encoding utf8

                # The cause worth reporting belongs to the FIRST attempt. A re-run that passes
                # produces no error output of its own, so read the test's own block out of the
                # file-level log instead — otherwise the row would land in "其他" with no summary.
                $cause = Get-FailureCategory -Output $singleOut
                if ($singleExit -eq 0) {
                    $replay = "PASS(重跑通过)"; $replayColor = "Green"
                    $firstCause = Get-FailureCategory -Output (Get-TestErrorLines -Output $output -TestName $testName)
                    $cause = [pscustomobject]@{
                        Category = "偶发失败(重跑通过)"
                        Summary  = if ($firstCause.Summary) { $firstCause.Summary } else { "首次执行失败、单独重跑通过，判定为偶发（服务端/网络抖动）。" }
                    }
                }
                else { $replay = "FAIL"; $replayColor = "Red" }

                $key = "$($file.Name)|$testName"
                $versionsText = "-"
                if ($testVersionInfo.ContainsKey($key)) { $versionsText = Format-Versions $testVersionInfo[$key] }

                Write-Step ("      结果: {0}" -f $replay) $replayColor
                if ($cause.Summary) { Write-Step ("      原因: [{0}] {1}" -f $cause.Category, $cause.Summary) "DarkYellow" }
                Write-Step ""

                Add-TsvRow -Path $failuresTsv -Fields @($file.Name, $testName, $replay, $cause.Category, $versionsText, $cause.Summary)
                $failedCases.Add([pscustomobject]@{
                        File     = $file.Name
                        Test     = $testName
                        Replay   = $replay
                        Category = $cause.Category
                        Versions = $versionsText
                        Message  = $cause.Summary
                    })
            }
        }
    }
    elseif ($status -ne "PASS" -and $SkipIndividual) {
        Write-Step "  （已跳过失败用例的单独重跑，因为指定了 -SkipIndividual）" "DarkGray"
        foreach ($failedName in @($testResults.GetEnumerator() | Where-Object { $_.Value -eq "FAIL" } | ForEach-Object { $_.Key })) {
            $key = "$($file.Name)|$failedName"
            $versionsText = "-"
            if ($testVersionInfo.ContainsKey($key)) { $versionsText = Format-Versions $testVersionInfo[$key] }
            Add-TsvRow -Path $failuresTsv -Fields @($file.Name, $failedName, "(未重跑)", "未分析", $versionsText, "")
            $failedCases.Add([pscustomobject]@{
                    File     = $file.Name
                    Test     = $failedName
                    Replay   = "(未重跑)"
                    Category = "未分析"
                    Versions = $versionsText
                    Message  = ""
                })
        }
    }

    Write-Step ""

    if ($exitCode -ne 0 -and $StopOnError) {
        Write-Step "遇到失败文件，按 -StopOnError 提前结束。" "Yellow"
        break
    }
}

if (-not $ReportOnly) {
    $finishedAt = Get-Date
    $wallClock = "{0:F1}s" -f ($finishedAt - $startedAt).TotalSeconds
}
$totalTestsRun = $totalPass + $totalFail + $totalSkip
$passRate = if ($totalTestsRun -gt 0) { [math]::Round(100 * $totalPass / [math]::Max(1, ($totalPass + $totalFail)), 2) } else { 0 }

# ---------------------------------------------------------------------------------------------
# Console summary
# ---------------------------------------------------------------------------------------------

Write-Step "============================================================" "Cyan"
Write-Step " 汇总" "Cyan"
Write-Step "============================================================" "Cyan"
$fileResults | Format-Table -AutoSize
Write-Step ("文件数        : {0}" -f $fileResults.Count) "White"
Write-Step ("用例总数      : {0}" -f $totalTestsRun) "White"
Write-Step ("通过          : {0}" -f $totalPass) "Green"
Write-Step ("失败          : {0}" -f $totalFail) "Red"
Write-Step ("跳过          : {0}" -f $totalSkip) "Yellow"
Write-Step ("编译失败文件  : {0}" -f $totalBuild) "Magenta"
Write-Step ("超时文件      : {0}" -f $totalTimeout) "Red"
Write-Step ("通过率        : {0}%  (通过 / (通过+失败))" -f $passRate) "White"
Write-Step ("总耗时        : {0}" -f $wallClock) "White"
Write-Step ""

if ($failedCases.Count -gt 0) {
    Write-Step "============================================================" "Cyan"
    Write-Step " 失败用例明细" "Cyan"
    Write-Step "============================================================" "Cyan"
    $failedCases | Format-Table File, Test, Replay, Category, Versions -AutoSize -Wrap
    Write-Step ""
}

# ---------------------------------------------------------------------------------------------
# Markdown report (Chinese)
# ---------------------------------------------------------------------------------------------

$md = [System.Collections.Generic.List[string]]::new()
$md.Add("# newtests 集成测试报告")
$md.Add("")
$md.Add("> 本报告由 ``newtests/run-tests.ps1`` 自动生成，数据来源为本次运行的原始 ``go test -v`` 日志与 ``test-logs/*.tsv`` 记录。")
if ($ReportOnly) {
    $md.Add(">")
    $md.Add("> 本报告以 ``-ReportOnly`` 从 ``$LogDir`` 重建，未重新执行测试；起止时间为日志文件时间戳。")
}
$md.Add("")
$md.Add("## 一、运行环境")
$md.Add("")
$md.Add("| 项目 | 值 |")
$md.Add("|---|---|")
$md.Add("| 测试套件 | newtests（v4.0 新模型 API 集成测试） |")
$md.Add("| API 基地址 | ``$apiBaseUrl`` |")
$md.Add("| SDK 版本 | ``$sdkVersion`` |")
$md.Add("| 单文件超时 | ``$Timeout`` |")
$md.Add("| 开始时间 | $(Get-Date $startedAt -Format 'yyyy-MM-dd HH:mm:ss') |")
$md.Add("| 结束时间 | $(Get-Date $finishedAt -Format 'yyyy-MM-dd HH:mm:ss') |")
$md.Add("| 总耗时 | $wallClock |")
$md.Add("| 日志目录 | ``$LogDir`` |")
$md.Add("| 失败用例重跑 | $(if ($SkipIndividual) { "未启用(-SkipIndividual)" } else { "已启用" }) |")
$md.Add("")

$md.Add("## 二、总览")
$md.Add("")
$md.Add("| 指标 | 数值 |")
$md.Add("|---|---|")
$md.Add("| 测试文件 | $($fileResults.Count) |")
$md.Add("| 用例总数（含跳过） | $totalTestsRun |")
$md.Add("| 通过 | $totalPass |")
$md.Add("| 失败 | $totalFail |")
$md.Add("| 跳过 | $totalSkip |")
$md.Add("| 编译失败文件 | $totalBuild |")
$md.Add("| 超时文件 | $totalTimeout |")
$md.Add("| 通过率 | $passRate% |")
$md.Add("")
if ($totalFail -eq 0 -and $totalBuild -eq 0 -and $totalTimeout -eq 0) {
    $md.Add("**结论：本次运行全部用例通过。**")
} else {
    $md.Add("**结论：本次运行存在失败项**，明细见第三、四节。")
}
$md.Add("")

$md.Add("## 三、逐文件结果")
$md.Add("")
$md.Add("| 文件 | 用例 | 通过 | 失败 | 跳过 | 状态 | 耗时 | 调用的 API 版本 |")
$md.Add("|---|---:|---:|---:|---:|---|---:|---|")
foreach ($r in $fileResults) {
    $md.Add("| ``$($r.File)`` | $($r.Total) | $($r.Pass) | $($r.Fail) | $($r.Skip) | $($r.Status) | $($r.Duration) | $($r.Versions) |")
}
$md.Add("")

$md.Add("## 四、失败用例明细")
$md.Add("")
if ($failedCases.Count -eq 0) {
    $md.Add("本次运行没有失败用例。")
} else {
    $md.Add("| # | 文件 | 用例 | 单独重跑 | 原因分类 | 调用的 API 版本 | 错误摘要 |")
    $md.Add("|---:|---|---|---|---|---|---|")
    $i = 0
    foreach ($f in $failedCases) {
        $i++
        $msg = ($f.Message -replace '\|', '\|')
        $md.Add("| $i | ``$($f.File)`` | ``$($f.Test)`` | $($f.Replay) | $($f.Category) | $($f.Versions) | $msg |")
    }
}
$md.Add("")

$md.Add("## 五、失败原因分类统计")
$md.Add("")
if ($failedCases.Count -eq 0) {
    $md.Add("无失败用例，无需分类。")
} else {
    $md.Add("| 原因分类 | 数量 |")
    $md.Add("|---|---:|")
    foreach ($group in ($failedCases | Group-Object Category |
            Sort-Object @{Expression = { $_.Count }; Descending = $true }, @{Expression = { $_.Name }; Descending = $false })) {
        $md.Add("| $($group.Name) | $($group.Count) |")
    }
}
$md.Add("")

# ---- Analysis: pass rate by API version (v3.0 vs v4.0) -------------------------------------
$md.Add("## 六、API 版本维度分析")
$md.Add("")
$md.Add("每个生成的 request 现在携带其 operation 自己的 API 版本前缀（v3.0 或 v4.0，取自规范中该 operation 的 ``APIVersion``）。")
$md.Add("下表按用例调用的 API 版本统计通过率，用于判断失败是否集中在某一个版本上。")
$md.Add("")
$md.Add("| 用例调用的 API 版本 | 用例数 | 通过 | 失败 | 通过率 |")
$md.Add("|---|---:|---:|---:|---:|")
$versionBuckets = [ordered]@{}
foreach ($kv in $allTestResults.GetEnumerator()) {
    # Skipped cases are not part of the pass rate, and must not create a bucket of their own.
    if ($kv.Value -eq "SKIP") { continue }
    $key = $kv.Key
    $versionsText = "-"
    if ($testVersionInfo.ContainsKey($key)) { $versionsText = Format-Versions $testVersionInfo[$key] }
    if (-not $versionBuckets.Contains($versionsText)) {
        $versionBuckets[$versionsText] = @{ Total = 0; Pass = 0; Fail = 0 }
    }
    $bucket = $versionBuckets[$versionsText]
    $bucket.Total++
    if ($kv.Value -eq "PASS") { $bucket.Pass++ } else { $bucket.Fail++ }
}
foreach ($name in ($versionBuckets.Keys | Sort-Object)) {
    $b = $versionBuckets[$name]
    if ($b.Total -eq 0) { continue }
    $rate = "{0:F2}" -f (100 * $b.Pass / $b.Total)
    $md.Add("| $name | $($b.Total) | $($b.Pass) | $($b.Fail) | $rate% |")
}
$md.Add("")

# ---- Auto-generated observations ------------------------------------------------------------
$md.Add("## 七、分析结论与建议")
$md.Add("")
$observations = [System.Collections.Generic.List[string]]::new()

if ($totalFail -eq 0 -and $totalBuild -eq 0 -and $totalTimeout -eq 0) {
    $observations.Add("- 本次运行 **$totalPass 个用例全部通过**（跳过 $totalSkip 个），未发现回归。")
} else {
    $observations.Add("- 共执行 $($fileResults.Count) 个文件 / $totalTestsRun 个用例，通过 $totalPass，失败 $totalFail，跳过 $totalSkip，通过率 **$passRate%**。")

    $topCategory = $failedCases | Group-Object Category | Sort-Object Count -Descending | Select-Object -First 1
    if ($topCategory) {
        $observations.Add("- 失败最集中的原因分类是 **$($topCategory.Name)**（$($topCategory.Count) 个用例），应优先排查。")
    }

    # Version concentration: is one API version responsible for most failures?
    $v3Fail = 0; $v3Total = 0; $v4Fail = 0; $v4Total = 0
    foreach ($kv in $allTestResults.GetEnumerator()) {
        if ($kv.Value -eq "SKIP") { continue }
        $versionsText = "-"
        if ($testVersionInfo.ContainsKey($kv.Key)) { $versionsText = Format-Versions $testVersionInfo[$kv.Key] }
        if ($versionsText -eq "v3.0") { $v3Total++; if ($kv.Value -eq "FAIL") { $v3Fail++ } }
        elseif ($versionsText -eq "v4.0") { $v4Total++; if ($kv.Value -eq "FAIL") { $v4Fail++ } }
    }
    if ($v3Total -gt 0 -or $v4Total -gt 0) {
        $v3Rate = if ($v3Total -gt 0) { "{0:F2}" -f (100 * $v3Fail / $v3Total) } else { "0.00" }
        $v4Rate = if ($v4Total -gt 0) { "{0:F2}" -f (100 * $v4Fail / $v4Total) } else { "0.00" }
        $observations.Add("- 仅调用 v3.0 接口的用例 $v3Total 个，失败率 $v3Rate%；仅调用 v4.0 接口的用例 $v4Total 个，失败率 $v4Rate%。")
        if ($v3Total -gt 0 -and $v4Total -gt 0 -and $v3Rate -gt ($v4Rate * 2) -and $v3Rate -gt 5) {
            $observations.Add("- 失败在 v3.0 用例上明显更密集，但**这并不表示版本前缀寻址有误**：同一批 v3.0 URL 重试即可返回 200，而同路径在 v4.0 下为 404（端点不存在），说明原因是这些 v3.0 端点的**服务端偶发超时**，而非请求打到了错误的版本。判定请以第四节的错误摘要与「单独重跑」列为准。")
        }
    }

    $replayPass = @($failedCases | Where-Object { $_.Replay -like "PASS*" }).Count
    $networkFail = @($failedCases | Where-Object { $_.Category -like "*网络错误*" }).Count
    $transportFail = $networkFail + $replayPass
    if ($transportFail -gt 0) {
        $observations.Add("- 失败中有 $transportFail 个属于**传输层问题**（$networkFail 个网络错误 + $replayPass 个重跑即通过），表现为 ``context deadline exceeded`` / ``TLS handshake timeout``，属于服务端或网络抖动，而非接口契约或请求模型缺陷；建议先按文件重跑以剔除这些噪声，再分析剩余失败。")
    }

    if ($replayPass -gt 0) {
        $observations.Add("- 有 $replayPass 个用例在单独重跑时通过，说明存在**偶发性（服务器抖动）**或共享远程文件状态的因素，建议确认测试之间的远程目录是否需要隔离。")
    }

    if ($totalBuild -gt 0) {
        $buildFiles = ($fileResults | Where-Object { $_.Status -eq "BUILD" } | ForEach-Object { $_.File }) -join ", "
        $observations.Add("- 有 $totalBuild 个文件**编译失败**（$buildFiles），这类文件的所有用例都未真正执行，需先修复编译问题再复测。")
    }
    if ($totalTimeout -gt 0) {
        $timeoutFiles = ($fileResults | Where-Object { $_.Status -eq "TIMEOUT" } | ForEach-Object { $_.File }) -join ", "
        $observations.Add("- 有 $totalTimeout 个文件**超时**（$timeoutFiles）。请提醒测试人员：可能是服务器出现异常，请检查服务端状态。")
    }
}

if ($skippedTests.Count -gt 0) {
    $observations.Add("- 有 $($skippedTests.Count) 个用例被跳过（``t.Skip``），多为多文件上传受限的 LightCells 组装/合并用例，不计入通过率。")
}
$observations.Add("- 详细日志：``$LogDir/results.tsv``（逐文件）、``$LogDir/failures.tsv``（逐失败用例）、``$LogDir/<文件>.log``（原始输出）。")

foreach ($line in $observations) { $md.Add($line) }
$md.Add("")

$md.Add("## 八、跳过的用例")
$md.Add("")
if ($skippedTests.Count -eq 0) {
    $md.Add("本次运行没有被跳过的用例。")
} else {
    $md.Add("| 文件 | 用例 |")
    $md.Add("|---|---|")
    foreach ($s in $skippedTests) { $md.Add("| ``$($s.File)`` | ``$($s.Test)`` |") }
}
$md.Add("")
$md.Add("---")
$md.Add("")
$md.Add("_报告由 ``newtests/run-tests.ps1`` 生成。_")

$reportFileInfo = New-Object System.Text.UTF8Encoding($false)
[System.IO.File]::WriteAllText($reportPath, ($md -join "`n") + "`n", $reportFileInfo)

Write-Step "============================================================" "Cyan"
Write-Step " 报告已生成: $reportPath" "Green"
Write-Step " 逐文件记录: $resultsTsv" "Green"
Write-Step " 失败记录  : $failuresTsv" "Green"
Write-Step "============================================================" "Cyan"

if ($FailIfAnyTestFails -and ($totalFail -gt 0 -or $totalBuild -gt 0 -or $totalTimeout -gt 0)) {
    exit 1
}
exit 0
