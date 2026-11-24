/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Walls">
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

// Walls Encapsulates the object that represents the walls of a 3-D chart.
type Walls struct {
	Floor
	// Gets the x coordinate of the left-bottom corner of Wall center in units of 1/4000 of chart's width after calls Chart.Calculate() method.
	CenterX *int32 `json:"CenterX,omitempty" xml:"CenterX"`
	// Gets the y coordinate of the left-bottom corner of Wall center in units of 1/4000 of chart's height after calls Chart.Calculate() method.
	CenterY *int32 `json:"CenterY,omitempty" xml:"CenterY"`
	// Gets the depth front to back in units of 1/4000 of chart's width after calls Chart.Calculate() method.
	Depth *int32 `json:"Depth,omitempty" xml:"Depth"`
	// Gets the height of top to bottom in units of 1/4000 of chart's height after calls Chart.Calculate() method.
	Height *int32 `json:"Height,omitempty" xml:"Height"`
	// Gets the width of left to right in units of 1/4000 of chart's width after calls Chart.Calculate() method.
	Width            *int32      `json:"Width,omitempty" xml:"Width"`
	BackgroundColor  *Color      `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
	FillFormat       *FillFormat `json:"FillFormat,omitempty" xml:"FillFormat"`
	ForegroundColor  *Color      `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
	Format           string      `json:"Format,omitempty" xml:"Format"`
	InvertIfNegative *bool       `json:"InvertIfNegative,omitempty" xml:"InvertIfNegative"`
	Transparency     *float64    `json:"Transparency,omitempty" xml:"Transparency"`
}
