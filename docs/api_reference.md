# API Reference

## Overview

The Aspose.Cells Cloud SDK for Go provides 464 API operations across 43 controllers, all accessible through a unified client interface.

The endpoint tables below list each operation's `Path` as declared in the specification. Every operation
is served under its own API version — the `APIVersion` the spec gives for that operation (`v3.0` or
`v4.0`) — so the request the SDK sends is `/{APIVersion}{Path}`: `/v4.0/cells/convert/spreadsheet`,
`/v3.0/cells/{name}`, and so on. The prefix is part of the generated request path; the client adds none.

## Client

### Initialization

```go
client := asposecellscloud.NewAsposeCellsCloudClient(
    clientId,      // string: Aspose Cloud Client ID
    clientSecret,  // string: Aspose Cloud Client Secret
    baseURL,       // string: API base URL (e.g., "https://api.aspose.cloud")
    opts...,       // ...AsposeCellsCloudClientOption: optional configuration
)
```

### Client Options

| Option | Description |
|--------|-------------|
| `WithTimeout(d time.Duration)` | Set request timeout (default: 30s) |
| `WithRetries(n int)` | Set retry count with exponential backoff |
| `WithHeader(key, value string)` | Set global request header |

### Executing Requests

```go
responses, err := client.Do(context.Background(), request)
// responses is []*RichResponse
// err is *SDKError on failure
```

### Response

```go
type RichResponse struct {
    StatusCode int                 // HTTP status code
    Headers    map[string][]string // Response headers
    Body       []byte              // Raw response body
}

// Parse JSON response
var result models.CellsCloudResponse
resp.GetJSON(&result)

// Get response as string
text := resp.ToString()

// Get raw bytes
data := resp.ToBytes()
```

## Controllers

### CellsController (40 operations)

Cell-level operations on worksheets.

| Operation | Method | Path |
|-----------|--------|------|
| `PostClearContents` | POST | `/cells/{name}/worksheets/{sheetName}/cells/clearcontents` |
| `PostClearFormats` | POST | `/cells/{name}/worksheets/{sheetName}/cells/clearformats` |
| `PostCopyCellIntoCell` | POST | `/cells/{name}/worksheets/{sheetName}/cells/{cellName}/copy` |
| `PostSetCellHtmlString` | POST | `/cells/{name}/worksheets/{sheetName}/cells/{cellName}/htmlstring` |
| `PostSetCellRangeValue` | POST | `/cells/{name}/worksheets/{sheetName}/cells` |
| `PostUpdateWorksheetRangeStyle` | POST | `/cells/{name}/worksheets/{sheetName}/cells/style` |
| `PostWorksheetMerge` | POST | `/cells/{name}/worksheets/{sheetName}/cells/merge` |
| `PostWorksheetUnmerge` | POST | `/cells/{name}/worksheets/{sheetName}/cells/unmerge` |
| `PostCellCharacters` | POST | `/cells/{name}/worksheets/{sheetName}/cells/{cellName}/characters` |
| `GetWorksheetColumns` | GET | `/cells/{name}/worksheets/{sheetName}/columns` |
| `GetWorksheetRows` | GET | `/cells/{name}/worksheets/{sheetName}/cells/rows` |
| `GetWorksheetCell` | GET | `/cells/{name}/worksheets/{sheetName}/cells/{cellName}` |
| `GetWorksheetCellStyle` | GET | `/cells/{name}/worksheets/{sheetName}/cells/{cellName}/style` |

### WorkbookController (25 operations)

Workbook-level operations.

| Operation | Method | Path |
|-----------|--------|------|
| `PostWorkbookSaveAs` | POST | `/cells/{name}/saveAs` |
| `PostWorkbookMerge` | POST | `/cells/{name}/merge` |
| `PostWorkbookSplit` | POST | `/cells/{name}/split` |
| `PostWorkbookProtect` | POST | `/cells/{name}/protection` |
| `PostWorkbookEncrypt` | POST | `/cells/{name}/encryption` |
| `PostWorkbookDecrypt` | POST | `/cells/{name}/decryption` |
| `PostWorkbookSettings` | POST | `/cells/{name}/settings` |
| `GetWorkbook` | GET | `/cells/{name}` |
| `GetWorkbookSettings` | GET | `/cells/{name}/settings` |
| `CreateWorkbook` | PUT | `/cells/{name}` |

### WorksheetsController (39 operations)

Worksheet management.

| Operation | Method | Path |
|-----------|--------|------|
| `PutAddNewWorksheet` | PUT | `/cells/{name}/worksheets/{sheetName}` |
| `DeleteWorksheet` | DELETE | `/cells/{name}/worksheets/{sheetName}` |
| `PostCopyWorksheet` | POST | `/cells/{name}/worksheets/{sheetName}/copy` |
| `PostRenameWorksheet` | POST | `/cells/{name}/worksheets/{sheetName}/rename` |
| `PostMoveWorksheet` | POST | `/cells/{name}/worksheets/{sheetName}/move` |
| `PostHideWorksheet` | POST | `/cells/{name}/worksheets/{sheetName}/hide` |
| `PostUnhideWorksheet` | POST | `/cells/{name}/worksheets/{sheetName}/unhide` |
| `GetWorksheet` | GET | `/cells/{name}/worksheets/{sheetName}` |
| `GetWorksheets` | GET | `/cells/{name}/worksheets` |

### ConversionController (42 operations)

Format conversion and export.

| Operation | Method | Path |
|-----------|--------|------|
| `PostWorkbookSaveAs` | POST | `/cells/{name}/saveAs` |
| `PutConvertWorkbook` | PUT | `/cells/convert` |
| `GetWorksheetWithFormat` | GET | `/cells/{name}/worksheets/{sheetName}` |
| `PostWorkbookExportAs` | POST | `/cells/{name}/export` |
| `PostWorkbookToPdf` | POST | `/cells/{name}/toPdf` |
| `PostWorkbookToHtml` | POST | `/cells/{name}/toHtml` |

### ChartsController (24 operations)

Chart creation and management.

| Operation | Method | Path |
|-----------|--------|------|
| `PutWorksheetAddChart` | PUT | `/cells/{name}/worksheets/{sheetName}/charts` |
| `DeleteWorksheetChart` | DELETE | `/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}` |
| `GetWorksheetChart` | GET | `/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}` |
| `GetWorksheetCharts` | GET | `/cells/{name}/worksheets/{sheetName}/charts` |
| `PostWorksheetChart` | POST | `/cells/{name}/worksheets/{sheetName}/charts/{chartIndex}` |

### DataProcessingController (23 operations)

Data import, export, merge, split.

| Operation | Method | Path |
|-----------|--------|------|
| `PostWorkbookMerge` | POST | `/cells/{name}/merge` |
| `PostWorkbookSplit` | POST | `/cells/{name}/split` |
| `PostImportData` | POST | `/cells/{name}/importdata` |
| `PostWorkbookProtect` | POST | `/cells/{name}/protection` |

### File & Storage Controllers

File management in cloud storage.

| Operation | Method | Path |
|-----------|--------|------|
| `UploadFile` | PUT | `/cells/storage/file/{path}` |
| `DownloadFile` | GET | `/cells/storage/file/{path}` |
| `CopyFile` | PUT | `/cells/storage/file/copy/{path}` |
| `MoveFile` | PUT | `/cells/storage/file/move/{path}` |
| `DeleteFile` | DELETE | `/cells/storage/file/{path}` |
| `CreateFolder` | PUT | `/cells/storage/folder/{path}` |
| `DeleteFolder` | DELETE | `/cells/storage/folder/{path}` |
| `GetFilesList` | GET | `/cells/storage/folder/{path}` |
| `GetDiscUsage` | GET | `/cells/storage/disc` |

### Other Notable Controllers

| Controller | Operations | Key Features |
|-----------|-----------|-------------|
| `TextProcessingController` | 24 | Add/extract/convert/trim text, word case |
| `PivotTablesController` | 21 | Pivot table CRUD and filtering |
| `LightCellsController` | 15 | Lightweight batch operations |
| `SearchController` | 14 | Text search and replace |
| `RangesController` | 14 | Named range operations |
| `AutoFilterController` | 13 | AutoFilter and date/custom filters |
| `ListObjectsController` | 13 | List objects (tables) management |
| `ConditionalFormattingsController` | 9 | Format conditions, data bars |
| `ShapesController` | 8 | Shape management |
| `PicturesController` | 7 | Picture insert/update/delete |
| `SparklineGroupsController` | 6 | Sparkline management |
| `AIController` | 5 | AI translation, summarization |
| `BatchController` | 5 | Batch convert/protect/lock/split |

## Common Parameters

These optional parameters are available across most operations:

| Parameter | Type | Description |
|-----------|------|-------------|
| `folder` | `string` | Remote folder path |
| `storageName` | `string` | Storage name (default: empty) |
| `password` | `string` | File password for encrypted files |
| `region` | `string` | Locale setting (e.g., `en-US`) |

## Error Handling

```go
responses, err := client.Do(context.Background(), request)
if err != nil {
    if sdkErr, ok := err.(*asposecellscloud.SDKError); ok {
        fmt.Printf("SDK Error [%d]: %s\n", sdkErr.Code, sdkErr.Message)
    }
    return
}
```
