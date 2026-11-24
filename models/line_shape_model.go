/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LineShape">
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

// LineShape Represents the line shape.
type LineShape struct {
	Shape
	// Gets and sets the begin arrow head length of the line.
	BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
	// Gets and sets the begin arrow head style of the line.
	BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
	// Gets and sets the begin arrow head width of the line.
	BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
	// Gets and sets the end arrow head length of the line.
	EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
	// Gets and sets the end arrow head style of the line.
	EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
	// Gets and sets the end arrow head width of the line.
	EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
	Link              *Link  `json:"link,omitempty" xml:"link"`
}
