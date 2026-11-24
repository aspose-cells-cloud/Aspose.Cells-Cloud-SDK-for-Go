/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CopyOptions">
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

// CopyOptions Represents the copy options.
type CopyOptions struct {
	// Indicates whether copying column width in unit of characters.
	ColumnCharacterWidth *bool `json:"ColumnCharacterWidth,omitempty" xml:"ColumnCharacterWidth"`
	// If the formula is not valid for the dest destination, only copy values.
	CopyInvalidFormulasAsValues *bool `json:"CopyInvalidFormulasAsValues,omitempty" xml:"CopyInvalidFormulasAsValues"`
	// Indicates whether copying the names.
	CopyNames *bool `json:"CopyNames,omitempty" xml:"CopyNames"`
	// Indicates whether extend ranges when copying the range to adjacent range.
	ExtendToAdjacentRange *bool `json:"ExtendToAdjacentRange,omitempty" xml:"ExtendToAdjacentRange"`
	// When copying the range in the same file and the chart refers to the source sheet,             False means the copied chart's data source will not be changed.             True means the copied chart's data source refers to the destination sheet.
	ReferToDestinationSheet *bool `json:"ReferToDestinationSheet,omitempty" xml:"ReferToDestinationSheet"`
	// In ms excel, when copying formulas which refer to other worksheets while copying a worksheet to another one,             the copied formulas should refer to source workbook.             However, for some situations user may need the copied formulas refer to worksheets with the same name             in the same workbook, such as when those worksheets have been copied before this copy operation,             then this property should be kept as true.
	ReferToSheetWithSameName *bool `json:"ReferToSheetWithSameName,omitempty" xml:"ReferToSheetWithSameName"`
	CopyTheme                *bool `json:"CopyTheme,omitempty" xml:"CopyTheme"`
}
