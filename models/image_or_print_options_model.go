/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ImageOrPrintOptions">
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

// ImageOrPrintOptions
type ImageOrPrintOptions struct {
	TextCrossType                     string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
	GridlineType                      string `json:"GridlineType,omitempty" xml:"GridlineType"`
	OutputBlankPageWhenNothingToPrint *bool  `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
	CheckWorkbookDefaultFont          *bool  `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
	DefaultFont                       string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
	IsOptimized                       *bool  `json:"IsOptimized,omitempty" xml:"IsOptimized"`
	PageCount                         *int32 `json:"PageCount,omitempty" xml:"PageCount"`
	PageIndex                         *int32 `json:"PageIndex,omitempty" xml:"PageIndex"`
	IsFontSubstitutionCharGranularity *bool  `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
	Transparent                       *bool  `json:"Transparent,omitempty" xml:"Transparent"`
	OnlyArea                          *bool  `json:"OnlyArea,omitempty" xml:"OnlyArea"`
	SVGFitToViewPort                  *bool  `json:"SVGFitToViewPort,omitempty" xml:"SVGFitToViewPort"`
	EmbededImageNameInSvg             string `json:"EmbededImageNameInSvg,omitempty" xml:"EmbededImageNameInSvg"`
	AllColumnsInOnePagePerSheet       *bool  `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
	PrintWithStatusDialog             *bool  `json:"PrintWithStatusDialog,omitempty" xml:"PrintWithStatusDialog"`
	HorizontalResolution              *int32 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
	VerticalResolution                *int32 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
	DefaultEditLanguage               string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
	TiffColorDepth                    string `json:"TiffColorDepth,omitempty" xml:"TiffColorDepth"`
	TiffCompression                   string `json:"TiffCompression,omitempty" xml:"TiffCompression"`
	PrintingPage                      string `json:"PrintingPage,omitempty" xml:"PrintingPage"`
	Quality                           *int32 `json:"Quality,omitempty" xml:"Quality"`
	ImageType                         string `json:"ImageType,omitempty" xml:"ImageType"`
	OnePagePerSheet                   *bool  `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
	TiffBinarizationMethod            string `json:"TiffBinarizationMethod,omitempty" xml:"TiffBinarizationMethod"`
}
