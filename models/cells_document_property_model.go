/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="CellsDocumentProperty">
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

// CellsDocumentProperty Cells document property.
type CellsDocumentProperty struct {
	// Returns the name of the property.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets or sets the value of the property.
	Value string `json:"Value,omitempty" xml:"Value"`
	// Indicates whether this property is linked to content
	IsLinkedToContent string `json:"IsLinkedToContent,omitempty" xml:"IsLinkedToContent"`
	// The linked content source.
	Source string `json:"Source,omitempty" xml:"Source"`
	// Gets the data type of the property.
	Type string `json:"Type,omitempty" xml:"Type"`
	// Returns true if this property does not have a name in the OLE2 storage and a   unique name was generated only for the public API.
	IsGeneratedName string `json:"IsGeneratedName,omitempty" xml:"IsGeneratedName"`
}
