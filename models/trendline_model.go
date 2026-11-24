/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Trendline">
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

// Trendline Represents a trendline in a chart.
type Trendline struct {
	Line
	// The class has a property named "link" of type "Link" that can be accessed and modified.
	Link *Link `json:"link,omitempty" xml:"link"`
	// Returns or sets the number of periods (or units on a scatter chart) that the trendline extends backward.                          The number of periods must be greater than or equal to zero.                         If the chart type is column ,the number of periods must be between 0 and 0.5
	Backward *float64 `json:"Backward,omitempty" xml:"Backward"`
	// Represents the DataLabels object for the specified series.
	DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
	// Represents if the equation for the trendline is displayed on the chart (in the same data label as the R-squared value). Setting this property to True automatically turns on data labels.
	DisplayEquation *bool `json:"DisplayEquation,omitempty" xml:"DisplayEquation"`
	// Represents if the R-squared value of the trendline is displayed on the chart (in the same data label as the equation). Setting this property to True automatically turns on data labels.
	DisplayRSquared *bool `json:"DisplayRSquared,omitempty" xml:"DisplayRSquared"`
	// Returns or sets the number of periods (or units on a scatter chart) that the trendline extends forward.                         The number of periods must be greater than or equal to zero.
	Forward *float64 `json:"Forward,omitempty" xml:"Forward"`
	// Returns or sets the point where the trendline crosses the value axis.
	Intercept *float64 `json:"Intercept,omitempty" xml:"Intercept"`
	// Returns if Microsoft Excel automatically determines the name of the trendline.
	IsNameAuto *bool `json:"IsNameAuto,omitempty" xml:"IsNameAuto"`
	// Gets the legend entry according to this trendline
	LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
	// Returns the name of the trendline.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Returns or sets the trendline order (an integer greater than 1) when the trendline type is Polynomial.                          The order must be between 2 and 6.
	Order *int32 `json:"Order,omitempty" xml:"Order"`
	// Returns or sets the period for the moving-average trendline.
	Period *int32 `json:"Period,omitempty" xml:"Period"`
	// Returns the trendline type.
	Type string `json:"Type,omitempty" xml:"Type"`
}
