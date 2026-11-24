/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PivotFilter">
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

// PivotFilter Represents a PivotFilter in PivotFilter Collection.
type PivotFilter struct {
	// Gets the autofilter of the pivot filter.
	AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
	// Gets the Evaluation Order of the pivot filter.
	EvaluationOrder *int32 `json:"EvaluationOrder,omitempty" xml:"EvaluationOrder"`
	// Gets the field index of the pivot filter.
	FieldIndex *int32 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
	// Gets the autofilter type of the pivot filter.
	FilterType string `json:"FilterType,omitempty" xml:"FilterType"`
	// Gets the measure field index of the pivot filter.
	ValueFieldIndex *int32 `json:"ValueFieldIndex,omitempty" xml:"ValueFieldIndex"`
	// Gets the member property field index of the pivot filter.
	MemberPropertyFieldIndex *int32 `json:"MemberPropertyFieldIndex,omitempty" xml:"MemberPropertyFieldIndex"`
	// Gets the name of the pivot filter.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets the string value1 of the label pivot filter.
	Value1 string `json:"Value1,omitempty" xml:"Value1"`
	// Gets the string value2 of the label pivot filter.
	Value2 string `json:"Value2,omitempty" xml:"Value2"`
	// A property that allows for setting and getting a Top10Filter object for filtering data.
	Top10Filter *Top10Filter `json:"Top10Filter,omitempty" xml:"Top10Filter"`
}
