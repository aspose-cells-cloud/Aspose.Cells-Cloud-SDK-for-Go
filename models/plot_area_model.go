/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PlotArea">
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

// PlotArea Encapsulates the object that represents the plot area in a chart.
type PlotArea struct {
	ChartFrame
	// Gets or sets the height of plot area in units of 1/4000 of the chart area.
	InnerHeight *int32 `json:"InnerHeight,omitempty" xml:"InnerHeight"`
	// Gets or sets the width  of plot area in units of 1/4000 of the chart area.
	InnerWidth *int32 `json:"InnerWidth,omitempty" xml:"InnerWidth"`
	// Gets or gets the x coordinate of the upper top corner of plot area in units of 1/4000 of the chart area.
	InnerX *int32 `json:"InnerX,omitempty" xml:"InnerX"`
	// Gets or gets the x coordinate of the upper top corner of plot area in units of 1/4000 of the chart area.
	InnerY *int32 `json:"InnerY,omitempty" xml:"InnerY"`
}
