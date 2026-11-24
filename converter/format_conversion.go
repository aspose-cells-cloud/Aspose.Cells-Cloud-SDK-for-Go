/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="format_conversion.go">
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

package converter

import (
	"context"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// Convert converts a local spreadsheet (src) to the target format and writes
// the result bytes to sink.
//
//	src  DataSource: local file / bytes / reader
//	sink DataSink:   write the converted file to a path, byte buffer, or URL
//	format:          target format, e.g. convert.FormatPDF (see format.go)
//
// Uses the v4.0 ConvertSpreadsheet endpoint
// (PUT /v4.0/cells/convert/spreadsheet?format=...).
func Convert(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, format string, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: src and sink must not be nil", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewConvertSpreadsheetRequest(format, "", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// ConvertToPDF converts a spreadsheet to PDF.
func ConvertToPDF(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatPDF, opts...)
}

// ConvertToCSV converts a spreadsheet to CSV.
func ConvertToCSV(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatCSV, opts...)
}

// ConvertToJSON converts a spreadsheet to JSON.
func ConvertToJSON(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatJSON, opts...)
}

// ConvertToHTML converts a spreadsheet to HTML.
func ConvertToHTML(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatHTML, opts...)
}

// ConvertToXlsx converts a spreadsheet to XLSX.
func ConvertToXlsx(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatXlsx, opts...)
}

// ConvertToPNG converts a spreadsheet to PNG.
func ConvertToPNG(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatPNG, opts...)
}

// ConvertToDocx converts a spreadsheet to DOCX.
func ConvertToDocx(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatDocx, opts...)
}

// ConvertToSQL converts a spreadsheet to SQL.
func ConvertToSQL(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, opts ...Option) error {
	return Convert(ctx, client, src, sink, FormatSQL, opts...)
}

// The v3.0-only component converts (ConvertWorksheet* / ConvertRangeToPDF /
// ConvertChartToPDF / ConvertTableToPDF) were removed during the v4.0 migration
// because the v4.0 API exposes no equivalent interface.

// Workbook exports a workbook from cloud storage to the target format and
// writes the downloaded bytes to sink. wf locates the cloud file; format is
// the target format (see the Format* constants, e.g. FormatPDF).
func Workbook(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, sink datasource.DataSink, format string, opts ...Option) error {

	if wf == nil || wf.Name == "" || sink == nil {
		return fmt.Errorf("%w: workbook name and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewExportSpreadsheetAsFormatRequest(format, wf.Name, reqOpts...)

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// Worksheet exports a single worksheet of a cloud workbook to the target
// format and writes the downloaded bytes to sink.
func Worksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, worksheet string, sink datasource.DataSink, format string, opts ...Option) error {

	if wf == nil || wf.Name == "" || worksheet == "" || sink == nil {
		return fmt.Errorf("%w: workbook name, worksheet and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewExportWorksheetAsFormatRequest(format, wf.Name, worksheet, reqOpts...)

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// The v3.0-only export operations WorksheetXML and RangeValues were removed
// during the v4.0 migration because the v4.0 API exposes no equivalent
// interface (exportxml and ranges/{...}/value both return 404 under v4.0).
