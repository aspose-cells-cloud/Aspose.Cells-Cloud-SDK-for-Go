/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Name">
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

// Name Represents a defined name for a range of cells.
type Name struct {
	LinkElement
	// Gets and sets the comment of the name.                         Only applies for Excel 2007.
	Comment string `json:"Comment,omitempty" xml:"Comment"`
	// Property Summary: WorksheetIndex is an optional integer property marked with the XmlElement attribute "worksheetindex".
	WorksheetIndex *int32 `json:"WorksheetIndex,omitempty" xml:"WorksheetIndex"`
	// Indicates whether this name is referred by other formulas.
	IsReferred *bool `json:"IsReferred,omitempty" xml:"IsReferred"`
	// Indicates whether the name is visible.
	IsVisible *bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// Gets or sets a R1C1 reference of the .
	R1C1RefersTo string `json:"R1C1RefersTo,omitempty" xml:"R1C1RefersTo"`
	// Returns or sets the formula that the name is defined to refer to, beginning with an equal sign.
	RefersTo string `json:"RefersTo,omitempty" xml:"RefersTo"`
	// Gets the name text of the object.
	Text string `json:"Text,omitempty" xml:"Text"`
}
