/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PivotTableOperateParameter">
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

import ()

// PivotTableOperateParameter Represents pivot table operate parameter.
type PivotTableOperateParameter struct {
	OperateParameter
	// Represents source data of pivot table.
	SourceData string `json:"SourceData,omitempty" xml:"SourceData"`
	// Represents start cell name of the pivot table.
	DestCellName string `json:"DestCellName,omitempty" xml:"DestCellName"`
	// Represents table name of pivot table.
	TableName string `json:"TableName,omitempty" xml:"TableName"`
	// Represents whether the same source is used.
	UseSameSource *bool `json:"UseSameSource,omitempty" xml:"UseSameSource"`
	// Represents pivot table index.
	PivotTableIndex *int32 `json:"PivotTableIndex,omitempty" xml:"PivotTableIndex"`
	// Represents pivot row fields.
	PivotFieldRows []interface{} `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
	// Represents pivot column fields.
	PivotFieldColumns []interface{} `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
	// Represents pivot data field.
	PivotFieldData []interface{} `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
}
