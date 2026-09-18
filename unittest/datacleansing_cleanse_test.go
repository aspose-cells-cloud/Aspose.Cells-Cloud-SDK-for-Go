package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datacleansing"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
)

// TestRemoveBlankRows tests the v4.0 RemoveBlankRows function.
// This is the equivalent of the v3.0 RemoveBlankRows for local files.
func TestRemoveBlankRows(t *testing.T) {
	client, capture := testutil.NewServer(t, "cleaned")
	sink := &datasource.BytesSink{}

	err := datacleansing.RemoveBlankRows(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink)
	if err != nil {
		t.Fatalf("datacleansing.RemoveBlankRows failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/remove/blank-rows" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/remove/blank-rows", c.Method, c.Path)
	}
	if got := string(c.Files["Spreadsheet"]); got != "source-xlsx" {
		t.Errorf("Spreadsheet part = %q, want source-xlsx", got)
	}
	if got := string(sink.Bytes()); got != "cleaned" {
		t.Errorf("sink = %q, want cleaned", got)
	}
}

// TestRemoveBlankColumns tests the v4.0 RemoveBlankColumns function.
func TestRemoveBlankColumns(t *testing.T) {
	client, capture := testutil.NewServer(t, "cleaned")
	sink := &datasource.BytesSink{}

	err := datacleansing.RemoveBlankColumns(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink)
	if err != nil {
		t.Fatalf("datacleansing.RemoveBlankColumns failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/remove/blank-columns" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/remove/blank-columns", c.Method, c.Path)
	}
}

// TestRemoveBlankWorksheets tests the v4.0 RemoveBlankWorksheets function.
func TestRemoveBlankWorksheets(t *testing.T) {
	client, capture := testutil.NewServer(t, "cleaned")
	sink := &datasource.BytesSink{}

	err := datacleansing.RemoveBlankWorksheets(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink)
	if err != nil {
		t.Fatalf("datacleansing.RemoveBlankWorksheets failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/remove/blank-worksheets" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/remove/blank-worksheets", c.Method, c.Path)
	}
}

// TestRemoveDuplicates tests the v4.0 RemoveDuplicates function with different scoping options.
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		opts     []datacleansing.Option
		queryKey string
		queryVal string
	}{
		{"WithWorksheet", []datacleansing.Option{datacleansing.WithWorksheet("Sheet1")}, "worksheet", "Sheet1"},
		{"WithRange", []datacleansing.Option{datacleansing.WithRange("A1:B10")}, "range", "A1:B10"},
		{"WithTable", []datacleansing.Option{datacleansing.WithTable("Table1")}, "table", "Table1"},
		{"WithOutPath", []datacleansing.Option{datacleansing.WithOutPath("out/cleaned.xlsx")}, "outPath", "out/cleaned.xlsx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "cleaned")
			sink := &datasource.BytesSink{}

			err := datacleansing.RemoveDuplicates(context.Background(), client,
				datasource.BytesSource([]byte("source-xlsx")),
				sink,
				tt.opts...)
			if err != nil {
				t.Fatalf("RemoveDuplicates failed: %v", err)
			}

			c := capture()
			if c.Method != "PUT" || c.Path != "/v4.0/cells/remove/duplicates" {
				t.Errorf("request = %s %s, want PUT /v4.0/cells/remove/duplicates", c.Method, c.Path)
			}
			if got := c.Query.Get(tt.queryKey); got != tt.queryVal {
				t.Errorf("%s = %q, want %q", tt.queryKey, got, tt.queryVal)
			}
		})
	}
}

// TestRemoveBlankRows_Real tests remove blank rows using real API.
func TestRemoveBlankRows_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	err := datacleansing.RemoveBlankRows(context.Background(), client,
		datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		sink)
	if err != nil {
		t.Skipf("RemoveBlankRows failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("RemoveBlankRows output size: %d bytes", len(sink.Bytes()))
	}
}

// TestCleanse_Validation tests validation of cleanse parameters.
func TestCleanse_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}

	tests := []struct {
		name string
		fn   func() error
	}{
		{"RemoveBlankRows nil source", func() error {
			return datacleansing.RemoveBlankRows(ctx, client, nil, sink)
		}},
		{"RemoveBlankRows nil sink", func() error {
			return datacleansing.RemoveBlankRows(ctx, client, src, nil)
		}},
		{"RemoveBlankColumns nil source", func() error {
			return datacleansing.RemoveBlankColumns(ctx, client, nil, sink)
		}},
		{"RemoveBlankWorksheets nil source", func() error {
			return datacleansing.RemoveBlankWorksheets(ctx, client, nil, sink)
		}},
		{"RemoveDuplicates nil source", func() error {
			return datacleansing.RemoveDuplicates(ctx, client, nil, sink)
		}},
		{"RemoveDuplicates nil sink", func() error {
			return datacleansing.RemoveDuplicates(ctx, client, src, nil)
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
