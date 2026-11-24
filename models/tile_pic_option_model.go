/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="TilePicOption">
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

// TilePicOption            Represents tile picture as texture.
type TilePicOption struct {
	// Gets or sets the X offset for tiling picture.
	OffsetX *float64 `json:"OffsetX,omitempty" xml:"OffsetX"`
	// Gets or sets the Y offset for tiling picture.
	OffsetY *float64 `json:"OffsetY,omitempty" xml:"OffsetY"`
	// Gets or sets the X scale for tiling picture.
	ScaleX *float64 `json:"ScaleX,omitempty" xml:"ScaleX"`
	// Gets or sets the Y scale for tiling picture.
	ScaleY *float64 `json:"ScaleY,omitempty" xml:"ScaleY"`
	// Gets or sets the alignment for tiling.
	AlignmentType string `json:"AlignmentType,omitempty" xml:"AlignmentType"`
	// Gets or sets the mirror type for tiling.
	MirrorType string `json:"MirrorType,omitempty" xml:"MirrorType"`
}
