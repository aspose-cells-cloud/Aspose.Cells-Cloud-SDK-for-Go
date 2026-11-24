/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="GlobalizationSettings">
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

// GlobalizationSettings Represents the globalization settings.
type GlobalizationSettings struct {
	// Gets or sets the globalization settings for Chart.
	ChartSettings *ChartGlobalizationSettings `json:"ChartSettings,omitempty" xml:"ChartSettings"`
	// Gets or sets the globalization settings for pivot table.
	PivotSettings *PivotGlobalizationSettings `json:"PivotSettings,omitempty" xml:"PivotSettings"`
	// Gets the separator for list, parameters of function, ...etc.
	ListSeparator string `json:"ListSeparator,omitempty" xml:"ListSeparator"`
	// Gets the separator for rows in array data in formula.
	RowSeparatorOfFormulaArray string `json:"RowSeparatorOfFormulaArray,omitempty" xml:"RowSeparatorOfFormulaArray"`
	// Gets the separator for the items in array's row data in formula.
	ColumnSeparatorOfFormulaArray string `json:"ColumnSeparatorOfFormulaArray,omitempty" xml:"ColumnSeparatorOfFormulaArray"`
}
