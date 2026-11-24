/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="options.go">
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

package datacleansing

import (
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// Option configures a high-level data cleansing operation.
type Option = sdkutil.ConfigOption

// WithCommonParameter passes an optional parameter to the underlying generated
// request. The name must match the operation parameter name in the spec.
func WithCommonParameter(name string, value interface{}) Option {
	return sdkutil.WithCommonParameter(name, value)
}

// WithQueryParameter appends an implicit query parameter to the request URL.
func WithQueryParameter(name, value string) Option {
	return sdkutil.WithQueryParameter(name, value)
}

// WithQueryParameters appends multiple implicit query parameters to the request
// URL.
func WithQueryParameters(params map[string]string) Option {
	return sdkutil.WithQueryParameters(params)
}

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.Option) Option {
	return sdkutil.WithRaw(opts...)
}

// --- Cleansing scoping ---

// WithWorksheet scopes duplicate removal to one worksheet (RemoveDuplicates).
func WithWorksheet(name string) Option {
	return WithCommonParameter("worksheet", name)
}

// WithRange scopes duplicate removal to one cell range, A1 style e.g. "E35:F40"
// (RemoveDuplicates).
func WithRange(cellArea string) Option {
	return WithCommonParameter("range", cellArea)
}

// WithTable scopes duplicate removal to a named table (RemoveDuplicates).
func WithTable(name string) Option {
	return WithCommonParameter("table", name)
}

// --- Output / credentials ---

// WithOutPath writes the cleaned workbook to this cloud path instead of
// returning it in the response body.
func WithOutPath(p string) Option {
	return WithCommonParameter("outPath", p)
}

// WithOutStorageName sets the cloud storage name for the output workbook.
func WithOutStorageName(s string) Option {
	return WithCommonParameter("outStorageName", s)
}

// WithPassword sets the workbook open password.
func WithPassword(p string) Option {
	return WithCommonParameter("password", p)
}

// WithRegion sets the region.
func WithRegion(r string) Option {
	return WithCommonParameter("region", r)
}
