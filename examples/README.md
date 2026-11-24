# Aspose.Cells Cloud Go SDK examples

These examples target the new-model SDK (module `asposecellscloud`, v4.0 API)
and replace the legacy `Example_*.go` samples written against the v26 SDK.

Each sample keeps its original file name and stays a standalone `main`
program. Every file carries a `//go:build ignore` tag so that it can be run
individually with `go run` without colliding with the other `main` functions in
`go build ./...`.

## Prerequisites

Set your credentials before running anything:

```bash
export ProductClientId=your-client-id
export ProductClientSecret=your-client-secret
```

`CellsCloudApiBaseUrl` is optional; when it is empty the SDK uses
`https://api.aspose.cloud/v4.0`.

## Running

Run the examples from this directory so the bundled fixture files resolve:

```bash
cd examples
go run ./Example_QuickStart.go
go run ./Example_ConvertExcel.go
go run ./Example_BatchProcessing.go
go run ./Example_CloudStorageManagement.go
go run ./Example_CreateReportFiles.go
go run ./Example_ExcelSplit.go
go run ./Example_Management.go
go run ./Example_MergeExcelFiles.go
go run ./Example_Protection.go
go run ./Example_SplitExcelFiles.go
go run ./Example_TextProcessing.go
```

On Windows the same commands work with backslashes, e.g.
`go run .\Example_QuickStart.go`.

## Sample map

| Legacy sample | New-model API |
| --- | --- |
| `Example_QuickStart.go` | `converter.Convert` |
| `Example_ConvertExcel.go` | `converter.Convert`, `converter.Workbook`, `converter.Worksheet` |
| `Example_BatchProcessing.go` | `requests.NewPostBatch*Request` |
| `Example_CloudStorageManagement.go` | `requests.New*FileRequest`, `requests.New*FolderRequest` |
| `Example_CreateReportFiles.go` | `requests` (smart markers, cell values, chart) |
| `Example_ExcelSplit.go` | `dataprocessing.SplitSpreadsheet` |
| `Example_Management.go` | `editor.Add/Rename/Move/DeleteWorksheet` |
| `Example_MergeExcelFiles.go` | `dataprocessing.MergeSpreadsheet`, `requests.NewPostWorkbooksMergeRequest` |
| `Example_Protection.go` | `requests.NewProtect/UnprotectSpreadsheetRequest` |
| `Example_SplitExcelFiles.go` | `dataprocessing.SplitSpreadsheet`, `dataprocessing.SplitRemoteSpreadsheet` |
| `Example_TextProcessing.go` | `requests.NewTrimCharacter/UpdateWordCase/AddText/ConvertTextRequest` |

## What changed in the new model

- Input and output are expressed with `datasource.DataSource` /
  `datasource.DataSink` (`FilePathSource`, `BytesSource`, `ReaderSource`,
  `FilePathSink`, `BytesSink`, `UrlPathSink`) instead of the
  `CellsCloudOption{OptionName: "LocalOutPath", ...}` convention.
- The convenience facades (`converter`, `dataprocessing`, `editor`,
  `datacleansing`, `searcher`, `reporting`) wrap the generated requests and
  write the response bytes to a sink.
- Raw operations use the generated `requests` package with
  `asposecellscloud.DoChecked`, which validates the request and checks the HTTP
  status in one step.
- The v4.0 local split endpoints return a zip archive; the samples extract it.
- The v4.0 local merge endpoint merges the worksheets of a single workbook; the
  legacy "map of local files" merge form is no longer available.
