/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="chart_frame.go">
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

// ChartFrame            Encapsulates the object that represents the frame object in a chart.
type ChartFrame struct {
	// Gets the area.
	Area *Area `json:"Area,omitempty" xml:"Area"`
	// True if the text in the object changes font size when the object size changes. The default value is True.
	AutoScaleFont *bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
	// Gets and sets the display mode of the background
	BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
	// Gets the border.
	Border *Line `json:"Border,omitempty" xml:"Border"`
	// Gets a  object of the specified ChartFrame object.
	Font *Font `json:"Font,omitempty" xml:"Font"`
	// Indicates whether the chart frame is automatic sized.
	IsAutomaticSize *bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
	// Indicates whether the size of the plot area size includes the tick marks, and the axis labels.             False specifies that the size shall determine the size of the plot area, the tick marks, and the axis labels.
	IsInnerMode *bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
	// True if the frame has a shadow.
	Shadow *bool `json:"Shadow,omitempty" xml:"Shadow"`
	// Gets or sets the width of frame in units of 1/4000 of the chart area.
	Width *int32 `json:"Width,omitempty" xml:"Width"`
	// Gets or sets the height of frame in units of 1/4000 of the chart area.
	Height *int32 `json:"Height,omitempty" xml:"Height"`
	// Gets or sets the x coordinate of the upper left corner in units of 1/4000 of the chart area.
	X *int32 `json:"X,omitempty" xml:"X"`
	// Gets or sets the y coordinate of the upper left corner in units of 1/4000 of the chart area.
	Y *int32 `json:"Y,omitempty" xml:"Y"`
}

// chartFrameMarker closes the ChartFrameLike family: only ChartFrame and the models
// embedding it carry this promoted method, so no unrelated type satisfies ChartFrameLike.
func (m ChartFrame) chartFrameMarker() {}
