/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ErrorBar">
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

// ErrorBar Represents error bar of data series.
type ErrorBar struct {
	Line
	// A property named "Link" of type "Link" that can be accessed and modified.
	Link *Link `json:"Link,omitempty" xml:"Link"`
	// Represents amount of error bar.                          The amount must be greater than or equal to zero.
	Amount *float64 `json:"Amount,omitempty" xml:"Amount"`
	// Represents error bar display type.
	DisplayType string `json:"DisplayType,omitempty" xml:"DisplayType"`
	// Represents negative error amount when error bar type is Custom.
	MinusValue string `json:"MinusValue,omitempty" xml:"MinusValue"`
	// Represents positive error amount when error bar type is Custom.
	PlusValue string `json:"PlusValue,omitempty" xml:"PlusValue"`
	// Indicates if formatting error bars with a T-top.
	ShowMarkerTTop *bool `json:"ShowMarkerTTop,omitempty" xml:"ShowMarkerTTop"`
	// Represents error bar amount type.
	Type string `json:"Type,omitempty" xml:"Type"`
}
