/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AggregateResultByColor">
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

// AggregateResultByColor
type AggregateResultByColor struct {
	AggregateOperation string   `json:"AggregateOperation,omitempty" xml:"AggregateOperation"`
	ColorName          string   `json:"ColorName,omitempty" xml:"ColorName"`
	Count              *int32   `json:"Count,omitempty" xml:"Count"`
	Sum                *float64 `json:"Sum,omitempty" xml:"Sum"`
	MaxValue           *float64 `json:"MaxValue,omitempty" xml:"MaxValue"`
	MinValue           *float64 `json:"MinValue,omitempty" xml:"MinValue"`
	AverageValue       *float64 `json:"AverageValue,omitempty" xml:"AverageValue"`
}
