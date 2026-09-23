/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="pivot_column.go">
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

package models

import "encoding/json"

// PivotColumn Represents pivot column for data table.
type PivotColumn struct {
	// Represents pivot column name.
	PivotColumnName string `json:"PivotColumnName,omitempty" xml:"PivotColumnName"`
	// Represents column name that sets the column's value to the value of the pivot column.
	ValueColumnNames   []string `json:"ValueColumnNames,omitempty" xml:"ValueColumnNames"`
	AppliedOperateType string   `json:"AppliedOperateType,omitempty" xml:"AppliedOperateType"`
}

// MarshalJSON fills AppliedOperateType with this type's own name -- the
// member the service's JSON converter dispatches on -- so callers passing a
// concrete operation never have to set it themselves. An explicit value is
// kept as-is.
func (m PivotColumn) MarshalJSON() ([]byte, error) {
	type plain PivotColumn
	p := plain(m)
	if p.AppliedOperateType == "" {
		p.AppliedOperateType = "PivotColumn"
	}
	return json.Marshal(p)
}
