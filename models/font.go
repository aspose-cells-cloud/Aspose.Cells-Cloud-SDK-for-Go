/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="font.go">
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

// Font            Encapsulates the font object used in a spreadsheet.
type Font struct {
	// Gets or sets the  of the font.
	Color *Color `json:"Color,omitempty" xml:"Color"`
	// Gets and sets the double size of the font.
	DoubleSize *float64 `json:"DoubleSize,omitempty" xml:"DoubleSize"`
	// Gets or sets a value indicating whether the font is bold.
	IsBold *bool `json:"IsBold,omitempty" xml:"IsBold"`
	// Gets or sets a value indicating whether the font is italic.
	IsItalic *bool `json:"IsItalic,omitempty" xml:"IsItalic"`
	// Gets or sets a value indicating whether the font is single strikeout.
	IsStrikeout *bool `json:"IsStrikeout,omitempty" xml:"IsStrikeout"`
	// Gets or sets a value indicating whether the font is subscript.
	IsSubscript *bool `json:"IsSubscript,omitempty" xml:"IsSubscript"`
	// Gets or sets a value indicating whether the font is super script.
	IsSuperscript *bool `json:"IsSuperscript,omitempty" xml:"IsSuperscript"`
	// Gets  or sets the name of the .
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets or sets the size of the font.
	Size *int32 `json:"Size,omitempty" xml:"Size"`
	// Gets or sets the font underline type.
	Underline string `json:"Underline,omitempty" xml:"Underline"`
}

// fontMarker closes the FontLike family: only Font and the models embedding it carry this
// promoted method, so no unrelated type satisfies FontLike.
func (m Font) fontMarker() {}
