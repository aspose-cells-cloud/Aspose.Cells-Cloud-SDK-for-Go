/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Worksheet">
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

// Worksheet            Encapsulates the object that represents a single worksheet.
type Worksheet struct {
	// Property Summary: Contains a list of links represented by the class link.
	Links []Link `json:"Links,omitempty" xml:"Links"`
	// Indicates if the specified worksheet is displayed from right to left instead of from left to right.             Default is false.
	DisplayRightToLeft *bool `json:"DisplayRightToLeft,omitempty" xml:"DisplayRightToLeft"`
	// True if zero values are displayed.
	DisplayZeros *bool `json:"DisplayZeros,omitempty" xml:"DisplayZeros"`
	// Represents first visible column index.
	FirstVisibleColumn *int32 `json:"FirstVisibleColumn,omitempty" xml:"FirstVisibleColumn"`
	// Represents first visible row index.
	FirstVisibleRow *int32 `json:"FirstVisibleRow,omitempty" xml:"FirstVisibleRow"`
	// Gets or sets the name of the worksheet.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets the index of sheet in the worksheet collection.
	Index *int32 `json:"Index,omitempty" xml:"Index"`
	// Gets or sets a value indicating whether the gridlines are visible.Default is true.
	IsGridlinesVisible *bool `json:"IsGridlinesVisible,omitempty" xml:"IsGridlinesVisible"`
	// Indicates whether to show outline.
	IsOutlineShown *bool `json:"IsOutlineShown,omitempty" xml:"IsOutlineShown"`
	// Indicates whether the specified worksheet is shown in normal view or page break preview.
	IsPageBreakPreview *bool `json:"IsPageBreakPreview,omitempty" xml:"IsPageBreakPreview"`
	// Represents if the worksheet is visible.
	IsVisible *bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// Indicates if the worksheet is protected.
	IsProtected *bool `json:"IsProtected,omitempty" xml:"IsProtected"`
	// Gets or sets a value indicating whether the worksheet will display row and column headers.             Default is true.
	IsRowColumnHeadersVisible *bool `json:"IsRowColumnHeadersVisible,omitempty" xml:"IsRowColumnHeadersVisible"`
	// Indicates whether the ruler is visible. This property is only applied for page break preview.
	IsRulerVisible *bool `json:"IsRulerVisible,omitempty" xml:"IsRulerVisible"`
	// Indicates whether this worksheet is selected when the workbook is opened.
	IsSelected *bool `json:"IsSelected,omitempty" xml:"IsSelected"`
	// Represents worksheet tab color.
	TabColor *Color `json:"TabColor,omitempty" xml:"TabColor"`
	// Indicates whether the Transition Formula Entry (Lotus compatibility) option is enabled.
	TransitionEntry *bool `json:"TransitionEntry,omitempty" xml:"TransitionEntry"`
	// Indicates whether the Transition Formula Evaluation (Lotus compatibility) option is enabled.
	TransitionEvaluation *bool `json:"TransitionEvaluation,omitempty" xml:"TransitionEvaluation"`
	// Represents worksheet type.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Gets and sets the view type.
	ViewType string `json:"ViewType,omitempty" xml:"ViewType"`
	// Indicates the visible state for this sheet.
	VisibilityType string `json:"VisibilityType,omitempty" xml:"VisibilityType"`
	// Represents the scaling factor in percentage. It should be between 10 and 400.
	Zoom *int32 `json:"Zoom,omitempty" xml:"Zoom"`
	// Gets the  collection.
	Cells *LinkElement `json:"Cells,omitempty" xml:"Cells"`
	// Gets a  collection
	Charts     *LinkElement `json:"Charts,omitempty" xml:"Charts"`
	AutoShapes *LinkElement `json:"AutoShapes,omitempty" xml:"AutoShapes"`
	// Represents a collection of  in a worksheet.
	OleObjects *LinkElement `json:"OleObjects,omitempty" xml:"OleObjects"`
	// Gets the  collection.
	Comments *LinkElement `json:"Comments,omitempty" xml:"Comments"`
	// Gets a  collection.
	Pictures    *LinkElement `json:"Pictures,omitempty" xml:"Pictures"`
	MergedCells *LinkElement `json:"MergedCells,omitempty" xml:"MergedCells"`
	// Gets the data validation setting collection in the worksheet.
	Validations *LinkElement `json:"Validations,omitempty" xml:"Validations"`
	// Gets the ConditionalFormattings in the worksheet.
	ConditionalFormattings *LinkElement `json:"ConditionalFormattings,omitempty" xml:"ConditionalFormattings"`
	// Gets the  collection.
	Hyperlinks *LinkElement `json:"Hyperlinks,omitempty" xml:"Hyperlinks"`
}
