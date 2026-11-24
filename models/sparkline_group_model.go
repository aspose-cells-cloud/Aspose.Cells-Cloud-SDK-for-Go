/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="SparklineGroup">
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

// SparklineGroup            is organized into sparkline group. A SparklineGroup contains a variable number of sparkline items.             A sparkline group specifies the type, display settings and axis settings for the sparklines.
type SparklineGroup struct {
	// Indicates whether to show data in hidden rows and columns.
	DisplayHidden *bool `json:"DisplayHidden,omitempty" xml:"DisplayHidden"`
	// Gets and sets the color of the first point of data in the sparkline group.
	FirstPointColor *CellsColor `json:"FirstPointColor,omitempty" xml:"FirstPointColor"`
	// Gets and sets the color of the highest points of data in the sparkline group.
	HighPointColor *CellsColor `json:"HighPointColor,omitempty" xml:"HighPointColor"`
	// Gets and sets the color of the horizontal axis in the sparkline group.
	HorizontalAxisColor *CellsColor `json:"HorizontalAxisColor,omitempty" xml:"HorizontalAxisColor"`
	// Represents the range that contains the date values for the sparkline data.
	HorizontalAxisDateRange string `json:"HorizontalAxisDateRange,omitempty" xml:"HorizontalAxisDateRange"`
	// Gets and sets the color of the last point of data in the sparkline group.
	LastPointColor *CellsColor `json:"LastPointColor,omitempty" xml:"LastPointColor"`
	// Gets and sets the line weight in each line sparkline in the sparkline group, in the unit of points.
	LineWeight *float64 `json:"LineWeight,omitempty" xml:"LineWeight"`
	// Gets and sets the color of the lowest points of data in the sparkline group.
	LowPointColor *CellsColor `json:"LowPointColor,omitempty" xml:"LowPointColor"`
	// Gets and sets the color of points in each line sparkline in the sparkline group.
	MarkersColor *CellsColor `json:"MarkersColor,omitempty" xml:"MarkersColor"`
	// Gets and sets the color of the negative values on the sparkline group.
	NegativePointsColor *CellsColor `json:"NegativePointsColor,omitempty" xml:"NegativePointsColor"`
	// Indicates how to plot empty cells.
	PlotEmptyCellsType string `json:"PlotEmptyCellsType,omitempty" xml:"PlotEmptyCellsType"`
	// Indicates whether the plot data is right to left.
	PlotRightToLeft *bool `json:"PlotRightToLeft,omitempty" xml:"PlotRightToLeft"`
	// Gets and sets the preset style type of the sparkline group.
	PresetStyle string `json:"PresetStyle,omitempty" xml:"PresetStyle"`
	// Gets and sets the color of the sparklines in the sparkline group.
	SeriesColor *CellsColor `json:"SeriesColor,omitempty" xml:"SeriesColor"`
	// Indicates whether to highlight the first point of data in the sparkline group.
	ShowFirstPoint *bool `json:"ShowFirstPoint,omitempty" xml:"ShowFirstPoint"`
	// Indicates whether to highlight the highest points of data in the sparkline group.
	ShowHighPoint *bool `json:"ShowHighPoint,omitempty" xml:"ShowHighPoint"`
	// Indicates whether to show the sparkline horizontal axis.             The horizontal axis appears if the sparkline has data that crosses the zero axis.
	ShowHorizontalAxis *bool `json:"ShowHorizontalAxis,omitempty" xml:"ShowHorizontalAxis"`
	// Indicates whether to highlight the last point of data in the sparkline group.
	ShowLastPoint *bool `json:"ShowLastPoint,omitempty" xml:"ShowLastPoint"`
	// Indicates whether to highlight the lowest points of data in the sparkline group.
	ShowLowPoint *bool `json:"ShowLowPoint,omitempty" xml:"ShowLowPoint"`
	// Indicates whether to highlight each point in each line sparkline in the sparkline group.
	ShowMarkers *bool `json:"ShowMarkers,omitempty" xml:"ShowMarkers"`
	// Indicates whether to highlight the negative values on the sparkline group with a different color or marker.
	ShowNegativePoints *bool `json:"ShowNegativePoints,omitempty" xml:"ShowNegativePoints"`
	// Gets the collection of  object.
	SparklineCollection []Sparkline `json:"SparklineCollection,omitempty" xml:"SparklineCollection"`
	// Indicates the sparkline type of the sparkline group.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Gets and sets the custom maximum value for the vertical axis.
	VerticalAxisMaxValue *float64 `json:"VerticalAxisMaxValue,omitempty" xml:"VerticalAxisMaxValue"`
	// Represents the vertical axis maximum value type.
	VerticalAxisMaxValueType string `json:"VerticalAxisMaxValueType,omitempty" xml:"VerticalAxisMaxValueType"`
	// Gets and sets the custom minimum value for the vertical axis.
	VerticalAxisMinValue *float64 `json:"VerticalAxisMinValue,omitempty" xml:"VerticalAxisMinValue"`
	// Represents the vertical axis minimum value type.
	VerticalAxisMinValueType string `json:"VerticalAxisMinValueType,omitempty" xml:"VerticalAxisMinValueType"`
}
