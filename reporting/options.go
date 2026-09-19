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
