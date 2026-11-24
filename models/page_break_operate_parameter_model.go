/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PageBreakOperateParameter">
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

// PageBreakOperateParameter Represents page break operate parameter.
type PageBreakOperateParameter struct {
	OperateParameter
	// Represents page break type.
	PageBreakType string `json:"PageBreakType,omitempty" xml:"PageBreakType"`
	// Represents page break index.
	Index *int32 `json:"Index,omitempty" xml:"Index"`
	// Represents row index of page break.
	Row *int32 `json:"Row,omitempty" xml:"Row"`
	// Represents column index of page break.
	Column *int32 `json:"Column,omitempty" xml:"Column"`
	// Represents start row index of page break.
	StartIndex *int32 `json:"StartIndex,omitempty" xml:"StartIndex"`
	// Represents end row index of page break.
	EndIndex *int32 `json:"EndIndex,omitempty" xml:"EndIndex"`
}
