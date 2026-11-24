/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataFillValue">
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

// DataFillValue Represents that the data is populated with the specified value.
type DataFillValue struct {
	// Represents default bool value.
	DefaultBoolean *bool `json:"DefaultBoolean,omitempty" xml:"DefaultBoolean"`
	// Represents default string value.
	DefaultString string `json:"DefaultString,omitempty" xml:"DefaultString"`
	// Represents default number value.
	DefaultNumber *int32 `json:"DefaultNumber,omitempty" xml:"DefaultNumber"`
	// Represents default double value.
	DefaultDouble *float64 `json:"DefaultDouble,omitempty" xml:"DefaultDouble"`
	// Represents default date value.
	DefaultDate string `json:"DefaultDate,omitempty" xml:"DefaultDate"`
}
