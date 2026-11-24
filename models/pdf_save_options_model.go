/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PdfSaveOptions">
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

// PdfSaveOptions
type PdfSaveOptions struct {
	SaveOptions
	DisplayDocTitle         *bool               `json:"DisplayDocTitle,omitempty" xml:"DisplayDocTitle"`
	ExportDocumentStructure *bool               `json:"ExportDocumentStructure,omitempty" xml:"ExportDocumentStructure"`
	EmfRenderSetting        string              `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
	CustomPropertiesExport  string              `json:"CustomPropertiesExport,omitempty" xml:"CustomPropertiesExport"`
	OptimizationType        string              `json:"OptimizationType,omitempty" xml:"OptimizationType"`
	Producer                string              `json:"Producer,omitempty" xml:"Producer"`
	PdfCompression          string              `json:"PdfCompression,omitempty" xml:"PdfCompression"`
	FontEncoding            string              `json:"FontEncoding,omitempty" xml:"FontEncoding"`
	Watermark               *RenderingWatermark `json:"Watermark,omitempty" xml:"Watermark"`
	CalculateFormula        *bool               `json:"CalculateFormula,omitempty" xml:"CalculateFormula"`
	CheckFontCompatibility  *bool               `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
	Compliance              string              `json:"Compliance,omitempty" xml:"Compliance"`
	DefaultFont             string              `json:"DefaultFont,omitempty" xml:"DefaultFont"`
	OnePagePerSheet         *bool               `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
	PrintingPageType        string              `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
	SecurityOptions         *PdfSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions"`
	DesiredPPI              *int32              `json:"desiredPPI,omitempty" xml:"desiredPPI"`
	JpegQuality             *int32              `json:"jpegQuality,omitempty" xml:"jpegQuality"`
	ImageType               string              `json:"ImageType,omitempty" xml:"ImageType"`
}
