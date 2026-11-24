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

package reporting

import (
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/requests"
)

// Option configures a high-level report/analysis operation.
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

// WithRaw passes any requests-level Option straight through to the generated
// request constructor.
func WithRaw(opts ...requests.Option) Option {
	return sdkutil.WithRaw(opts...)
}

// WithWorksheet restricts an analysis/calculation to a worksheet.
func WithWorksheet(w string) Option {
	return WithCommonParameter("worksheet", w)
}

// WithRange restricts an analysis/calculation to a cell range (A1 style).
func WithRange(r string) Option {
	return WithCommonParameter("range", r)
}

// WithOperation sets the aggregation/calculation operation (e.g. count/sum).
func WithOperation(op string) Option {
	return WithCommonParameter("operation", op)
}

// WithColorPosition sets the color position (e.g. fill/font) for
// AggregateByColor.
func WithColorPosition(p string) Option {
	return WithCommonParameter("colorPosition", p)
}

// WithSheetName scopes an AI analysis/calculation to a worksheet by name.
func WithSheetName(s string) Option {
	return WithCommonParameter("sheetName", s)
}

// WithNewColumnName names the category column produced by
// CategorizeSpreadsheet.
func WithNewColumnName(n string) Option {
	return WithCommonParameter("newColumnName", n)
}

// WithPassword sets the workbook open password.
func WithPassword(p string) Option {
	return WithCommonParameter("password", p)
}

// WithRegion sets the region.
func WithRegion(r string) Option {
	return WithCommonParameter("region", r)
}
