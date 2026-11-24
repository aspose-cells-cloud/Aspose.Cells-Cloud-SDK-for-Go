/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="report.go">
*   Copyright (c) 2026 Aspose.Cells Cloud
* </copyright>
* <summary>
*   Permission is hereby granted, free of charge, to any person obtaining a copy
*  of this software and associated documentation files (the "Software"), to deal
*  in the Software without restriction, including without limitation the rights
*  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  copies of the Software, and to permit persons to whom the Software is
*  furnished to do so, subject to the following conditions:
*
*  The above copyright notice and this permission notice shall be included in all
*  copies or substantial portions of the Software.
*
*  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*  SOFTWARE.
* </summary>
-------------------------------------------------------------------------------------------------------------------- **/

// Package reporting provides AI-based report analysis and calculation
// helpers on top of the Cells Cloud v4.0 API.
//
// The v4.0 AI endpoints used here are live:
//
//   - PUT /v4.0/cells/ai/report/analysis        (ReportAnalysis)
//   - PUT /v4.0/cells/ai/summarize/spreadsheet  (Summarize)
//   - PUT /v4.0/cells/calculate/aggergate/color (AggregateByColor)
//   - PUT /v4.0/cells/calculate/math            (MathCalculate)
//
// ReportAnalysis and Summarize return a markdown text body; AggregateByColor
// returns a JSON body. Decode them with (*RichResponse).ToBytes / .GetJSON.
package reporting

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// ReportAnalysis performs an AI report analysis on a spreadsheet and returns
// the raw response body (markdown text describing the workbook).
func ReportAnalysis(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, opts ...Option) (*asposecellscloud.RichResponse, error) {

	if src == nil {
		return nil, fmt.Errorf("%w: source is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	req := requests.NewReportAIAnalysisRequest("", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	return sdkutil.DoChecked(ctx, client, req)
}

// Summarize generates an AI summary of a spreadsheet and writes the result
// (markdown text) to sink.
func Summarize(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewSummarizeSpreadsheetRequest("", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// AggregateByColor aggregates cells by color in a spreadsheet and returns the
// raw response body. Use WithWorksheet/WithRange/WithOperation/
// WithColorPosition to scope the aggregation.
func AggregateByColor(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, opts ...Option) (*asposecellscloud.RichResponse, error) {

	if src == nil {
		return nil, fmt.Errorf("%w: source is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	req := requests.NewAggregateCellsByColorRequest("", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	return sdkutil.DoChecked(ctx, client, req)
}

// MathCalculate runs a math calculation over value (e.g. a cell range) in a
// spreadsheet and returns the raw response body. operation is one of "Add",
// "Minus", "Multiply", "Divide" or "Percentage". Use WithWorksheet /
// WithRange to scope the calculation.
func MathCalculate(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, operation, value string, opts ...Option) (*asposecellscloud.RichResponse, error) {

	if src == nil {
		return nil, fmt.Errorf("%w: source is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	req := requests.NewMathCalculateRequest(operation, "", value, cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	return sdkutil.DoChecked(ctx, client, req)
}

// CalculateFormula evaluates formula against worksheet of a local spreadsheet
// and returns the raw response body (the computed result stream). Use
// WithRegion / WithPassword to pass optional parameters.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/calculate/formula.
func CalculateFormula(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, worksheet, formula string, opts ...Option) (*asposecellscloud.RichResponse, error) {

	if src == nil {
		return nil, fmt.Errorf("%w: source is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	req := requests.NewCalculationFormulaRequest(formula, "", worksheet, cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	return sdkutil.DoChecked(ctx, client, req)
}

// CategorizeSpreadsheet uses AI to classify the values of targetColumn into
// logical groups and returns the resulting spreadsheet as the raw response
// body. Use WithSheetName to scope the column to a worksheet and
// WithNewColumnName to name the generated category column.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/ai/categorize/spreadsheet.
func CategorizeSpreadsheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, targetColumn string, opts ...Option) (*asposecellscloud.RichResponse, error) {

	if src == nil {
		return nil, fmt.Errorf("%w: source is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	req := requests.NewCategorizeSpreadsheetRequest("", targetColumn, cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	return sdkutil.DoChecked(ctx, client, req)
}
