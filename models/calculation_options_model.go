/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CalculationOptions">
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

// CalculationOptions            Represents options for calculation.
type CalculationOptions struct {
	// Specifies the stack size for calculating cells recursively.
	CalcStackSize *int32 `json:"CalcStackSize,omitempty" xml:"CalcStackSize"`
	// Indicates whether errors encountered while calculating formulas should be ignored.             The error may be unsupported function, external links, etc.             The default value is true.
	IgnoreError *bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
	// Specifies the strategy for processing precision of calculation.
	PrecisionStrategy string `json:"PrecisionStrategy,omitempty" xml:"PrecisionStrategy"`
	// Indicates whether calculate the dependent cells recursively when calculating one cell and it depends on other cells.             The default value is true.
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive"`
	// The custom formula calculation engine to extend the default calculation engine of Aspose.Cells.
	CustomEngine *AbstractCalculationEngine `json:"CustomEngine,omitempty" xml:"CustomEngine"`
	// The monitor for user to track the progress of formula calculation.
	CalculationMonitor *AbstractCalculationMonitor `json:"CalculationMonitor,omitempty" xml:"CalculationMonitor"`
	// Specifies the data sources for external links used in formulas.
	LinkedDataSources []interface{} `json:"LinkedDataSources,omitempty" xml:"LinkedDataSources"`
}
