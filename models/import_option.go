/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="import_option.go">
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

// ImportOption Class Summary: This class covers the features of artificial intelligence, machine learning, deep learning, and natural language processing.
type ImportOption struct {
	// Destination worksheet name.
	DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
	// A boolean property `IsInsert` that can be both read and written.
	IsInsert       *bool       `json:"IsInsert,omitempty" xml:"IsInsert"`
	ImportDataType string      `json:"ImportDataType,omitempty" xml:"ImportDataType"`
	DataSource     *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	Source         *FileSource `json:"Source,omitempty" xml:"Source"`
	// SeparatorString is the column delimiter for CSV imports. The live
	// importdata endpoint requires it for ImportDataType "CSVData".
	SeparatorString string `json:"SeparatorString,omitempty" xml:"SeparatorString"`
	// ConvertNumericData converts the imported data to numeric types when
	// applicable.
	ConvertNumericData *bool `json:"ConvertNumericData,omitempty" xml:"ConvertNumericData"`
}

// importOptionMarker closes the ImportOptionLike family: only ImportOption and the models
// embedding it carry this promoted method, so no unrelated type satisfies ImportOptionLike.
func (m ImportOption) importOptionMarker() {}
