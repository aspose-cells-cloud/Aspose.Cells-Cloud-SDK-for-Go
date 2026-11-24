/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="FormulaSettings">
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

// FormulaSettings Settings of formulas and calculation.
type FormulaSettings struct {
	// Indicates whether the application is required to perform a full calculation when the workbook is opened.
	CalculateOnOpen *bool `json:"CalculateOnOpen,omitempty" xml:"CalculateOnOpen"`
	// Indicates whether recalculate the workbook before saving the document, when in manual calculation mode.
	CalculateOnSave *bool `json:"CalculateOnSave,omitempty" xml:"CalculateOnSave"`
	// Indicates whether calculates all formulas every time when a calculation is triggered.
	ForceFullCalculation *bool `json:"ForceFullCalculation,omitempty" xml:"ForceFullCalculation"`
	// Gets or sets the mode for workbook calculation in ms excel.
	CalculationMode string `json:"CalculationMode,omitempty" xml:"CalculationMode"`
	// Specifies the version of the calculation engine used to calculate values in the workbook.
	CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
	// Indicates whether enable iterative calculation to resolve circular references.
	EnableIterativeCalculation *bool `json:"EnableIterativeCalculation,omitempty" xml:"EnableIterativeCalculation"`
	// The maximum iterations to resolve a circular reference.
	MaxIteration *int32 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
	// The maximum change to resolve a circular reference.
	MaxChange *float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
	// Whether the precision of calculated result be set as they are displayed while calculating formulas
	PrecisionAsDisplayed *bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
	// Whether enable calculation chain for formulas. Default is false.
	EnableCalculationChain *bool `json:"EnableCalculationChain,omitempty" xml:"EnableCalculationChain"`
	// Indicates whether preserve those spaces and line breaks that are padded between formula tokens             while getting and setting formulas.             Default value is false.
	PreservePaddingSpaces *bool `json:"PreservePaddingSpaces,omitempty" xml:"PreservePaddingSpaces"`
}
