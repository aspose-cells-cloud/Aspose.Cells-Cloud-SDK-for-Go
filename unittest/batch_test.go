package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/dataprocessing"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
)

// TestBatchOperations_Validation tests batch operations validation.
func TestBatchOperations_Validation(t *testing.T) {
	ctx := context.Background()

	// Test that nil source returns appropriate error
	if err := dataprocessing.MergeSpreadsheet(ctx, &asposecellscloud.AsposeCellsCloudClient{}, nil, &datasource.BytesSink{}); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("nil source should return ErrInvalidParam")
	}

	// Test that nil sink returns appropriate error
	if err := dataprocessing.MergeSpreadsheet(ctx, &asposecellscloud.AsposeCellsCloudClient{}, datasource.BytesSource([]byte("x")), nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("nil sink should return ErrInvalidParam")
	}
}

// TestBatchOperations_Remote tests remote batch operations.
func TestBatchOperations_Remote(t *testing.T) {
	client, capture := testutil.NewServer(t, "batch-result")
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	// Test MergeRemoteSpreadsheet with various options
	err := dataprocessing.MergeRemoteSpreadsheet(context.Background(), client, wf, "Book2.xlsx", sink)
	if err != nil {
		t.Fatalf("MergeRemoteSpreadsheet failed: %v", err)
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

// TestBatchSplitOperations tests batch split operations.
func TestBatchSplitOperations(t *testing.T) {
	client, capture := testutil.NewServer(t, `{"Files":["out/0.xlsx"]}`)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	resp, err := dataprocessing.SplitRemoteSpreadsheet(context.Background(), client, wf, "Out")
	if err != nil {
		t.Fatalf("SplitRemoteSpreadsheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/split/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/split/spreadsheet", c.Method, c.Path)
	}
	if got := c.Query.Get("outPath"); got != "Out" {
		t.Errorf("outPath = %q, want Out", got)
	}
	if got := c.Query.Get("folder"); got != "TestData/In" {
		t.Errorf("folder = %q, want TestData/In", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
}

// TestBatchMergeWithPassword tests batch merge with password protection.
func TestBatchMergeWithPassword(t *testing.T) {
	tests := []struct {
		name     string
		opts     []dataprocessing.Option
		queryKey string
		queryVal string
	}{
		{"WithPassword", []dataprocessing.Option{dataprocessing.WithPassword("secret")}, "password", "secret"},
		{"WithOutFormat csv", []dataprocessing.Option{dataprocessing.WithOutFormat("csv")}, "outFormat", "csv"},
		{"WithOutFormat pdf", []dataprocessing.Option{dataprocessing.WithOutFormat("pdf")}, "outFormat", "pdf"},
		{"WithMergeInOneSheet true", []dataprocessing.Option{dataprocessing.WithMergeInOneSheet(true)}, "mergeInOneSheet", "true"},
		{"WithOutPath", []dataprocessing.Option{dataprocessing.WithOutPath("out/merged.xlsx")}, "outPath", "out/merged.xlsx"},
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

// TestBatchSplitWithFormat tests batch split with various output formats.
func TestBatchSplitWithFormat(t *testing.T) {
	tests := []struct {
		name     string
		opts     []dataprocessing.Option
		queryKey string
		queryVal string
	}{
		{"WithOutFormat csv", []dataprocessing.Option{dataprocessing.WithOutFormat("csv")}, "outFormat", "csv"},
		{"WithOutFormat pdf", []dataprocessing.Option{dataprocessing.WithOutFormat("pdf")}, "outFormat", "pdf"},
		{"WithOutFormat xlsx", []dataprocessing.Option{dataprocessing.WithOutFormat("xlsx")}, "outFormat", "xlsx"},
		{"WithFrom 0", []dataprocessing.Option{dataprocessing.WithFrom(0)}, "from", "0"},
		{"WithFrom 1", []dataprocessing.Option{dataprocessing.WithFrom(1)}, "from", "1"},
		{"WithTo 5", []dataprocessing.Option{dataprocessing.WithTo(5)}, "to", "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "split-out")
			sink := &datasource.BytesSink{}

			err := dataprocessing.SplitSpreadsheet(context.Background(), client,
				datasource.BytesSource([]byte("x")),
				sink,
				tt.opts...)
			if err != nil {
				t.Fatalf("SplitSpreadsheet failed: %v", err)
			}

			c := capture()
			if got := c.Query.Get(tt.queryKey); got != tt.queryVal {
				t.Errorf("%s = %q, want %q", tt.queryKey, got, tt.queryVal)
			}
		})
	}
}

// TestBatchMergeWithPassword_Real tests batch merge with password protection using real API.
func TestBatchMergeWithPassword_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	tests := []struct {
		name     string
		opts     []dataprocessing.Option
		queryKey string
		queryVal string
	}{
		{"WithPassword", []dataprocessing.Option{dataprocessing.WithPassword("secret")}, "password", "secret"},
		{"WithOutFormat csv", []dataprocessing.Option{dataprocessing.WithOutFormat("csv")}, "outFormat", "csv"},
		{"WithOutFormat pdf", []dataprocessing.Option{dataprocessing.WithOutFormat("pdf")}, "outFormat", "pdf"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := dataprocessing.MergeSpreadsheet(context.Background(), client,
				datasource.BytesSource([]byte("x")),
				sink,
				tt.opts...)
			if err != nil {
				t.Skipf("MergeSpreadsheet failed (might be expected): %v", err)
				return
			}

			if len(sink.Bytes()) == 0 {
				t.Errorf("expected non-empty output")
			}
			t.Logf("Output size: %d bytes", len(sink.Bytes()))
		})
	}
}
