/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Cells">
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

// Cells Encapsulates a collection of cell relevant objects, such as Aspose.Cells.Cell, Aspose.Cells.Row, ...etc.
type Cells struct {
	LinkElement
	// Maximum row index of cell which contains data or style.
	MaxRow *int32 `json:"MaxRow,omitempty" xml:"MaxRow"`
	// Maximum column index of those cells that have been instantiated in the collection(does not include the column                         where style is defined for the whole column but no cell has been instantiated in it).
	MaxColumn *int32 `json:"MaxColumn,omitempty" xml:"MaxColumn"`
	// The class has a public property "CellCount" of type integer that can be read and modified.
	CellCount *int32 `json:"CellCount,omitempty" xml:"CellCount"`
	// Gets the collection of  objects that represents the individual rows in this worksheet.
	Rows *LinkElement `json:"Rows,omitempty" xml:"Rows"`
	// Gets the collection of  objects that represents the individual columns in this worksheet.
	Columns  *LinkElement  `json:"Columns,omitempty" xml:"Columns"`
	CellList []LinkElement `json:"CellList,omitempty" xml:"CellList"`
}
