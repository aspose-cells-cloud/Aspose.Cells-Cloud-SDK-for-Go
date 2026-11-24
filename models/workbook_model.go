/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Workbook">
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

// Workbook            Represents a root object to create an Excel spreadsheet.
type Workbook struct {
	// Gets and sets the current file name.
	FileName string `json:"FileName,omitempty" xml:"FileName"`
	// A property of type List Link  named Links is specified to be serialized as an XmlElement with the tag "link".
	Links []Link `json:"Links,omitempty" xml:"Links"`
	// Gets the  collection in the spreadsheet.
	Worksheets *LinkElement `json:"Worksheets,omitempty" xml:"Worksheets"`
	// Gets or sets the default  object of the workbook.
	DefaultStyle       *LinkElement `json:"DefaultStyle,omitempty" xml:"DefaultStyle"`
	DocumentProperties *LinkElement `json:"DocumentProperties,omitempty" xml:"DocumentProperties"`
	Names              *LinkElement `json:"Names,omitempty" xml:"Names"`
	// Represents the workbook settings.
	Settings         *LinkElement `json:"Settings,omitempty" xml:"Settings"`
	IsWriteProtected string       `json:"IsWriteProtected,omitempty" xml:"IsWriteProtected"`
	IsProtected      string       `json:"IsProtected,omitempty" xml:"IsProtected"`
	IsEncryption     string       `json:"IsEncryption,omitempty" xml:"IsEncryption"`
	Password         string       `json:"Password,omitempty" xml:"Password"`
}
