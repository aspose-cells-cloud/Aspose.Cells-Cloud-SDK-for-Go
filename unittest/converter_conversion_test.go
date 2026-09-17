package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/converter"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
)

// TestConvert_Spreadsheet tests the v4.0 Convert function for local files.
// This is the equivalent of the v3.0 ConvertSpreadsheet for local files.
func TestConvert_Spreadsheet(t *testing.T) {
	tests := []struct {
		name   string
		format string
		fn     func(context.Context, *asposecellscloud.AsposeCellsCloudClient, datasource.DataSource, datasource.DataSink, ...converter.Option) error
	}{
		{"PDF", converter.FormatPDF, converter.ConvertToPDF},
		{"CSV", converter.FormatCSV, converter.ConvertToCSV},
		{"JSON", converter.FormatJSON, converter.ConvertToJSON},
		{"HTML", converter.FormatHTML, converter.ConvertToHTML},
		{"XLSX", converter.FormatXlsx, converter.ConvertToXlsx},
		{"PNG", converter.FormatPNG, converter.ConvertToPNG},
		{"DOCX", converter.FormatDocx, converter.ConvertToDocx},
		{"SQL", converter.FormatSQL, converter.ConvertToSQL},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "converted")
			sink := &datasource.BytesSink{}

			err := tt.fn(context.Background(), client,
				datasource.BytesSource([]byte("source-xlsx")),
				sink)
			if err != nil {
				t.Fatalf("%s failed: %v", tt.name, err)
			}

			c := capture()
			if c.Method != "PUT" || c.Path != "/v4.0/cells/convert/spreadsheet" {
				t.Errorf("request = %s %s, want PUT /v4.0/cells/convert/spreadsheet", c.Method, c.Path)
			}
			if got := c.Query.Get("format"); got != tt.format {
				t.Errorf("format = %q, want %q", got, tt.format)
			}
			if got := string(c.Files["Spreadsheet"]); got != "source-xlsx" {
				t.Errorf("Spreadsheet part = %q, want source-xlsx", got)
			}
		})
	}
}

// TestConvert_Workbook tests the v4.0 Workbook function for cloud files.
// This is the equivalent of the v3.0 cloud workbook conversion.
func TestConvert_Workbook(t *testing.T) {
	client, capture := testutil.NewServer(t, "converted")
	sink := &datasource.BytesSink{}

	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}
	err := converter.Workbook(context.Background(), client, wf, sink, converter.FormatPDF)
	if err != nil {
		t.Fatalf("converter.Workbook failed: %v", err)
	}

	c := capture()
	if c.Method != "GET" || c.Path != "/v4.0/cells/Book1.xlsx" {
		t.Errorf("request = %s %s, want GET /v4.0/cells/Book1.xlsx", c.Method, c.Path)
	}
	if got := c.Query.Get("format"); got != "pdf" {
		t.Errorf("format = %q, want pdf", got)
	}
	if got := c.Query.Get("folder"); got != "TestData/In" {
		t.Errorf("folder = %q, want TestData/In", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
}

// TestConvert_Worksheet tests the v4.0 Worksheet function for cloud worksheet.
// This is the equivalent of the v3.0 cloud worksheet conversion.
func TestConvert_Worksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "converted")
	sink := &datasource.BytesSink{}

	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}
	err := converter.Worksheet(context.Background(), client, wf, "Sheet1", sink, converter.FormatCSV)
	if err != nil {
		t.Fatalf("converter.Worksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "GET" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1" {
		t.Errorf("request = %s %s, want GET /v4.0/cells/Book1.xlsx/worksheets/Sheet1", c.Method, c.Path)
	}
	if got := c.Query.Get("format"); got != "csv" {
		t.Errorf("format = %q, want csv", got)
	}
}

// TestConvert_Validation tests validation of convert parameters.
func TestConvert_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}

	tests := []struct {
		name string
		fn   func() error
	}{
		{"Convert nil source", func() error {
			return converter.Convert(ctx, client, nil, sink, converter.FormatPDF)
		}},
		{"Convert nil sink", func() error {
			return converter.Convert(ctx, client, src, nil, converter.FormatPDF)
		}},
		{"ConvertToPDF nil source", func() error {
			return converter.ConvertToPDF(ctx, client, nil, sink)
		}},
		{"Workbook nil workbook", func() error {
			return converter.Workbook(ctx, client, nil, sink, converter.FormatPDF)
		}},
		{"Workbook empty name", func() error {
			return converter.Workbook(ctx, client, &asposecellscloud.WorkbookRef{}, sink, converter.FormatPDF)
		}},
		{"Worksheet nil workbook", func() error {
			return converter.Worksheet(ctx, client, nil, "Sheet1", sink, converter.FormatPDF)
		}},
		{"Worksheet empty name", func() error {
			return converter.Worksheet(ctx, client, &asposecellscloud.WorkbookRef{}, "Sheet1", sink, converter.FormatPDF)
		}},
		{"Worksheet empty sheet", func() error {
			return converter.Worksheet(ctx, client, &asposecellscloud.WorkbookRef{Name: "B.xlsx"}, "", sink, converter.FormatPDF)
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
