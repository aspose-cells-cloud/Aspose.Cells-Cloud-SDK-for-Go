# Development Guide

## Project Overview

This SDK is auto-generated from `aspose.cells.cloud.specification.json`, a comprehensive API specification containing 464 operations and 484 data models.

## Directory Structure

```
cells.cloud-sdk-go-dev/
├── aspose.cells.cloud.specification.json   # Master API specification
├── models/                                  # Generated Go data models
│   └── *.go                                # One file per struct (484 files)
├── requests/                                # Generated API request wrappers
│   └── *_request.go                        # One file per operation (460 files)
├── integrationtests/                        # Integration test suite
│   └── *_test.go                           # One file per test group (46 files)
├── TestingData/                             # Test configuration (JSON)
│   └── */*.json
├── testdata/                                # Test data files (XLSX, CSV, etc.)
├── generate_models.py                       # Model generation script
├── generate_requests.py                     # Request generation script
├── generate_tests.py                        # Test generation script
├── split_models.py                          # Model file splitting script
└── references/                              # Generation rules
    ├── model_generation_rules.md
    ├── request_generation_rules.md
    ├── config_update_rules.md
    ├── test_generation_rules.md
    ├── code_check_rules.md
    └── build_error_fix_rules.md
```

## Code Generation

### Prerequisites

- Python 3.6+
- The API specification file: `aspose.cells.cloud.specification.json`
- Generation scripts in the project root

### Generate Models

```bash
python3 generate_models.py
```

Reads the `Models` array from the specification and writes every type into the single file
`models/model_cells.go`. That monolith is an intermediate artifact — run `split_models.py` next,
which splits it into one file per model type and deletes it.

Type mapping rules:

| Spec DataType | Go Type |
|--------------|---------|
| `String` | `string` |
| `Long` | `*int64` |
| `Integer` | `*int32` |
| `Boolean` | `*bool` |
| `Float`/`Double` | `*float64` |
| `DateTime` | `time.Time` |
| `Byte` | `*int32` |
| `Class` (with Reference) | `*ReferencedType` |
| `Class` (with Reference, model is `IsAbstract`) | `ReferencedType` |
| `Class` (with Reference, model has subclasses) | `ReferencedTypeLike` |
| `Container` (with Reference) | `[]ReferencedType` |
| `Object` | `map[string]interface{}` |
| `Any` | `interface{}` |

Scalar primitives are **pointers** because every generated field carries an unconditional
`omitempty`: `nil` means "unset" and is dropped from the JSON, while a zero that was set explicitly
is still emitted. Strings, `time.Time`, and slices stay value types.

`Byte` is deliberately **not** `[]byte`. The specification uses it only for
`Color.A/R/G/B`, which hold 0–255 colour components, not binary blobs — Go's `encoding/json`
renders `[]byte` as a base64 *string*, so the old mapping put `{"R":"MA=="}` on the wire where the
service expects a number.

A `Class` reference to an **abstract** model (`IsAbstract: true`) is a different case: the model
becomes a Go interface, and a pointer to an interface can never be satisfied, so the field is
typed as the interface itself. An abstract parent is likewise not embedded, so a child declares
the properties it inherits instead of skipping them — otherwise a discriminator such as
`AppliedOperateType` disappears. Each such child gets a generated `MarshalJSON` that fills that
discriminator in from its own type name (the service's converter routes by it and rejects
non-members), so callers pass the data only. See
[`references/model_generation_rules.md`](../references/model_generation_rules.md).

A `Class` reference to a **concrete** model that has subclasses is the third case, and the one Go
makes awkward: `ImportOption` is a struct with nine subclasses, so a `*ImportIntArrayOption` is not
assignable to a `*ImportOption`. Each of the 19 such bases gets a generated family interface
(`models.ImportOptionLike`, `models.ShapeLike`, …) whose marker method is declared on the base and
promoted to every subclass by embedding. Request parameters take that interface rather than a
pointer, so `requests.WithCommonParameter("importOption", &models.ImportIntArrayOption{…})` — a
call that used to drop the value and put a null body on the wire — works. Model *properties* keep
the concrete base type, because callers read them back. The subclasses of a concrete base are told
apart by ordinary data properties (`ImportDataType`), which the caller sets; only an abstract
parent has a derivable discriminator.

### Generate Request Files

```bash
python3 generate_requests.py
```

Reads `Operations` array and generates request wrapper files in `requests/`. Each operation gets:

- `{Name}Request` struct with required + optional fields
- `New{Name}Request()` constructor with validation
- Interface methods: `GetMethod()`, `GetPath()`, `GetQueryParameters()`, `GetJSONBody()`, `GetMultipartForm()`, `GetHeaderParameters()`, `Description()`

### Generate Test Files

```bash
python3 generate_tests.py
```

Reads `TestingData/` JSON files and generates test functions in `integrationtests/`. Matches test parameters against the API spec to correctly categorize required vs optional parameters.

### Split Model Files

```bash
python3 split_models.py
```

Splits the monolithic `models/model_cells.go` produced by `generate_models.py` into individual
files (one struct per file, in `models/`) and deletes the monolith. The two scripts together are
the model-generation step; `models/model_cells.go` is never committed.

## Key Design Patterns

### Request Interface

All request types implement the `RequestOption` interface defined in `request.go`:

```go
type RequestOption interface {
    GetMethod() string
    GetHeaderParameters() map[string]string
    GetPath() string
    GetQueryParameters() url.Values
    GetJSONBody() interface{}
    GetMultipartForm() map[string]interface{}
}
```

### Optional Parameters

Optional parameters are handled generically via `WithCommonParameter`:

```go
func WithCommonParameter(key string, value interface{}) RequestOption {
    return optionFunc(func(c *requestConfig) {
        if c.Params == nil {
            c.Params = make(map[string]interface{})
        }
        c.Params[key] = value
    })
}
```

### Extra Query Parameters (Implicit Extension)

Every request supports custom query parameters on top of the operation's declared query parameters — an implicit special feature for passing extra attributes such as save options.

Set at construction time:

```go
requests.NewConvertSpreadsheetRequest("pdf", "../testdata/Book1.xlsx",
    requests.WithQueryParameter("SaveOptions", `{"CalcMode":"Manual"}`),
    requests.WithQueryParameters(map[string]string{"CheckExcelRestriction": "false"}),
)
```

Or after construction via instance methods:

```go
req := requests.NewConvertSpreadsheetRequest("pdf", "../testdata/Book1.xlsx")
req.AddQueryParameter("CustomAttr", "abc")
req.AddQueryParameters(map[string]string{"CustomAttr2": "def"})
```

These are appended last in `GetQueryParameters()` and merge into the final request URL.

### Parameter Type Rules

| Parameter Type | Required | Optional |
|---------------|----------|----------|
| `string` | `string` (value) | `string` (value) |
| `int` | `int` (value) | `*int` (pointer) |
| `bool` | `bool` (value) | `*bool` (pointer) |
| `float64` | `float64` (value) | `*float64` (pointer) |
| Struct/Class | `*models.Type` (pointer) | `*models.Type` (pointer) |
| Struct/Class whose model has subclasses | `models.Type` (abstract) / `models.TypeLike` (concrete) | same |
| Slice/Container | `[]models.Type` (value) | `[]models.Type` (value) |

## Adding New APIs

1. Add the operation to `aspose.cells.cloud.specification.json`
2. If new models are needed, add them to the `Models` array
3. Run `python3 generate_models.py` then `python3 split_models.py` to generate model files
4. Run `python3 generate_requests.py` to generate request files
5. Add test data to `TestingData/` directory
6. Run `python3 generate_tests.py` to generate test files
7. Run `go build ./...` and `go vet ./...` to verify

## Building

```bash
# Build SDK
go build ./...

# Run static analysis
go vet ./...

# Build integration tests
cd integrationtests && go build ./...
```

## Contributing

1. Follow existing code patterns and naming conventions
2. All struct fields use PascalCase (`CellName`, not `cell_name`)
3. JSON tags match the original spec parameter names
4. Required parameters are validated in constructors
5. Optional pointer parameters have nil checks before dereferencing
6. Run `go vet ./...` before submitting changes
