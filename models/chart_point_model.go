/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ChartPoint">
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

// ChartPoint Represents a single point in a series in a chart.
type ChartPoint struct {
	LinkElement
	// Gets the area.
	Area *Area `json:"Area,omitempty" xml:"Area"`
	// Gets the border.
	Border *Line `json:"Border,omitempty" xml:"Border"`
	// Returns a DataLabels object that represents the data label associated with the point.
	DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
	// The distance of an open pie slice from the center of the pie chart is expressed as a percentage of the pie diameter.
	Explosion *int32 `json:"Explosion,omitempty" xml:"Explosion"`
	// Gets the marker.
	Marker *Marker `json:"Marker,omitempty" xml:"Marker"`
	// True if the chartpoint has a shadow.
	Shadow *bool `json:"Shadow,omitempty" xml:"Shadow"`
	// Gets or sets the X value of the chart point.
	XValue map[string]interface{} `json:"XValue,omitempty" xml:"XValue"`
	// Gets or sets the Y value of the chart point.
	YValue map[string]interface{} `json:"YValue,omitempty" xml:"YValue"`
	// Gets or sets a value indicates whether this data points is in the second pie or bar on a pie of pie or bar of pie chart.
	IsInSecondaryPlot *bool `json:"IsInSecondaryPlot,omitempty" xml:"IsInSecondaryPlot"`
}
