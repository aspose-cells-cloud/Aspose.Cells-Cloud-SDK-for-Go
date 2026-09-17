package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/dataprocessing"
	"asposecellscloud/internal/testutil"
)

// TestImportData_CSV tests the v4.0 ImportCSV function which is the
// equivalent of the v3.0 ImportDataIntoSpreadsheet for CSV files.
func TestImportData_CSV(t *testing.T) {
	client, capture := testutil.NewServer(t, "imported")
	sink := &datasource.BytesSink{}

	err := dataprocessing.ImportCSV(context.Background(), client,
		datasource.BytesSource([]byte("A,B,C\n1,2,3")),
		datasource.BytesSource([]byte("template-xlsx")),
		sink,
		"Sheet1",
		"E3")
	if err != nil {
		t.Fatalf("dataprocessing.ImportCSV failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/import/data/csv" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/import/data/csv", c.Method, c.Path)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if got := c.Query.Get("startcell"); got != "E3" {
		t.Errorf("startcell = %q, want E3", got)
	}
	if got := string(c.Files["datafile"]); got != "A,B,C\n1,2,3" {
		t.Errorf("datafile part = %q, want CSV data", got)
	}
}

// TestImportData_JSON tests the v4.0 ImportJSON function which is the
// equivalent of the v3.0 importjson endpoint for JSON data.
func TestImportData_JSON(t *testing.T) {
	client, capture := testutil.NewServer(t, "imported")
	sink := &datasource.BytesSink{}

	err := dataprocessing.ImportJSON(context.Background(), client,
		datasource.BytesSource([]byte(`[{"A":1,"B":2}]`)),
		datasource.BytesSource([]byte("template-xlsx")),
		sink,
		"Sheet1",
		"A1")
	if err != nil {
		t.Fatalf("dataprocessing.ImportJSON failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/import/data/json" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/import/data/json", c.Method, c.Path)
	}
	if got := string(c.Files["datafile"]); got != `[{"A":1,"B":2}]` {
		t.Errorf("datafile part = %q, want JSON data", got)
	}
}

// TestImportData_XML tests the v4.0 ImportXML function which is the
// equivalent of the v3.0 importxml endpoint for XML data.
func TestImportData_XML(t *testing.T) {
	client, capture := testutil.NewServer(t, "imported")
	sink := &datasource.BytesSink{}

	err := dataprocessing.ImportXML(context.Background(), client,
		datasource.BytesSource([]byte(`<root><row><A>1</A></row></root>`)),
		datasource.BytesSource([]byte("template-xlsx")),
		sink,
		"Sheet1",
		"A1")
	if err != nil {
		t.Fatalf("dataprocessing.ImportXML failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/import/data/xml" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/import/data/xml", c.Method, c.Path)
	}
	if got := string(c.Files["datafile"]); got != `<root><row><A>1</A></row></root>` {
		t.Errorf("datafile part = %q, want XML data", got)
	}
}

// TestImportData_Generic tests the v4.0 ImportData function which infers
// the data format from the content.
func TestImportData_Generic(t *testing.T) {
	client, capture := testutil.NewServer(t, "imported")
	sink := &datasource.BytesSink{}

	err := dataprocessing.ImportData(context.Background(), client,
		datasource.BytesSource([]byte("A,B\n1,2")),
		datasource.BytesSource([]byte("template-xlsx")),
		sink,
		"Sheet1",
		"B2")
	if err != nil {
		t.Fatalf("dataprocessing.ImportData failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/import/data" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/import/data", c.Method, c.Path)
	}
}

// TestImportData_Validation tests validation of import parameters.
func TestImportData_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}

	tests := []struct {
		name string
		fn   func() error
	}{
		{"nil datafile", func() error {
			return dataprocessing.ImportCSV(ctx, client, nil, src, sink, "S", "A1")
		}},
		{"nil template", func() error {
			return dataprocessing.ImportCSV(ctx, client, src, nil, sink, "S", "A1")
		}},
		{"nil sink", func() error {
			return dataprocessing.ImportCSV(ctx, client, src, src, nil, "S", "A1")
		}},
		{"empty worksheet", func() error {
			return dataprocessing.ImportCSV(ctx, client, src, src, sink, "", "A1")
		}},
		{"empty startcell", func() error {
			return dataprocessing.ImportCSV(ctx, client, src, src, sink, "S", "")
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.fn(), asposecellscloud.ErrInvalidParam) {
				t.Errorf("expected ErrInvalidParam")
			}
		})
	}
}
