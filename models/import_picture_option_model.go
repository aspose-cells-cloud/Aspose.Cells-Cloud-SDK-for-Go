/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ImportPictureOption">
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

// ImportPictureOption Class summary: The features of the topic were explored, discussing its components, patterns, and significance.
type ImportPictureOption struct {
	ImportOption
	// A property "UpperLeftRow" of type integer with a public getter and setter is defined.
	UpperLeftRow     *int32 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
	UpperLeftColumn  *int32 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
	LowerRightRow    *int32 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
	LowerRightColumn *int32 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
	Filename         string `json:"Filename,omitempty" xml:"Filename"`
	// base64
	Data string `json:"Data,omitempty" xml:"Data"`
}
