/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataCleansingRequest">
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

// DataCleansingRequest A data cleansing request for an Spreadsheet file.
type DataCleansingRequest struct {
	// Spreadsheet files that require data cleaning.
	File *FileInfo `json:"File,omitempty" xml:"File"`
	// Whether check restriction of Spreadsheet file when user modify cells related objects.
	CheckExcelRestriction *bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
	// The regional settings for workbook.
	Region string `json:"Region,omitempty" xml:"Region"`
	// finish to data cleansing, outfile`s file format.
	OutFileFormat string `json:"OutFileFormat,omitempty" xml:"OutFileFormat"`
	// Data cleansing content
	DataCleansing *DataCleansing `json:"DataCleansing,omitempty" xml:"DataCleansing"`
}
