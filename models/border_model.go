/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Border">
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

// Border            Encapsulates the object that represents the cell border.
type Border struct {
	// Gets or sets the cell border type.
	LineStyle string `json:"LineStyle,omitempty" xml:"LineStyle"`
	// Gets or sets the  of the border.
	Color *Color `json:"Color,omitempty" xml:"Color"`
	// This class has a property called "BorderType" of type string that can be both get and set.
	BorderType string `json:"BorderType,omitempty" xml:"BorderType"`
	// Gets and sets the theme color of the border.
	ThemeColor *ThemeColor `json:"ThemeColor,omitempty" xml:"ThemeColor"`
	// Gets and sets the color with a 32-bit ARGB value.
	ArgbColor *int32 `json:"ArgbColor,omitempty" xml:"ArgbColor"`
}
