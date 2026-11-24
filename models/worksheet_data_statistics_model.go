/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="WorksheetDataStatistics">
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

// WorksheetDataStatistics Represents worksheet data statistics.
type WorksheetDataStatistics struct {
	// Represents worksheet name.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents chart number.
	ChartsCount *int32 `json:"ChartsCount,omitempty" xml:"ChartsCount"`
	// Represents list object number.
	TablesCount *int32 `json:"TablesCount,omitempty" xml:"TablesCount"`
	// Represents pivot table number.
	PivotTablesCount *int32 `json:"PivotTablesCount,omitempty" xml:"PivotTablesCount"`
	// Represents shape number.
	ShapesCount *int32 `json:"ShapesCount,omitempty" xml:"ShapesCount"`
	// Represents shape number.
	HyperlinksCount *int32 `json:"HyperlinksCount,omitempty" xml:"HyperlinksCount"`
	// Represents hyperlink number.
	QueryTablesCount *int32 `json:"QueryTablesCount,omitempty" xml:"QueryTablesCount"`
	// Represents query table number.
	CellsCount *int32 `json:"CellsCount,omitempty" xml:"CellsCount"`
	// Represents cell number.
	CellsCountInTable *int32 `json:"CellsCountInTable,omitempty" xml:"CellsCountInTable"`
	// Represents formula number.
	CellsCountIsFormula *int32 `json:"CellsCountIsFormula,omitempty" xml:"CellsCountIsFormula"`
}
