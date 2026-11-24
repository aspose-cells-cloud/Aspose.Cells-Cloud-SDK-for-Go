/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Series">
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

// Series Encapsulates the object that represents a single data series in a chart.
type Series struct {
	LinkElement
	// Represents the background area of Series object.
	Area *Area `json:"Area,omitempty" xml:"Area"`
	// Gets or sets the 3D shape type used with the 3-D bar or column chart.
	Bar3DShapeType string `json:"Bar3DShapeType,omitempty" xml:"Bar3DShapeType"`
	// Represents border of Series object.
	Border *Line `json:"Border,omitempty" xml:"Border"`
	// Gets or sets the scale factor for bubbles in the specified chart group.                          It can be an integer value from 0 (zero) to 300,                          corresponding to a percentage of the default size.                         Applies only to bubble charts.
	BubbleScale *int32 `json:"BubbleScale,omitempty" xml:"BubbleScale"`
	// Gets or sets the bubble sizes values of the chart series.
	BubbleSizes string `json:"BubbleSizes,omitempty" xml:"BubbleSizes"`
	// Gets the number of the data values.
	CountOfDataValues *int32 `json:"CountOfDataValues,omitempty" xml:"CountOfDataValues"`
	// Represents the DataLabels object for the specified ASeries.
	DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
	// Gets the series's name that displays on the chart graph.
	DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
	// Returns or sets the size of the hole in a doughnut chart group.                          The hole size is expressed as a percentage of the chart size, between 10 and 90 percent.
	DoughnutHoleSize *int32 `json:"DoughnutHoleSize,omitempty" xml:"DoughnutHoleSize"`
	// Returns a  object that represents the down bars on a line chart.                         Applies only to line charts.
	DownBars *DropBars `json:"DownBars,omitempty" xml:"DownBars"`
	// Returns a  object that represents the drop lines for a series on the line chart or area chart.                         Applies only to line chart or area charts.
	DropLines *Line `json:"DropLines,omitempty" xml:"DropLines"`
	// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
	Explosion *int32 `json:"Explosion,omitempty" xml:"Explosion"`
	// Gets or sets the angle of the first pie-chart or doughnut-chart slice, in degrees (clockwise from vertical).                          Applies only to pie, 3-D pie, and doughnut charts, 0 to 360.
	FirstSliceAngle *int32 `json:"FirstSliceAngle,omitempty" xml:"FirstSliceAngle"`
	// Returns or sets the space between bar or column clusters, as a percentage of the bar or column width.                         The value of this property must be between 0 and 500.
	GapWidth *int32 `json:"GapWidth,omitempty" xml:"GapWidth"`
	// True if the series has a three-dimensional appearance.                          Applies only to bubble charts.
	Has3DEffect *bool `json:"Has3DEffect,omitempty" xml:"Has3DEffect"`
	// True if the chart has drop lines.                         Applies only to line chart or area charts.
	HasDropLines *bool `json:"HasDropLines,omitempty" xml:"HasDropLines"`
	// True if the line chart has high-low lines.                           Applies only to line charts.
	HasHiLoLines *bool `json:"HasHiLoLines,omitempty" xml:"HasHiLoLines"`
	// True if the series has leader lines.
	HasLeaderLines *bool `json:"HasLeaderLines,omitempty" xml:"HasLeaderLines"`
	// True if a radar chart has category axis labels. Applies only to radar charts.
	HasRadarAxisLabels *bool `json:"HasRadarAxisLabels,omitempty" xml:"HasRadarAxisLabels"`
	// True if a stacked column chart or bar chart has series lines or                         if a Pie of Pie chart or Bar of Pie chart has connector lines between the two sections.                          Applies only to stacked column charts, bar charts, Pie of Pie charts, or Bar of Pie charts.
	HasSeriesLines *bool `json:"HasSeriesLines,omitempty" xml:"HasSeriesLines"`
	// True if a line chart has up and down bars.                         Applies only to line charts.
	HasUpDownBars *bool `json:"HasUpDownBars,omitempty" xml:"HasUpDownBars"`
	// Returns a HiLoLines object that represents the high-low lines for a series on a line chart.                          Applies only to line charts.
	HiLoLines *Line `json:"HiLoLines,omitempty" xml:"HiLoLines"`
	// Indicates whether the threshold value is automatic.
	IsAutoSplit *bool `json:"IsAutoSplit,omitempty" xml:"IsAutoSplit"`
	// Represents if the color of points is varied.                          The chart must contain only one series.
	IsColorVaried *bool `json:"IsColorVaried,omitempty" xml:"IsColorVaried"`
	// Represents leader lines on a chart. Leader lines connect data labels to data points.                          This object isn’t a collection; there’s no object that represents a single leader line.
	LeaderLines *Line `json:"LeaderLines,omitempty" xml:"LeaderLines"`
	// Gets the legend entry according to this series.
	LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
	// Gets the marker.
	Marker *Marker `json:"Marker,omitempty" xml:"Marker"`
	// Gets or sets the name of the data series.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Specifies how bars and columns are positioned.                         Can be a value between – 100 and 100.                          Applies only to 2-D bar and 2-D column charts.
	Overlap *int32 `json:"Overlap,omitempty" xml:"Overlap"`
	// Indicates if this series is plotted on second value axis.
	PlotOnSecondAxis *bool `json:"PlotOnSecondAxis,omitempty" xml:"PlotOnSecondAxis"`
	// Gets the collection of points in a series in a chart.
	Points *LinkElement `json:"Points,omitempty" xml:"Points"`
	// Returns or sets the size of the secondary section of either a pie of pie chart or a bar of pie chart,                          as a percentage of the size of the primary pie.                         Can be a value from 5 to 200.
	SecondPlotSize *int32 `json:"SecondPlotSize,omitempty" xml:"SecondPlotSize"`
	// Returns a SeriesLines object that represents the series lines for a stacked bar chart or a stacked column chart.                         Applies only to stacked bar and stacked column charts.
	SeriesLines *Line `json:"SeriesLines,omitempty" xml:"SeriesLines"`
	// True if the series has a shadow.
	Shadow *bool `json:"Shadow,omitempty" xml:"Shadow"`
	// True if negative bubbles are shown for the chart group. Valid only for bubble charts.
	ShowNegativeBubbles *bool `json:"ShowNegativeBubbles,omitempty" xml:"ShowNegativeBubbles"`
	// Gets or sets what the bubble size represents on a bubble chart.
	SizeRepresents string `json:"SizeRepresents,omitempty" xml:"SizeRepresents"`
	// Represents curve smoothing.                          True if curve smoothing is turned on for the line chart or scatter chart.                         Applies only to line and scatter connected by lines charts.
	Smooth *bool `json:"Smooth,omitempty" xml:"Smooth"`
	// Returns or sets a value that how to determine which data points are in the second pie or bar on a pie of pie or bar of                         pie chart.
	SplitType string `json:"SplitType,omitempty" xml:"SplitType"`
	// Returns or sets a value that shall be used to determine which data points are in the second pie or bar on                         a pie of pie or bar of pie chart.
	SplitValue *float64 `json:"SplitValue,omitempty" xml:"SplitValue"`
	// Returns an object that represents a collection of all the trendlines for the series.
	TrendLines *Trendlines `json:"TrendLines,omitempty" xml:"TrendLines"`
	// Gets or sets a data series' type.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Returns an DropBars object that represents the up bars on a line chart.                         Applies only to line charts.
	UpBars *DropBars `json:"UpBars,omitempty" xml:"UpBars"`
	// Represents the data of the chart series.
	Values string `json:"Values,omitempty" xml:"Values"`
	// Represents X direction error bar of the series.
	XErrorBar *ErrorBar `json:"XErrorBar,omitempty" xml:"XErrorBar"`
	// Represents the x values of the chart series.
	XValues string `json:"XValues,omitempty" xml:"XValues"`
	// Represents Y direction error bar of the series.
	YErrorBar *ErrorBar `json:"YErrorBar,omitempty" xml:"YErrorBar"`
}
