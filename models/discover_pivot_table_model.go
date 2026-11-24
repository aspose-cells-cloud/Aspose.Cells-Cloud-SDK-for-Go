/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DiscoverPivotTable">
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

// DiscoverPivotTable Represents a pivot table, which is a pivot table created based on data analysis of a table.
type DiscoverPivotTable struct {
	// Represents pivot table name.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents pivot table title.
	Title string `json:"Title,omitempty" xml:"Title"`
	// Represents pivot table data range.
	DataRange string `json:"DataRange,omitempty" xml:"DataRange"`
	// Represents row fields in a PivotTable report.
	PivotFieldRows []interface{} `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
	// Represents column fields in a PivotTable report.
	PivotFieldColumns []interface{} `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
	// Represents data fields in a PivotTable report.
	PivotFieldData []interface{} `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
	// Represents pivot table thumbnail. Base64String
	Thumbnail string `json:"Thumbnail,omitempty" xml:"Thumbnail"`
}
