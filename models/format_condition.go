/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="format_condition.go">
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

// FormatCondition Represents conditional formatting condition.
type FormatCondition struct {
	LinkElement
	// The priority of this conditional formatting rule. This value is used to determine which                         format should be evaluated and rendered. Lower numeric values are higher priority than                         higher numeric values, where '1' is the highest priority.
	Priority *int32 `json:"Priority,omitempty" xml:"Priority"`
	// Gets and sets whether the conditional format Type.
	Type string `json:"Type,omitempty" xml:"Type"`
	// True, no rules with lower priority may be applied over this rule, when this rule evaluates to true.                         Only applies for Excel 2007;
	StopIfTrue *bool `json:"StopIfTrue,omitempty" xml:"StopIfTrue"`
	// Get the conditional formatting's "AboveAverage" instance.                         The default instance's rule highlights cells that are                          above the average for all values in the range.                         Valid only for type = AboveAverage.
	AboveAverage *AboveAverage `json:"AboveAverage,omitempty" xml:"AboveAverage"`
	// Get the conditional formatting's "ColorScale" instance.                         The default instance is a "green-yellow-red" 3ColorScale .                         Valid only for type = ColorScale.
	ColorScale *ColorScale `json:"ColorScale,omitempty" xml:"ColorScale"`
	// Get the conditional formatting's "DataBar" instance.                         The default instance's color is blue.                         Valid only for type is DataBar.
	DataBar *DataBar `json:"DataBar,omitempty" xml:"DataBar"`
	// Gets and sets the value or expression associated with conditional formatting.
	Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
	// Gets and sets the value or expression associated with conditional formatting.
	Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
	// Get the conditional formatting's "IconSet" instance.                         The default instance's IconSetType is TrafficLights31.                         Valid only for type = IconSet.
	IconSet *IconSet `json:"IconSet,omitempty" xml:"IconSet"`
	// Gets and sets the conditional format operator type.
	Operator string `json:"Operator,omitempty" xml:"Operator"`
	// Gets or setts style of conditional formatted cell ranges.
	Style *Style `json:"Style,omitempty" xml:"Style"`
	// The text value in a "text contains" conditional formatting rule.                          Valid only for type = containsText, notContainsText, beginsWith and endsWith.                         The default value is null.
	Text string `json:"Text,omitempty" xml:"Text"`
	// The applicable time period in a "date occurring…" conditional formatting rule.                          Valid only for type = timePeriod.                         The default value is TimePeriodType.Today.
	TimePeriod string `json:"TimePeriod,omitempty" xml:"TimePeriod"`
	// Get the conditional formatting's "Top10" instance.                         The default instance's rule highlights cells whose                         values fall in the top 10 bracket.                         Valid only for type is Top10.
	Top10 *Top10 `json:"Top10,omitempty" xml:"Top10"`
}

// formatConditionMarker closes the FormatConditionLike family: only FormatCondition and the
// models embedding it carry this promoted method, so no unrelated type satisfies
// FormatConditionLike.
func (m FormatCondition) formatConditionMarker() {}
