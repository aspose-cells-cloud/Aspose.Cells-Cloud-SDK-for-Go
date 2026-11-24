/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Title">
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

// Title Encapsulates the object that represents the title of chart or axis.
type Title struct {
	ChartFrame
	// Represents whether the title is visible.
	IsVisible *bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// A property named LinkedSource with both set and get accessors is defined.
	LinkedSource  string `json:"LinkedSource,omitempty" xml:"LinkedSource"`
	RotationAngle *int32 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	// Gets or sets the text of display unit label.
	Text                    string `json:"Text,omitempty" xml:"Text"`
	TextDirection           string `json:"TextDirection,omitempty" xml:"TextDirection"`
	TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
	TextVerticalAlignment   string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}
