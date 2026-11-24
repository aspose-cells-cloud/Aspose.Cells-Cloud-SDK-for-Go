/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AddTextOptions">
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

// AddTextOptions Class summary: The features of this class include hands-on learning activities, group projects, interactive discussions, guest speakers, and field trips to real-world applications.
type AddTextOptions struct {
	// The class has a public property named "Name" with a getter and setter method.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents data source.  There are three types of data, they are CloudFileSystem, RequestFiles, HttpUri.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents file information. Include of filename, filesize, and file content(base64String).
	FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
	// Specifies the range of cells within the worksheet where the spreadsheet operations will be performed. This parameter allows users to define the exact area to be processed, ensuring that operations are applied only to the designated cells.
	ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
	// Add text content.
	Text string `json:"Text,omitempty" xml:"Text"`
	// Represents where text should be inserted or selected in the spreadsheet.
	SelectPoistion string `json:"SelectPoistion,omitempty" xml:"SelectPoistion"`
	// Selected text of cell where text should be inserted or selected in the spreadsheet.
	SelectText string `json:"SelectText,omitempty" xml:"SelectText"`
	// Indicates whether empty cells should be skipped during processing.
	SkipEmptyCells *bool `json:"SkipEmptyCells,omitempty" xml:"SkipEmptyCells"`
}
