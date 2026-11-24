/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="SplitTextOptions">
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

// SplitTextOptions Appliance features: 1. Automatic defrost system 2. Energy-efficient LED lighting 3. Adjustable glass shelves 4. Ice and water dispenser with filtration system
type SplitTextOptions struct {
	// The property "Name" is a publicly accessible and overridable property of type string in the class.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Represents data source.  There are three types of data, they are CloudFileSystem, RequestFiles, HttpUri.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents file information. Include of filename, filesize, and file content(base64String).
	FileInfo                       *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
	Worksheet                      string    `json:"Worksheet,omitempty" xml:"Worksheet"`
	Range                          string    `json:"Range,omitempty" xml:"Range"`
	SplitDelimitersType            string    `json:"SplitDelimitersType,omitempty" xml:"SplitDelimitersType"`
	CustomDelimiter                string    `json:"CustomDelimiter,omitempty" xml:"CustomDelimiter"`
	KeepDelimitersInResultingCells *bool     `json:"KeepDelimitersInResultingCells,omitempty" xml:"KeepDelimitersInResultingCells"`
	KeepDelimitersPosition         string    `json:"KeepDelimitersPosition,omitempty" xml:"KeepDelimitersPosition"`
	HowToSplit                     string    `json:"HowToSplit,omitempty" xml:"HowToSplit"`
}
