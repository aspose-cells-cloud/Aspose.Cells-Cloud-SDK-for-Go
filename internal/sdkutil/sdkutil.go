/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="sdkutil.go">
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

// Package sdkutil provides shared helpers for the high-level feature packages
// (converter, reporting, editor, searcher, datacleansing, dataprocessing).
// It is internal to avoid polluting the public API surface.
package sdkutil

import (
	"context"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/requests"
)

// DoChecked executes a single request and verifies its HTTP status code,
// returning the first response on success. It delegates to the canonical
// implementation on the root package so the logic lives in exactly one place.
func DoChecked(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, req asposecellscloud.RequestOption) (*asposecellscloud.RichResponse, error) {
	return asposecellscloud.DoChecked(ctx, client, req)
}

// WriteToSink writes data to a DataSink.
func WriteToSink(sink datasource.DataSink, data []byte) error {
	w, err := sink.Write()
	if err != nil {
		return err
	}
	defer w.Close()
	_, err = w.Write(data)
	return err
}

// WorkbookParams expands a WorkbookRef into request-level options.
func WorkbookParams(wf *asposecellscloud.WorkbookRef) []requests.Option {
	var out []requests.Option
	if wf == nil {
		return out
	}
	if wf.Folder != "" {
		out = append(out, requests.WithCommonParameter("folder", wf.Folder))
	}
	if wf.StorageName != "" {
		out = append(out, requests.WithCommonParameter("storageName", wf.StorageName))
	}
	return out
}

// Config holds the common option state shared by all high-level packages.
type Config struct {
	ReqOpts []requests.Option
}

// ConfigOption configures a Config.
type ConfigOption func(*Config)

// Apply applies a slice of ConfigOptions to cfg.
func Apply(cfg *Config, opts []ConfigOption) {
	for _, o := range opts {
		if o != nil {
			o(cfg)
		}
	}
}

// WithCommonParameter passes an optional parameter to the underlying generated request.
func WithCommonParameter(name string, value interface{}) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, requests.WithCommonParameter(name, value))
	}
}

// WithQueryParameter appends a custom query parameter to the request URL.
func WithQueryParameter(name, value string) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, requests.WithQueryParameter(name, value))
	}
}

// WithQueryParameters appends multiple custom query parameters.
func WithQueryParameters(params map[string]string) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, requests.WithQueryParameters(params))
	}
}

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.Option) ConfigOption {
	return func(c *Config) {
		c.ReqOpts = append(c.ReqOpts, opts...)
	}
}

// SpreadsheetSetter is implemented by request types that accept a spreadsheet
// file via SetSpreadsheetBytes. Using this interface avoids brittle type
// switches over concrete request types.
type SpreadsheetSetter interface {
	SetSpreadsheetBytes(data []byte, name string)
}

// DatafileSetter is implemented by request types that accept a data file
// via SetDatafileBytes (e.g. import endpoints).
type DatafileSetter interface {
	SetDatafileBytes(data []byte, name string)
}
