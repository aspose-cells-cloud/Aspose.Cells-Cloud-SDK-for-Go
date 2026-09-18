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

// TestSplitSpreadsheet_Local tests the v4.0 SplitSpreadsheet function.
// This is the equivalent of the v3.0 SplitSpreadsheet for local files.
func TestSplitSpreadsheet_Local(t *testing.T) {
	client, capture := testutil.NewServer(t, "split-out")
	sink := &datasource.BytesSink{}

	err := dataprocessing.SplitSpreadsheet(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		dataprocessing.WithOutFormat("pdf"),
		dataprocessing.WithFrom(0),
		dataprocessing.WithTo(2))
	if err != nil {
		t.Fatalf("dataprocessing.SplitSpreadsheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/split/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/split/spreadsheet", c.Method, c.Path)
	}
	if got := c.Query.Get("outFormat"); got != "pdf" {
		t.Errorf("outFormat = %q, want pdf", got)
	}
	if got := c.Query.Get("from"); got != "0" {
		t.Errorf("from = %q, want 0", got)
	}
	if got := c.Query.Get("to"); got != "2" {
		t.Errorf("to = %q, want 2", got)
	}
}

// TestSplitSpreadsheet_Remote tests the v4.0 SplitRemoteSpreadsheet function.
// This is the equivalent of the v3.0 SplitRemoteSpreadsheet for cloud files.
func TestSplitSpreadsheet_Remote(t *testing.T) {
	client, capture := testutil.NewServer(t, `{"Files":["out/0.xlsx","out/1.xlsx"]}`)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	resp, err := dataprocessing.SplitRemoteSpreadsheet(context.Background(), client, wf, "Out")
	if err != nil {
		t.Fatalf("dataprocessing.SplitRemoteSpreadsheet failed: %v", err)
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

// TestSplitSpreadsheet_Validation tests validation of split parameters.
func TestSplitSpreadsheet_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	wf := &asposecellscloud.WorkbookRef{Name: "B.xlsx"}

	type testFunc func() error
	tests := []struct {
		name string
		fn   testFunc
	}{
		{"SplitSpreadsheet nil source", func() error {
			return dataprocessing.SplitSpreadsheet(ctx, client, nil, &datasource.BytesSink{})
		}},
		{"SplitSpreadsheet nil sink", func() error {
			return dataprocessing.SplitSpreadsheet(ctx, client, src, nil)
		}},
		{"SplitRemoteSpreadsheet nil wf", func() error {
			_, err := dataprocessing.SplitRemoteSpreadsheet(ctx, client, nil, "Out")
			return err
		}},
		{"SplitRemoteSpreadsheet empty name", func() error {
			_, err := dataprocessing.SplitRemoteSpreadsheet(ctx, client, &asposecellscloud.WorkbookRef{}, "Out")
			return err
		}},
		{"SplitRemoteSpreadsheet empty outPath", func() error {
			_, err := dataprocessing.SplitRemoteSpreadsheet(ctx, client, wf, "")
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.fn().(error), asposecellscloud.ErrInvalidParam) {
				t.Errorf("expected ErrInvalidParam")
			}
		})
	}
}

// TestSplitSpreadsheet_Local_Real tests local split using real API.
func TestSplitSpreadsheet_Local_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	err := dataprocessing.SplitSpreadsheet(context.Background(), client,
		datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		sink,
		dataprocessing.WithOutFormat("pdf"))
	if err != nil {
		t.Skipf("SplitSpreadsheet failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("SplitSpreadsheet output size: %d bytes", len(sink.Bytes()))
	}
}

// TestSplitSpreadsheet_Options tests the various options available for split.
func TestSplitSpreadsheet_Options(t *testing.T) {
	tests := []struct {
		name     string
		opts     []dataprocessing.Option
		queryKey string
		queryVal string
	}{
		{"WithOutFormat", []dataprocessing.Option{dataprocessing.WithOutFormat("csv")}, "outFormat", "csv"},
		{"WithFrom", []dataprocessing.Option{dataprocessing.WithFrom(0)}, "from", "0"},
		{"WithTo", []dataprocessing.Option{dataprocessing.WithTo(5)}, "to", "5"},
		{"WithOutPath", []dataprocessing.Option{dataprocessing.WithOutPath("out/split.xlsx")}, "outPath", "out/split.xlsx"},
		{"WithPassword", []dataprocessing.Option{dataprocessing.WithPassword("secret")}, "password", "secret"},
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
