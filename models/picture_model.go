/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Picture">
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

// Picture Encapsulates the object that represents a single picture in a spreadsheet.
type Picture struct {
	Shape
	// Represents the  of the border line of a picture.
	BorderLineColor *Color `json:"BorderLineColor,omitempty" xml:"BorderLineColor"`
	// Gets or sets the weight of the border line of a picture in units of pt.
	BorderWeight *float64 `json:"BorderWeight,omitempty" xml:"BorderWeight"`
	// Gets the original height of the picture.
	OriginalHeight *int32 `json:"OriginalHeight,omitempty" xml:"OriginalHeight"`
	// Gets the original width of the picture.
	OriginalWidth *int32 `json:"OriginalWidth,omitempty" xml:"OriginalWidth"`
	// This class has a property called "ImageFormat" that allows getting and setting a string value.
	ImageFormat string `json:"ImageFormat,omitempty" xml:"ImageFormat"`
	// Gets or sets the path and name of the source file for the linked image.
	SourceFullName string `json:"SourceFullName,omitempty" xml:"SourceFullName"`
	Link           *Link  `json:"link,omitempty" xml:"link"`
}
