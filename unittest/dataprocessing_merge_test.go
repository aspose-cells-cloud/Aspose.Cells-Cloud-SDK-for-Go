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

// TestMergeSpreadsheet_Local tests the v4.0 MergeSpreadsheet function.
// This is the equivalent of the v3.0 MergeSpreadsheet for local files.
func TestMergeSpreadsheet_Local(t *testing.T) {
	client, capture := testutil.NewServer(t, "merged")
	sink := &datasource.BytesSink{}

	err := dataprocessing.MergeSpreadsheet(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		dataprocessing.WithOutFormat("pdf"),
		dataprocessing.WithMergeInOneSheet(true))
	if err != nil {
		t.Fatalf("dataprocessing.MergeSpreadsheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/merge/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/merge/spreadsheet", c.Method, c.Path)
	}
	if got := c.Query.Get("outFormat"); got != "pdf" {
		t.Errorf("outFormat = %q, want pdf", got)
	}
	if got := c.Query.Get("mergeInOneSheet"); got != "true" {
		t.Errorf("mergeInOneSheet = %q, want true", got)
	}
	if got := string(c.Files["Spreadsheet"]); got != "source-xlsx" {
		t.Errorf("Spreadsheet part = %q, want source-xlsx", got)
	}
}

// TestMergeSpreadsheet_Remote tests the v4.0 MergeRemoteSpreadsheet function.
// This is the equivalent of the v3.0 MergeRemoteSpreadsheet for cloud files.
func TestMergeSpreadsheet_Remote(t *testing.T) {
	client, capture := testutil.NewServer(t, "merged")
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	err := dataprocessing.MergeRemoteSpreadsheet(context.Background(), client, wf, "Book2.xlsx", sink)
	if err != nil {
		t.Fatalf("dataprocessing.MergeRemoteSpreadsheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/merge/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/merge/spreadsheet", c.Method, c.Path)
	}
	if got := c.Query.Get("mergedSpreadsheet"); got != "Book2.xlsx" {
		t.Errorf("mergedSpreadsheet = %q, want Book2.xlsx", got)
	}
	if got := c.Query.Get("folder"); got != "TestData/In" {
		t.Errorf("folder = %q, want TestData/In", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
}

// TestMergeSpreadsheet_Validation tests validation of merge parameters.
func TestMergeSpreadsheet_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "B.xlsx"}

	tests := []struct {
		name string
		fn   func() error
	}{
		{"MergeSpreadsheet nil source", func() error {
			return dataprocessing.MergeSpreadsheet(ctx, client, nil, sink)
		}},
		{"MergeSpreadsheet nil sink", func() error {
			return dataprocessing.MergeSpreadsheet(ctx, client, src, nil)
		}},
		{"MergeRemoteSpreadsheet nil wf", func() error {
			return dataprocessing.MergeRemoteSpreadsheet(ctx, client, nil, "B2.xlsx", sink)
		}},
		{"MergeRemoteSpreadsheet empty name", func() error {
			return dataprocessing.MergeRemoteSpreadsheet(ctx, client, &asposecellscloud.WorkbookRef{}, "B2.xlsx", sink)
		}},
		{"MergeRemoteSpreadsheet nil sink", func() error {
			return dataprocessing.MergeRemoteSpreadsheet(ctx, client, wf, "B2.xlsx", nil)
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

// TestMergeSpreadsheet_Options tests the various options available for merge.
func TestMergeSpreadsheet_Options(t *testing.T) {
	tests := []struct {
		name     string
		opts     []dataprocessing.Option
		queryKey string
		queryVal string
	}{
		{"WithOutFormat", []dataprocessing.Option{dataprocessing.WithOutFormat("csv")}, "outFormat", "csv"},
		{"WithMergeInOneSheet true", []dataprocessing.Option{dataprocessing.WithMergeInOneSheet(true)}, "mergeInOneSheet", "true"},
		{"WithMergeInOneSheet false", []dataprocessing.Option{dataprocessing.WithMergeInOneSheet(false)}, "mergeInOneSheet", "false"},
		{"WithOutPath", []dataprocessing.Option{dataprocessing.WithOutPath("out/merged.xlsx")}, "outPath", "out/merged.xlsx"},
		{"WithPassword", []dataprocessing.Option{dataprocessing.WithPassword("secret")}, "password", "secret"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "merged")
			sink := &datasource.BytesSink{}

			err := dataprocessing.MergeSpreadsheet(context.Background(), client,
				datasource.BytesSource([]byte("x")),
				sink,
				tt.opts...)
			if err != nil {
				t.Fatalf("MergeSpreadsheet failed: %v", err)
			}

			c := capture()
			if got := c.Query.Get(tt.queryKey); got != tt.queryVal {
				t.Errorf("%s = %q, want %q", tt.queryKey, got, tt.queryVal)
			}
		})
	}
}
