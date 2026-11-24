/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CellValue">
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

// CellValue Represents the cell value and corresponding type.
type CellValue struct {
	// A public integer property called "rowIndex" with both getter and setter.
	RowIndex    *int32 `json:"rowIndex,omitempty" xml:"rowIndex"`
	ColumnIndex *int32 `json:"columnIndex,omitempty" xml:"columnIndex"`
	Type        string `json:"type,omitempty" xml:"type"`
	Value       string `json:"value,omitempty" xml:"value"`
	Formula     string `json:"formula,omitempty" xml:"formula"`
	Style       *Style `json:"style,omitempty" xml:"style"`
}
