/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AnalyzedTableDescription">
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

// AnalyzedTableDescription Represents analyzed table description.
type AnalyzedTableDescription struct {
	// Represents table name.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents worksheet name which is where the table is located.
	SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
	// Represents analyzed description about table columns.
	Columns []AnalyzedColumnDescription `json:"Columns,omitempty" xml:"Columns"`
	// Represents date columns list.
	DateColumns []int32 `json:"DateColumns,omitempty" xml:"DateColumns"`
	// Represents number columns list.
	NumberColumns []int32 `json:"NumberColumns,omitempty" xml:"NumberColumns"`
	// Represents string columns list.
	TextColumns []int32 `json:"TextColumns,omitempty" xml:"TextColumns"`
	// Represents exception columns list.
	ExceptionColumns []int32 `json:"ExceptionColumns,omitempty" xml:"ExceptionColumns"`
	// Represents there is a table header in the table.
	HasTableHeaderRow *bool `json:"HasTableHeaderRow,omitempty" xml:"HasTableHeaderRow"`
	// Represents there is a total row in the table.
	HasTableTotalRow *bool `json:"HasTableTotalRow,omitempty" xml:"HasTableTotalRow"`
	// Represents the column index as the start data column.
	StartDataColumnIndex *int32 `json:"StartDataColumnIndex,omitempty" xml:"StartDataColumnIndex"`
	// Represents the column index as the end data column.
	EndDataColumnIndex *int32 `json:"EndDataColumnIndex,omitempty" xml:"EndDataColumnIndex"`
	// Represents the row index as the start data row.
	StartDataRowIndex *int32 `json:"StartDataRowIndex,omitempty" xml:"StartDataRowIndex"`
	// Represents the row index as the end data row.
	EndDataRowIndex *int32 `json:"EndDataRowIndex,omitempty" xml:"EndDataRowIndex"`
	// Represents table thumbnail. Base64String
	Thumbnail string `json:"Thumbnail,omitempty" xml:"Thumbnail"`
	// Represents a collection of charts, which is a collection of charts created based on data analysis of a table.
	DiscoverCharts []DiscoverChart `json:"DiscoverCharts,omitempty" xml:"DiscoverCharts"`
	// Represents a collection of pivot tables, which is a collection of pivot tables created based on data analysis of a table.
	DiscoverPivotTables []DiscoverPivotTable `json:"DiscoverPivotTables,omitempty" xml:"DiscoverPivotTables"`
}
