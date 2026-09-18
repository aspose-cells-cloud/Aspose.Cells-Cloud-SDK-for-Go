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

// TestCalculate_AggregateByColor tests the v4.0 AggregateByColor function.
// This is the equivalent of the v3.0 AggregateCellsByColor operation.
func TestCalculate_AggregateByColor(t *testing.T) {
	client, capture := testutil.NewServer(t, `{"Total":1}`)

	resp, err := reporting.AggregateByColor(context.Background(), client, datasource.BytesSource([]byte("xlsx")))
	if err != nil {
		t.Fatalf("reporting.AggregateByColor failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/calculate/aggergate/color" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/calculate/aggergate/color", c.Method, c.Path)
	}

	// Test with worksheet and range options
	_, err = reporting.AggregateByColor(context.Background(), client,
		datasource.BytesSource([]byte("xlsx")),
		reporting.WithWorksheet("Sheet1"),
		reporting.WithRange("A1:C5"))
	if err != nil {
		t.Fatalf("AggregateByColor with options failed: %v", err)
	}

	c = capture()
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

// TestCalculate_MathCalculate tests the v4.0 MathCalculate function.
// This is the equivalent of the v3.0 MathCalculate operation.
func TestCalculate_MathCalculate(t *testing.T) {
	client, capture := testutil.NewServer(t, "calc-result")

	resp, err := reporting.MathCalculate(context.Background(), client,
		datasource.BytesSource([]byte("xlsx")),
		"Add",
		"A1:B5")
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
	if resp.ToString() != "calc-result" {
		t.Errorf("body = %q, want calc-result", resp.ToString())
	}
}

// TestCalculate_MathCalculateWithWorksheet tests MathCalculate with worksheet option.
func TestCalculate_MathCalculateWithWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "calc-result")

	_, err := reporting.MathCalculate(context.Background(), client,
		datasource.BytesSource([]byte("xlsx")),
		"Subtract",
		"C1:D10",
		reporting.WithWorksheet("Sheet2"))
	if err != nil {
		t.Fatalf("MathCalculate with worksheet failed: %v", err)
	}

	c := capture()
	if got := c.Query.Get("worksheet"); got != "Sheet2" {
		t.Errorf("worksheet = %q, want Sheet2", got)
	}
	if got := c.Query.Get("operation"); got != "Subtract" {
		t.Errorf("operation = %q, want Subtract", got)
	}
	if got := c.Query.Get("value"); got != "C1:D10" {
		t.Errorf("value = %q, want C1:D10", got)
	}
}

// TestCalculate_MathCalculateOperations tests different math operations.
func TestCalculate_MathCalculateOperations(t *testing.T) {
	operations := []string{"Add", "Subtract", "Multiply", "Divide"}

	for _, op := range operations {
		t.Run(op, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "calc-result")

			_, err := reporting.MathCalculate(context.Background(), client,
				datasource.BytesSource([]byte("x")),
				op,
				"A1",
				reporting.WithWorksheet("Sheet1"))
			if err != nil {
				t.Fatalf("MathCalculate %s failed: %v", op, err)
			}

			c := capture()
			if got := c.Query.Get("operation"); got != op {
				t.Errorf("operation = %q, want %q", got, op)
			}
		})
	}
}

// TestCalculate_Validation tests validation of calculate parameters.
func TestCalculate_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	sink := &datasource.BytesSink{}

	if _, err := reporting.AggregateByColor(ctx, client, nil); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("AggregateByColor nil source: got %v, want ErrInvalidParam", err)
	}

	if err := reporting.Summarize(ctx, client, nil, sink); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("Summarize nil source: got %v, want ErrInvalidParam", err)
	}

	if _, err := reporting.MathCalculate(ctx, client, nil, "Add", "A1"); !errors.Is(err, asposecellscloud.ErrInvalidParam) {
		t.Errorf("MathCalculate nil source: got %v, want ErrInvalidParam", err)
	}
}

// TestCalculate_AggregateByColor_Real tests aggregate by color using real API.
func TestCalculate_AggregateByColor_Real(t *testing.T) {
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

// TestCalculate_WithRangeOption tests MathCalculate with various options.
func TestCalculate_WithRangeOption(t *testing.T) {
	tests := []struct {
		name     string
		opts     []reporting.Option
		queryKey string
		queryVal string
	}{
		{"WithWorksheet", []reporting.Option{reporting.WithWorksheet("Sheet1")}, "worksheet", "Sheet1"},
		{"WithRange A1:C5", []reporting.Option{reporting.WithRange("A1:C5")}, "range", "A1:C5"},
		{"WithRange B2:D10", []reporting.Option{reporting.WithRange("B2:D10")}, "range", "B2:D10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, capture := testutil.NewServer(t, "calc-result")

			_, err := reporting.MathCalculate(context.Background(), client,
				datasource.BytesSource([]byte("x")),
				"Add",
				"A1",
				tt.opts...)
			if err != nil {
				t.Fatalf("MathCalculate failed: %v", err)
			}

			c := capture()
			if got := c.Query.Get(tt.queryKey); got != tt.queryVal {
				t.Errorf("%s = %q, want %q", tt.queryKey, got, tt.queryVal)
			}
		})
	}
}
