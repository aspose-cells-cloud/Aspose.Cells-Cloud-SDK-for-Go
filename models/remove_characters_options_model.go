/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="RemoveCharactersOptions">
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

// RemoveCharactersOptions Class summary: The class features discussing effective communication strategies, developing problem-solving skills, and increasing self-awareness for personal growth.
type RemoveCharactersOptions struct {
	// Name property with a getter and setter override that returns a string.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents data source.  There are three types of data, they are CloudFileSystem, RequestFiles, HttpUri.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents file information. Include of filename, filesize, and file content(base64String).
	FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
	// Specifies the range of cells within the worksheet where the spreadsheet operations will be performed. This parameter allows users to define the exact area to be processed, ensuring that operations are applied only to the designated cells.
	ScopeOptions                *ScopeOptions                `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
	RemoveCharactersByCharacter *RemoveCharactersByCharacter `json:"RemoveCharactersByCharacter,omitempty" xml:"RemoveCharactersByCharacter"`
	RemoveCharactersByPosition  *RemoveCharactersByPosition  `json:"RemoveCharactersByPosition,omitempty" xml:"RemoveCharactersByPosition"`
}
