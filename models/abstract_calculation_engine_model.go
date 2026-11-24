/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AbstractCalculationEngine">
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

// AbstractCalculationEngine Represents user's custom calculation engine to extend the default calculation engine of Aspose.Cells.
type AbstractCalculationEngine struct {
	// Indicates whether this engine needs the literal text of parameter while doing calculation. Default value is false.
	IsParamLiteralRequired *bool `json:"IsParamLiteralRequired,omitempty" xml:"IsParamLiteralRequired"`
	// Indicates whether this engine needs the parameter to be calculated in array mode. Default value is false.             If  is required when calculating custom             functions, this property needs to be set as true.
	IsParamArrayModeRequired *bool `json:"IsParamArrayModeRequired,omitempty" xml:"IsParamArrayModeRequired"`
	// Whether built-in functions that have been supported by the built-in engine             should be checked and processed by this implementation.             Default is false.             If user needs to change the calculation logic of some built-in functions, this property should be set as true.             Otherwise please leave this property as false for performance consideration.
	ProcessBuiltInFunctions *bool `json:"ProcessBuiltInFunctions,omitempty" xml:"ProcessBuiltInFunctions"`
}
