/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LineFormat">
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

// LineFormat Represents all setting of the line.
type LineFormat struct {
	FillFormat
	// Gets and sets the begin arrow length type of the line.
	BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
	// Gets and sets the begin arrow type of the line.
	BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
	// Gets and sets the begin arrow width type of the line.
	BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
	// Specifies the ending caps.
	CapType string `json:"CapType,omitempty" xml:"CapType"`
	// Specifies the line compound type.
	CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
	// Specifies the line dash type.
	DashStyle string `json:"DashStyle,omitempty" xml:"DashStyle"`
	// Gets and sets the end arrow length type of the line.
	EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
	// Gets and sets the end arrow type of the line.
	EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
	// Gets and sets the end arrow width type of the line.
	EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
	// Specifies the line join type.
	JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
	// Gets or sets the weight of the line in unit of points.
	Weight *float64 `json:"Weight,omitempty" xml:"Weight"`
}
