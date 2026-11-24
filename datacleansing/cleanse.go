/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="cleanse.go">
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

// Package datacleansing provides helpers that clean a spreadsheet by removing
// blank rows, blank columns, blank worksheets and duplicate rows.
//
// Cleansing is exposed as local multipart operations: the spreadsheet is
// uploaded, the server removes the targeted content and returns the cleaned
// workbook as file bytes, which are written to the sink.
package datacleansing

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// RemoveBlankRows removes blank rows from a spreadsheet and writes the cleaned
// workbook to sink.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/blank-rows.
func RemoveBlankRows(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.Option) asposecellscloud.RequestOption {
			return requests.NewRemoveSpreadsheetBlankRowsRequest("", opts...)
		}, opts...)
}

// RemoveBlankColumns removes blank columns from a spreadsheet and writes the
// cleaned workbook to sink.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/blank-columns.
func RemoveBlankColumns(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.Option) asposecellscloud.RequestOption {
			return requests.NewRemoveSpreadsheetBlankColumnsRequest("", opts...)
		}, opts...)
}

// RemoveBlankWorksheets removes blank worksheets from a spreadsheet and writes
// the cleaned workbook to sink.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/blank-worksheets.
func RemoveBlankWorksheets(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.Option) asposecellscloud.RequestOption {
			return requests.NewRemoveSpreadsheetBlankWorksheetsRequest("", opts...)
		}, opts...)
}

// RemoveDuplicates removes duplicate rows within a worksheet, cell range or
// table of a spreadsheet and writes the cleaned workbook to sink. Scope the
// removal with WithWorksheet / WithRange / WithTable.
//
// Uses the v4.0 endpoint PUT /v4.0/cells/remove/duplicates.
func RemoveDuplicates(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return clean(ctx, client, src, sink,
		func(opts ...requests.Option) asposecellscloud.RequestOption {
			return requests.NewRemoveDuplicatesRequest("", opts...)
		}, opts...)
}

// clean uploads a spreadsheet, runs the cleansing request and writes the
// cleaned workbook to sink.
func clean(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink,
	builder func(...requests.Option) asposecellscloud.RequestOption, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := builder(cfg.ReqOpts...)

	// Attach the spreadsheet to the multipart "Spreadsheet" field.
	if setter, ok := req.(sdkutil.SpreadsheetSetter); ok {
		setter.SetSpreadsheetBytes(data, "Spreadsheet")
	}

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}
