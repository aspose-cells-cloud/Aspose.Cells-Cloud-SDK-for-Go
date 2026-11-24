/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="TrimContentOptions">
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

// TrimContentOptions
type TrimContentOptions struct {
	// Represents data source.  There are three types of data, they are CloudFileSystem, RequestFiles, HttpUri.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents file information. Include of filename, filesize, and file content(base64String).
	FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
	// Trim Content
	TrimContent string `json:"TrimContent,omitempty" xml:"TrimContent"`
	// If the trim leading value is true, the trim content before and after cell values will be deleted.
	TrimLeading *bool `json:"TrimLeading,omitempty" xml:"TrimLeading"`
	// If the trim trailing value is true, the trim content before and after cell values will be deleted.
	TrimTrailing *bool `json:"TrimTrailing,omitempty" xml:"TrimTrailing"`
	// When the trim space between word to 1 parameter is true, it enables the removal of extra spaces between words within a cell, ensuring that only a single space is maintained between words.
	TrimSpaceBetweenWordTo1 *bool `json:"TrimSpaceBetweenWordTo1,omitempty" xml:"TrimSpaceBetweenWordTo1"`
	TrimNonBreakingSpaces   *bool `json:"TrimNonBreakingSpaces,omitempty" xml:"TrimNonBreakingSpaces"`
	// When this parameter is enabled (set to True), it deletes extra line breaks within the selected range, ensuring that only necessary line breaks are retained.
	RemoveExtraLineBreaks *bool `json:"RemoveExtraLineBreaks,omitempty" xml:"RemoveExtraLineBreaks"`
	// When this parameter is enabled (set to True), it removes all line breaks within the selected range, resulting in a continuous block of text without any line breaks.
	RemoveAllLineBreaks *bool `json:"RemoveAllLineBreaks,omitempty" xml:"RemoveAllLineBreaks"`
	// Specifies the range of cells within the worksheet where the spreadsheet operations will be performed. This parameter allows users to define the exact area to be processed, ensuring that operations are applied only to the designated cells.
	ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
}
