/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CreatePivotTableRequest">
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

// CreatePivotTableRequest Indicates create pivot table request
type CreatePivotTableRequest struct {
	// Pivot table name
	Name string `json:"Name,omitempty" xml:"Name"`
	// The data for the new PivotTable cache.
	SourceData string `json:"SourceData,omitempty" xml:"SourceData"`
	// The cell in the upper-left corner of the PivotTable report's destination range.
	DestCellName string `json:"DestCellName,omitempty" xml:"DestCellName"`
	// Indicates whether using same data source when another existing pivot table has used this data source.If the property is true, it will save memory.
	UseSameSource *bool `json:"UseSameSource,omitempty" xml:"UseSameSource"`
	// Represents row fields in a PivotTable report.
	PivotFieldRows []interface{} `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
	// Represents column fields in a PivotTable report.
	PivotFieldColumns []interface{} `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
	// Represents data fields in a PivotTable report.
	PivotFieldData []interface{} `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
}
