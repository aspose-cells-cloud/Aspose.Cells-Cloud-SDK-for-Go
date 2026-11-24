/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="TextOptions">
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

// TextOptions Represents the text options.
type TextOptions struct {
	Font
	// Represents fill format.
	Fill *FillFormat `json:"Fill,omitempty" xml:"Fill"`
	// Represents kerning.
	Kerning *float64 `json:"Kerning,omitempty" xml:"Kerning"`
	// Represents outline format.
	Outline *LineFormat `json:"Outline,omitempty" xml:"Outline"`
	// Represents shadow effect.
	Shadow *ShadowEffect `json:"Shadow,omitempty" xml:"Shadow"`
	// Represents spacing.
	Spacing *float64 `json:"Spacing,omitempty" xml:"Spacing"`
	// Represents under line color.
	UnderlineColor *CellsColor `json:"UnderlineColor,omitempty" xml:"UnderlineColor"`
}
