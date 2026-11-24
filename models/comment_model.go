/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Comment">
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

// Comment Encapsulates the object that represents a cell comment.
type Comment struct {
	LinkElement
	// Property: CellName attribute with XmlElement tag in the class.
	CellName string `json:"CellName,omitempty" xml:"CellName"`
	// Gets and sets Name of the original comment author
	Author string `json:"Author,omitempty" xml:"Author"`
	// Gets and sets the html string which contains data and some formats in this comment.
	HtmlNote string `json:"HtmlNote,omitempty" xml:"HtmlNote"`
	// Represents the content of comment.
	Note string `json:"Note,omitempty" xml:"Note"`
	// Indicates if size of comment is adjusted automatically according to its content.
	AutoSize *bool `json:"AutoSize,omitempty" xml:"AutoSize"`
	// Represents if the comment is visible or not.
	IsVisible *bool `json:"IsVisible,omitempty" xml:"IsVisible"`
	// Represents the width of the comment, in unit of pixels.
	Width *int32 `json:"Width,omitempty" xml:"Width"`
	// Represents the Height of the comment, in unit of pixels.
	Height *int32 `json:"Height,omitempty" xml:"Height"`
	// Gets and sets the text horizontal alignment type of the comment.
	TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
	// Gets and sets the text orientation type of the comment.
	TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
	// Gets and sets the text vertical alignment type of the comment.
	TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}
