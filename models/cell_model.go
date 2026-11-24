/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Cell">
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

// Cell Encapsulates the object that represents a single Workbook cell.
type Cell struct {
	LinkElement
	// Gets the name of the cell.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets row number (zero based) of the cell.
	Row *int32 `json:"Row,omitempty" xml:"Row"`
	// Gets column number (zero based) of the cell.
	Column *int32 `json:"Column,omitempty" xml:"Column"`
	// Gets the value contained in this cell.
	Value string `json:"Value,omitempty" xml:"Value"`
	// Represents cell value type.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Gets or sets a formula of the .
	Formula string `json:"Formula,omitempty" xml:"Formula"`
	// Represents if the specified cell contains formula.
	IsFormula *bool `json:"IsFormula,omitempty" xml:"IsFormula"`
	// Checks if a cell is part of a merged range or not.
	IsMerged *bool `json:"IsMerged,omitempty" xml:"IsMerged"`
	// Indicates the cell's formula is and array formula                          and it is the first cell of the array.
	IsArrayHeader *bool `json:"IsArrayHeader,omitempty" xml:"IsArrayHeader"`
	// Indicates whether the cell formula is an array formula.
	IsInArray *bool `json:"IsInArray,omitempty" xml:"IsInArray"`
	// Checks if the value of this cell is an error.
	IsErrorValue *bool `json:"IsErrorValue,omitempty" xml:"IsErrorValue"`
	// Indicates whether this cell is part of table formula.
	IsInTable *bool `json:"IsInTable,omitempty" xml:"IsInTable"`
	// Indicates if the cell's style is set. If return false, it means this cell has a default cell format.
	IsStyleSet *bool `json:"IsStyleSet,omitempty" xml:"IsStyleSet"`
	// Gets and sets the html string which contains data and some formats in this cell.
	HtmlString string `json:"HtmlString,omitempty" xml:"HtmlString"`
	// This class property represents a style element with the specified XML element name.
	Style *LinkElement `json:"Style,omitempty" xml:"Style"`
	// Gets the parent worksheet.
	Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
}
