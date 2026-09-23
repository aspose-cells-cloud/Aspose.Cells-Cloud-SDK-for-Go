package unittest_test

import (
	"encoding/json"
	"testing"

	asposecellscloud "asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

// The specification spells inheritance with ParentName: ImportOption has nine
// subclasses (ImportIntArrayOption, ImportCSVDataOption, ...), Shape has 22,
// SaveOptions has 19, FormatCondition has one. Go has no subclass relation between
// structs -- embedding is not a subtype relation -- so a *ImportIntArrayOption is not
// assignable to a *ImportOption.
//
// That made every request parameter taking such a model unusable from a caller who
// holds a subclass. PostImportData asserted
//
//	if val, exist := cfg.Params["importOption"].(*models.ImportOption); exist {
//
// which never matches for a subclass: the value was dropped, the request carried a
// null body, and the service answered HTTP 400 "Error reading JObject from Json".
// The same assertion sat in PostWorkbookSaveAs and PutWorksheetShape, and the two
// positional parameters (PostWorksheetShape.dto, PutWorksheetConditionalFormatting
// .formatcondition) could not even be called with a subclass.
//
// Every concrete base now gets a generated `<Base>Like` interface -- the marker
// method is declared on the base and promoted to each subclass through embedding, so
// the family is closed and the parameter keeps a real type instead of interface{}.
// These tests pin the parameters and the JSON body the service sees; the live suite
// (newtests) exercises the ImportData and Shape endpoints end to end.

// The families, and that membership is transitive: PdfSaveOptions inherits through
// PaginatedSaveOptions, three levels below SaveOptions.
var (
	_ models.ImportOptionLike    = &models.ImportOption{}
	_ models.ImportOptionLike    = &models.ImportIntArrayOption{}
	_ models.ImportOptionLike    = &models.ImportCSVDataOption{}
	_ models.SaveOptionsLike     = &models.SaveOptions{}
	_ models.SaveOptionsLike     = &models.PaginatedSaveOptions{}
	_ models.SaveOptionsLike     = &models.PdfSaveOptions{}
	_ models.ShapeLike           = &models.Shape{}
	_ models.ShapeLike           = &models.ArcShape{}
	_ models.FormatConditionLike = &models.FormatCondition{}
	_ models.FormatConditionLike = &models.StyleFormatCondition{}
)

// decodeBody marshals a request body the way the client does and returns it both raw
// and decoded, so a test can assert on fields and on "not null" without repeating the
// unmarshalling.
func decodeBody(t *testing.T, body interface{}) (string, map[string]interface{}) {
	t.Helper()

	raw, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("unmarshal request body: %v (body was %s)", err, raw)
	}
	return string(raw), decoded
}

// TestPostImportDataRequestCarriesAConcreteImportOption is the reported defect: the
// caller holds an ImportIntArrayOption, which is an ImportOption by the
// specification but not by Go.
func TestPostImportDataRequestCarriesAConcreteImportOption(t *testing.T) {
	option := &models.ImportIntArrayOption{}
	option.DestinationWorksheet = "Sheet1"
	option.ImportDataType = "IntArray"
	option.FirstRow = asposecellscloud.Int32Ptr(3)
	option.FirstColumn = asposecellscloud.Int32Ptr(1)
	option.IsVertical = asposecellscloud.BoolPtr(true)
	option.IsInsert = asposecellscloud.BoolPtr(true)
	option.Data = []interface{}{int64(1), int64(2), int64(3), int64(4)}

	request := requests.NewPostImportDataRequest("Book1.xlsx",
		requests.WithCommonParameter("importOption", option),
		requests.WithCommonParameter("folder", "TestData/In"),
	)

	raw, decoded := decodeBody(t, request.GetJSONBody())
	if raw == "null" {
		t.Fatalf("request body is null: the concrete ImportOption was dropped")
	}
	// The subclass's own properties.
	if decoded["FirstRow"] != float64(3) || decoded["FirstColumn"] != float64(1) {
		t.Fatalf("body lost the ImportIntArrayOption position: %s", raw)
	}
	if decoded["IsVertical"] != true || decoded["IsInsert"] != true {
		t.Fatalf("body lost the ImportIntArrayOption flags: %s", raw)
	}
	if data, ok := decoded["Data"].([]interface{}); !ok || len(data) != 4 {
		t.Fatalf("body lost the imported data: %s", raw)
	}
	// And the properties it inherits from ImportOption.
	if decoded["DestinationWorksheet"] != "Sheet1" || decoded["ImportDataType"] != "IntArray" {
		t.Fatalf("body lost the inherited ImportOption properties: %s", raw)
	}
}

// TestPostImportDataRequestStillCarriesTheBaseImportOption keeps the previously
// working call working: the base is a member of its own family.
func TestPostImportDataRequestStillCarriesTheBaseImportOption(t *testing.T) {
	request := requests.NewPostImportDataRequest("Book1.xlsx",
		requests.WithCommonParameter("importOption", &models.ImportOption{
			DestinationWorksheet: "Sheet1",
			ImportDataType:       "IntArray",
		}),
	)

	raw, decoded := decodeBody(t, request.GetJSONBody())
	if decoded["DestinationWorksheet"] != "Sheet1" || decoded["ImportDataType"] != "IntArray" {
		t.Fatalf("body = %s, want the base ImportOption properties", raw)
	}
}

// TestPutWorksheetShapeRequestCarriesAConcreteShape covers the third assertion the
// generator emitted.
func TestPutWorksheetShapeRequestCarriesAConcreteShape(t *testing.T) {
	shape := &models.ArcShape{}
	shape.Name = "Arc1"
	shape.Text = "hello"

	request := requests.NewPutWorksheetShapeRequest("Book1.xlsx", "Sheet1",
		requests.WithCommonParameter("shapeDTO", shape),
		requests.WithCommonParameter("DrawingType", "arc"),
		requests.WithCommonParameter("upperLeftRow", asposecellscloud.Int32Ptr(1)),
	)

	raw, decoded := decodeBody(t, request.GetJSONBody())
	if decoded["Name"] != "Arc1" || decoded["Text"] != "hello" {
		t.Fatalf("body = %s, want the ArcShape properties", raw)
	}
}

// TestPostWorkbookSaveAsRequestCarriesATwoLevelDescendant proves the marker method is
// promoted through more than one level of embedding: PdfSaveOptions ->
// PaginatedSaveOptions -> SaveOptions.
func TestPostWorkbookSaveAsRequestCarriesATwoLevelDescendant(t *testing.T) {
	saveOptions := &models.PdfSaveOptions{}
	saveOptions.SaveFormat = "pdf" // declared on SaveOptions, two levels up
	saveOptions.DisplayDocTitle = asposecellscloud.BoolPtr(true)

	request := requests.NewPostWorkbookSaveAsRequest("Book1.xlsx", "OutResult/Book1.pdf",
		requests.WithCommonParameter("saveOptions", saveOptions),
	)

	raw, decoded := decodeBody(t, request.GetJSONBody())
	if decoded["SaveFormat"] != "pdf" {
		t.Fatalf("body = %s, want the inherited SaveFormat", raw)
	}
	if decoded["DisplayDocTitle"] != true {
		t.Fatalf("body = %s, want the PdfSaveOptions property", raw)
	}
}

// TestPutWorksheetConditionalFormattingRequestAcceptsAConcreteCondition covers a
// positional parameter, which used to reject a subclass at compile time.
func TestPutWorksheetConditionalFormattingRequestAcceptsAConcreteCondition(t *testing.T) {
	condition := &models.StyleFormatCondition{}
	condition.Type = "CellValue"
	condition.Formula1 = "=A1>10"

	request := requests.NewPutWorksheetConditionalFormattingRequest(
		"A1:B2", condition, "Book1.xlsx", "Sheet1")

	raw, decoded := decodeBody(t, request.GetJSONBody())
	if decoded["Type"] != "CellValue" || decoded["Formula1"] != "=A1>10" {
		t.Fatalf("body = %s, want the StyleFormatCondition properties", raw)
	}
}

// TestPostWorksheetShapeRequestAcceptsAnyShapeFamilyMember is the other positional
// parameter, and pins that the inherited Shape properties travel with the subclass.
func TestPostWorksheetShapeRequestAcceptsAnyShapeFamilyMember(t *testing.T) {
	shape := &models.Oval{}
	shape.Name = "Oval1"
	shape.UpperLeftRow = asposecellscloud.Int32Ptr(2)

	request := requests.NewPostWorksheetShapeRequest(shape, "Book1.xlsx", 0, "Sheet1")

	raw, decoded := decodeBody(t, request.GetJSONBody())
	if decoded["Name"] != "Oval1" || decoded["UpperLeftRow"] != float64(2) {
		t.Fatalf("body = %s, want the Oval properties", raw)
	}
}

// TestShapeFamilyKeepsConcreteFieldAccess documents why the models themselves are
// unchanged: a property typed as the base stays concrete, so its own fields can be
// read and written. Only request parameters -- write-only inputs -- take the family
// interface.
func TestShapeFamilyKeepsConcreteFieldAccess(t *testing.T) {
	shape := &models.ArcShape{}
	shape.Name = "Arc1"

	// Field access through the embedded base still works.
	shape.UpperLeftColumn = asposecellscloud.Int32Ptr(4)
	if shape.Shape.UpperLeftColumn == nil || *shape.Shape.UpperLeftColumn != 4 {
		t.Fatalf("the embedded Shape did not receive UpperLeftColumn")
	}
}
