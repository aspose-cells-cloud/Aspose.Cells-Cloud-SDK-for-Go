/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Axis">
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

// Axis Encapsulates the object that represents an axis of chart.
type Axis struct {
	LinkElement
	// Gets the .
	Area *Area `json:"Area,omitempty" xml:"Area"`
	// Represents if the value axis crosses the category axis between categories.
	AxisBetweenCategories *bool `json:"AxisBetweenCategories,omitempty" xml:"AxisBetweenCategories"`
	// Gets the appearance of an Axis.
	AxisLine *Line `json:"AxisLine,omitempty" xml:"AxisLine"`
	// Represents the base unit scale for the category axis.
	BaseUnitScale string `json:"BaseUnitScale,omitempty" xml:"BaseUnitScale"`
	// Represents the category axis type.
	CategoryType string `json:"CategoryType,omitempty" xml:"CategoryType"`
	// Represents the point on the value axis where the category axis crosses it.
	CrossAt *float64 `json:"CrossAt,omitempty" xml:"CrossAt"`
	// Represents the  on the specified axis where the other axis crosses.
	CrossType string `json:"CrossType,omitempty" xml:"CrossType"`
	// Represents the unit label for the specified axis.
	DisplayUnit string `json:"DisplayUnit,omitempty" xml:"DisplayUnit"`
	// Represents a unit label on an axis in the specified chart.                          Unit labels are useful for charting large values— for example, in the millions or billions.
	DisplayUnitLabel *DisplayUnitLabel `json:"DisplayUnitLabel,omitempty" xml:"DisplayUnitLabel"`
	// Indicates whether the labels shall be shown as multi level.
	HasMultiLevelLabels *bool `json:"HasMultiLevelLabels,omitempty" xml:"HasMultiLevelLabels"`
	// Indicates whether the major unit of the axis is automatically assigned.
	IsAutomaticMajorUnit *bool `json:"IsAutomaticMajorUnit,omitempty" xml:"IsAutomaticMajorUnit"`
	// Indicates whether the max value is automatically assigned.
	IsAutomaticMaxValue *bool `json:"IsAutomaticMaxValue,omitempty" xml:"IsAutomaticMaxValue"`
	// Indicates whether the minor unit of the axis is automatically assigned.
	IsAutomaticMinorUnit *bool `json:"IsAutomaticMinorUnit,omitempty" xml:"IsAutomaticMinorUnit"`
	// Indicates whether the min value is automatically assigned.
	IsAutomaticMinValue *bool `json:"IsAutomaticMinValue,omitempty" xml:"IsAutomaticMinValue"`
	// Represents if the display unit label is shown on the specified axis.
	IsDisplayUnitLabelShown *bool `json:"IsDisplayUnitLabelShown,omitempty" xml:"IsDisplayUnitLabelShown"`
	// Represents if the value axis scale type is logarithmic or not.
	IsLogarithmic *bool `json:"IsLogarithmic,omitempty" xml:"IsLogarithmic"`
	// Represents if Microsoft Excel plots data points from last to first.
	IsPlotOrderReversed *bool `json:"IsPlotOrderReversed,omitempty" xml:"IsPlotOrderReversed"`
	// Represents if the axis is visible.
	IsVisible *bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// Represents the logarithmic base. Default value is 10.Only applies for Excel2007.
	LogBase *float64 `json:"LogBase,omitempty" xml:"LogBase"`
	// Represents major gridlines on a chart axis.
	MajorGridLines *Line `json:"MajorGridLines,omitempty" xml:"MajorGridLines"`
	// Represents the type of major tick mark for the specified axis.
	MajorTickMark string `json:"MajorTickMark,omitempty" xml:"MajorTickMark"`
	// Represents the major units for the axis.
	MajorUnit *float64 `json:"MajorUnit,omitempty" xml:"MajorUnit"`
	// Represents the major unit scale for the category axis.
	MajorUnitScale string `json:"MajorUnitScale,omitempty" xml:"MajorUnitScale"`
	// Represents the maximum value on the value axis.
	MaxValue *float64 `json:"MaxValue,omitempty" xml:"MaxValue"`
	// Represents minor gridlines on a chart axis.
	MinorGridLines *Line `json:"MinorGridLines,omitempty" xml:"MinorGridLines"`
	// Represents the type of minor tick mark for the specified axis.
	MinorTickMark string `json:"MinorTickMark,omitempty" xml:"MinorTickMark"`
	// Represents the minor units for the axis.
	MinorUnit *float64 `json:"MinorUnit,omitempty" xml:"MinorUnit"`
	// Represents the major unit scale for the category axis.
	MinorUnitScale string `json:"MinorUnitScale,omitempty" xml:"MinorUnitScale"`
	// Represents the minimum value on the value axis.
	MinValue *float64 `json:"MinValue,omitempty" xml:"MinValue"`
	// Represents the position of tick-mark labels on the specified axis.
	TickLabelPosition string `json:"TickLabelPosition,omitempty" xml:"TickLabelPosition"`
	// Returns a  object that represents the tick-mark labels for the specified axis.
	TickLabels *TickLabels `json:"TickLabels,omitempty" xml:"TickLabels"`
	// Represents the number of categories or series between tick-mark labels. Applies only to category and series axes.
	TickLabelSpacing *int32 `json:"TickLabelSpacing,omitempty" xml:"TickLabelSpacing"`
	// Returns or sets the number of categories or series between tick marks. Applies only to category and series axes.
	TickMarkSpacing *int32 `json:"TickMarkSpacing,omitempty" xml:"TickMarkSpacing"`
	// Gets the axis' title.
	Title *Title `json:"Title,omitempty" xml:"Title"`
}
