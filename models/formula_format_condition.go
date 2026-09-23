/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="formula_format_condition.go">
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

// FormulaFormatCondition Class summary: This class covers various features related to data visualization and graphical representation of data sets.
type FormulaFormatCondition struct {
	// Gets and sets the value or expression associated with conditional formatting.
	Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
	// Gets and sets the value or expression associated with conditional formatting.
	Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
	// Gets and sets the conditional format operator type.
	Operator string `json:"Operator,omitempty" xml:"Operator"`
}

// formulaFormatConditionMarker closes the FormulaFormatConditionLike family: only
// FormulaFormatCondition and the models embedding it carry this promoted method, so no
// unrelated type satisfies FormulaFormatConditionLike.
func (m FormulaFormatCondition) formulaFormatConditionMarker() {}
