/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ColorScale">
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

// ColorScale Describe the ColorScale conditional formatting rule. This conditional formatting    rule creates a gradated color scale on the cells.
type ColorScale struct {
	// Get or set this ColorScale's max value object.             Cannot set null or CFValueObject with type FormatConditionValueType.Min to it.
	MaxCfvo *ConditionalFormattingValue `json:"MaxCfvo,omitempty" xml:"MaxCfvo"`
	// Get or set the gradient color for the maximum value in the range.
	MaxColor *Color `json:"MaxColor,omitempty" xml:"MaxColor"`
	// Get or set this ColorScale's mid value object.             Cannot set CFValueObject with type FormatConditionValueType.Max or FormatConditionValueType.Min to it.
	MidCfvo *ConditionalFormattingValue `json:"MidCfvo,omitempty" xml:"MidCfvo"`
	// Get or set the gradient color for the middle value in the range.
	MidColor *Color `json:"MidColor,omitempty" xml:"MidColor"`
	// Get or set this ColorScale's min value object.             Cannot set null or CFValueObject with type FormatConditionValueType.Max to it.
	MinCfvo *ConditionalFormattingValue `json:"MinCfvo,omitempty" xml:"MinCfvo"`
	// Get or set the gradient color for the minimum value in the range.
	MinColor *Color `json:"MinColor,omitempty" xml:"MinColor"`
}
