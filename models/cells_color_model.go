/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CellsColor">
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

// CellsColor            Represents all types of color.
type CellsColor struct {
	// Gets and sets the RGB color.
	Color *Color `json:"Color,omitempty" xml:"Color"`
	// Gets and sets the color index in the color palette. Only applies of indexed color.
	ColorIndex *int32 `json:"ColorIndex,omitempty" xml:"ColorIndex"`
	// Gets and set the color which should apply to cell or shape.
	IsShapeColor *bool `json:"IsShapeColor,omitempty" xml:"IsShapeColor"`
	// Set the tint of the shape color
	Tint *float64 `json:"tint,omitempty" xml:"tint"`
	// Gets and sets the color from a 32-bit ARGB value.
	Argb *int32 `json:"Argb,omitempty" xml:"Argb"`
	// Gets the theme color. Only applies for theme color type.
	ThemeColor *ThemeColor `json:"ThemeColor,omitempty" xml:"ThemeColor"`
	// The color type.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Gets and sets transparency as a value from 0.0 (opaque) through 1.0 (clear).
	Transparency *float64 `json:"Transparency,omitempty" xml:"Transparency"`
}
