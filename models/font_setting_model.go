/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="FontSetting">
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

// FontSetting            Represents a range of characters within the cell text.
type FontSetting struct {
	// Returns the font of this object.
	Font *Font `json:"Font,omitempty" xml:"Font"`
	// Gets the length of the characters.
	Length *int32 `json:"Length,omitempty" xml:"Length"`
	// Gets the start index of the characters.
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex"`
	// Returns the text options.
	TextOptions *TextOptions `json:"TextOptions,omitempty" xml:"TextOptions"`
	// Gets the type of text node.
	Type string `json:"Type,omitempty" xml:"Type"`
}
