/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ConditionalFormattingValue">
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

// ConditionalFormattingValue            Describes the values of the interpolation points in a gradient scale, dataBar or iconSet.
type ConditionalFormattingValue struct {
	// Get or set the Greater Than Or Equal flag.              Use only for icon sets, determines whether this threshold value uses              the greater than or equal to operator.              'false' indicates 'greater than' is used instead of 'greater than or equal to'.             Default value is true.
	IsGTE *bool `json:"IsGTE,omitempty" xml:"IsGTE"`
	// Get or set the type of this conditional formatting value object.             Setting the type to FormatConditionValueType.Min or FormatConditionValueType.Max              will auto set "Value" to null.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Get or set the value of this conditional formatting value object.             It should be used in conjunction with Type.
	Value map[string]interface{} `json:"Value,omitempty" xml:"Value"`
}
