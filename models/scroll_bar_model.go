/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ScrollBar">
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

// ScrollBar Represents a scroll bar object.
type ScrollBar struct {
	Shape
	// Gets or sets the current value.
	CurrentValue *int32 `json:"CurrentValue,omitempty" xml:"CurrentValue"`
	// Gets or sets the amount that the scroll bar or spinner is incremented a line scroll.
	IncrementalChange *int32 `json:"IncrementalChange,omitempty" xml:"IncrementalChange"`
	// Indicates whether this is a horizontal scroll bar.
	IsHorizontal *bool `json:"IsHorizontal,omitempty" xml:"IsHorizontal"`
	// Gets or sets the maximum value of a scroll bar or spinner range.
	Max *int32 `json:"Max,omitempty" xml:"Max"`
	// Gets or sets the minimum value of a scroll bar or spinner range.
	Min *int32 `json:"Min,omitempty" xml:"Min"`
	// Gets or sets page change
	PageChange *int32 `json:"PageChange,omitempty" xml:"PageChange"`
	// Indicates whether the shape has 3-D shading.
	Shadow *bool `json:"Shadow,omitempty" xml:"Shadow"`
	Link   *Link `json:"link,omitempty" xml:"link"`
}
