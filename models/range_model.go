/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Range">
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

// Range
type Range struct {
	ColumnCount *int32   `json:"ColumnCount,omitempty" xml:"ColumnCount"`
	ColumnWidth *float64 `json:"ColumnWidth,omitempty" xml:"ColumnWidth"`
	FirstColumn *int32   `json:"FirstColumn,omitempty" xml:"FirstColumn"`
	FirstRow    *int32   `json:"FirstRow,omitempty" xml:"FirstRow"`
	Name        string   `json:"Name,omitempty" xml:"Name"`
	RefersTo    string   `json:"RefersTo,omitempty" xml:"RefersTo"`
	RowCount    *int32   `json:"RowCount,omitempty" xml:"RowCount"`
	RowHeight   *float64 `json:"RowHeight,omitempty" xml:"RowHeight"`
	Worksheet   string   `json:"Worksheet,omitempty" xml:"Worksheet"`
}
