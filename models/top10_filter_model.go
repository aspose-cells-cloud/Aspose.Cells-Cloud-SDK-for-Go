/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Top10Filter">
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

// Top10Filter            Represents the top 10 filter.
type Top10Filter struct {
	// An integer property named FieldIndex that can be accessed and modified.
	FieldIndex *int32 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
	//
	Criteria string `json:"Criteria,omitempty" xml:"Criteria"`
	// Indicates whether the items is percent.
	IsPercent *bool `json:"IsPercent,omitempty" xml:"IsPercent"`
	// Indicates whether it's top filter.
	IsTop *bool `json:"IsTop,omitempty" xml:"IsTop"`
	// Gets and sets the items of the filter.
	Items *int32 `json:"Items,omitempty" xml:"Items"`
}
