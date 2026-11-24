/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ImageSaveOptions">
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

// ImageSaveOptions
type ImageSaveOptions struct {
	SaveOptions
	ChartImageType        string `json:"ChartImageType,omitempty" xml:"ChartImageType"`
	EmbededImageNameInSvg string `json:"EmbededImageNameInSvg,omitempty" xml:"EmbededImageNameInSvg"`
	HorizontalResolution  *int32 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
	ImageFormat           string `json:"ImageFormat,omitempty" xml:"ImageFormat"`
	IsCellAutoFit         *bool  `json:"IsCellAutoFit,omitempty" xml:"IsCellAutoFit"`
	OnePagePerSheet       *bool  `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
	OnlyArea              *bool  `json:"OnlyArea,omitempty" xml:"OnlyArea"`
	PrintingPage          string `json:"PrintingPage,omitempty" xml:"PrintingPage"`
	PrintWithStatusDialog *bool  `json:"PrintWithStatusDialog,omitempty" xml:"PrintWithStatusDialog"`
	Quality               *int32 `json:"Quality,omitempty" xml:"Quality"`
	TiffCompression       string `json:"TiffCompression,omitempty" xml:"TiffCompression"`
	VerticalResolution    *int32 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
}
