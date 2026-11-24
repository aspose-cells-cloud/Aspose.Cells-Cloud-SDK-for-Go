/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LoadOptions">
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

// LoadOptions Represents the options of loading the file.
type LoadOptions struct {
	// This class has a public property named "ConvertNumericData" of type string that can be read from and written to.
	ConvertNumericData string `json:"ConvertNumericData,omitempty" xml:"ConvertNumericData"`
	// Gets and sets the interrupt monitor.
	InterruptMonitor string `json:"InterruptMonitor,omitempty" xml:"InterruptMonitor"`
	// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
	LanguageCode    string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
	LoadDataOptions string `json:"LoadDataOptions,omitempty" xml:"LoadDataOptions"`
	// Gets the load format.
	LoadFormat                 string `json:"LoadFormat,omitempty" xml:"LoadFormat"`
	OnlyLoadDocumentProperties string `json:"OnlyLoadDocumentProperties,omitempty" xml:"OnlyLoadDocumentProperties"`
	// Indicates whether parsing the formula when reading the file.
	ParsingFormulaOnOpen string `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
	// Gets and set the password of the workbook.
	Password string `json:"Password,omitempty" xml:"Password"`
	// Gets or sets the system regional settings based on CountryCode at the time the file was loaded.
	Region string `json:"Region,omitempty" xml:"Region"`
	// Sets the default standard font name
	StandardFont string `json:"StandardFont,omitempty" xml:"StandardFont"`
	// Sets the default standard font size.
	StandardFontSize *float64 `json:"StandardFontSize,omitempty" xml:"StandardFontSize"`
}
