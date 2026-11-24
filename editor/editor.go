/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="editor.go">
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

// Package editor provides worksheet-management helpers on top of the
// Cells Cloud v4.0 API.
//
// The v4.0 API exposes worksheet editing as local-multipart operations: the
// spreadsheet is uploaded, the server applies the edit and returns the updated
// workbook as file bytes, which are written to the sink. The cloud WorkbookRef
// style editing from v3.0 has no v4.0 equivalent.
package editor

import (
	"context"
	"encoding/json"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// Worksheet describes one worksheet returned by ListWorksheets.
type Worksheet struct {
	// Name is the worksheet name.
	Name string
	// Type is the sheet type (e.g. "Worksheet").
	Type string
}

// worksheetInfo mirrors the live worksheet list response item.
type worksheetInfo struct {
	WorksheetName string `json:"WorksheetName"`
	SheetType     string `json:"SheetType"`
}

// AddWorksheet adds a new worksheet to a spreadsheet and writes the updated
// workbook to sink. Position the sheet with WithPosition.
func AddWorksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, sheetName string, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	if sheetName == "" {
		return fmt.Errorf("%w: sheet name is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	reqOpts := append([]requests.Option{requests.WithCommonParameter("sheetName", sheetName)}, cfg.ReqOpts...)
	req := requests.NewAddWorksheetToSpreadsheetRequest("", reqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// DeleteWorksheet deletes a worksheet from a spreadsheet and writes the updated
// workbook to sink.
func DeleteWorksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, sheetName string, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	if sheetName == "" {
		return fmt.Errorf("%w: sheet name is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewDeleteWorksheetFromSpreadsheetRequest(sheetName, "", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// RenameWorksheet renames a worksheet in a spreadsheet and writes the updated
// workbook to sink.
func RenameWorksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, oldName, newName string, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	if oldName == "" || newName == "" {
		return fmt.Errorf("%w: old and new sheet names are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewRenameWorksheetInSpreadsheetRequest(oldName, "", newName, cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// MoveWorksheet moves a worksheet to the given zero-based position in a
// spreadsheet and writes the updated workbook to sink.
func MoveWorksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, worksheet string, position int, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	if worksheet == "" {
		return fmt.Errorf("%w: worksheet name is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewMoveWorksheetInSpreadsheetRequest(position, "", worksheet, cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// ListWorksheets returns the worksheets of a spreadsheet.
func ListWorksheets(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, opts ...Option) ([]Worksheet, error) {

	if src == nil {
		return nil, fmt.Errorf("%w: source is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	req := requests.NewGetWorksheetsWithLocalSpreadsheetRequest("", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return nil, err
	}
	var out []worksheetInfo
	if err := json.Unmarshal(resp.Body, &out); err != nil {
		return nil, err
	}
	ws := make([]Worksheet, 0, len(out))
	for _, w := range out {
		ws = append(ws, Worksheet{Name: w.WorksheetName, Type: w.SheetType})
	}
	return ws, nil
}

// CreateSpreadsheet creates a new empty workbook and writes it to sink.
// Optionally initialize it from a cloud template file with WithTemplate.
func CreateSpreadsheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	sink datasource.DataSink, opts ...Option) error {

	if sink == nil {
		return fmt.Errorf("%w: sink is required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	req := requests.NewCreateSpreadsheetRequest(cfg.ReqOpts...)

	resp, err := sdkutil.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}
