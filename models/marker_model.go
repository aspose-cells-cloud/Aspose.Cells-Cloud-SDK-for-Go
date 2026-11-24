/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Marker">
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

// Marker            Represents the marker in a line chart, scatter chart, or radar chart.
type Marker struct {
	// Gets the border.
	Border *Line `json:"Border,omitempty" xml:"Border"`
	// Gets the area.
	Area *Area `json:"Area,omitempty" xml:"Area"`
	// Represents the marker size in unit of points. Applies to line chart, scatter chart, or radar chart.
	MarkerSize *int32 `json:"MarkerSize,omitempty" xml:"MarkerSize"`
	// Represents the marker style. Applies to line chart, scatter chart, or radar chart.
	MarkerStyle string `json:"MarkerStyle,omitempty" xml:"MarkerStyle"`
}
