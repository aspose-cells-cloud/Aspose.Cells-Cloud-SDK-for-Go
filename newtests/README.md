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

Each of these was a defect in the new model rather than a porting choice, so each
is called out in a comment at the point of use. Items 2 and 3 were subsequently
**fixed in the model generator** and now carry no workaround; what remains below
is the record of what the port had to work around at the time.

1. **13 tests are skipped** — `TestLightCells_PostAssemble_*` (9) and
   `TestLightCells_PostMerge_*` (4). Each uploads two workbooks at once
   (`assemblytest.xlsx` + `datasource.xlsx`). The legacy `PostAssembleRequest.File`
   and `PostMergeRequest.File` were a `map[string]string` and every entry became
   its own multipart part; the v4.0 requests expose a single `File string`, so a
   multi-file assembly cannot be expressed. These tests are emitted with `t.Skip`
   and a comment rather than being silently reduced to one file. The test
   functions still exist, so the count stays at 509.
2. **A polymorphic operate could not be attached to its step.** The specification
   models `AppliedStep.AppliedOperate` as an abstract `AppliedOperate` with three
   concrete children (`UnpivotColumn`, `PivotColumn`, `MergeQueries`), and each
   child inherits an `AppliedOperateType` that the service's own JSON converter
   dispatches on. Two model defects made that unusable: the field was declared
   `*models.AppliedOperate` — a *pointer to an interface*, which no concrete model
   can ever satisfy — and because an abstract parent becomes an interface it is
   not embedded, so the children silently lost `AppliedOperateType`. The legacy
   `TestPostDataTransformation` is an empty stub, so the body here is hand-written;
   it used to build a `UnpivotColumn` and drop it on the floor, and the request
   failed with an opaque `HTTP 500 ... Object reference not set to an instance of
   an object` (the service cannot deserialize the operate without its
   discriminator). It is **fixed**, in three parts: the field is now the
   `AppliedOperate` interface itself, so a `*UnpivotColumn` assigns to it
   directly; the children declare their inherited `AppliedOperateType`; and each
   such child fills that discriminator in from its own type name on
   serialization, so the test just builds a `UnpivotColumn` and hands it over. See
   `unittest/models_applied_operate_test.go` in the SDK repository for the pinned
   wire shape, and note the value is a real enum — the service answers a
   non-member with `Could not convert '...' to AppliedOperateType`.
3. **`models.Color.A/R/G/B` were declared `[]byte`** (they hold 0–255 colour
   components; the legacy values were `int64`). This was a defect in the model,
   not a porting choice: `encoding/json` serializes `[]byte` as a base64
   **string**, so a colour filter went out as `{"R":"MA=="}` where the service
   expects a number. It is **fixed** — the generator's `Byte` mapping is now
   `*int32` like `Integer` — so the port simply sets
   `asposecellscloud.Int32Ptr(48)` and no longer needs a workaround. See
   `unittest/models_color_test.go` in the SDK repository for the pinned
   behaviour.
4. **Upload failures now abort the test.** The legacy upload preamble discarded
   its error, so a missing remote file produced confusing downstream failures.
   `mustUploadFile` returning a non-nil error triggers `t.Fatal`.
5. **Runtime differences are expected.** The legacy suite targeted v3.0; this one
   targets v4.0. The port guarantees that the code compiles and that the API
   calls are equivalent — it does not guarantee every test returns 2xx. Individual
   endpoint behaviour may legitimately differ between versions.
6. **A subclass could not be passed where its base was expected.** The legacy
   suite builds `&models.ImportOption{}` for `PostImportData`,
   `&models.SaveOptions{}` for `PostWorkbookSaveAs` and `&models.Shape{}` for
   `PutWorksheetShape`, and the generated requests asserted exactly those types
   (`cfg.Params["importOption"].(*models.ImportOption)`). A subclass — the shape
   the endpoints actually document, e.g. `ImportIntArrayOption` — failed the
   assertion, so the value was **dropped**, the request went out with a null body,
   and the service answered `HTTP 400 Error reading JObject from Json`. Go has no
   subtype relation between structs, so the SDK **fixed** it by generating a family
   interface per concrete base (`models.ImportOptionLike`, `models.ShapeLike`, …),
   which the base and every subclass satisfy by the promoted marker method; the
   parameters now take that interface. See
   `unittest/requests_polymorphic_param_test.go` in the SDK repository. The three
   tests here whose parameter is now a subclass (`TestWorkbookController_`
   `PostImportData`, `TestShapesController_PutWorksheetShape`,
   `TestConversion30_WorkbookSaveAs_pdf_...`) are deliberate: they are the live
   coverage of the family, see Maintenance.

## Maintenance

`newtests/` is generated output. Change `tools/convert_integrationtests.py` and
regenerate rather than hand-editing these files — hand edits are lost on the next
run.

**Known hand-added setup:** three bodies have no counterpart in the legacy suite,
so the converter cannot reproduce them. Re-add them after a regeneration:

- the colour-filter setups in `api_cells_autofiltercontroller_test.go` and
  `api_cells_rangescontroller_test.go` — regenerating emits `nil` for the
  `ColorFilterRequest` argument instead, and the two tests silently stop covering
  the colour path;
- the family-typed parameters noted in point 6 — `TestWorkbookController_`
  `PostImportData` hands over an `ImportIntArrayOption`, `TestShapesController_`
  `PutWorksheetShape` an `ArcShape`, and `TestConversion30_WorkbookSaveAs_pdf_...`
  a `PdfSaveOptions`. The legacy suite uses the base type for all three, so
  regenerating emits `&models.ImportOption{}` / `&models.Shape{}` /
  `&models.SaveOptions{}`: the tests still compile and pass, but the polymorphic
  parameters stop being covered by a live run;
- the `AppliedOperate` assignment in
  `api_cells_dataprocessingcontroller_test.go` (`TestDataProcessingController_`
  `PostDataTransformation`) — regenerating drops the operate, and the request goes
  back to failing with `HTTP 500 Object reference not set to an instance of an
  object`. Its discriminator no longer needs re-adding: `models.UnpivotColumn`
  supplies it on serialization.

**The converter does not run from this repository as-is.** It reads
`aspose.cells.cloud.specification.json`, which is not checked in here (it lives in
the SDK development repository, alongside `references/` and the port
directories); points 2 and 3 above mean the script would also need the model
generator's fix to be present in `models/` to emit correct code. Regenerating
therefore means running it from the development repository — this copy of
`newtests/` is the delivered output.
