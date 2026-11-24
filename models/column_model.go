/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Column">
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

// Column Represents a single column in a worksheet.
type Column struct {
	LinkElement
	// Gets the group level of the column.
	GroupLevel *int32 `json:"GroupLevel,omitempty" xml:"GroupLevel"`
	// Gets the index of this column.
	Index *int32 `json:"Index,omitempty" xml:"Index"`
	// Indicates whether the column is hidden.
	IsHidden *bool `json:"IsHidden,omitempty" xml:"IsHidden"`
	// Gets and sets the column width in unit of characters.
	Width *float64 `json:"Width,omitempty" xml:"Width"`
	// Gets the style of this column.
	Style *LinkElement `json:"Style,omitempty" xml:"Style"`
}
