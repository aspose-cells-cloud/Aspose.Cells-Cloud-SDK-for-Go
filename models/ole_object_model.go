/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="OleObject">
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

// OleObject Represents an OleObject in a worksheet.
type OleObject struct {
	Shape
	// True if the specified object is displayed as an icon                          and the image will not be auto changed.
	DisplayAsIcon *bool `json:"DisplayAsIcon,omitempty" xml:"DisplayAsIcon"`
	// Gets and sets the file type of the embedded ole object data
	FileFormatType string `json:"FileFormatType,omitempty" xml:"FileFormatType"`
	// Gets or sets the path and name of the source file for the linked image.
	ImageSourceFullName string `json:"ImageSourceFullName,omitempty" xml:"ImageSourceFullName"`
	// True indicates that the size of the ole object will be auto changed as the size of snapshot of the embedded content                         when the ole object is activated.
	IsAutoSize *bool `json:"IsAutoSize,omitempty" xml:"IsAutoSize"`
	// Returns true if the OleObject links to the file.
	IsLink *bool `json:"IsLink,omitempty" xml:"IsLink"`
	// Gets or sets the ProgID of the OLE object.
	ProgID string `json:"ProgID,omitempty" xml:"ProgID"`
	// Returns the source full name of the source file for the linked OLE object.
	SourceFullName string `json:"SourceFullName,omitempty" xml:"SourceFullName"`
	Link           *Link  `json:"link,omitempty" xml:"link"`
}
