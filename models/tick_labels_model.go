/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="TickLabels">
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

// TickLabels Represents the tick-mark labels associated with tick marks on a chart axis.
type TickLabels struct {
	LinkElement
	// True if the text in the object changes font size when the object size changes. The default value is True.
	AutoScaleFont *bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
	// Gets and sets the display mode of the background
	BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
	// Returns a  object that represents the font of the specified TickLabels object.
	Font *Font `json:"Font,omitempty" xml:"Font"`
	// Represents the format number for the TickLabels object.
	Number *int32 `json:"Number,omitempty" xml:"Number"`
	// Represents the format string for the TickLabels object.
	NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
	// True if the number format is linked to the cells                          (so that the number format changes in the labels when it changes in the cells).
	NumberFormatLinked *bool `json:"NumberFormatLinked,omitempty" xml:"NumberFormatLinked"`
	// Gets and sets the distance of labels from the axis.
	Offset *int32 `json:"Offset,omitempty" xml:"Offset"`
	// Represents text rotation angle in clockwise.
	RotationAngle *int32 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
	// Represents text reading order.
	TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
	// Represents text reading order.
	ReadingOrder string `json:"ReadingOrder,omitempty" xml:"ReadingOrder"`
	// Gets and sets the direction of text.
	DirectionType string `json:"DirectionType,omitempty" xml:"DirectionType"`
}
