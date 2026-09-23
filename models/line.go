/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="line.go">
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

// Line Encapsulates the object that represents the line format.
type Line struct {
	// Specifies the length of the arrowhead for the begin of a line.
	BeginArrowLength string `json:"BeginArrowLength,omitempty" xml:"BeginArrowLength"`
	// Specifies the width of the arrowhead for the begin of a line.
	BeginArrowWidth string `json:"BeginArrowWidth,omitempty" xml:"BeginArrowWidth"`
	// Specifies an arrowhead for the begin of a line.
	BeginType string `json:"BeginType,omitempty" xml:"BeginType"`
	// Specifies the ending caps.
	CapType string `json:"CapType,omitempty" xml:"CapType"`
	// Represents the  of the line.
	Color *Color `json:"Color,omitempty" xml:"Color"`
	// Specifies the compound line type
	CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
	// Specifies the dash line type
	DashType string `json:"DashType,omitempty" xml:"DashType"`
	// Specifies the length of the arrowhead for the end of a line.
	EndArrowLength string `json:"EndArrowLength,omitempty" xml:"EndArrowLength"`
	// Specifies the width of the arrowhead for the end of a line.
	EndArrowWidth string `json:"EndArrowWidth,omitempty" xml:"EndArrowWidth"`
	// Specifies an arrowhead for the end of a line.
	EndType string `json:"EndType,omitempty" xml:"EndType"`
	// Represents gradient fill.
	GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
	// Indicates whether this line style is auto assigned.
	IsAuto *bool `json:"IsAuto,omitempty" xml:"IsAuto"`
	// Indicates whether the color of line is automatic assigned.
	IsAutomaticColor *bool `json:"IsAutomaticColor,omitempty" xml:"IsAutomaticColor"`
	// Represents whether the line is visible.
	IsVisible *bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// Specifies the joining caps.
	JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
	// Represents the style of the line.
	Style string `json:"Style,omitempty" xml:"Style"`
	// Returns or sets the degree of transparency of the line as a value from 0.0 (opaque) through 1.0 (clear).
	Transparency *float64 `json:"Transparency,omitempty" xml:"Transparency"`
	// Gets or sets the  of the line.
	Weight string `json:"Weight,omitempty" xml:"Weight"`
	// Gets or sets the weight of the line in unit of points.
	WeightPt *float64 `json:"WeightPt,omitempty" xml:"WeightPt"`
}

// lineMarker closes the LineLike family: only Line and the models embedding it carry this
// promoted method, so no unrelated type satisfies LineLike.
func (m Line) lineMarker() {}
