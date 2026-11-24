/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Style">
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

// Style            Represents display style of excel document,such as font,color,alignment,border,etc.            The Style object contains all style attributes (font, number format, alignment, and so on) as properties.
type Style struct {
	// Gets a  object.
	Font *Font `json:"Font,omitempty" xml:"Font"`
	// Gets or sets the name of the style.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets and sets the culture-dependent pattern string for number format.             If no number format has been set for this object, null will be returned.             If number format is builtin, the pattern string corresponding to the builtin number will be returned.
	CultureCustom string `json:"CultureCustom,omitempty" xml:"CultureCustom"`
	// Represents the custom number format string of this style object.             If the custom number format is not set(For example, the number format is builtin), "" will be returned.
	Custom string `json:"Custom,omitempty" xml:"Custom"`
	// Gets or sets a style's background color.
	BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
	// Gets or sets a style's foreground color.
	ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
	// Represents if the formula will be hidden when the worksheet is protected.
	IsFormulaHidden *bool `json:"IsFormulaHidden,omitempty" xml:"IsFormulaHidden"`
	// Indicates whether the number format is a date format.
	IsDateTime *bool `json:"IsDateTime,omitempty" xml:"IsDateTime"`
	// Gets or sets a value indicating whether the text within a cell is wrapped.
	IsTextWrapped *bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
	// Indicates whether the cell shading is a gradient pattern.
	IsGradient *bool `json:"IsGradient,omitempty" xml:"IsGradient"`
	// Gets or sets a value indicating whether a cell can be modified or not.
	IsLocked *bool `json:"IsLocked,omitempty" xml:"IsLocked"`
	// Indicates whether the number format is a percent format.
	IsPercent *bool `json:"IsPercent,omitempty" xml:"IsPercent"`
	// Represents if text automatically shrinks to fit in the available column width.
	ShrinkToFit *bool `json:"ShrinkToFit,omitempty" xml:"ShrinkToFit"`
	// Represents the indent level for the cell or range. Can only be an integer from 0 to 250.
	IndentLevel *int32 `json:"IndentLevel,omitempty" xml:"IndentLevel"`
	// Gets or sets the display format of numbers and dates. The formatting patterns are different for different regions.
	Number *int32 `json:"Number,omitempty" xml:"Number"`
	// Represents text rotation angle.
	RotationAngle *int32 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	// Gets or sets the cell background pattern type.
	Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
	// Represents text reading order.
	TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
	// Gets or sets the vertical alignment type of the text in a cell.
	VerticalAlignment string `json:"VerticalAlignment,omitempty" xml:"VerticalAlignment"`
	// Gets or sets the horizontal alignment type of the text in a cell.
	HorizontalAlignment string `json:"HorizontalAlignment,omitempty" xml:"HorizontalAlignment"`
	// A public property named `BorderCollection` that is a list of `Border` objects.
	BorderCollection []Border `json:"BorderCollection,omitempty" xml:"BorderCollection"`
	// Gets and sets the background theme color.
	BackgroundThemeColor *ThemeColor `json:"BackgroundThemeColor,omitempty" xml:"BackgroundThemeColor"`
	// Gets and sets the foreground theme color.
	ForegroundThemeColor *ThemeColor `json:"ForegroundThemeColor,omitempty" xml:"ForegroundThemeColor"`
}
