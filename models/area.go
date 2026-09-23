/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="area.go">
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

// Area            Encapsulates the object that represents an area format.
type Area struct {
	// Gets or sets the background  of the .
	BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
	// Represents a  object that contains fill formatting properties for the specified chart or shape.
	FillFormat *FillFormat `json:"FillFormat,omitempty" xml:"FillFormat"`
	// Gets or sets the foreground .
	ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
	// This class has a property named "Format" of type string with both getter and setter methods.
	Format string `json:"Format,omitempty" xml:"Format"`
	// If the property is true and the value of chart point is a negative number,             the foreground color and background color will be exchanged.
	InvertIfNegative *bool `json:"InvertIfNegative,omitempty" xml:"InvertIfNegative"`
	// Returns or sets the degree of transparency of the area as a value from 0.0 (opaque) through 1.0 (clear).
	Transparency *float64 `json:"Transparency,omitempty" xml:"Transparency"`
}

// areaMarker closes the AreaLike family: only Area and the models embedding it carry this
// promoted method, so no unrelated type satisfies AreaLike.
func (m Area) areaMarker() {}
