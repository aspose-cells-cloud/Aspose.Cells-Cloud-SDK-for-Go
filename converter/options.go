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

package converter

import (
	"encoding/json"

	"asposecellscloud"
	"asposecellscloud/internal/sdkutil"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

// Option configures a high-level convert/export operation.
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

// WithQueryParameters appends multiple implicit query parameters.
func WithQueryParameters(m map[string]string) Option {
	return sdkutil.WithQueryParameters(m)
}

// WithRaw passes any requests-level Option straight through to the generated
// request constructor. Use it to reach parameters not covered above.
func WithRaw(opts ...requests.Option) Option {
	return sdkutil.WithRaw(opts...)
}

// --- Conversion / export parameter group ---

// WithPassword sets the workbook open password.
func WithPassword(p string) Option {
	return WithCommonParameter("password", p)
}

// WithRegion sets the region.
func WithRegion(r string) Option {
	return WithCommonParameter("region", r)
}

// WithFontsLocation sets the location of fonts used for conversion.
func WithFontsLocation(l string) Option {
	return WithCommonParameter("FontsLocation", l)
}

// WithOutPath writes the converted/exported file to this cloud path instead of
// returning it in the response body.
func WithOutPath(p string) Option {
	return WithCommonParameter("outPath", p)
}

// WithStorageName sets the cloud storage name.
func WithStorageName(s string) Option {
	return WithCommonParameter("storageName", s)
}

// WithFolder sets the cloud folder that contains the workbook (cloud Workbook /
// Worksheet export).
func WithFolder(f string) Option {
	return WithCommonParameter("folder", f)
}

// WithAutoRowsFit fits rows automatically before conversion.
func WithAutoRowsFit(v bool) Option {
	return WithCommonParameter("AutoRowsFit", asposecellscloud.BoolPtr(v))
}

// WithAutoColumnsFit fits columns automatically before conversion.
func WithAutoColumnsFit(v bool) Option {
	return WithCommonParameter("AutoColumnsFit", asposecellscloud.BoolPtr(v))
}

// WithCheckExcelRestriction checks the Excel restriction before conversion.
func WithCheckExcelRestriction(v bool) Option {
	return WithCommonParameter("checkExcelRestriction", asposecellscloud.BoolPtr(v))
}

// WithOnePagePerSheet fits each worksheet onto one page for PDF output.
func WithOnePagePerSheet(v bool) Option {
	return WithCommonParameter("onePagePerSheet", asposecellscloud.BoolPtr(v))
}

// WithSaveOptions attaches a SaveOptions model (serialized to the implicit
// "SaveOptions" query parameter) for fine-grained save control.
func WithSaveOptions(so *models.SaveOptions) Option {
	return func(c *sdkutil.Config) {
		b, err := json.Marshal(so)
		if err != nil {
			return
		}
		c.ReqOpts = append(c.ReqOpts, requests.WithQueryParameter("SaveOptions", string(b)))
	}
}
