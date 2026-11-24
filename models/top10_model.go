/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Top10">
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

// Top10 Describe the Top10 conditional formatting rule. This conditional formatting     rule highlights cells whose values fall in the top N or bottom N bracket,     as specified.
type Top10 struct {
	// Get or set whether a "top/bottom n" rule is a "bottom n" rule.             Default value is false.
	IsBottom *bool `json:"IsBottom,omitempty" xml:"IsBottom"`
	// Get or set whether a "top/bottom n" rule is a "top/bottom n percent" rule.             Default value is false.
	IsPercent *bool `json:"IsPercent,omitempty" xml:"IsPercent"`
	// Get or set the value of "n" in a "top/bottom n" conditional formatting rule.             If IsPercent is true, the value must between 0 and 100.             Otherwise it must between 0 and 1000.             Default value is 10.
	Rank *int32 `json:"Rank,omitempty" xml:"Rank"`
}
