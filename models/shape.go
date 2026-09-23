/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="shape.go">
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

// Shape Represents the msodrawing object.
type Shape struct {
	LinkElement
	// Gets and sets the name of the shape.
	Name string `json:"Name,omitempty" xml:"Name"`
	// Gets mso drawing type.
	MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
	// Gets and sets the auto shape type.
	AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
	// Represents the way the drawing object is attached to the cells below it.                         The property controls the placement of an object on a worksheet.
	Placement string `json:"Placement,omitempty" xml:"Placement"`
	// Represents upper left corner row index.
	UpperLeftRow *int32 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
	// Represents the vertical offset of shape from its top row, in unit of pixels.
	Top *int32 `json:"Top,omitempty" xml:"Top"`
	// Represents upper left corner column index.
	UpperLeftColumn *int32 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
	// Represents the horizontal offset of shape from its left column, in unit of pixels.
	Left *int32 `json:"Left,omitempty" xml:"Left"`
	// Represents lower right corner row index.
	LowerRightRow *int32 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
	// Represents the width of the shape's vertical offset from its lower bottom corner row, in unit of pixels.
	Bottom *int32 `json:"Bottom,omitempty" xml:"Bottom"`
	// Represents lower right corner column index.
	LowerRightColumn *int32 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
	// Represents the width of the shape's horizontal  offset from its lower right corner column, in unit of pixels.
	Right *int32 `json:"Right,omitempty" xml:"Right"`
	// Represents the width of shape, in unit of pixels.
	Width *int32 `json:"Width,omitempty" xml:"Width"`
	// Represents the height of shape, in unit of pixel.
	Height *int32 `json:"Height,omitempty" xml:"Height"`
	// Gets and sets the horizontal offset of shape from worksheet left border,in unit of pixels.
	X *int32 `json:"X,omitempty" xml:"X"`
	// Gets and sets the vertical offset of shape from worksheet top border,in unit of pixels.
	Y *int32 `json:"Y,omitempty" xml:"Y"`
	// Gets and sets the rotation of the shape.
	RotationAngle *float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	// Gets and sets the html string which contains data and some formats in this textbox.
	HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
	// Represents the string in this TextBox object.
	Text string `json:"Text,omitempty" xml:"Text"`
	// Returns or sets the descriptive (alternative) text string of the  object.
	AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
	// Gets and sets the text horizontal alignment type of the shape.
	TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
	// Gets and sets the text horizontal overflow type of the shape which contains text.
	TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
	// Gets and sets the text orientation type of the shape.
	TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
	// Gets and sets the text vertical alignment type of the shape.
	TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
	// Gets and sets the text vertical overflow type of the shape which contains text.
	TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
	// Indicates whether the shape is a group.
	IsGroup *bool `json:"IsGroup,omitempty" xml:"IsGroup"`
	// Indicates whether the object is visible.
	IsHidden *bool `json:"IsHidden,omitempty" xml:"IsHidden"`
	// True means that don't allow changes in aspect ratio.
	IsLockAspectRatio *bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
	// True if the object is locked, False if the object can be modified when the sheet is protected.
	IsLocked *bool `json:"IsLocked,omitempty" xml:"IsLocked"`
	// True if the object is printable
	IsPrintable *bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
	// Gets and sets the text wrapped type of the shape which contains text.
	IsTextWrapped *bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
	// Indicates whether this shape is a word art.
	IsWordArt *bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
	// Gets or sets the worksheet range linked to the control's value.
	LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
	// Returns the position of a shape in the z-order.
	ZOrderPosition *int32 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
	// Represents the font of shape.
	Font *Font `json:"Font,omitempty" xml:"Font"`
	// Gets the hyperlink of the shape.
	Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
}

// shapeMarker closes the ShapeLike family: only Shape and the models embedding it carry
// this promoted method, so no unrelated type satisfies ShapeLike.
func (m Shape) shapeMarker() {}
