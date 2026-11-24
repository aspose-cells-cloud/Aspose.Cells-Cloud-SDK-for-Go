/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PicFormatOption">
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

// PicFormatOption            Represents picture format option
type PicFormatOption struct {
	// Gets or sets the picture fill type.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Gets or sets how many the picture stack and scale with.
	Scale *float64 `json:"Scale,omitempty" xml:"Scale"`
	// Gets or sets the left offset for stretching picture.
	Left *float64 `json:"Left,omitempty" xml:"Left"`
	// Gets or sets the right offset for stretching picture.
	Right *float64 `json:"Right,omitempty" xml:"Right"`
	// Gets or sets the top offset for stretching picture.
	Top *float64 `json:"Top,omitempty" xml:"Top"`
	// Gets or sets the bottom offset for stretching picture.
	Bottom *float64 `json:"Bottom,omitempty" xml:"Bottom"`
}
