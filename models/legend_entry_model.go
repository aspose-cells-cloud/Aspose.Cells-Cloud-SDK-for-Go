/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="LegendEntry">
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

// LegendEntry Represents a legend entry in a chart legend.
type LegendEntry struct {
	LinkElement
	// True if the text in the object changes font size when the object size changes.                          The default value is True.
	AutoScaleFont *bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
	// Gets and sets the display mode of the background
	BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
	// Gets a  object of the specified ChartFrame object.
	Font *Font `json:"Font,omitempty" xml:"Font"`
	// Gets and sets whether the legend entry is deleted.
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted"`
}
