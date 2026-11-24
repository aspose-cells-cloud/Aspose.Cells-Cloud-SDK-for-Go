/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="TextWaterMarkerRequest">
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

// TextWaterMarkerRequest Indicates text water marker request.
type TextWaterMarkerRequest struct {
	// A property named Text of type string that can be accessed and modified.
	Text string `json:"Text,omitempty" xml:"Text"`
	// Indicates font name.
	FontName string `json:"FontName,omitempty" xml:"FontName"`
	// Indicates font size.
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize"`
	// Indicates image height.
	Height *int32 `json:"Height,omitempty" xml:"Height"`
	// Indicates image width.
	Width            *int32 `json:"Width,omitempty" xml:"Width"`
	ImageAdaptOption string `json:"ImageAdaptOption,omitempty" xml:"ImageAdaptOption"`
}
