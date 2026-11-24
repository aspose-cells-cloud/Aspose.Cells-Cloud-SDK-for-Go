/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AutoFitterOptions">
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

// AutoFitterOptions Represents all auto fitter options.
type AutoFitterOptions struct {
	// Gets and set the type of auto fitting row height of merged cells.
	AutoFitMergedCellsType string `json:"AutoFitMergedCellsType,omitempty" xml:"AutoFitMergedCellsType"`
	// Ignores the hidden rows/columns.
	IgnoreHidden *bool `json:"IgnoreHidden,omitempty" xml:"IgnoreHidden"`
	// Indicates whether only fit the rows which height are not customed.
	OnlyAuto *bool `json:"OnlyAuto,omitempty" xml:"OnlyAuto"`
	// Gets or sets default edit language.
	DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
	// Gets and sets the max row height(in unit of Point) when autofitting rows.
	MaxRowHeight *float64 `json:"MaxRowHeight,omitempty" xml:"MaxRowHeight"`
	// Gets and sets the type of auto fitting wrapped text.
	AutoFitWrappedTextType string `json:"AutoFitWrappedTextType,omitempty" xml:"AutoFitWrappedTextType"`
	// Gets and sets the formatted strategy.
	FormatStrategy string `json:"FormatStrategy,omitempty" xml:"FormatStrategy"`
	// Indicates whether fit for rendering purpose.
	ForRendering *bool `json:"ForRendering,omitempty" xml:"ForRendering"`
}
