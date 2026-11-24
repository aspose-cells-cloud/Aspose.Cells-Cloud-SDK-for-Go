/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Hyperlink">
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

// Hyperlink Encapsulates the object that represents a hyperlink.
type Hyperlink struct {
	LinkElement
	// Represents the address of a hyperlink.
	Address string `json:"Address,omitempty" xml:"Address"`
	// Gets the range of hyperlink.
	Area *CellArea `json:"Area,omitempty" xml:"Area"`
	// Returns or sets the ScreenTip text for the specified hyperlink.
	ScreenTip string `json:"ScreenTip,omitempty" xml:"ScreenTip"`
	// Represents the text to be displayed for the specified hyperlink. The default value is the address of the hyperlink.
	TextToDisplay string `json:"TextToDisplay,omitempty" xml:"TextToDisplay"`
	// Gets the link type.
	LinkType string `json:"LinkType,omitempty" xml:"LinkType"`
}
