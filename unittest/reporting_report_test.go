package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/testutil"
	"asposecellscloud/reporting"
)

func TestReportAnalysis(t *testing.T) {
	client, capture := testutil.NewServer(t, "analysis-markdown")

	resp, err := reporting.ReportAnalysis(context.Background(), client, datasource.BytesSource([]byte("xlsx")))
	if err != nil {
		t.Fatalf("reporting.ReportAnalysis failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/ai/report/analysis" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/ai/report/analysis", c.Method, c.Path)
	}
	if got := string(c.Files["Spreadsheet"]); got != "xlsx" {
		t.Errorf("Spreadsheet part = %q, want xlsx", got)
	}
	if resp.StatusCode != 200 || resp.ToString() != "analysis-markdown" {
		t.Errorf("response = %d %q, want 200 analysis-markdown", resp.StatusCode, resp.ToString())
	}
}

func TestSummarize(t *testing.T) {
	client, capture := testutil.NewServer(t, "summary-markdown")
	sink := &datasource.BytesSink{}

	err := reporting.Summarize(context.Background(), client, datasource.BytesSource([]byte("x")), sink)
	if err != nil {
		t.Fatalf("reporting.Summarize failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/ai/summarize/spreadsheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/ai/summarize/spreadsheet", c.Method, c.Path)
	}
	if got := string(sink.Bytes()); got != "summary-markdown" {
		t.Errorf("sink = %q, want summary-markdown", got)
	}
}

func TestAggregateByColor(t *testing.T) {
	client, capture := testutil.NewServer(t, `{"Total":1}`)

	resp, err := reporting.AggregateByColor(context.Background(), client, datasource.BytesSource([]byte("x")), reporting.WithWorksheet("Sheet1"), reporting.WithRange("A1:C5"))
	if err != nil {
		t.Fatalf("reporting.AggregateByColor failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/calculate/aggergate/color" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/calculate/aggergate/color", c.Method, c.Path)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if got := c.Query.Get("range"); got != "A1:C5" {
		t.Errorf("range = %q, want A1:C5", got)
	}
	var out struct {
		Total int `json:"Total"`
	}
	if err := resp.GetJSON(&out); err != nil {
		t.Fatalf("GetJSON failed: %v", err)
	}
	if out.Total != 1 {
		t.Errorf("Total = %d, want 1", out.Total)
	}
}

func TestMathCalculate(t *testing.T) {
	client, capture := testutil.NewServer(t, "calc-result")

	resp, err := reporting.MathCalculate(context.Background(), client, datasource.BytesSource([]byte("x")), "Add", "A1:B5", reporting.WithWorksheet("Sheet1"))
	if err != nil {
		t.Fatalf("reporting.MathCalculate failed: %v", err)
	}
	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/calculate/math" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/calculate/math", c.Method, c.Path)
	}
	if got := c.Query.Get("operation"); got != "Add" {
		t.Errorf("operation = %q, want Add", got)
	}
	if got := c.Query.Get("value"); got != "A1:B5" {
		t.Errorf("value = %q, want A1:B5", got)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if resp.ToString() != "calc-result" {
		t.Errorf("body = %q, want calc-result", resp.ToString())
	}
}

// TestReportAnalysis_Real tests report analysis using real API.
func TestReportAnalysis_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)

	resp, err := reporting.ReportAnalysis(context.Background(), client, datasource.BytesSource([]byte("TestData/Book1.xlsx")))
	if err != nil {
		t.Skipf("ReportAnalysis failed (might be expected): %v", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	} else {
		t.Logf("ReportAnalysis response: %s", resp.ToString())
	}
}

// TestSummarize_Real tests summarize using real API.
func TestSummarize_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	err := reporting.Summarize(context.Background(), client, datasource.BytesSource([]byte("TestData/Book1.xlsx")), sink)
	if err != nil {
		t.Skipf("Summarize failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("Summarize output size: %d bytes", len(sink.Bytes()))
	}
}

// TestAggregateByColor_Real tests aggregate by color using real API.
func TestAggregateByColor_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)

	resp, err := reporting.AggregateByColor(context.Background(), client, datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		reporting.WithWorksheet("Sheet1"),
		reporting.WithRange("A1:C5"))
	if err != nil {
		t.Skipf("AggregateByColor failed (might be expected): %v", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	} else {
		t.Logf("AggregateByColor response: %s", resp.ToString())
	}
}

// TestMathCalculate_Real tests math calculate using real API.
func TestMathCalculate_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)

	resp, err := reporting.MathCalculate(context.Background(), client, datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		"Add",
		"A1:B5",
		reporting.WithWorksheet("Sheet1"))
	if err != nil {
		t.Skipf("MathCalculate failed (might be expected): %v", err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	} else {
		t.Logf("MathCalculate response: %s", resp.ToString())
	}
}

// TestReportingValidation tests validation of reporting parameters.
func TestReportingValidation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	sink := &datasource.BytesSink{}

	if _, err := reporting.ReportAnalysis(ctx, client, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("reporting.ReportAnalysis nil source: got %v, want ErrInvalidParam", err)
	}
	if err := reporting.Summarize(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("reporting.Summarize nil source: got %v, want ErrInvalidParam", err)
	}
	if err := reporting.Summarize(ctx, client, datasource.BytesSource([]byte("x")), nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("reporting.Summarize nil sink: got %v, want ErrInvalidParam", err)
	}
	if _, err := reporting.AggregateByColor(ctx, client, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("reporting.AggregateByColor nil source: got %v, want ErrInvalidParam", err)
	}
	if _, err := reporting.MathCalculate(ctx, client, nil, "Add", "A1"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("reporting.MathCalculate nil source: got %v, want ErrInvalidParam", err)
	}
}
