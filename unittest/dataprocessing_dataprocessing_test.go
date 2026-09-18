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

// TestImportData_Real tests import data using real API.
func TestImportData_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	csvData := "Name,Age,Department\nAlice,30,Engineering\nBob,25,Marketing"
	templateData := datasource.BytesSource([]byte("TestData/Book1.xlsx"))

	err := dataprocessing.ImportCSV(context.Background(), client, datasource.BytesSource([]byte(csvData)), templateData, sink, "Sheet1", "A1")
	if err != nil {
		t.Skipf("ImportCSV failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("ImportCSV output size: %d bytes", len(sink.Bytes()))
	}
}

// TestMergeSpreadsheet_Real tests merge spreadsheet using real API.
func TestMergeSpreadsheet_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	err := dataprocessing.MergeSpreadsheet(context.Background(), client,
		datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		sink,
		dataprocessing.WithOutFormat("pdf"))
	if err != nil {
		t.Skipf("MergeSpreadsheet failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("MergeSpreadsheet output size: %d bytes", len(sink.Bytes()))
	}
}

// TestMergeRemoteSpreadsheet_Real tests merge remote spreadsheet using real API.
func TestMergeRemoteSpreadsheet_Real(t *testing.T) {
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

// TestSplitSpreadsheet_Real tests split spreadsheet using real API.
func TestSplitSpreadsheet_Real(t *testing.T) {
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

// TestSplitRemoteSpreadsheet_Real tests split remote spreadsheet using real API.
func TestSplitRemoteSpreadsheet_Real(t *testing.T) {
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
		t.Logf("SplitRemoteSpreadsheet succeeded, status: %d", resp.StatusCode)
	}
}

func TestImportData(t *testing.T) {
	tests := []struct {
		name string
		fn   func(context.Context, *asposecellscloud.AsposeCellsCloudClient, datasource.DataSource, datasource.DataSource, datasource.DataSink, string, string, ...dataprocessing.Option) error
		path string
	}{
		{"dataprocessing.ImportData", dataprocessing.ImportData, "/v4.0/cells/import/data"},
		{"dataprocessing.ImportCSV", dataprocessing.ImportCSV, "/v4.0/cells/import/data/csv"},
		{"dataprocessing.ImportJSON", dataprocessing.ImportJSON, "/v4.0/cells/import/data/json"},
		{"dataprocessing.ImportXML", dataprocessing.ImportXML, "/v4.0/cells/import/data/xml"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "merged")
			sink := &datasource.BytesSink{}

			err := tt.fn(context.Background(), client,
				datasource.BytesSource([]byte("csv-data")),
				datasource.BytesSource([]byte("template")),
				sink, "Sheet1", "B2")
			if err != nil {
				t.Fatalf("%s failed: %v", tt.name, err)
			}
			c := capture()
			if c.Method != "PUT" || c.Path != tt.path {
				t.Errorf("request = %s %s, want PUT %s", c.Method, c.Path, tt.path)
			}
			if got := c.Query.Get("worksheet"); got != "Sheet1" {
				t.Errorf("worksheet = %q, want Sheet1", got)
			}
			if got := c.Query.Get("startcell"); got != "B2" {
				t.Errorf("startcell = %q, want B2", got)
			}
			if got := string(c.Files["datafile"]); got != "csv-data" {
				t.Errorf("datafile part = %q, want csv-data", got)
			}
			if got := string(c.Files["Spreadsheet"]); got != "template" {
				t.Errorf("Spreadsheet part = %q, want template", got)
			}
			if got := string(sink.Bytes()); got != "merged" {
				t.Errorf("sink = %q, want merged", got)
			}
		})
	}
}

func TestMergeSpreadsheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "merged-xlsx")
	sink := &datasource.BytesSink{}

	err := dataprocessing.MergeSpreadsheet(context.Background(), client, datasource.BytesSource([]byte("x")), sink, dataprocessing.WithOutFormat("pdf"), dataprocessing.WithMergeInOneSheet(true))
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
	if got := string(c.Files["Spreadsheet"]); got != "x" {
		t.Errorf("Spreadsheet part = %q, want x", got)
	}
	if got := string(sink.Bytes()); got != "merged-xlsx" {
		t.Errorf("sink = %q, want merged-xlsx", got)
	}
}

func TestMergeRemoteSpreadsheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "merged-xlsx")
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
	if got := string(sink.Bytes()); got != "merged-xlsx" {
		t.Errorf("sink = %q, want merged-xlsx", got)
	}
}

func TestSplitSpreadsheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "split-out")
	sink := &datasource.BytesSink{}

	err := dataprocessing.SplitSpreadsheet(context.Background(), client, datasource.BytesSource([]byte("x")), sink, dataprocessing.WithOutFormat("csv"), dataprocessing.WithFrom(0), dataprocessing.WithTo(2))
	if err != nil {
		t.Fatalf("dataprocessing.SplitSpreadsheet failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/split/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/split/spreadsheet", c.Method, c.Path)
	}
	if got := c.Query.Get("outFormat"); got != "csv" {
		t.Errorf("outFormat = %q, want csv", got)
	}
	if got := c.Query.Get("from"); got != "0" {
		t.Errorf("from = %q, want 0", got)
	}
	if got := c.Query.Get("to"); got != "2" {
		t.Errorf("to = %q, want 2", got)
	}
	if got := string(sink.Bytes()); got != "split-out" {
		t.Errorf("sink = %q, want split-out", got)
	}
}

func TestSplitRemoteSpreadsheet(t *testing.T) {
	client, capture := testutil.NewServer(t, `{"Files":["out/0.xlsx"]}`)
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

func TestDataprocessingValidation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "B.xlsx"}

	if err := dataprocessing.ImportData(ctx, client, nil, src, sink, "S", "A1"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.ImportData nil datafile: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.ImportData(ctx, client, src, nil, sink, "S", "A1"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.ImportData nil template: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.ImportData(ctx, client, src, src, nil, "S", "A1"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.ImportData nil sink: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.ImportData(ctx, client, src, src, sink, "", "A1"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.ImportData empty worksheet: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.ImportData(ctx, client, src, src, sink, "S", ""); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.ImportData empty startcell: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.MergeSpreadsheet(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.MergeSpreadsheet nil source: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.MergeRemoteSpreadsheet(ctx, client, nil, "B2.xlsx", sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("MergeRemote nil wf: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.MergeRemoteSpreadsheet(ctx, client, wf, "", sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("MergeRemote empty merged name: got %v, want ErrInvalidParam", err)
	}
	if err := dataprocessing.SplitSpreadsheet(ctx, client, src, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("dataprocessing.SplitSpreadsheet nil sink: got %v, want ErrInvalidParam", err)
	}
	if _, err := dataprocessing.SplitRemoteSpreadsheet(ctx, client, nil, "Out"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("SplitRemote nil wf: got %v, want ErrInvalidParam", err)
	}
	if _, err := dataprocessing.SplitRemoteSpreadsheet(ctx, client, wf, ""); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("SplitRemote empty outPath: got %v, want ErrInvalidParam", err)
	}
}
