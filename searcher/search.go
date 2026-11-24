/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="search.go">
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

// Package searcher provides find-and-replace helpers for local and cloud
// spreadsheets.
//
// The v4.0 API exposes search/replace through the content endpoints:
//
//   - Local files:    PUT /v4.0/cells/search/content  and
//     PUT /v4.0/cells/replace/content
//   - Cloud files:    PUT /v4.0/cells/{name}/search/content          (replace)
//     PUT /v4.0/cells/{name}/worksheets/{ws}/search/content
//     PUT /v4.0/cells/{name}/worksheets/{ws}/ranges/{area}/...
//
// Local Search requires a worksheet name (the live endpoint returns 400 without
// one). Local Replace returns the modified workbook as file bytes, so it does
// not round-trip through cloud storage. Cloud Replace modifies the stored file
// in place.
package searcher

import (
	"context"
	"encoding/json"
	"fmt"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// TextItem is a single text match returned by the search/content endpoints.
type TextItem struct {
	// Text is the matched text value.
	Text string
	// Worksheet names the worksheet that contains the match.
	Worksheet string
	// Position locates the match, e.g. "Cell:E35".
	Position string
}

// searchResponse mirrors the live search/content response body:
// {"TextItems":[{"Filename":"...","Worksheet":"...","Position":"...","Content":"..."}]}
type searchResponse struct {
	TextItems []struct {
		Filename  string `json:"Filename"`
		Worksheet string `json:"Worksheet"`
		Position  string `json:"Position"`
		Content   string `json:"Content"`
	} `json:"TextItems"`
}

// Search searches a local spreadsheet for text within the given worksheet and
// returns the matching text items. The live v4.0 endpoint requires the
// worksheet name.
func Search(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, worksheet, text string, opts ...Option) ([]TextItem, error) {

	if src == nil || worksheet == "" || text == "" {
		return nil, fmt.Errorf("%w: source, worksheet and text are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return nil, err
	}

	reqOpts := append([]requests.Option{requests.WithCommonParameter("worksheet", worksheet)}, cfg.ReqOpts...)
	req := requests.NewSearchSpreadsheetContentRequest(text, "", reqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return nil, err
	}
	return decodeTextItems(resp)
}

// Replace replaces text in a local spreadsheet and writes the resulting
// workbook to sink. Scope the replacement with WithWorksheet /
// WithCellArea. The live endpoint returns the modified file directly.
func Replace(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	src datasource.DataSource, sink datasource.DataSink, oldValue, newValue string, opts ...Option) error {

	if src == nil || sink == nil {
		return fmt.Errorf("%w: source and sink are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	data, err := src.ByteData()
	if err != nil {
		return err
	}

	req := requests.NewReplaceSpreadsheetContentRequest(newValue, oldValue, "", cfg.ReqOpts...)
	req.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return err
	}
	return sdkutil.WriteToSink(sink, resp.Body)
}

// SearchWorksheet searches text in a worksheet of a cloud workbook and returns
// the matching text items.
func SearchWorksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, worksheet, text string, opts ...Option) ([]TextItem, error) {

	if wf == nil || wf.Name == "" || worksheet == "" || text == "" {
		return nil, fmt.Errorf("%w: workbook name, worksheet and text are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewSearchContentInRemoteWorksheetRequest(wf.Name, text, worksheet, reqOpts...)

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return nil, err
	}
	return decodeTextItems(resp)
}

// SearchRange searches text in a cell range of a cloud worksheet and returns
// the matching text items. cellArea uses A1 style, e.g. "E35:F40".
func SearchRange(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, worksheet, cellArea, text string, opts ...Option) ([]TextItem, error) {

	if wf == nil || wf.Name == "" || worksheet == "" || cellArea == "" || text == "" {
		return nil, fmt.Errorf("%w: workbook name, worksheet, cell area and text are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewSearchContentInRemoteRangeRequest(cellArea, wf.Name, text, worksheet, reqOpts...)

	resp, err := asposecellscloud.DoChecked(ctx, client, req)
	if err != nil {
		return nil, err
	}
	return decodeTextItems(resp)
}

// ReplaceWorkbook replaces text in a cloud workbook in place.
func ReplaceWorkbook(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, oldValue, newValue string, opts ...Option) error {

	if wf == nil || wf.Name == "" || oldValue == "" || newValue == "" {
		return fmt.Errorf("%w: workbook name, old value and new value are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewReplaceContentInRemoteSpreadsheetRequest(wf.Name, newValue, oldValue, reqOpts...)

	_, err := asposecellscloud.DoChecked(ctx, client, req)
	return err
}

// ReplaceWorksheet replaces text in a worksheet of a cloud workbook in place.
func ReplaceWorksheet(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, worksheet, oldValue, newValue string, opts ...Option) error {

	if wf == nil || wf.Name == "" || worksheet == "" || oldValue == "" || newValue == "" {
		return fmt.Errorf("%w: workbook name, worksheet, old value and new value are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewReplaceContentInRemoteWorksheetRequest(wf.Name, newValue, oldValue, worksheet, reqOpts...)

	_, err := asposecellscloud.DoChecked(ctx, client, req)
	return err
}

// ReplaceRange replaces text in a cell range of a cloud worksheet in place.
// cellArea uses A1 style, e.g. "E35:F40".
func ReplaceRange(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient,
	wf *asposecellscloud.WorkbookRef, worksheet, cellArea, oldValue, newValue string, opts ...Option) error {

	if wf == nil || wf.Name == "" || worksheet == "" || cellArea == "" || oldValue == "" || newValue == "" {
		return fmt.Errorf("%w: workbook name, worksheet, cell area, old value and new value are required", asposecellscloud.ErrInvalidParam)
	}
	cfg := &sdkutil.Config{}
	sdkutil.Apply(cfg, opts)

	reqOpts := append(sdkutil.WorkbookParams(wf), cfg.ReqOpts...)
	req := requests.NewReplaceContentInRemoteRangeRequest(cellArea, wf.Name, newValue, oldValue, worksheet, reqOpts...)

	_, err := asposecellscloud.DoChecked(ctx, client, req)
	return err
}

// decodeTextItems extracts text items from a search/content response body.
func decodeTextItems(resp *asposecellscloud.RichResponse) ([]TextItem, error) {
	var out searchResponse
	if err := json.Unmarshal(resp.Body, &out); err != nil {
		return nil, err
	}
	items := make([]TextItem, 0, len(out.TextItems))
	for _, it := range out.TextItems {
		items = append(items, TextItem{Text: it.Content, Worksheet: it.Worksheet, Position: it.Position})
	}
	return items, nil
}
