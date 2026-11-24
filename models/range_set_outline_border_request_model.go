/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="RangeSetOutlineBorderRequest">
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

// RangeSetOutlineBorderRequest Indicates range set outline border request.
type RangeSetOutlineBorderRequest struct {
	// Encapsulates the object that represents a range of cells within a spreadsheet.
	Range *Range `json:"Range,omitempty" xml:"Range"`
	// LeftBorder, RightBorder, TopBorder, BottomBorder, DiagonalDown, DiagonalUp, Vertical and Horizontal.
	BorderEdge string `json:"borderEdge,omitempty" xml:"borderEdge"`
	// None, Thin, Medium, Dashed, Dotted, Thick, Double, Hair, MediumDashed, DashDot, MediumDashDot, DashDotDot, MediumDashDotDot and SlantedDashDot.
	BorderStyle string `json:"borderStyle,omitempty" xml:"borderStyle"`
	// Border color.
	BorderColor *Color `json:"borderColor,omitempty" xml:"borderColor"`
}
