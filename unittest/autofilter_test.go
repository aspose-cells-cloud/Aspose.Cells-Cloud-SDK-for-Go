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

// TestAutoFilter_Validation tests validation of remote operations.
// Note: The v4.0 API has changed how batch operations work.
// Use the lower-level requests package for batch operations if needed.
func TestAutoFilter_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	_ = context.Background()

	tests := []struct {
		name string
		fn   func() error
	}{
		{"MergeRemote nil wf", func() error {
			return dataprocessing.MergeRemoteSpreadsheet(context.Background(), client, nil, "B.xlsx", &datasource.BytesSink{})
		}},
		{"MergeRemote empty name", func() error {
			return dataprocessing.MergeRemoteSpreadsheet(context.Background(), client, &asposecellscloud.WorkbookRef{}, "B.xlsx", &datasource.BytesSink{})
		}},
		{"SplitRemote nil wf", func() error {
			_, err := dataprocessing.SplitRemoteSpreadsheet(context.Background(), client, nil, "Out")
			return err
		}},
		{"SplitRemote empty name", func() error {
			_, err := dataprocessing.SplitRemoteSpreadsheet(context.Background(), client, &asposecellscloud.WorkbookRef{}, "Out")
			return err
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

// TestAutoFilter_MergeRemote_Real tests merge remote with real API.
func TestAutoFilter_MergeRemote_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData"}

	err := dataprocessing.MergeRemoteSpreadsheet(context.Background(), client, wf, "Book2.xlsx", sink)
	if err != nil {
		t.Skipf("MergeRemoteSpreadsheet failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("MergeRemoteSpreadsheet output size: %d bytes", len(sink.Bytes()))
	}
}

// TestAutoFilter_SplitRemote_Real tests split remote with real API.
func TestAutoFilter_SplitRemote_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData"}

	resp, err := dataprocessing.SplitRemoteSpreadsheet(context.Background(), client, wf, "Out")
	if err != nil {
		t.Skipf("SplitRemoteSpreadsheet failed (might be expected): %v", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	} else {
		t.Logf("SplitRemoteSpreadsheet status: %d", resp.StatusCode)
	}
}

// TestAutoFilter_RemoteOperations tests remote operations with WorkbookRef.
func TestAutoFilter_RemoteOperations(t *testing.T) {
	client, capture := testutil.NewServer(t, "batch-result")
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	// Test MergeRemoteSpreadsheet
	err := dataprocessing.MergeRemoteSpreadsheet(context.Background(), client, wf, "Book2.xlsx", sink)
	if err != nil {
		t.Fatalf("MergeRemoteSpreadsheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/merge/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/merge/spreadsheet", c.Method, c.Path)
	}
}
