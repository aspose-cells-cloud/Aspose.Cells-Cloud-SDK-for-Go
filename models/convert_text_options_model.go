/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ConvertTextOptions">
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

// ConvertTextOptions Class summary: The features of the new smartphone include a high-resolution display, multiple camera lenses for versatile photography, a fast processor for seamless performance, and a larger battery for extended usage time.
type ConvertTextOptions struct {
	// The class has a public property called "Name" of type string that can be accessed and modified.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents data source.  There are three types of data, they are CloudFileSystem, RequestFiles, HttpUri.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents file information. Include of filename, filesize, and file content(base64String).
	FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
	// Specifies the range of cells within the worksheet where the spreadsheet operations will be performed. This parameter allows users to define the exact area to be processed, ensuring that operations are applied only to the designated cells.
	ScopeOptions     *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
	ConvertTextType  string        `json:"ConvertTextType,omitempty" xml:"ConvertTextType"`
	SourceCharacters string        `json:"SourceCharacters,omitempty" xml:"SourceCharacters"`
	TargetCharacters string        `json:"TargetCharacters,omitempty" xml:"TargetCharacters"`
}
