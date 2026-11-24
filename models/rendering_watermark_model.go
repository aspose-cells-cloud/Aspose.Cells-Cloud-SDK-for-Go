/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="RenderingWatermark">
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

// RenderingWatermark
type RenderingWatermark struct {
	Rotation           *float64       `json:"Rotation,omitempty" xml:"Rotation"`
	ScaleToPagePercent *int32         `json:"ScaleToPagePercent,omitempty" xml:"ScaleToPagePercent"`
	Opacity            *float64       `json:"Opacity,omitempty" xml:"Opacity"`
	IsBackground       *bool          `json:"IsBackground,omitempty" xml:"IsBackground"`
	Text               string         `json:"Text,omitempty" xml:"Text"`
	Font               *RenderingFont `json:"Font,omitempty" xml:"Font"`
	Image              []interface{}  `json:"Image,omitempty" xml:"Image"`
	HAlignment         string         `json:"HAlignment,omitempty" xml:"HAlignment"`
	VAlignment         string         `json:"VAlignment,omitempty" xml:"VAlignment"`
	OffsetX            *float64       `json:"OffsetX,omitempty" xml:"OffsetX"`
	OffsetY            *float64       `json:"OffsetY,omitempty" xml:"OffsetY"`
}
