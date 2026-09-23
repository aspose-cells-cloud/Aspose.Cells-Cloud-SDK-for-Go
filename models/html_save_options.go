/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="html_save_options.go">
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

// HtmlSaveOptions
type HtmlSaveOptions struct {
	SaveOptions
	ExportPageHeaders               *bool                `json:"ExportPageHeaders,omitempty" xml:"ExportPageHeaders"`
	ExportPageFooters               *bool                `json:"ExportPageFooters,omitempty" xml:"ExportPageFooters"`
	ExportRowColumnHeadings         *bool                `json:"ExportRowColumnHeadings,omitempty" xml:"ExportRowColumnHeadings"`
	ShowAllSheets                   *bool                `json:"ShowAllSheets,omitempty" xml:"ShowAllSheets"`
	ImageOptions                    *ImageOrPrintOptions `json:"ImageOptions,omitempty" xml:"ImageOptions"`
	SaveAsSingleFile                *bool                `json:"SaveAsSingleFile,omitempty" xml:"SaveAsSingleFile"`
	ExportHiddenWorksheet           *bool                `json:"ExportHiddenWorksheet,omitempty" xml:"ExportHiddenWorksheet"`
	ExportGridLines                 *bool                `json:"ExportGridLines,omitempty" xml:"ExportGridLines"`
	PresentationPreference          *bool                `json:"PresentationPreference,omitempty" xml:"PresentationPreference"`
	CellCssPrefix                   string               `json:"CellCssPrefix,omitempty" xml:"CellCssPrefix"`
	TableCssId                      string               `json:"TableCssId,omitempty" xml:"TableCssId"`
	IsFullPathLink                  *bool                `json:"IsFullPathLink,omitempty" xml:"IsFullPathLink"`
	ExportWorksheetCSSSeparately    *bool                `json:"ExportWorksheetCSSSeparately,omitempty" xml:"ExportWorksheetCSSSeparately"`
	ExportSimilarBorderStyle        *bool                `json:"ExportSimilarBorderStyle,omitempty" xml:"ExportSimilarBorderStyle"`
	MergeEmptyTdForcely             *bool                `json:"MergeEmptyTdForcely,omitempty" xml:"MergeEmptyTdForcely"`
	ExportCellCoordinate            *bool                `json:"ExportCellCoordinate,omitempty" xml:"ExportCellCoordinate"`
	ExportExtraHeadings             *bool                `json:"ExportExtraHeadings,omitempty" xml:"ExportExtraHeadings"`
	ExportHeadings                  *bool                `json:"ExportHeadings,omitempty" xml:"ExportHeadings"`
	ExportFormula                   *bool                `json:"ExportFormula,omitempty" xml:"ExportFormula"`
	AddTooltipText                  *bool                `json:"AddTooltipText,omitempty" xml:"AddTooltipText"`
	ExportBogusRowData              *bool                `json:"ExportBogusRowData,omitempty" xml:"ExportBogusRowData"`
	ExcludeUnusedStyles             *bool                `json:"ExcludeUnusedStyles,omitempty" xml:"ExcludeUnusedStyles"`
	ExportDocumentProperties        *bool                `json:"ExportDocumentProperties,omitempty" xml:"ExportDocumentProperties"`
	ExportWorksheetProperties       *bool                `json:"ExportWorksheetProperties,omitempty" xml:"ExportWorksheetProperties"`
	ExportWorkbookProperties        *bool                `json:"ExportWorkbookProperties,omitempty" xml:"ExportWorkbookProperties"`
	ExportFrameScriptsAndProperties *bool                `json:"ExportFrameScriptsAndProperties,omitempty" xml:"ExportFrameScriptsAndProperties"`
	AttachedFilesDirectory          string               `json:"AttachedFilesDirectory,omitempty" xml:"AttachedFilesDirectory"`
	AttachedFilesUrlPrefix          string               `json:"AttachedFilesUrlPrefix,omitempty" xml:"AttachedFilesUrlPrefix"`
	Encoding                        string               `json:"Encoding,omitempty" xml:"Encoding"`
	ExportActiveWorksheetOnly       *bool                `json:"ExportActiveWorksheetOnly,omitempty" xml:"ExportActiveWorksheetOnly"`
	ExportChartImageFormat          string               `json:"ExportChartImageFormat,omitempty" xml:"ExportChartImageFormat"`
	ExportImagesAsBase64            *bool                `json:"ExportImagesAsBase64,omitempty" xml:"ExportImagesAsBase64"`
	HiddenColDisplayType            string               `json:"HiddenColDisplayType,omitempty" xml:"HiddenColDisplayType"`
	HiddenRowDisplayType            string               `json:"HiddenRowDisplayType,omitempty" xml:"HiddenRowDisplayType"`
	HtmlCrossStringType             string               `json:"HtmlCrossStringType,omitempty" xml:"HtmlCrossStringType"`
	IsExpImageToTempDir             *bool                `json:"IsExpImageToTempDir,omitempty" xml:"IsExpImageToTempDir"`
	PageTitle                       string               `json:"PageTitle,omitempty" xml:"PageTitle"`
	ParseHtmlTagInCell              *bool                `json:"ParseHtmlTagInCell,omitempty" xml:"ParseHtmlTagInCell"`
	CellNameAttribute               string               `json:"CellNameAttribute,omitempty" xml:"CellNameAttribute"`
}

// htmlSaveOptionsMarker closes the HtmlSaveOptionsLike family: only HtmlSaveOptions and the
// models embedding it carry this promoted method, so no unrelated type satisfies
// HtmlSaveOptionsLike.
func (m HtmlSaveOptions) htmlSaveOptionsMarker() {}
