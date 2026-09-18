# newtests — integration tests for the v4.0 (new model) SDK

These are the **live** integration tests for the `asposecellscloud` client in the
repository root. They are a one-to-one port of the legacy suite in
[`../integrationtests`](../integrationtests): the same 47 files, the same **509**
test functions, the same test names, and the same `TestData` inputs — rewritten
against the new-model API (`requests.NewXxxRequest(...)` +
`asposecellscloud.DoChecked(...)` instead of `new(XxxRequest)` +
`CellsApi.Xxx(...)`).

Unlike the legacy suite, every test here performs real HTTP calls against the
Aspose.Cells Cloud service. There are no mocks or fixtures.

## Requirements

| Environment variable | Meaning |
|---|---|
| `CellsCloudClientId` | client id from <https://dashboard.aspose.cloud> |
| `CellsCloudClientSecret` | client secret |
| `CellsCloudApiBaseUrl` | e.g. `https://api.aspose.cloud` |

The client defaults to API version **v4.0** (the legacy suite pinned `v3.0`).

## Running

```bash
export CellsCloudClientId=...
export CellsCloudClientSecret=...
export CellsCloudApiBaseUrl=https://api.aspose.cloud

cd newtests
go test ./...                          # the whole suite
go test ./... -run TestAddText -v      # a single group
go test ./... -run TestAddText_AddText -v
```

`TestData` is a symlink to `../integrationtests/TestData`, so both suites share
one copy of the 61 sample workbooks. Run the tests from inside `newtests/` so the
relative `TestData/...` paths used by the tests resolve.

Each passing test prints `<n>\t<TestName>`, matching the legacy output format:

```
=== RUN   TestAddText_AddText
1	TestAddText_AddText
--- PASS: TestAddText_AddText (0.83s)
```

## How the port was produced

`../tools/convert_integrationtests.py` generated these files from the legacy
suite. It is kept in the repository so the port can be reproduced or re-run
after the specification changes:

```bash
python3 tools/convert_integrationtests.py   # from the repository root
```

The script reads four sources: the OpenAPI specification
(`aspose.cells.cloud.specification.json`) for parameter names and required-ness,
`requests/*.go` for the constructor argument order and option value types,
`models/*.go` for body-model field types, and the legacy tests themselves. It
re-runs `gofmt -w newtests/` at the end, so its output is already formatted.

**Note:** constructor positional parameters do *not* follow the specification's
parameter order (for `ConvertSpreadsheet`, the spec lists `[Spreadsheet, format]`
but the constructor is `(format, Spreadsheet)`). The converter therefore matches
parameters **by name**, never by position.

### Mapping rules

| Legacy | New model |
|---|---|
| `request := new(XxxRequest)` + field assignment | `requests.NewXxxRequest(<required, by name>, requests.WithCommonParameter("<spec name>", <value>)...)` |
| `_, httpResponse, err := CellsApi.Xxx(request)` | `_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)` |
| `if err != nil {…} else if httpResponse.StatusCode < 200 \|\| > 299 {t.Fail()} else {…}` | `if err != nil {…} else {…}` — `DoChecked` already fails on a non-2xx status |
| `var f = new(Font); f.Size = int64(16)` | `f := &models.Font{}; f.Size = asposecellscloud.Int32Ptr(16)` |
| `new(UploadFileRequest)` + `CellsApi.UploadFile(...)` | `mustUploadFile(t, path, localFile, storageName)` |
| `GetBaseTest().CellsApi` | `GetBaseTest().Client` |

`base_test.go` carries the shared helpers: `GetBaseTest`, `GetTestNumber`, the
legacy `FromBealoonToString` / `FromIntToString` helpers, `mustUploadFile`, and
`intPtr`. The last one exists because the SDK's `ptr.go` provides
`Int32Ptr`/`Int64Ptr`/`Float64Ptr`/`BoolPtr` but no `*int` helper, while 104
generated option parameters take `*int`.

## Known differences from the legacy suite

These are **deliberate**, and none of them are fixed here — they are defects in
the new model, out of scope for a test port. Each is called out in a comment at
the point of use.

1. **13 tests are skipped** — `TestLightCells_PostAssemble_*` (9) and
   `TestLightCells_PostMerge_*` (4). Each uploads two workbooks at once
   (`assemblytest.xlsx` + `datasource.xlsx`). The legacy `PostAssembleRequest.File`
   and `PostMergeRequest.File` were a `map[string]string` and every entry became
   its own multipart part; the v4.0 requests expose a single `File string`, so a
   multi-file assembly cannot be expressed. These tests are emitted with `t.Skip`
   and a comment rather than being silently reduced to one file. The test
   functions still exist, so the count stays at 509.
2. **`AppliedStep.AppliedOperate` cannot be set.** It is declared
   `*models.AppliedOperate` where `AppliedOperate` is an interface, so no value
   can ever satisfy it. Affects
   `TestDataProcessingController_PostDataTransformation`. The assignment is
   dropped; the rest of the request is sent as before.
3. **`UnpivotColumn.AppliedOperateType` no longer exists** in the new model, so
   that one assignment is dropped. Affects the same test as above.
4. **`models.Color.A/R/G/B` are declared `[]byte`** (they hold 0–255 colour
   components). The legacy values were `int64`. The port writes
   `[]byte{48}` so the value survives the declared type; the wire encoding may
   still differ from what the server expects.
5. **Upload failures now abort the test.** The legacy upload preamble discarded
   its error, so a missing remote file produced confusing downstream failures.
   `mustUploadFile` returning a non-nil error triggers `t.Fatal`.
6. **Runtime differences are expected.** The legacy suite targeted v3.0; this one
   targets v4.0. The port guarantees that the code compiles and that the API
   calls are equivalent — it does not guarantee every test returns 2xx. Individual
   endpoint behaviour may legitimately differ between versions.

## Maintenance

`newtests/` is generated output. Change `tools/convert_integrationtests.py` and
regenerate rather than hand-editing these files — hand edits are lost on the next
run.
