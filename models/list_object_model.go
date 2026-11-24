/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ListObject">
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

// ListObject Represents a list object on a worksheet.            The ListObject object is a member of the ListObjects collection.             The ListObjects collection contains all the list objects on a worksheet.
type ListObject struct {
	LinkElement
	// Gets auto filter.
	AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
	// Gets and sets the display name.
	DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
	// Gets the start column of the range.
	StartColumn *int32 `json:"StartColumn,omitempty" xml:"StartColumn"`
	// Gets the start row of the range.
	StartRow *int32 `json:"StartRow,omitempty" xml:"StartRow"`
	// Gets the end column of the range.
	EndColumn *int32 `json:"EndColumn,omitempty" xml:"EndColumn"`
	// Gets the end  row of the range.
	EndRow *int32 `json:"EndRow,omitempty" xml:"EndRow"`
	// Gets ListColumns of the ListObject.
	ListColumns []ListColumn `json:"ListColumns,omitempty" xml:"ListColumns"`
	// Gets and sets whether this ListObject show header row.
	ShowHeaderRow *bool `json:"ShowHeaderRow,omitempty" xml:"ShowHeaderRow"`
	// Indicates whether column stripe formatting is applied.
	ShowTableStyleColumnStripes *bool `json:"ShowTableStyleColumnStripes,omitempty" xml:"ShowTableStyleColumnStripes"`
	// Indicates whether the first column in the table should have the style applied.
	ShowTableStyleFirstColumn *bool `json:"ShowTableStyleFirstColumn,omitempty" xml:"ShowTableStyleFirstColumn"`
	// Indicates whether the last column in the table should have the style applied.
	ShowTableStyleLastColumn *bool `json:"ShowTableStyleLastColumn,omitempty" xml:"ShowTableStyleLastColumn"`
	// Indicates whether row stripe formatting is applied.
	ShowTableStyleRowStripes *bool `json:"ShowTableStyleRowStripes,omitempty" xml:"ShowTableStyleRowStripes"`
	// Gets and sets whether this ListObject show total row.
	ShowTotals *bool `json:"ShowTotals,omitempty" xml:"ShowTotals"`
	// Gets and sets the table style name.
	TableStyleName string `json:"TableStyleName,omitempty" xml:"TableStyleName"`
	// Gets and the built-in table style.
	TableStyleType string `json:"TableStyleType,omitempty" xml:"TableStyleType"`
	// Gets the data range of the ListObject.
	DataRange *Range `json:"DataRange,omitempty" xml:"DataRange"`
	// Gets the data source type of the table.
	DataSourceType string `json:"DataSourceType,omitempty" xml:"DataSourceType"`
	// Gets and sets the comment of the table.
	Comment string `json:"Comment,omitempty" xml:"Comment"`
	// Gets an  used for this list.
	XmlMap *XmlMap `json:"XmlMap,omitempty" xml:"XmlMap"`
	// Gets and sets the alternative text.
	AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
	// Gets and sets the alternative description.
	AlternativeDescription string `json:"AlternativeDescription,omitempty" xml:"AlternativeDescription"`
}
