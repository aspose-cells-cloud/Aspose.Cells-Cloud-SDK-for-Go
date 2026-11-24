/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataLabels">
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

// DataLabels Encapsulates a collection of all the DataLabel objects for the specified Series.
type DataLabels struct {
	ChartFrame
	// Indicates the text is auto generated.
	IsAutoText *bool `json:"IsAutoText,omitempty" xml:"IsAutoText"`
	// A nullable boolean property "IsDeleted" indicating whether an object has been deleted.
	IsDeleted    *bool  `json:"IsDeleted,omitempty" xml:"IsDeleted"`
	LinkedSource string `json:"LinkedSource,omitempty" xml:"LinkedSource"`
	// Gets and sets the built-in number format.
	Number *int32 `json:"Number,omitempty" xml:"Number"`
	// Represents the format string for the DataLabels object.
	NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
	// True if the number format is linked to the cells                          (so that the number format changes in the labels when it changes in the cells).
	NumberFormatLinked *bool `json:"NumberFormatLinked,omitempty" xml:"NumberFormatLinked"`
	// Represents the position of the data label.
	Position      string `json:"Position,omitempty" xml:"Position"`
	RotationAngle *int32 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	// Gets or sets the separator type used for the data labels on a chart.
	Separator string `json:"Separator,omitempty" xml:"Separator"`
	// Represents a specified chart's data label percentage value display behavior. True displays the percentage value. False to hide.
	ShowBubbleSize *bool `json:"ShowBubbleSize,omitempty" xml:"ShowBubbleSize"`
	// Represents a specified chart's data label category name display behavior.True to display the category name for the data labels on a chart. False to hide.
	ShowCategoryName *bool `json:"ShowCategoryName,omitempty" xml:"ShowCategoryName"`
	// Represents a specified chart's data label legend key display behavior.                         True if the data label legend key is visible.
	ShowLegendKey *bool `json:"ShowLegendKey,omitempty" xml:"ShowLegendKey"`
	// Represents a specified chart's data label percentage value display behavior. True displays the percentage value. False to hide.
	ShowPercentage *bool `json:"ShowPercentage,omitempty" xml:"ShowPercentage"`
	// Returns or sets a Boolean to indicate the series name display behavior for the data labels on a chart.                         True to show the series name. False to hide.
	ShowSeriesName *bool `json:"ShowSeriesName,omitempty" xml:"ShowSeriesName"`
	// Represents a specified chart's data label values display behavior. True displays the values. False to hide.
	ShowValue *bool `json:"ShowValue,omitempty" xml:"ShowValue"`
	// Gets or sets the text of data label.
	Text                    string `json:"Text,omitempty" xml:"Text"`
	TextDirection           string `json:"TextDirection,omitempty" xml:"TextDirection"`
	TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
	TextVerticalAlignment   string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}
