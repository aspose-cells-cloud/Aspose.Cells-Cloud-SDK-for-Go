/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="model_cells">
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

package asposecellscloud
import "time"

type DiscUsage struct {
 
    UsedSize int64 `json:"UsedSize,omitempty" xml:"UsedSize"`
    TotalSize int64 `json:"TotalSize,omitempty" xml:"TotalSize"`
}


type ObjectExist struct {
 
    Exists bool `json:"Exists,omitempty" xml:"Exists"`
    IsFolder bool `json:"IsFolder,omitempty" xml:"IsFolder"`
}


type ObjectExistsExtensions interface {
}


type FileVersion struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        IsFolder bool `json:"IsFolder,omitempty" xml:"IsFolder"`
        ModifiedDate  time.Time `json:"ModifiedDate,omitempty" xml:"ModifiedDate"`
        Size int64 `json:"Size,omitempty" xml:"Size"`
        Path string `json:"Path,omitempty" xml:"Path"`
 
    VersionId string `json:"VersionId,omitempty" xml:"VersionId"`
    IsLatest bool `json:"IsLatest,omitempty" xml:"IsLatest"`
}


type StorageExist struct {
 
    Exists bool `json:"Exists,omitempty" xml:"Exists"`
}


type FileVersions struct {
 
    Value []FileVersion `json:"Value,omitempty" xml:"Value"`
}


type FilesList struct {
 
    Value []StorageFile `json:"Value,omitempty" xml:"Value"`
}


type FilesUploadResult struct {
 
    Uploaded []string `json:"Uploaded,omitempty" xml:"Uploaded"`
    Errors []Error `json:"Errors,omitempty" xml:"Errors"`
}


type StorageFile struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    IsFolder bool `json:"IsFolder,omitempty" xml:"IsFolder"`
    ModifiedDate  time.Time `json:"ModifiedDate,omitempty" xml:"ModifiedDate"`
    Size int64 `json:"Size,omitempty" xml:"Size"`
    Path string `json:"Path,omitempty" xml:"Path"`
}


type GoogleDriveStorageFile struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        IsFolder bool `json:"IsFolder,omitempty" xml:"IsFolder"`
        ModifiedDate  time.Time `json:"ModifiedDate,omitempty" xml:"ModifiedDate"`
        Size int64 `json:"Size,omitempty" xml:"Size"`
        Path string `json:"Path,omitempty" xml:"Path"`
 
    MimeType string `json:"MimeType,omitempty" xml:"MimeType"`
}


type AggregateResultByColor struct {
 
    AggregateOperation string `json:"AggregateOperation,omitempty" xml:"AggregateOperation"`
    ColorName string `json:"ColorName,omitempty" xml:"ColorName"`
    Count int64 `json:"Count,omitempty" xml:"Count"`
    Sum float64 `json:"Sum,omitempty" xml:"Sum"`
    MaxValue float64 `json:"MaxValue,omitempty" xml:"MaxValue"`
    MinValue float64 `json:"MinValue,omitempty" xml:"MinValue"`
    AverageValue float64 `json:"AverageValue,omitempty" xml:"AverageValue"`
}


type BrokenLink struct {
 
    Filename string `json:"Filename,omitempty" xml:"Filename"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Position string `json:"Position,omitempty" xml:"Position"`
    LinkAddress string `json:"LinkAddress,omitempty" xml:"LinkAddress"`
}


type CellArea struct {
 
    EndColumn int64 `json:"EndColumn,omitempty" xml:"EndColumn"`
    EndRow int64 `json:"EndRow,omitempty" xml:"EndRow"`
    StartColumn int64 `json:"StartColumn,omitempty" xml:"StartColumn"`
    StartRow int64 `json:"StartRow,omitempty" xml:"StartRow"`
}


type CellsCloudFileInfo struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Size int64 `json:"Size,omitempty" xml:"Size"`
    Folder string `json:"Folder,omitempty" xml:"Folder"`
    Storage string `json:"Storage,omitempty" xml:"Storage"`
}


type CellsCloudPublicKey struct {
 
    Exponent string `json:"Exponent,omitempty" xml:"Exponent"`
    Modulus string `json:"Modulus,omitempty" xml:"Modulus"`
}


type Color struct {
 
    A int64 `json:"A,omitempty" xml:"A"`
    R int64 `json:"R,omitempty" xml:"R"`
    G int64 `json:"G,omitempty" xml:"G"`
    B int64 `json:"B,omitempty" xml:"B"`
}


type PdfSecurityOptions struct {
 
    AnnotationsPermission bool `json:"AnnotationsPermission,omitempty" xml:"AnnotationsPermission"`
    AssembleDocumentPermission bool `json:"AssembleDocumentPermission,omitempty" xml:"AssembleDocumentPermission"`
    ExtractContentPermission bool `json:"ExtractContentPermission,omitempty" xml:"ExtractContentPermission"`
    FillFormsPermission bool `json:"FillFormsPermission,omitempty" xml:"FillFormsPermission"`
    FullQualityPrintPermission bool `json:"FullQualityPrintPermission,omitempty" xml:"FullQualityPrintPermission"`
    ModifyDocumentPermission bool `json:"ModifyDocumentPermission,omitempty" xml:"ModifyDocumentPermission"`
    OwnerPassword string `json:"OwnerPassword,omitempty" xml:"OwnerPassword"`
    PrintPermission bool `json:"PrintPermission,omitempty" xml:"PrintPermission"`
    UserPassword string `json:"UserPassword,omitempty" xml:"UserPassword"`
}


type Range struct {
 
    ColumnCount int64 `json:"ColumnCount,omitempty" xml:"ColumnCount"`
    ColumnWidth float64 `json:"ColumnWidth,omitempty" xml:"ColumnWidth"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    Name string `json:"Name,omitempty" xml:"Name"`
    RefersTo string `json:"RefersTo,omitempty" xml:"RefersTo"`
    RowCount int64 `json:"RowCount,omitempty" xml:"RowCount"`
    RowHeight float64 `json:"RowHeight,omitempty" xml:"RowHeight"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
}


type RemoveCharactersByPosition struct {
 
    TheFirstNCharacters int64 `json:"TheFirstNCharacters,omitempty" xml:"TheFirstNCharacters"`
    TheLastNCharacters int64 `json:"TheLastNCharacters,omitempty" xml:"TheLastNCharacters"`
    AllCharactersBeforeText string `json:"AllCharactersBeforeText,omitempty" xml:"AllCharactersBeforeText"`
    AllCharactersAfterText string `json:"AllCharactersAfterText,omitempty" xml:"AllCharactersAfterText"`
}


type SaveResult struct {
 
    Documents []CellsCloudFileInfo `json:"Documents,omitempty" xml:"Documents"`
}


type PaginatedSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
    CheckWorkbookDefaultFont bool `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
    CheckFontCompatibility bool `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
    IsFontSubstitutionCharGranularity bool `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
    OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
    AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
    IgnoreError bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
    OutputBlankPageWhenNothingToPrint bool `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
    PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
    PageCount int64 `json:"PageCount,omitempty" xml:"PageCount"`
    PrintingPageType string `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
    GridlineType string `json:"GridlineType,omitempty" xml:"GridlineType"`
    TextCrossType string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
    DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
    EmfRenderSetting string `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
}


type SpreadsheetTemplate struct {
 
}


type TextItem struct {
 
    Filename string `json:"Filename,omitempty" xml:"Filename"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Position string `json:"Position,omitempty" xml:"Position"`
    Content string `json:"Content,omitempty" xml:"Content"`
}


type DbfSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportAsString bool `json:"ExportAsString,omitempty" xml:"ExportAsString"`
}


type DifSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
}


type DocxSaveOptions struct {
     
        DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
        CheckWorkbookDefaultFont bool `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
        CheckFontCompatibility bool `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
        IsFontSubstitutionCharGranularity bool `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
        OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
        AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
        IgnoreError bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
        OutputBlankPageWhenNothingToPrint bool `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
        PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
        PageCount int64 `json:"PageCount,omitempty" xml:"PageCount"`
        PrintingPageType string `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
        GridlineType string `json:"GridlineType,omitempty" xml:"GridlineType"`
        TextCrossType string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
        DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
        EmfRenderSetting string `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
 
}


type EbookSaveOptions struct {
     
        ExportPageHeaders bool `json:"ExportPageHeaders,omitempty" xml:"ExportPageHeaders"`
        ExportPageFooters bool `json:"ExportPageFooters,omitempty" xml:"ExportPageFooters"`
        ExportRowColumnHeadings bool `json:"ExportRowColumnHeadings,omitempty" xml:"ExportRowColumnHeadings"`
        ShowAllSheets bool `json:"ShowAllSheets,omitempty" xml:"ShowAllSheets"`
        ImageOptions *ImageOrPrintOptions `json:"ImageOptions,omitempty" xml:"ImageOptions"`
        SaveAsSingleFile bool `json:"SaveAsSingleFile,omitempty" xml:"SaveAsSingleFile"`
        ExportHiddenWorksheet bool `json:"ExportHiddenWorksheet,omitempty" xml:"ExportHiddenWorksheet"`
        ExportGridLines bool `json:"ExportGridLines,omitempty" xml:"ExportGridLines"`
        PresentationPreference bool `json:"PresentationPreference,omitempty" xml:"PresentationPreference"`
        CellCssPrefix string `json:"CellCssPrefix,omitempty" xml:"CellCssPrefix"`
        TableCssId string `json:"TableCssId,omitempty" xml:"TableCssId"`
        IsFullPathLink bool `json:"IsFullPathLink,omitempty" xml:"IsFullPathLink"`
        ExportWorksheetCSSSeparately bool `json:"ExportWorksheetCSSSeparately,omitempty" xml:"ExportWorksheetCSSSeparately"`
        ExportSimilarBorderStyle bool `json:"ExportSimilarBorderStyle,omitempty" xml:"ExportSimilarBorderStyle"`
        MergeEmptyTdForcely bool `json:"MergeEmptyTdForcely,omitempty" xml:"MergeEmptyTdForcely"`
        ExportCellCoordinate bool `json:"ExportCellCoordinate,omitempty" xml:"ExportCellCoordinate"`
        ExportExtraHeadings bool `json:"ExportExtraHeadings,omitempty" xml:"ExportExtraHeadings"`
        ExportHeadings bool `json:"ExportHeadings,omitempty" xml:"ExportHeadings"`
        ExportFormula bool `json:"ExportFormula,omitempty" xml:"ExportFormula"`
        AddTooltipText bool `json:"AddTooltipText,omitempty" xml:"AddTooltipText"`
        ExportBogusRowData bool `json:"ExportBogusRowData,omitempty" xml:"ExportBogusRowData"`
        ExcludeUnusedStyles bool `json:"ExcludeUnusedStyles,omitempty" xml:"ExcludeUnusedStyles"`
        ExportDocumentProperties bool `json:"ExportDocumentProperties,omitempty" xml:"ExportDocumentProperties"`
        ExportWorksheetProperties bool `json:"ExportWorksheetProperties,omitempty" xml:"ExportWorksheetProperties"`
        ExportWorkbookProperties bool `json:"ExportWorkbookProperties,omitempty" xml:"ExportWorkbookProperties"`
        ExportFrameScriptsAndProperties bool `json:"ExportFrameScriptsAndProperties,omitempty" xml:"ExportFrameScriptsAndProperties"`
        AttachedFilesDirectory string `json:"AttachedFilesDirectory,omitempty" xml:"AttachedFilesDirectory"`
        AttachedFilesUrlPrefix string `json:"AttachedFilesUrlPrefix,omitempty" xml:"AttachedFilesUrlPrefix"`
        Encoding string `json:"Encoding,omitempty" xml:"Encoding"`
        ExportActiveWorksheetOnly bool `json:"ExportActiveWorksheetOnly,omitempty" xml:"ExportActiveWorksheetOnly"`
        ExportChartImageFormat string `json:"ExportChartImageFormat,omitempty" xml:"ExportChartImageFormat"`
        ExportImagesAsBase64 bool `json:"ExportImagesAsBase64,omitempty" xml:"ExportImagesAsBase64"`
        HiddenColDisplayType string `json:"HiddenColDisplayType,omitempty" xml:"HiddenColDisplayType"`
        HiddenRowDisplayType string `json:"HiddenRowDisplayType,omitempty" xml:"HiddenRowDisplayType"`
        HtmlCrossStringType string `json:"HtmlCrossStringType,omitempty" xml:"HtmlCrossStringType"`
        IsExpImageToTempDir bool `json:"IsExpImageToTempDir,omitempty" xml:"IsExpImageToTempDir"`
        PageTitle string `json:"PageTitle,omitempty" xml:"PageTitle"`
        ParseHtmlTagInCell bool `json:"ParseHtmlTagInCell,omitempty" xml:"ParseHtmlTagInCell"`
        CellNameAttribute string `json:"CellNameAttribute,omitempty" xml:"CellNameAttribute"`
 
}


type HtmlSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportPageHeaders bool `json:"ExportPageHeaders,omitempty" xml:"ExportPageHeaders"`
    ExportPageFooters bool `json:"ExportPageFooters,omitempty" xml:"ExportPageFooters"`
    ExportRowColumnHeadings bool `json:"ExportRowColumnHeadings,omitempty" xml:"ExportRowColumnHeadings"`
    ShowAllSheets bool `json:"ShowAllSheets,omitempty" xml:"ShowAllSheets"`
    ImageOptions *ImageOrPrintOptions `json:"ImageOptions,omitempty" xml:"ImageOptions"`
    SaveAsSingleFile bool `json:"SaveAsSingleFile,omitempty" xml:"SaveAsSingleFile"`
    ExportHiddenWorksheet bool `json:"ExportHiddenWorksheet,omitempty" xml:"ExportHiddenWorksheet"`
    ExportGridLines bool `json:"ExportGridLines,omitempty" xml:"ExportGridLines"`
    PresentationPreference bool `json:"PresentationPreference,omitempty" xml:"PresentationPreference"`
    CellCssPrefix string `json:"CellCssPrefix,omitempty" xml:"CellCssPrefix"`
    TableCssId string `json:"TableCssId,omitempty" xml:"TableCssId"`
    IsFullPathLink bool `json:"IsFullPathLink,omitempty" xml:"IsFullPathLink"`
    ExportWorksheetCSSSeparately bool `json:"ExportWorksheetCSSSeparately,omitempty" xml:"ExportWorksheetCSSSeparately"`
    ExportSimilarBorderStyle bool `json:"ExportSimilarBorderStyle,omitempty" xml:"ExportSimilarBorderStyle"`
    MergeEmptyTdForcely bool `json:"MergeEmptyTdForcely,omitempty" xml:"MergeEmptyTdForcely"`
    ExportCellCoordinate bool `json:"ExportCellCoordinate,omitempty" xml:"ExportCellCoordinate"`
    ExportExtraHeadings bool `json:"ExportExtraHeadings,omitempty" xml:"ExportExtraHeadings"`
    ExportHeadings bool `json:"ExportHeadings,omitempty" xml:"ExportHeadings"`
    ExportFormula bool `json:"ExportFormula,omitempty" xml:"ExportFormula"`
    AddTooltipText bool `json:"AddTooltipText,omitempty" xml:"AddTooltipText"`
    ExportBogusRowData bool `json:"ExportBogusRowData,omitempty" xml:"ExportBogusRowData"`
    ExcludeUnusedStyles bool `json:"ExcludeUnusedStyles,omitempty" xml:"ExcludeUnusedStyles"`
    ExportDocumentProperties bool `json:"ExportDocumentProperties,omitempty" xml:"ExportDocumentProperties"`
    ExportWorksheetProperties bool `json:"ExportWorksheetProperties,omitempty" xml:"ExportWorksheetProperties"`
    ExportWorkbookProperties bool `json:"ExportWorkbookProperties,omitempty" xml:"ExportWorkbookProperties"`
    ExportFrameScriptsAndProperties bool `json:"ExportFrameScriptsAndProperties,omitempty" xml:"ExportFrameScriptsAndProperties"`
    AttachedFilesDirectory string `json:"AttachedFilesDirectory,omitempty" xml:"AttachedFilesDirectory"`
    AttachedFilesUrlPrefix string `json:"AttachedFilesUrlPrefix,omitempty" xml:"AttachedFilesUrlPrefix"`
    Encoding string `json:"Encoding,omitempty" xml:"Encoding"`
    ExportActiveWorksheetOnly bool `json:"ExportActiveWorksheetOnly,omitempty" xml:"ExportActiveWorksheetOnly"`
    ExportChartImageFormat string `json:"ExportChartImageFormat,omitempty" xml:"ExportChartImageFormat"`
    ExportImagesAsBase64 bool `json:"ExportImagesAsBase64,omitempty" xml:"ExportImagesAsBase64"`
    HiddenColDisplayType string `json:"HiddenColDisplayType,omitempty" xml:"HiddenColDisplayType"`
    HiddenRowDisplayType string `json:"HiddenRowDisplayType,omitempty" xml:"HiddenRowDisplayType"`
    HtmlCrossStringType string `json:"HtmlCrossStringType,omitempty" xml:"HtmlCrossStringType"`
    IsExpImageToTempDir bool `json:"IsExpImageToTempDir,omitempty" xml:"IsExpImageToTempDir"`
    PageTitle string `json:"PageTitle,omitempty" xml:"PageTitle"`
    ParseHtmlTagInCell bool `json:"ParseHtmlTagInCell,omitempty" xml:"ParseHtmlTagInCell"`
    CellNameAttribute string `json:"CellNameAttribute,omitempty" xml:"CellNameAttribute"`
}


type ImageSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ChartImageType string `json:"ChartImageType,omitempty" xml:"ChartImageType"`
    EmbededImageNameInSvg string `json:"EmbededImageNameInSvg,omitempty" xml:"EmbededImageNameInSvg"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
    ImageFormat string `json:"ImageFormat,omitempty" xml:"ImageFormat"`
    IsCellAutoFit bool `json:"IsCellAutoFit,omitempty" xml:"IsCellAutoFit"`
    OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
    OnlyArea bool `json:"OnlyArea,omitempty" xml:"OnlyArea"`
    PrintingPage string `json:"PrintingPage,omitempty" xml:"PrintingPage"`
    PrintWithStatusDialog bool `json:"PrintWithStatusDialog,omitempty" xml:"PrintWithStatusDialog"`
    Quality int64 `json:"Quality,omitempty" xml:"Quality"`
    TiffCompression string `json:"TiffCompression,omitempty" xml:"TiffCompression"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
}


type JsonSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportArea *CellArea `json:"ExportArea,omitempty" xml:"ExportArea"`
    HasHeaderRow bool `json:"HasHeaderRow,omitempty" xml:"HasHeaderRow"`
    ExportAsString bool `json:"ExportAsString,omitempty" xml:"ExportAsString"`
    Indent string `json:"Indent,omitempty" xml:"Indent"`
}


type MarkdownSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    Encoding string `json:"Encoding,omitempty" xml:"Encoding"`
    FormatStrategy string `json:"FormatStrategy,omitempty" xml:"FormatStrategy"`
    LineSeparator string `json:"LineSeparator,omitempty" xml:"LineSeparator"`
}


type MHtmlSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportPageHeaders bool `json:"ExportPageHeaders,omitempty" xml:"ExportPageHeaders"`
    ExportPageFooters bool `json:"ExportPageFooters,omitempty" xml:"ExportPageFooters"`
    ExportRowColumnHeadings bool `json:"ExportRowColumnHeadings,omitempty" xml:"ExportRowColumnHeadings"`
    ShowAllSheets bool `json:"ShowAllSheets,omitempty" xml:"ShowAllSheets"`
    ImageOptions *ImageOrPrintOptions `json:"ImageOptions,omitempty" xml:"ImageOptions"`
    SaveAsSingleFile bool `json:"SaveAsSingleFile,omitempty" xml:"SaveAsSingleFile"`
    ExportHiddenWorksheet bool `json:"ExportHiddenWorksheet,omitempty" xml:"ExportHiddenWorksheet"`
    ExportGridLines bool `json:"ExportGridLines,omitempty" xml:"ExportGridLines"`
    PresentationPreference bool `json:"PresentationPreference,omitempty" xml:"PresentationPreference"`
    CellCssPrefix string `json:"CellCssPrefix,omitempty" xml:"CellCssPrefix"`
    TableCssId string `json:"TableCssId,omitempty" xml:"TableCssId"`
    IsFullPathLink bool `json:"IsFullPathLink,omitempty" xml:"IsFullPathLink"`
    ExportWorksheetCSSSeparately bool `json:"ExportWorksheetCSSSeparately,omitempty" xml:"ExportWorksheetCSSSeparately"`
    ExportSimilarBorderStyle bool `json:"ExportSimilarBorderStyle,omitempty" xml:"ExportSimilarBorderStyle"`
    MergeEmptyTdForcely bool `json:"MergeEmptyTdForcely,omitempty" xml:"MergeEmptyTdForcely"`
    ExportCellCoordinate bool `json:"ExportCellCoordinate,omitempty" xml:"ExportCellCoordinate"`
    ExportExtraHeadings bool `json:"ExportExtraHeadings,omitempty" xml:"ExportExtraHeadings"`
    ExportHeadings bool `json:"ExportHeadings,omitempty" xml:"ExportHeadings"`
    ExportFormula bool `json:"ExportFormula,omitempty" xml:"ExportFormula"`
    AddTooltipText bool `json:"AddTooltipText,omitempty" xml:"AddTooltipText"`
    ExportBogusRowData bool `json:"ExportBogusRowData,omitempty" xml:"ExportBogusRowData"`
    ExcludeUnusedStyles bool `json:"ExcludeUnusedStyles,omitempty" xml:"ExcludeUnusedStyles"`
    ExportDocumentProperties bool `json:"ExportDocumentProperties,omitempty" xml:"ExportDocumentProperties"`
    ExportWorksheetProperties bool `json:"ExportWorksheetProperties,omitempty" xml:"ExportWorksheetProperties"`
    ExportWorkbookProperties bool `json:"ExportWorkbookProperties,omitempty" xml:"ExportWorkbookProperties"`
    ExportFrameScriptsAndProperties bool `json:"ExportFrameScriptsAndProperties,omitempty" xml:"ExportFrameScriptsAndProperties"`
    AttachedFilesDirectory string `json:"AttachedFilesDirectory,omitempty" xml:"AttachedFilesDirectory"`
    AttachedFilesUrlPrefix string `json:"AttachedFilesUrlPrefix,omitempty" xml:"AttachedFilesUrlPrefix"`
    Encoding string `json:"Encoding,omitempty" xml:"Encoding"`
    ExportActiveWorksheetOnly bool `json:"ExportActiveWorksheetOnly,omitempty" xml:"ExportActiveWorksheetOnly"`
    ExportChartImageFormat string `json:"ExportChartImageFormat,omitempty" xml:"ExportChartImageFormat"`
    ExportImagesAsBase64 bool `json:"ExportImagesAsBase64,omitempty" xml:"ExportImagesAsBase64"`
    HiddenColDisplayType string `json:"HiddenColDisplayType,omitempty" xml:"HiddenColDisplayType"`
    HiddenRowDisplayType string `json:"HiddenRowDisplayType,omitempty" xml:"HiddenRowDisplayType"`
    HtmlCrossStringType string `json:"HtmlCrossStringType,omitempty" xml:"HtmlCrossStringType"`
    IsExpImageToTempDir bool `json:"IsExpImageToTempDir,omitempty" xml:"IsExpImageToTempDir"`
    PageTitle string `json:"PageTitle,omitempty" xml:"PageTitle"`
    ParseHtmlTagInCell bool `json:"ParseHtmlTagInCell,omitempty" xml:"ParseHtmlTagInCell"`
    CellNameAttribute string `json:"CellNameAttribute,omitempty" xml:"CellNameAttribute"`
}


type OdsSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    GeneratorType string `json:"GeneratorType,omitempty" xml:"GeneratorType"`
    OdfStrictVersion string `json:"OdfStrictVersion,omitempty" xml:"OdfStrictVersion"`
    IgnorePivotTables bool `json:"IgnorePivotTables,omitempty" xml:"IgnorePivotTables"`
}


type OoxmlSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportCellName bool `json:"ExportCellName,omitempty" xml:"ExportCellName"`
    UpdateZoom bool `json:"UpdateZoom,omitempty" xml:"UpdateZoom"`
    EnableZip64 bool `json:"EnableZip64,omitempty" xml:"EnableZip64"`
    EmbedOoxmlAsOleObject bool `json:"EmbedOoxmlAsOleObject,omitempty" xml:"EmbedOoxmlAsOleObject"`
    CompressionType string `json:"CompressionType,omitempty" xml:"CompressionType"`
}


type PclSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    FontFullName string `json:"fontFullName,omitempty" xml:"fontFullName"`
    FontPclName string `json:"fontPclName,omitempty" xml:"fontPclName"`
}


type PdfSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    DisplayDocTitle bool `json:"DisplayDocTitle,omitempty" xml:"DisplayDocTitle"`
    ExportDocumentStructure bool `json:"ExportDocumentStructure,omitempty" xml:"ExportDocumentStructure"`
    EmfRenderSetting string `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
    CustomPropertiesExport string `json:"CustomPropertiesExport,omitempty" xml:"CustomPropertiesExport"`
    OptimizationType string `json:"OptimizationType,omitempty" xml:"OptimizationType"`
    Producer string `json:"Producer,omitempty" xml:"Producer"`
    PdfCompression string `json:"PdfCompression,omitempty" xml:"PdfCompression"`
    FontEncoding string `json:"FontEncoding,omitempty" xml:"FontEncoding"`
    Watermark *RenderingWatermark `json:"Watermark,omitempty" xml:"Watermark"`
    CalculateFormula bool `json:"CalculateFormula,omitempty" xml:"CalculateFormula"`
    CheckFontCompatibility bool `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
    Compliance string `json:"Compliance,omitempty" xml:"Compliance"`
    DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
    OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
    PrintingPageType string `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
    SecurityOptions *PdfSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions"`
    DesiredPPI int64 `json:"desiredPPI,omitempty" xml:"desiredPPI"`
    JpegQuality int64 `json:"jpegQuality,omitempty" xml:"jpegQuality"`
    ImageType string `json:"ImageType,omitempty" xml:"ImageType"`
}


type PptxSaveOptions struct {
     
        DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
        CheckWorkbookDefaultFont bool `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
        CheckFontCompatibility bool `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
        IsFontSubstitutionCharGranularity bool `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
        OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
        AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
        IgnoreError bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
        OutputBlankPageWhenNothingToPrint bool `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
        PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
        PageCount int64 `json:"PageCount,omitempty" xml:"PageCount"`
        PrintingPageType string `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
        GridlineType string `json:"GridlineType,omitempty" xml:"GridlineType"`
        TextCrossType string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
        DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
        EmfRenderSetting string `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
 
    IgnoreHiddenRows bool `json:"IgnoreHiddenRows,omitempty" xml:"IgnoreHiddenRows"`
    AdjustFontSizeForRowType string `json:"AdjustFontSizeForRowType,omitempty" xml:"AdjustFontSizeForRowType"`
    ExportViewType string `json:"ExportViewType,omitempty" xml:"ExportViewType"`
}


type SaveOptions struct {
 
    SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
    CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
    ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
    CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
    EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
    RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
    SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
    ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
    MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
    SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
    EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
}


type SaveOptionsData struct {
 
    SaveOptions *SaveOptions `json:"SaveOptions,omitempty" xml:"SaveOptions"`
    Filename string `json:"Filename,omitempty" xml:"Filename"`
    StorageName string `json:"StorageName,omitempty" xml:"StorageName"`
}


type SpreadsheetML2003SaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportColumnIndexOfCell bool `json:"ExportColumnIndexOfCell,omitempty" xml:"ExportColumnIndexOfCell"`
    IsIndentedFormatting bool `json:"IsIndentedFormatting,omitempty" xml:"IsIndentedFormatting"`
    LimitAsXls bool `json:"LimitAsXls,omitempty" xml:"LimitAsXls"`
}


type SqlScriptSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    CheckIfTableExists bool `json:"CheckIfTableExists,omitempty" xml:"CheckIfTableExists"`
    ColumnTypeMap string `json:"ColumnTypeMap,omitempty" xml:"ColumnTypeMap"`
    CheckAllDataForColumnType bool `json:"CheckAllDataForColumnType,omitempty" xml:"CheckAllDataForColumnType"`
    AddBlankLineBetweenRows bool `json:"AddBlankLineBetweenRows,omitempty" xml:"AddBlankLineBetweenRows"`
    Separator string `json:"Separator,omitempty" xml:"Separator"`
    OperatorType string `json:"OperatorType,omitempty" xml:"OperatorType"`
    PrimaryKey int64 `json:"PrimaryKey,omitempty" xml:"PrimaryKey"`
    CreateTable bool `json:"CreateTable,omitempty" xml:"CreateTable"`
    IdName string `json:"IdName,omitempty" xml:"IdName"`
    StartId int64 `json:"StartId,omitempty" xml:"StartId"`
    TableName string `json:"TableName,omitempty" xml:"TableName"`
    ExportAsString bool `json:"ExportAsString,omitempty" xml:"ExportAsString"`
    ExportArea *CellArea `json:"ExportArea,omitempty" xml:"ExportArea"`
    HasHeaderRow bool `json:"HasHeaderRow,omitempty" xml:"HasHeaderRow"`
}


type SvgSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    SheetIndex int64 `json:"SheetIndex,omitempty" xml:"SheetIndex"`
    ChartImageType string `json:"ChartImageType,omitempty" xml:"ChartImageType"`
    EmbededImageNameInSvg string `json:"EmbededImageNameInSvg,omitempty" xml:"EmbededImageNameInSvg"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
    ImageFormat string `json:"ImageFormat,omitempty" xml:"ImageFormat"`
    IsCellAutoFit bool `json:"IsCellAutoFit,omitempty" xml:"IsCellAutoFit"`
    OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
    OnlyArea bool `json:"OnlyArea,omitempty" xml:"OnlyArea"`
    PrintingPage string `json:"PrintingPage,omitempty" xml:"PrintingPage"`
    PrintWithStatusDialog bool `json:"PrintWithStatusDialog,omitempty" xml:"PrintWithStatusDialog"`
    Quality int64 `json:"Quality,omitempty" xml:"Quality"`
    TiffCompression string `json:"TiffCompression,omitempty" xml:"TiffCompression"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
}


type TxtSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    QuoteType string `json:"QuoteType,omitempty" xml:"QuoteType"`
    Separator string `json:"Separator,omitempty" xml:"Separator"`
    SeparatorString string `json:"SeparatorString,omitempty" xml:"SeparatorString"`
    AlwaysQuoted bool `json:"AlwaysQuoted,omitempty" xml:"AlwaysQuoted"`
}


type XlsbSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    ExportAllColumnIndexes bool `json:"ExportAllColumnIndexes,omitempty" xml:"ExportAllColumnIndexes"`
    CompressionType string `json:"CompressionType,omitempty" xml:"CompressionType"`
}


type XlsSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    MatchColor bool `json:"MatchColor,omitempty" xml:"MatchColor"`
    WpsCompatibility bool `json:"WpsCompatibility,omitempty" xml:"WpsCompatibility"`
}


type XmlSaveOptions struct {
     
        SaveFormat string `json:"SaveFormat,omitempty" xml:"SaveFormat"`
        CachedFileFolder string `json:"CachedFileFolder,omitempty" xml:"CachedFileFolder"`
        ClearData bool `json:"ClearData,omitempty" xml:"ClearData"`
        CreateDirectory bool `json:"CreateDirectory,omitempty" xml:"CreateDirectory"`
        EnableHTTPCompression bool `json:"EnableHTTPCompression,omitempty" xml:"EnableHTTPCompression"`
        RefreshChartCache bool `json:"RefreshChartCache,omitempty" xml:"RefreshChartCache"`
        SortNames bool `json:"SortNames,omitempty" xml:"SortNames"`
        ValidateMergedAreas bool `json:"ValidateMergedAreas,omitempty" xml:"ValidateMergedAreas"`
        MergeAreas bool `json:"MergeAreas,omitempty" xml:"MergeAreas"`
        SortExternalNames bool `json:"SortExternalNames,omitempty" xml:"SortExternalNames"`
        CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
        UpdateSmartArt bool `json:"UpdateSmartArt,omitempty" xml:"UpdateSmartArt"`
        EncryptDocumentProperties bool `json:"EncryptDocumentProperties,omitempty" xml:"EncryptDocumentProperties"`
 
    SheetIndexes []int64 `json:"SheetIndexes,omitempty" xml:"SheetIndexes"`
    ExportArea *CellArea `json:"ExportArea,omitempty" xml:"ExportArea"`
    HasHeaderRow bool `json:"HasHeaderRow,omitempty" xml:"HasHeaderRow"`
    XmlMapName string `json:"XmlMapName,omitempty" xml:"XmlMapName"`
    SheetNameAsElementName bool `json:"SheetNameAsElementName,omitempty" xml:"SheetNameAsElementName"`
    DataAsAttribute bool `json:"DataAsAttribute,omitempty" xml:"DataAsAttribute"`
}


type XpsSaveOptions struct {
     
        DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
        CheckWorkbookDefaultFont bool `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
        CheckFontCompatibility bool `json:"CheckFontCompatibility,omitempty" xml:"CheckFontCompatibility"`
        IsFontSubstitutionCharGranularity bool `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
        OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
        AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
        IgnoreError bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
        OutputBlankPageWhenNothingToPrint bool `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
        PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
        PageCount int64 `json:"PageCount,omitempty" xml:"PageCount"`
        PrintingPageType string `json:"PrintingPageType,omitempty" xml:"PrintingPageType"`
        GridlineType string `json:"GridlineType,omitempty" xml:"GridlineType"`
        TextCrossType string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
        DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
        EmfRenderSetting string `json:"EmfRenderSetting,omitempty" xml:"EmfRenderSetting"`
 
}


type AggregateResultByColorResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    AggregateResults []AggregateResultByColor `json:"AggregateResults,omitempty" xml:"AggregateResults"`
}


type BrokenLinksResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    BrokenLinks []BrokenLink `json:"BrokenLinks,omitempty" xml:"BrokenLinks"`
}


type CellsCloudFileInfoResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    FileInfo *CellsCloudFileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
}


type CellsCloudPublicKeyResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    CellsCloudPublicKey *CellsCloudPublicKey `json:"CellsCloudPublicKey,omitempty" xml:"CellsCloudPublicKey"`
}


type CellsCloudResponse struct {
 
    Code int64 `json:"Code,omitempty" xml:"Code"`
    Status string `json:"Status,omitempty" xml:"Status"`
}


type SaveResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    SaveResult *SaveResult `json:"SaveResult,omitempty" xml:"SaveResult"`
}


type SearchResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    TextItems []TextItem `json:"TextItems,omitempty" xml:"TextItems"`
}


type ImageOrPrintOptions struct {
 
    TextCrossType string `json:"TextCrossType,omitempty" xml:"TextCrossType"`
    GridlineType string `json:"GridlineType,omitempty" xml:"GridlineType"`
    OutputBlankPageWhenNothingToPrint bool `json:"OutputBlankPageWhenNothingToPrint,omitempty" xml:"OutputBlankPageWhenNothingToPrint"`
    CheckWorkbookDefaultFont bool `json:"CheckWorkbookDefaultFont,omitempty" xml:"CheckWorkbookDefaultFont"`
    DefaultFont string `json:"DefaultFont,omitempty" xml:"DefaultFont"`
    IsOptimized bool `json:"IsOptimized,omitempty" xml:"IsOptimized"`
    PageCount int64 `json:"PageCount,omitempty" xml:"PageCount"`
    PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
    IsFontSubstitutionCharGranularity bool `json:"IsFontSubstitutionCharGranularity,omitempty" xml:"IsFontSubstitutionCharGranularity"`
    Transparent bool `json:"Transparent,omitempty" xml:"Transparent"`
    OnlyArea bool `json:"OnlyArea,omitempty" xml:"OnlyArea"`
    SVGFitToViewPort bool `json:"SVGFitToViewPort,omitempty" xml:"SVGFitToViewPort"`
    EmbededImageNameInSvg string `json:"EmbededImageNameInSvg,omitempty" xml:"EmbededImageNameInSvg"`
    AllColumnsInOnePagePerSheet bool `json:"AllColumnsInOnePagePerSheet,omitempty" xml:"AllColumnsInOnePagePerSheet"`
    PrintWithStatusDialog bool `json:"PrintWithStatusDialog,omitempty" xml:"PrintWithStatusDialog"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
    DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
    TiffColorDepth string `json:"TiffColorDepth,omitempty" xml:"TiffColorDepth"`
    TiffCompression string `json:"TiffCompression,omitempty" xml:"TiffCompression"`
    PrintingPage string `json:"PrintingPage,omitempty" xml:"PrintingPage"`
    Quality int64 `json:"Quality,omitempty" xml:"Quality"`
    ImageType string `json:"ImageType,omitempty" xml:"ImageType"`
    OnePagePerSheet bool `json:"OnePagePerSheet,omitempty" xml:"OnePagePerSheet"`
    TiffBinarizationMethod string `json:"TiffBinarizationMethod,omitempty" xml:"TiffBinarizationMethod"`
}


type RenderingFont struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Size float64 `json:"Size,omitempty" xml:"Size"`
    Bold bool `json:"Bold,omitempty" xml:"Bold"`
    Italic bool `json:"Italic,omitempty" xml:"Italic"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
}


type RenderingWatermark struct {
 
    Rotation float64 `json:"Rotation,omitempty" xml:"Rotation"`
    ScaleToPagePercent int64 `json:"ScaleToPagePercent,omitempty" xml:"ScaleToPagePercent"`
    Opacity float64 `json:"Opacity,omitempty" xml:"Opacity"`
    IsBackground bool `json:"IsBackground,omitempty" xml:"IsBackground"`
    Text string `json:"Text,omitempty" xml:"Text"`
    Font *RenderingFont `json:"Font,omitempty" xml:"Font"`
    Image []int64 `json:"Image,omitempty" xml:"Image"`
    HAlignment string `json:"HAlignment,omitempty" xml:"HAlignment"`
    VAlignment string `json:"VAlignment,omitempty" xml:"VAlignment"`
    OffsetX float64 `json:"OffsetX,omitempty" xml:"OffsetX"`
    OffsetY float64 `json:"OffsetY,omitempty" xml:"OffsetY"`
}


type Error struct {
 
}


type ErrorDetails struct {
 
}


type AboveAverage struct {
 
    IsAboveAverage bool `json:"IsAboveAverage,omitempty" xml:"IsAboveAverage"`
    IsEqualAverage bool `json:"IsEqualAverage,omitempty" xml:"IsEqualAverage"`
    StdDev int64 `json:"StdDev,omitempty" xml:"StdDev"`
}


type AbstractCalculationEngine struct {
 
    IsParamLiteralRequired bool `json:"IsParamLiteralRequired,omitempty" xml:"IsParamLiteralRequired"`
    IsParamArrayModeRequired bool `json:"IsParamArrayModeRequired,omitempty" xml:"IsParamArrayModeRequired"`
    ProcessBuiltInFunctions bool `json:"ProcessBuiltInFunctions,omitempty" xml:"ProcessBuiltInFunctions"`
}


type AbstractCalculationMonitor struct {
 
    OriginalValue *interface{} `json:"OriginalValue,omitempty" xml:"OriginalValue"`
    ValueChanged bool `json:"ValueChanged,omitempty" xml:"ValueChanged"`
    CalculatedValue *interface{} `json:"CalculatedValue,omitempty" xml:"CalculatedValue"`
}


type AutoFilter struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    FilterColumns []FilterColumn `json:"FilterColumns,omitempty" xml:"FilterColumns"`
    Range_ string `json:"Range,omitempty" xml:"Range"`
    Sorter *DataSorter `json:"Sorter,omitempty" xml:"Sorter"`
    ShowFilterButton bool `json:"ShowFilterButton,omitempty" xml:"ShowFilterButton"`
}


type AutoFitterOptions struct {
 
    AutoFitMergedCellsType string `json:"AutoFitMergedCellsType,omitempty" xml:"AutoFitMergedCellsType"`
    IgnoreHidden bool `json:"IgnoreHidden,omitempty" xml:"IgnoreHidden"`
    OnlyAuto bool `json:"OnlyAuto,omitempty" xml:"OnlyAuto"`
    DefaultEditLanguage string `json:"DefaultEditLanguage,omitempty" xml:"DefaultEditLanguage"`
    MaxRowHeight float64 `json:"MaxRowHeight,omitempty" xml:"MaxRowHeight"`
    AutoFitWrappedTextType string `json:"AutoFitWrappedTextType,omitempty" xml:"AutoFitWrappedTextType"`
    FormatStrategy string `json:"FormatStrategy,omitempty" xml:"FormatStrategy"`
    ForRendering bool `json:"ForRendering,omitempty" xml:"ForRendering"`
}


type Border struct {
 
    LineStyle string `json:"LineStyle,omitempty" xml:"LineStyle"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
    BorderType string `json:"BorderType,omitempty" xml:"BorderType"`
    ThemeColor *ThemeColor `json:"ThemeColor,omitempty" xml:"ThemeColor"`
    ArgbColor int64 `json:"ArgbColor,omitempty" xml:"ArgbColor"`
}


type CalculationOptions struct {
 
    CalcStackSize int64 `json:"CalcStackSize,omitempty" xml:"CalcStackSize"`
    IgnoreError bool `json:"IgnoreError,omitempty" xml:"IgnoreError"`
    PrecisionStrategy string `json:"PrecisionStrategy,omitempty" xml:"PrecisionStrategy"`
    Recursive bool `json:"Recursive,omitempty" xml:"Recursive"`
    CustomEngine *AbstractCalculationEngine `json:"CustomEngine,omitempty" xml:"CustomEngine"`
    CalculationMonitor *AbstractCalculationMonitor `json:"CalculationMonitor,omitempty" xml:"CalculationMonitor"`
    LinkedDataSources []Workbook `json:"LinkedDataSources,omitempty" xml:"LinkedDataSources"`
}


type Cell struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Row int64 `json:"Row,omitempty" xml:"Row"`
    Column int64 `json:"Column,omitempty" xml:"Column"`
    Value string `json:"Value,omitempty" xml:"Value"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Formula string `json:"Formula,omitempty" xml:"Formula"`
    IsFormula bool `json:"IsFormula,omitempty" xml:"IsFormula"`
    IsMerged bool `json:"IsMerged,omitempty" xml:"IsMerged"`
    IsArrayHeader bool `json:"IsArrayHeader,omitempty" xml:"IsArrayHeader"`
    IsInArray bool `json:"IsInArray,omitempty" xml:"IsInArray"`
    IsErrorValue bool `json:"IsErrorValue,omitempty" xml:"IsErrorValue"`
    IsInTable bool `json:"IsInTable,omitempty" xml:"IsInTable"`
    IsStyleSet bool `json:"IsStyleSet,omitempty" xml:"IsStyleSet"`
    HtmlString string `json:"HtmlString,omitempty" xml:"HtmlString"`
    Style *LinkElement `json:"Style,omitempty" xml:"Style"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
}


type Cells struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    MaxRow int64 `json:"MaxRow,omitempty" xml:"MaxRow"`
    MaxColumn int64 `json:"MaxColumn,omitempty" xml:"MaxColumn"`
    CellCount int64 `json:"CellCount,omitempty" xml:"CellCount"`
    Rows *LinkElement `json:"Rows,omitempty" xml:"Rows"`
    Columns *LinkElement `json:"Columns,omitempty" xml:"Columns"`
    CellList []LinkElement `json:"CellList,omitempty" xml:"CellList"`
}


type CellsColor struct {
 
    Color *Color `json:"Color,omitempty" xml:"Color"`
    ColorIndex int64 `json:"ColorIndex,omitempty" xml:"ColorIndex"`
    IsShapeColor bool `json:"IsShapeColor,omitempty" xml:"IsShapeColor"`
    Tint float64 `json:"tint,omitempty" xml:"tint"`
    Argb int64 `json:"Argb,omitempty" xml:"Argb"`
    ThemeColor *ThemeColor `json:"ThemeColor,omitempty" xml:"ThemeColor"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}


type CellsDocumentProperties struct {
 
    DocumentPropertyList []CellsDocumentProperty `json:"DocumentPropertyList,omitempty" xml:"DocumentPropertyList"`
}


type CellsDocumentProperty struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Value string `json:"Value,omitempty" xml:"Value"`
    IsLinkedToContent string `json:"IsLinkedToContent,omitempty" xml:"IsLinkedToContent"`
    Source string `json:"Source,omitempty" xml:"Source"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    IsGeneratedName string `json:"IsGeneratedName,omitempty" xml:"IsGeneratedName"`
}


type ColorFilter struct {
 
    FilterByFillColor bool `json:"FilterByFillColor,omitempty" xml:"FilterByFillColor"`
    Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
    Color *CellsColor `json:"Color,omitempty" xml:"Color"`
    ForegroundColorColor *CellsColor `json:"ForegroundColorColor,omitempty" xml:"ForegroundColorColor"`
    BackgroundColor *CellsColor `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
}


type ColorScale struct {
 
    MaxCfvo *ConditionalFormattingValue `json:"MaxCfvo,omitempty" xml:"MaxCfvo"`
    MaxColor *Color `json:"MaxColor,omitempty" xml:"MaxColor"`
    MidCfvo *ConditionalFormattingValue `json:"MidCfvo,omitempty" xml:"MidCfvo"`
    MidColor *Color `json:"MidColor,omitempty" xml:"MidColor"`
    MinCfvo *ConditionalFormattingValue `json:"MinCfvo,omitempty" xml:"MinCfvo"`
    MinColor *Color `json:"MinColor,omitempty" xml:"MinColor"`
}


type Column struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    GroupLevel int64 `json:"GroupLevel,omitempty" xml:"GroupLevel"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
    Width float64 `json:"Width,omitempty" xml:"Width"`
    Style *LinkElement `json:"Style,omitempty" xml:"Style"`
}


type Columns struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    MaxColumn int64 `json:"MaxColumn,omitempty" xml:"MaxColumn"`
    ColumnsCount int64 `json:"ColumnsCount,omitempty" xml:"ColumnsCount"`
    ColumnsList []LinkElement `json:"ColumnsList,omitempty" xml:"ColumnsList"`
}


type Comment struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    CellName string `json:"CellName,omitempty" xml:"CellName"`
    Author string `json:"Author,omitempty" xml:"Author"`
    HtmlNote string `json:"HtmlNote,omitempty" xml:"HtmlNote"`
    Note string `json:"Note,omitempty" xml:"Note"`
    AutoSize bool `json:"AutoSize,omitempty" xml:"AutoSize"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    Width int64 `json:"Width,omitempty" xml:"Width"`
    Height int64 `json:"Height,omitempty" xml:"Height"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}


type Comments struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    CommentList []LinkElement `json:"CommentList,omitempty" xml:"CommentList"`
}


type ConditionalFormatting struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Sqref string `json:"sqref,omitempty" xml:"sqref"`
    FormatConditions []FormatCondition `json:"FormatConditions,omitempty" xml:"FormatConditions"`
}


type ConditionalFormattingIcon struct {
 
    ImageData string `json:"ImageData,omitempty" xml:"ImageData"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type ConditionalFormattings struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Count int64 `json:"Count,omitempty" xml:"Count"`
    ConditionalFormattingList []ConditionalFormatting `json:"ConditionalFormattingList,omitempty" xml:"ConditionalFormattingList"`
}


type ConditionalFormattingValue struct {
 
    IsGTE bool `json:"IsGTE,omitempty" xml:"IsGTE"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Value *interface{} `json:"Value,omitempty" xml:"Value"`
}


type CopyOptions struct {
 
    ColumnCharacterWidth bool `json:"ColumnCharacterWidth,omitempty" xml:"ColumnCharacterWidth"`
    CopyInvalidFormulasAsValues bool `json:"CopyInvalidFormulasAsValues,omitempty" xml:"CopyInvalidFormulasAsValues"`
    CopyNames bool `json:"CopyNames,omitempty" xml:"CopyNames"`
    ExtendToAdjacentRange bool `json:"ExtendToAdjacentRange,omitempty" xml:"ExtendToAdjacentRange"`
    ReferToDestinationSheet bool `json:"ReferToDestinationSheet,omitempty" xml:"ReferToDestinationSheet"`
    ReferToSheetWithSameName bool `json:"ReferToSheetWithSameName,omitempty" xml:"ReferToSheetWithSameName"`
    CopyTheme bool `json:"CopyTheme,omitempty" xml:"CopyTheme"`
}


type CriteriaMultipleFilter struct {
     
 
    Criteria string `json:"Criteria,omitempty" xml:"Criteria"`
}


type CustomFilter struct {
 
    Criteria *interface{} `json:"Criteria,omitempty" xml:"Criteria"`
    FilterOperatorType string `json:"FilterOperatorType,omitempty" xml:"FilterOperatorType"`
}


type DataBar struct {
 
    AxisColor *Color `json:"AxisColor,omitempty" xml:"AxisColor"`
    AxisPosition string `json:"AxisPosition,omitempty" xml:"AxisPosition"`
    BarBorder *DataBarBorder `json:"BarBorder,omitempty" xml:"BarBorder"`
    BarFillType string `json:"BarFillType,omitempty" xml:"BarFillType"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
    Direction string `json:"Direction,omitempty" xml:"Direction"`
    MaxCfvo *ConditionalFormattingValue `json:"MaxCfvo,omitempty" xml:"MaxCfvo"`
    MaxLength int64 `json:"MaxLength,omitempty" xml:"MaxLength"`
    MinCfvo *ConditionalFormattingValue `json:"MinCfvo,omitempty" xml:"MinCfvo"`
    MinLength int64 `json:"MinLength,omitempty" xml:"MinLength"`
    NegativeBarFormat *NegativeBarFormat `json:"NegativeBarFormat,omitempty" xml:"NegativeBarFormat"`
    ShowValue bool `json:"ShowValue,omitempty" xml:"ShowValue"`
}


type DataBarBorder struct {
 
    Color *Color `json:"Color,omitempty" xml:"Color"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type DataCleansing struct {
 
    Ranges []Range `json:"Ranges,omitempty" xml:"Ranges"`
    NeedFillData bool `json:"NeedFillData,omitempty" xml:"NeedFillData"`
    DataFill *DataFill `json:"DataFill,omitempty" xml:"DataFill"`
}


type DataColumnFillValue struct {
 
    ColumnIndex int64 `json:"ColumnIndex,omitempty" xml:"ColumnIndex"`
    DataFillValue *DataFillValue `json:"DataFillValue,omitempty" xml:"DataFillValue"`
}


type DataFill struct {
 
    Ranges []Range `json:"Ranges,omitempty" xml:"Ranges"`
    DataFillDefaultValue *DataFillValue `json:"DataFillDefaultValue,omitempty" xml:"DataFillDefaultValue"`
    DataColumnFillValueList []DataColumnFillValue `json:"DataColumnFillValueList,omitempty" xml:"DataColumnFillValueList"`
}


type DataFillValue struct {
 
    DefaultBoolean bool `json:"DefaultBoolean,omitempty" xml:"DefaultBoolean"`
    DefaultString string `json:"DefaultString,omitempty" xml:"DefaultString"`
    DefaultNumber int64 `json:"DefaultNumber,omitempty" xml:"DefaultNumber"`
    DefaultDouble float64 `json:"DefaultDouble,omitempty" xml:"DefaultDouble"`
    DefaultDate string `json:"DefaultDate,omitempty" xml:"DefaultDate"`
}


type DataItem struct {
 
    DataItemType string `json:"DataItemType,omitempty" xml:"DataItemType"`
    Value string `json:"Value,omitempty" xml:"Value"`
}


type DataSorter struct {
 
    CaseSensitive bool `json:"CaseSensitive,omitempty" xml:"CaseSensitive"`
    HasHeaders bool `json:"HasHeaders,omitempty" xml:"HasHeaders"`
    KeyList []SortKey `json:"KeyList,omitempty" xml:"KeyList"`
    SortLeftToRight bool `json:"SortLeftToRight,omitempty" xml:"SortLeftToRight"`
    SortAsNumber bool `json:"SortAsNumber,omitempty" xml:"SortAsNumber"`
    Keys []DataSorterKey `json:"Keys,omitempty" xml:"Keys"`
}


type DataSorterKey struct {
 
    Order string `json:"Order,omitempty" xml:"Order"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    IconSetType string `json:"IconSetType,omitempty" xml:"IconSetType"`
    IconId int64 `json:"IconId,omitempty" xml:"IconId"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
}


type DataSource struct {
 
    DataSourceType string `json:"DataSourceType,omitempty" xml:"DataSourceType"`
    DataPath string `json:"DataPath,omitempty" xml:"DataPath"`
}


type DateTimeGroupItem struct {
     
 
    DateTimeGroupingType string `json:"DateTimeGroupingType,omitempty" xml:"DateTimeGroupingType"`
    Day int64 `json:"Day,omitempty" xml:"Day"`
    Hour int64 `json:"Hour,omitempty" xml:"Hour"`
    Minute int64 `json:"Minute,omitempty" xml:"Minute"`
    Month int64 `json:"Month,omitempty" xml:"Month"`
    Second int64 `json:"Second,omitempty" xml:"Second"`
    Year int64 `json:"Year,omitempty" xml:"Year"`
}


type DeduplicationRegion struct {
 
    Ranges []Range `json:"Ranges,omitempty" xml:"Ranges"`
    WorksheetNameList []string `json:"WorksheetNameList,omitempty" xml:"WorksheetNameList"`
}


type DynamicFilter struct {
 
    DynamicFilterType string `json:"DynamicFilterType,omitempty" xml:"DynamicFilterType"`
    MaxValue *interface{} `json:"MaxValue,omitempty" xml:"MaxValue"`
    Value *interface{} `json:"Value,omitempty" xml:"Value"`
}


type FileInfo struct {
 
    Filename string `json:"Filename,omitempty" xml:"Filename"`
    FileSize int64 `json:"FileSize,omitempty" xml:"FileSize"`
    FileContent string `json:"FileContent,omitempty" xml:"FileContent"`
}


type FileSource struct {
 
    FileSourceType string `json:"FileSourceType,omitempty" xml:"FileSourceType"`
    FilePath string `json:"FilePath,omitempty" xml:"FilePath"`
}


type FilesResult struct {
 
    Files []FileInfo `json:"Files,omitempty" xml:"Files"`
}


type FilterColumn struct {
 
    FieldIndex int64 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
    FilterType string `json:"FilterType,omitempty" xml:"FilterType"`
    MultipleFilters *MultipleFilters `json:"MultipleFilters,omitempty" xml:"MultipleFilters"`
    ColorFilter *ColorFilter `json:"ColorFilter,omitempty" xml:"ColorFilter"`
    CustomFilters []CustomFilter `json:"CustomFilters,omitempty" xml:"CustomFilters"`
    DynamicFilter *DynamicFilter `json:"DynamicFilter,omitempty" xml:"DynamicFilter"`
    IconFilter *IconFilter `json:"IconFilter,omitempty" xml:"IconFilter"`
    Top10Filter *Top10Filter `json:"Top10Filter,omitempty" xml:"Top10Filter"`
    Visibledropdown string `json:"Visibledropdown,omitempty" xml:"Visibledropdown"`
}


type Font struct {
 
    Color *Color `json:"Color,omitempty" xml:"Color"`
    DoubleSize float64 `json:"DoubleSize,omitempty" xml:"DoubleSize"`
    IsBold bool `json:"IsBold,omitempty" xml:"IsBold"`
    IsItalic bool `json:"IsItalic,omitempty" xml:"IsItalic"`
    IsStrikeout bool `json:"IsStrikeout,omitempty" xml:"IsStrikeout"`
    IsSubscript bool `json:"IsSubscript,omitempty" xml:"IsSubscript"`
    IsSuperscript bool `json:"IsSuperscript,omitempty" xml:"IsSuperscript"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Size int64 `json:"Size,omitempty" xml:"Size"`
    Underline string `json:"Underline,omitempty" xml:"Underline"`
}


type FontSetting struct {
 
    Font *Font `json:"Font,omitempty" xml:"Font"`
    Length int64 `json:"Length,omitempty" xml:"Length"`
    StartIndex int64 `json:"StartIndex,omitempty" xml:"StartIndex"`
    TextOptions *TextOptions `json:"TextOptions,omitempty" xml:"TextOptions"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type FormatCondition struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Priority int64 `json:"Priority,omitempty" xml:"Priority"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    StopIfTrue bool `json:"StopIfTrue,omitempty" xml:"StopIfTrue"`
    AboveAverage *AboveAverage `json:"AboveAverage,omitempty" xml:"AboveAverage"`
    ColorScale *ColorScale `json:"ColorScale,omitempty" xml:"ColorScale"`
    DataBar *DataBar `json:"DataBar,omitempty" xml:"DataBar"`
    Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
    Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
    IconSet *IconSet `json:"IconSet,omitempty" xml:"IconSet"`
    Operator string `json:"Operator,omitempty" xml:"Operator"`
    Style *Style `json:"Style,omitempty" xml:"Style"`
    Text string `json:"Text,omitempty" xml:"Text"`
    TimePeriod string `json:"TimePeriod,omitempty" xml:"TimePeriod"`
    Top10 *Top10 `json:"Top10,omitempty" xml:"Top10"`
}


type FormulaFormatCondition struct {
 
    Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
    Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
    Operator string `json:"Operator,omitempty" xml:"Operator"`
}


type FormulaSettings struct {
 
    CalculateOnOpen bool `json:"CalculateOnOpen,omitempty" xml:"CalculateOnOpen"`
    CalculateOnSave bool `json:"CalculateOnSave,omitempty" xml:"CalculateOnSave"`
    ForceFullCalculation bool `json:"ForceFullCalculation,omitempty" xml:"ForceFullCalculation"`
    CalculationMode string `json:"CalculationMode,omitempty" xml:"CalculationMode"`
    CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
    EnableIterativeCalculation bool `json:"EnableIterativeCalculation,omitempty" xml:"EnableIterativeCalculation"`
    MaxIteration int64 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
    MaxChange float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
    PrecisionAsDisplayed bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
    EnableCalculationChain bool `json:"EnableCalculationChain,omitempty" xml:"EnableCalculationChain"`
    PreservePaddingSpaces bool `json:"PreservePaddingSpaces,omitempty" xml:"PreservePaddingSpaces"`
}


type GlobalizationSettings struct {
 
    ChartSettings *ChartGlobalizationSettings `json:"ChartSettings,omitempty" xml:"ChartSettings"`
    PivotSettings *PivotGlobalizationSettings `json:"PivotSettings,omitempty" xml:"PivotSettings"`
    ListSeparator string `json:"ListSeparator,omitempty" xml:"ListSeparator"`
    RowSeparatorOfFormulaArray string `json:"RowSeparatorOfFormulaArray,omitempty" xml:"RowSeparatorOfFormulaArray"`
    ColumnSeparatorOfFormulaArray string `json:"ColumnSeparatorOfFormulaArray,omitempty" xml:"ColumnSeparatorOfFormulaArray"`
}


type HorizontalPageBreak struct {
 
    Row int64 `json:"Row,omitempty" xml:"Row"`
    EndColumn int64 `json:"EndColumn,omitempty" xml:"EndColumn"`
    StartColumn int64 `json:"StartColumn,omitempty" xml:"StartColumn"`
}


type HorizontalPageBreaks struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
}


type Hyperlink struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Address string `json:"Address,omitempty" xml:"Address"`
    Area *CellArea `json:"Area,omitempty" xml:"Area"`
    ScreenTip string `json:"ScreenTip,omitempty" xml:"ScreenTip"`
    TextToDisplay string `json:"TextToDisplay,omitempty" xml:"TextToDisplay"`
    LinkType string `json:"LinkType,omitempty" xml:"LinkType"`
}


type Hyperlinks struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Count int64 `json:"Count,omitempty" xml:"Count"`
    HyperlinkList []LinkElement `json:"HyperlinkList,omitempty" xml:"HyperlinkList"`
}


type IconFilter struct {
 
    IconId int64 `json:"IconId,omitempty" xml:"IconId"`
    IconSetType string `json:"IconSetType,omitempty" xml:"IconSetType"`
}


type IconSet struct {
 
    CfIcons []ConditionalFormattingIcon `json:"CfIcons,omitempty" xml:"CfIcons"`
    Cfvos []ConditionalFormattingValue `json:"Cfvos,omitempty" xml:"Cfvos"`
    IsCustom bool `json:"IsCustom,omitempty" xml:"IsCustom"`
    Reverse bool `json:"Reverse,omitempty" xml:"Reverse"`
    ShowValue bool `json:"ShowValue,omitempty" xml:"ShowValue"`
    IconSetType string `json:"IconSetType,omitempty" xml:"IconSetType"`
}


type Link struct {
 
    Href string `json:"Href,omitempty" xml:"Href"`
    Rel string `json:"Rel,omitempty" xml:"Rel"`
    Title string `json:"Title,omitempty" xml:"Title"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type LinkElement struct {
 
    Link *Link `json:"link,omitempty" xml:"link"`
}


type LoadOptions struct {
 
    ConvertNumericData string `json:"ConvertNumericData,omitempty" xml:"ConvertNumericData"`
    InterruptMonitor string `json:"InterruptMonitor,omitempty" xml:"InterruptMonitor"`
    LanguageCode string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
    LoadDataOptions string `json:"LoadDataOptions,omitempty" xml:"LoadDataOptions"`
    LoadFormat string `json:"LoadFormat,omitempty" xml:"LoadFormat"`
    OnlyLoadDocumentProperties string `json:"OnlyLoadDocumentProperties,omitempty" xml:"OnlyLoadDocumentProperties"`
    ParsingFormulaOnOpen string `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
    Password string `json:"Password,omitempty" xml:"Password"`
    Region string `json:"Region,omitempty" xml:"Region"`
    StandardFont string `json:"StandardFont,omitempty" xml:"StandardFont"`
    StandardFontSize float64 `json:"StandardFontSize,omitempty" xml:"StandardFontSize"`
}


type MergedCell struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    EndColumn int64 `json:"EndColumn,omitempty" xml:"EndColumn"`
    EndRow int64 `json:"EndRow,omitempty" xml:"EndRow"`
    StartColumn int64 `json:"StartColumn,omitempty" xml:"StartColumn"`
    StartRow int64 `json:"StartRow,omitempty" xml:"StartRow"`
}


type MergedCells struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Count int64 `json:"Count,omitempty" xml:"Count"`
    MergedCellList []LinkElement `json:"MergedCellList,omitempty" xml:"MergedCellList"`
}


type MultipleFilter interface {
}


type MultipleFilters struct {
 
    MatchBlank bool `json:"MatchBlank,omitempty" xml:"MatchBlank"`
    MultipleFilterList []MultipleFilter `json:"MultipleFilterList,omitempty" xml:"MultipleFilterList"`
}


type Name struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Comment string `json:"Comment,omitempty" xml:"Comment"`
    WorksheetIndex int64 `json:"WorksheetIndex,omitempty" xml:"WorksheetIndex"`
    IsReferred bool `json:"IsReferred,omitempty" xml:"IsReferred"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    R1C1RefersTo string `json:"R1C1RefersTo,omitempty" xml:"R1C1RefersTo"`
    RefersTo string `json:"RefersTo,omitempty" xml:"RefersTo"`
    Text string `json:"Text,omitempty" xml:"Text"`
}


type Names struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Count int64 `json:"Count,omitempty" xml:"Count"`
    NameList []LinkElement `json:"NameList,omitempty" xml:"NameList"`
}


type NegativeBarFormat struct {
 
    BorderColor *Color `json:"BorderColor,omitempty" xml:"BorderColor"`
    BorderColorType string `json:"BorderColorType,omitempty" xml:"BorderColorType"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
    ColorType string `json:"ColorType,omitempty" xml:"ColorType"`
}


type PageSection struct {
 
    Section int64 `json:"Section,omitempty" xml:"Section"`
    Context string `json:"Context,omitempty" xml:"Context"`
    Picture string `json:"Picture,omitempty" xml:"Picture"`
    FisrtPageContext string `json:"FisrtPageContext,omitempty" xml:"FisrtPageContext"`
    EvenPageContext string `json:"EvenPageContext,omitempty" xml:"EvenPageContext"`
}


type PageSetup struct {
 
    BlackAndWhite bool `json:"BlackAndWhite,omitempty" xml:"BlackAndWhite"`
    BottomMargin float64 `json:"BottomMargin,omitempty" xml:"BottomMargin"`
    CenterHorizontally bool `json:"CenterHorizontally,omitempty" xml:"CenterHorizontally"`
    CenterVertically bool `json:"CenterVertically,omitempty" xml:"CenterVertically"`
    FirstPageNumber int64 `json:"FirstPageNumber,omitempty" xml:"FirstPageNumber"`
    FitToPagesTall int64 `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall"`
    FitToPagesWide int64 `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide"`
    FooterMargin float64 `json:"FooterMargin,omitempty" xml:"FooterMargin"`
    HeaderMargin float64 `json:"HeaderMargin,omitempty" xml:"HeaderMargin"`
    IsAutoFirstPageNumber bool `json:"IsAutoFirstPageNumber,omitempty" xml:"IsAutoFirstPageNumber"`
    IsHFAlignMargins bool `json:"IsHFAlignMargins,omitempty" xml:"IsHFAlignMargins"`
    IsHFDiffFirst bool `json:"IsHFDiffFirst,omitempty" xml:"IsHFDiffFirst"`
    IsHFDiffOddEven bool `json:"IsHFDiffOddEven,omitempty" xml:"IsHFDiffOddEven"`
    IsHFScaleWithDoc bool `json:"IsHFScaleWithDoc,omitempty" xml:"IsHFScaleWithDoc"`
    IsPercentScale bool `json:"IsPercentScale,omitempty" xml:"IsPercentScale"`
    LeftMargin float64 `json:"LeftMargin,omitempty" xml:"LeftMargin"`
    Order string `json:"Order,omitempty" xml:"Order"`
    Orientation string `json:"Orientation,omitempty" xml:"Orientation"`
    PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
    PrintArea string `json:"PrintArea,omitempty" xml:"PrintArea"`
    PrintComments string `json:"PrintComments,omitempty" xml:"PrintComments"`
    PrintCopies int64 `json:"PrintCopies,omitempty" xml:"PrintCopies"`
    PrintDraft bool `json:"PrintDraft,omitempty" xml:"PrintDraft"`
    PrintErrors string `json:"PrintErrors,omitempty" xml:"PrintErrors"`
    PrintGridlines bool `json:"PrintGridlines,omitempty" xml:"PrintGridlines"`
    PrintHeadings bool `json:"PrintHeadings,omitempty" xml:"PrintHeadings"`
    PrintQuality int64 `json:"PrintQuality,omitempty" xml:"PrintQuality"`
    PrintTitleColumns string `json:"PrintTitleColumns,omitempty" xml:"PrintTitleColumns"`
    PrintTitleRows string `json:"PrintTitleRows,omitempty" xml:"PrintTitleRows"`
    RightMargin float64 `json:"RightMargin,omitempty" xml:"RightMargin"`
    TopMargin float64 `json:"TopMargin,omitempty" xml:"TopMargin"`
    Zoom int64 `json:"Zoom,omitempty" xml:"Zoom"`
    Header []PageSection `json:"Header,omitempty" xml:"Header"`
    Footer []PageSection `json:"Footer,omitempty" xml:"Footer"`
}


type PasteOptions struct {
 
    OnlyVisibleCells bool `json:"OnlyVisibleCells,omitempty" xml:"OnlyVisibleCells"`
    PasteType string `json:"PasteType,omitempty" xml:"PasteType"`
    SkipBlanks bool `json:"SkipBlanks,omitempty" xml:"SkipBlanks"`
    Transpose bool `json:"Transpose,omitempty" xml:"Transpose"`
}


type Protection struct {
 
    AllowDeletingColumn bool `json:"AllowDeletingColumn,omitempty" xml:"AllowDeletingColumn"`
    AllowDeletingRow bool `json:"AllowDeletingRow,omitempty" xml:"AllowDeletingRow"`
    AllowFiltering bool `json:"AllowFiltering,omitempty" xml:"AllowFiltering"`
    AllowFormattingCell bool `json:"AllowFormattingCell,omitempty" xml:"AllowFormattingCell"`
    AllowFormattingColumn bool `json:"AllowFormattingColumn,omitempty" xml:"AllowFormattingColumn"`
    AllowFormattingRow bool `json:"AllowFormattingRow,omitempty" xml:"AllowFormattingRow"`
    AllowInsertingColumn bool `json:"AllowInsertingColumn,omitempty" xml:"AllowInsertingColumn"`
    AllowInsertingHyperlink bool `json:"AllowInsertingHyperlink,omitempty" xml:"AllowInsertingHyperlink"`
    AllowInsertingRow bool `json:"AllowInsertingRow,omitempty" xml:"AllowInsertingRow"`
    AllowSorting bool `json:"AllowSorting,omitempty" xml:"AllowSorting"`
    AllowUsingPivotTable bool `json:"AllowUsingPivotTable,omitempty" xml:"AllowUsingPivotTable"`
    AllowEditingContent bool `json:"AllowEditingContent,omitempty" xml:"AllowEditingContent"`
    AllowEditingObject bool `json:"AllowEditingObject,omitempty" xml:"AllowEditingObject"`
    AllowEditingScenario bool `json:"AllowEditingScenario,omitempty" xml:"AllowEditingScenario"`
    Password string `json:"Password,omitempty" xml:"Password"`
    AllowSelectingLockedCell bool `json:"AllowSelectingLockedCell,omitempty" xml:"AllowSelectingLockedCell"`
    AllowSelectingUnlockedCell bool `json:"AllowSelectingUnlockedCell,omitempty" xml:"AllowSelectingUnlockedCell"`
}


type ProtectSheetParameter struct {
 
    ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
    Password string `json:"Password,omitempty" xml:"Password"`
    AllowEditArea []string `json:"AllowEditArea,omitempty" xml:"AllowEditArea"`
    AllowDeletingColumn string `json:"AllowDeletingColumn,omitempty" xml:"AllowDeletingColumn"`
    AllowDeletingRow string `json:"AllowDeletingRow,omitempty" xml:"AllowDeletingRow"`
    AllowFiltering string `json:"AllowFiltering,omitempty" xml:"AllowFiltering"`
    AllowFormattingCell string `json:"AllowFormattingCell,omitempty" xml:"AllowFormattingCell"`
    AllowFormattingColumn string `json:"AllowFormattingColumn,omitempty" xml:"AllowFormattingColumn"`
    AllowFormattingRow string `json:"AllowFormattingRow,omitempty" xml:"AllowFormattingRow"`
    AllowInsertingColumn string `json:"AllowInsertingColumn,omitempty" xml:"AllowInsertingColumn"`
    AllowInsertingHyperlink string `json:"AllowInsertingHyperlink,omitempty" xml:"AllowInsertingHyperlink"`
    AllowInsertingRow string `json:"AllowInsertingRow,omitempty" xml:"AllowInsertingRow"`
    AllowSelectingLockedCell string `json:"AllowSelectingLockedCell,omitempty" xml:"AllowSelectingLockedCell"`
    AllowSelectingUnlockedCell string `json:"AllowSelectingUnlockedCell,omitempty" xml:"AllowSelectingUnlockedCell"`
    AllowSorting string `json:"AllowSorting,omitempty" xml:"AllowSorting"`
    AllowUsingPivotTable string `json:"AllowUsingPivotTable,omitempty" xml:"AllowUsingPivotTable"`
}


type Ranges struct {
 
    RangeList []Range `json:"RangeList,omitempty" xml:"RangeList"`
}


type Row struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    GroupLevel int64 `json:"GroupLevel,omitempty" xml:"GroupLevel"`
    Height float64 `json:"Height,omitempty" xml:"Height"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    IsBlank bool `json:"IsBlank,omitempty" xml:"IsBlank"`
    IsHeightMatched bool `json:"IsHeightMatched,omitempty" xml:"IsHeightMatched"`
    IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
    Style *LinkElement `json:"Style,omitempty" xml:"Style"`
}


type Rows struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    MaxRow int64 `json:"MaxRow,omitempty" xml:"MaxRow"`
    RowsCount int64 `json:"RowsCount,omitempty" xml:"RowsCount"`
    RowsList []LinkElement `json:"RowsList,omitempty" xml:"RowsList"`
}


type SingleValue struct {
 
    Value string `json:"Value,omitempty" xml:"Value"`
    ValueType string `json:"ValueType,omitempty" xml:"ValueType"`
}


type SortKey struct {
 
    Key int64 `json:"Key,omitempty" xml:"Key"`
    SortOrder string `json:"SortOrder,omitempty" xml:"SortOrder"`
    CustomList []string `json:"CustomList,omitempty" xml:"CustomList"`
    Order string `json:"Order,omitempty" xml:"Order"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type SplitResult struct {
 
    Documents []CellsCloudFileInfo `json:"Documents,omitempty" xml:"Documents"`
}


type Style struct {
 
    Font *Font `json:"Font,omitempty" xml:"Font"`
    Name string `json:"Name,omitempty" xml:"Name"`
    CultureCustom string `json:"CultureCustom,omitempty" xml:"CultureCustom"`
    Custom string `json:"Custom,omitempty" xml:"Custom"`
    BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
    ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
    IsFormulaHidden bool `json:"IsFormulaHidden,omitempty" xml:"IsFormulaHidden"`
    IsDateTime bool `json:"IsDateTime,omitempty" xml:"IsDateTime"`
    IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
    IsGradient bool `json:"IsGradient,omitempty" xml:"IsGradient"`
    IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
    IsPercent bool `json:"IsPercent,omitempty" xml:"IsPercent"`
    ShrinkToFit bool `json:"ShrinkToFit,omitempty" xml:"ShrinkToFit"`
    IndentLevel int64 `json:"IndentLevel,omitempty" xml:"IndentLevel"`
    Number int64 `json:"Number,omitempty" xml:"Number"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
    TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
    VerticalAlignment string `json:"VerticalAlignment,omitempty" xml:"VerticalAlignment"`
    HorizontalAlignment string `json:"HorizontalAlignment,omitempty" xml:"HorizontalAlignment"`
    BorderCollection []Border `json:"BorderCollection,omitempty" xml:"BorderCollection"`
    BackgroundThemeColor *ThemeColor `json:"BackgroundThemeColor,omitempty" xml:"BackgroundThemeColor"`
    ForegroundThemeColor *ThemeColor `json:"ForegroundThemeColor,omitempty" xml:"ForegroundThemeColor"`
}


type StyleFormatCondition struct {
     
        Priority int64 `json:"Priority,omitempty" xml:"Priority"`
        Type_ string `json:"Type,omitempty" xml:"Type"`
        StopIfTrue bool `json:"StopIfTrue,omitempty" xml:"StopIfTrue"`
        AboveAverage *AboveAverage `json:"AboveAverage,omitempty" xml:"AboveAverage"`
        ColorScale *ColorScale `json:"ColorScale,omitempty" xml:"ColorScale"`
        DataBar *DataBar `json:"DataBar,omitempty" xml:"DataBar"`
        Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
        Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
        IconSet *IconSet `json:"IconSet,omitempty" xml:"IconSet"`
        Operator string `json:"Operator,omitempty" xml:"Operator"`
        Style *Style `json:"Style,omitempty" xml:"Style"`
        Text string `json:"Text,omitempty" xml:"Text"`
        TimePeriod string `json:"TimePeriod,omitempty" xml:"TimePeriod"`
        Top10 *Top10 `json:"Top10,omitempty" xml:"Top10"`
 
}


type Styles struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    StyleList []Style `json:"StyleList,omitempty" xml:"StyleList"`
}


type TextFormatCondition struct {
     
        Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
        Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
        Operator string `json:"Operator,omitempty" xml:"Operator"`
 
    Text string `json:"Text,omitempty" xml:"Text"`
}


type TextItems struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    TextItemList []TextItem `json:"TextItemList,omitempty" xml:"TextItemList"`
}


type TextOptions struct {
     
        Color *Color `json:"Color,omitempty" xml:"Color"`
        DoubleSize float64 `json:"DoubleSize,omitempty" xml:"DoubleSize"`
        IsBold bool `json:"IsBold,omitempty" xml:"IsBold"`
        IsItalic bool `json:"IsItalic,omitempty" xml:"IsItalic"`
        IsStrikeout bool `json:"IsStrikeout,omitempty" xml:"IsStrikeout"`
        IsSubscript bool `json:"IsSubscript,omitempty" xml:"IsSubscript"`
        IsSuperscript bool `json:"IsSuperscript,omitempty" xml:"IsSuperscript"`
        Name string `json:"Name,omitempty" xml:"Name"`
        Size int64 `json:"Size,omitempty" xml:"Size"`
        Underline string `json:"Underline,omitempty" xml:"Underline"`
 
    Fill *FillFormat `json:"Fill,omitempty" xml:"Fill"`
    Kerning float64 `json:"Kerning,omitempty" xml:"Kerning"`
    Outline *LineFormat `json:"Outline,omitempty" xml:"Outline"`
    Shadow *ShadowEffect `json:"Shadow,omitempty" xml:"Shadow"`
    Spacing float64 `json:"Spacing,omitempty" xml:"Spacing"`
    UnderlineColor *CellsColor `json:"UnderlineColor,omitempty" xml:"UnderlineColor"`
}


type ThemeColor struct {
 
    ColorType string `json:"ColorType,omitempty" xml:"ColorType"`
    Tint float64 `json:"Tint,omitempty" xml:"Tint"`
}


type TimePeriodFormatCondition struct {
 
    TimePeriod string `json:"TimePeriod,omitempty" xml:"TimePeriod"`
}


type Top10 struct {
 
    IsBottom bool `json:"IsBottom,omitempty" xml:"IsBottom"`
    IsPercent bool `json:"IsPercent,omitempty" xml:"IsPercent"`
    Rank int64 `json:"Rank,omitempty" xml:"Rank"`
}


type Top10Filter struct {
 
    FieldIndex int64 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
    Criteria string `json:"Criteria,omitempty" xml:"Criteria"`
    IsPercent bool `json:"IsPercent,omitempty" xml:"IsPercent"`
    IsTop bool `json:"IsTop,omitempty" xml:"IsTop"`
    Items int64 `json:"Items,omitempty" xml:"Items"`
}


type Validation struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AlertStyle string `json:"AlertStyle,omitempty" xml:"AlertStyle"`
    AreaList []CellArea `json:"AreaList,omitempty" xml:"AreaList"`
    ErrorMessage string `json:"ErrorMessage,omitempty" xml:"ErrorMessage"`
    ErrorTitle string `json:"ErrorTitle,omitempty" xml:"ErrorTitle"`
    Formula1 string `json:"Formula1,omitempty" xml:"Formula1"`
    Formula2 string `json:"Formula2,omitempty" xml:"Formula2"`
    IgnoreBlank bool `json:"IgnoreBlank,omitempty" xml:"IgnoreBlank"`
    InCellDropDown bool `json:"InCellDropDown,omitempty" xml:"InCellDropDown"`
    InputMessage string `json:"InputMessage,omitempty" xml:"InputMessage"`
    InputTitle string `json:"InputTitle,omitempty" xml:"InputTitle"`
    Operator string `json:"Operator,omitempty" xml:"Operator"`
    ShowError bool `json:"ShowError,omitempty" xml:"ShowError"`
    ShowInput bool `json:"ShowInput,omitempty" xml:"ShowInput"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Value1 string `json:"Value1,omitempty" xml:"Value1"`
    Value2 string `json:"Value2,omitempty" xml:"Value2"`
}


type Validations struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Count int64 `json:"Count,omitempty" xml:"Count"`
    ValidationList []LinkElement `json:"ValidationList,omitempty" xml:"ValidationList"`
}


type VerticalPageBreak struct {
 
    Column int64 `json:"Column,omitempty" xml:"Column"`
    EndRow int64 `json:"EndRow,omitempty" xml:"EndRow"`
    StartRow int64 `json:"StartRow,omitempty" xml:"StartRow"`
}


type VerticalPageBreaks struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
}


type Workbook struct {
 
    FileName string `json:"FileName,omitempty" xml:"FileName"`
    Links []Link `json:"Links,omitempty" xml:"Links"`
    Worksheets *LinkElement `json:"Worksheets,omitempty" xml:"Worksheets"`
    DefaultStyle *LinkElement `json:"DefaultStyle,omitempty" xml:"DefaultStyle"`
    DocumentProperties *LinkElement `json:"DocumentProperties,omitempty" xml:"DocumentProperties"`
    Names *LinkElement `json:"Names,omitempty" xml:"Names"`
    Settings *LinkElement `json:"Settings,omitempty" xml:"Settings"`
    IsWriteProtected string `json:"IsWriteProtected,omitempty" xml:"IsWriteProtected"`
    IsProtected string `json:"IsProtected,omitempty" xml:"IsProtected"`
    IsEncryption string `json:"IsEncryption,omitempty" xml:"IsEncryption"`
    Password string `json:"Password,omitempty" xml:"Password"`
}


type WorkbookSettings struct {
 
    AutoCompressPictures bool `json:"AutoCompressPictures,omitempty" xml:"AutoCompressPictures"`
    AutoRecover bool `json:"AutoRecover,omitempty" xml:"AutoRecover"`
    BuildVersion string `json:"BuildVersion,omitempty" xml:"BuildVersion"`
    CalcMode string `json:"CalcMode,omitempty" xml:"CalcMode"`
    CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
    CheckComptiliblity bool `json:"CheckComptiliblity,omitempty" xml:"CheckComptiliblity"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    CrashSave bool `json:"CrashSave,omitempty" xml:"CrashSave"`
    CreateCalcChain bool `json:"CreateCalcChain,omitempty" xml:"CreateCalcChain"`
    DataExtractLoad bool `json:"DataExtractLoad,omitempty" xml:"DataExtractLoad"`
    Date1904 bool `json:"Date1904,omitempty" xml:"Date1904"`
    DisplayDrawingObjects string `json:"DisplayDrawingObjects,omitempty" xml:"DisplayDrawingObjects"`
    EnableMacros bool `json:"EnableMacros,omitempty" xml:"EnableMacros"`
    FirstVisibleTab int64 `json:"FirstVisibleTab,omitempty" xml:"FirstVisibleTab"`
    HidePivotFieldList bool `json:"HidePivotFieldList,omitempty" xml:"HidePivotFieldList"`
    IsDefaultEncrypted bool `json:"IsDefaultEncrypted,omitempty" xml:"IsDefaultEncrypted"`
    IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
    IsHScrollBarVisible bool `json:"IsHScrollBarVisible,omitempty" xml:"IsHScrollBarVisible"`
    IsMinimized bool `json:"IsMinimized,omitempty" xml:"IsMinimized"`
    IsVScrollBarVisible bool `json:"IsVScrollBarVisible,omitempty" xml:"IsVScrollBarVisible"`
    Iteration bool `json:"Iteration,omitempty" xml:"Iteration"`
    LanguageCode string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
    MaxChange float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
    MaxIteration int64 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
    MemorySetting string `json:"MemorySetting,omitempty" xml:"MemorySetting"`
    NumberDecimalSeparator string `json:"NumberDecimalSeparator,omitempty" xml:"NumberDecimalSeparator"`
    NumberGroupSeparator string `json:"NumberGroupSeparator,omitempty" xml:"NumberGroupSeparator"`
    ParsingFormulaOnOpen bool `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
    PrecisionAsDisplayed bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
    RecalculateBeforeSave bool `json:"RecalculateBeforeSave,omitempty" xml:"RecalculateBeforeSave"`
    ReCalculateOnOpen bool `json:"ReCalculateOnOpen,omitempty" xml:"ReCalculateOnOpen"`
    RecommendReadOnly bool `json:"RecommendReadOnly,omitempty" xml:"RecommendReadOnly"`
    Region string `json:"Region,omitempty" xml:"Region"`
    RemovePersonalInformation bool `json:"RemovePersonalInformation,omitempty" xml:"RemovePersonalInformation"`
    RepairLoad bool `json:"RepairLoad,omitempty" xml:"RepairLoad"`
    Shared bool `json:"Shared,omitempty" xml:"Shared"`
    SheetTabBarWidth int64 `json:"SheetTabBarWidth,omitempty" xml:"SheetTabBarWidth"`
    ShowTabs bool `json:"ShowTabs,omitempty" xml:"ShowTabs"`
    UpdateAdjacentCellsBorder bool `json:"UpdateAdjacentCellsBorder,omitempty" xml:"UpdateAdjacentCellsBorder"`
    UpdateLinksType string `json:"UpdateLinksType,omitempty" xml:"UpdateLinksType"`
    WindowHeight float64 `json:"WindowHeight,omitempty" xml:"WindowHeight"`
    WindowLeft float64 `json:"WindowLeft,omitempty" xml:"WindowLeft"`
    WindowTop float64 `json:"WindowTop,omitempty" xml:"WindowTop"`
    WindowWidth float64 `json:"WindowWidth,omitempty" xml:"WindowWidth"`
    Author string `json:"Author,omitempty" xml:"Author"`
    CheckCustomNumberFormat bool `json:"CheckCustomNumberFormat,omitempty" xml:"CheckCustomNumberFormat"`
    ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
    GlobalizationSettings *GlobalizationSettings `json:"GlobalizationSettings,omitempty" xml:"GlobalizationSettings"`
    Password string `json:"Password,omitempty" xml:"Password"`
    WriteProtection *WriteProtection `json:"WriteProtection,omitempty" xml:"WriteProtection"`
    IsEncrypted bool `json:"IsEncrypted,omitempty" xml:"IsEncrypted"`
    IsProtected bool `json:"IsProtected,omitempty" xml:"IsProtected"`
    MaxRow int64 `json:"MaxRow,omitempty" xml:"MaxRow"`
    MaxColumn int64 `json:"MaxColumn,omitempty" xml:"MaxColumn"`
    SignificantDigits int64 `json:"SignificantDigits,omitempty" xml:"SignificantDigits"`
    CheckCompatibility bool `json:"CheckCompatibility,omitempty" xml:"CheckCompatibility"`
    PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
    MaxRowsOfSharedFormula int64 `json:"MaxRowsOfSharedFormula,omitempty" xml:"MaxRowsOfSharedFormula"`
    Compliance string `json:"Compliance,omitempty" xml:"Compliance"`
    QuotePrefixToStyle bool `json:"QuotePrefixToStyle,omitempty" xml:"QuotePrefixToStyle"`
    FormulaSettings *FormulaSettings `json:"FormulaSettings,omitempty" xml:"FormulaSettings"`
    ForceFullCalculate bool `json:"ForceFullCalculate,omitempty" xml:"ForceFullCalculate"`
}


type Worksheet struct {
 
    Links []Link `json:"Links,omitempty" xml:"Links"`
    DisplayRightToLeft bool `json:"DisplayRightToLeft,omitempty" xml:"DisplayRightToLeft"`
    DisplayZeros bool `json:"DisplayZeros,omitempty" xml:"DisplayZeros"`
    FirstVisibleColumn int64 `json:"FirstVisibleColumn,omitempty" xml:"FirstVisibleColumn"`
    FirstVisibleRow int64 `json:"FirstVisibleRow,omitempty" xml:"FirstVisibleRow"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    IsGridlinesVisible bool `json:"IsGridlinesVisible,omitempty" xml:"IsGridlinesVisible"`
    IsOutlineShown bool `json:"IsOutlineShown,omitempty" xml:"IsOutlineShown"`
    IsPageBreakPreview bool `json:"IsPageBreakPreview,omitempty" xml:"IsPageBreakPreview"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    IsProtected bool `json:"IsProtected,omitempty" xml:"IsProtected"`
    IsRowColumnHeadersVisible bool `json:"IsRowColumnHeadersVisible,omitempty" xml:"IsRowColumnHeadersVisible"`
    IsRulerVisible bool `json:"IsRulerVisible,omitempty" xml:"IsRulerVisible"`
    IsSelected bool `json:"IsSelected,omitempty" xml:"IsSelected"`
    TabColor *Color `json:"TabColor,omitempty" xml:"TabColor"`
    TransitionEntry bool `json:"TransitionEntry,omitempty" xml:"TransitionEntry"`
    TransitionEvaluation bool `json:"TransitionEvaluation,omitempty" xml:"TransitionEvaluation"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    ViewType string `json:"ViewType,omitempty" xml:"ViewType"`
    VisibilityType string `json:"VisibilityType,omitempty" xml:"VisibilityType"`
    Zoom int64 `json:"Zoom,omitempty" xml:"Zoom"`
    Cells *LinkElement `json:"Cells,omitempty" xml:"Cells"`
    Charts *LinkElement `json:"Charts,omitempty" xml:"Charts"`
    AutoShapes *LinkElement `json:"AutoShapes,omitempty" xml:"AutoShapes"`
    OleObjects *LinkElement `json:"OleObjects,omitempty" xml:"OleObjects"`
    Comments *LinkElement `json:"Comments,omitempty" xml:"Comments"`
    Pictures *LinkElement `json:"Pictures,omitempty" xml:"Pictures"`
    MergedCells *LinkElement `json:"MergedCells,omitempty" xml:"MergedCells"`
    Validations *LinkElement `json:"Validations,omitempty" xml:"Validations"`
    ConditionalFormattings *LinkElement `json:"ConditionalFormattings,omitempty" xml:"ConditionalFormattings"`
    Hyperlinks *LinkElement `json:"Hyperlinks,omitempty" xml:"Hyperlinks"`
}


type Worksheets struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    WorksheetList []LinkElement `json:"WorksheetList,omitempty" xml:"WorksheetList"`
}


type WriteProtection struct {
 
    Author string `json:"Author,omitempty" xml:"Author"`
    RecommendReadOnly bool `json:"RecommendReadOnly,omitempty" xml:"RecommendReadOnly"`
    IsWriteProtected bool `json:"IsWriteProtected,omitempty" xml:"IsWriteProtected"`
    Password string `json:"Password,omitempty" xml:"Password"`
}


type XmlDataBinding struct {
 
    Url string `json:"Url,omitempty" xml:"Url"`
}


type XmlMap struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    RootElementName string `json:"RootElementName,omitempty" xml:"RootElementName"`
    DataBinding *XmlDataBinding `json:"DataBinding,omitempty" xml:"DataBinding"`
}


type CellsObjectOperateTaskParameter struct {
     
 
    OperateObject *OperateObject `json:"OperateObject,omitempty" xml:"OperateObject"`
    OperateParameter *OperateParameter `json:"OperateParameter,omitempty" xml:"OperateParameter"`
    DestinationDataSource *DataSource `json:"DestinationDataSource,omitempty" xml:"DestinationDataSource"`
    DestinationWorkbook *FileSource `json:"DestinationWorkbook,omitempty" xml:"DestinationWorkbook"`
}


type ConvertTaskParameter struct {
     
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    DestinationFile string `json:"DestinationFile,omitempty" xml:"DestinationFile"`
    Region string `json:"Region,omitempty" xml:"Region"`
    SaveOptions *SaveOptions `json:"SaveOptions,omitempty" xml:"SaveOptions"`
}


type ConvertWorksheetTaskParameter struct {
     
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    Sheet string `json:"Sheet,omitempty" xml:"Sheet"`
    TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
    Target *FileSource `json:"Target,omitempty" xml:"Target"`
    Format string `json:"Format,omitempty" xml:"Format"`
    Area string `json:"Area,omitempty" xml:"Area"`
    PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
}


type ImportDataTaskParameter struct {
     
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    ImportOption *ImportOption `json:"ImportOption,omitempty" xml:"ImportOption"`
    TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
    DestinationWorkbook *FileSource `json:"DestinationWorkbook,omitempty" xml:"DestinationWorkbook"`
}


type ResultDestination struct {
 
    DestinationType string `json:"DestinationType,omitempty" xml:"DestinationType"`
    InputFile string `json:"InputFile,omitempty" xml:"InputFile"`
    OutputFile string `json:"OutputFile,omitempty" xml:"OutputFile"`
}


type SaveFilesToCloudResult struct {
     
        Description string `json:"Description,omitempty" xml:"Description"`
        OutFileList []DataSource `json:"OutFileList,omitempty" xml:"OutFileList"`
 
    SavedFiles []Link `json:"SavedFiles,omitempty" xml:"SavedFiles"`
}


type SaveResultTaskParameter struct {
     
 
    ResultSource string `json:"ResultSource,omitempty" xml:"ResultSource"`
    ResultDestination *ResultDestination `json:"ResultDestination,omitempty" xml:"ResultDestination"`
}


type SmartMarkerTaskParameter struct {
     
 
    SourceWorkbook *FileSource `json:"SourceWorkbook,omitempty" xml:"SourceWorkbook"`
    DestinationWorkbook *FileSource `json:"DestinationWorkbook,omitempty" xml:"DestinationWorkbook"`
    XmlFile *FileSource `json:"xmlFile,omitempty" xml:"xmlFile"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
    XMLFileDataSource *DataSource `json:"XMLFileDataSource,omitempty" xml:"XMLFileDataSource"`
}


type SplitWorkbookTaskParameter struct {
     
 
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    DestinationFilePosition *FileSource `json:"DestinationFilePosition,omitempty" xml:"DestinationFilePosition"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
    DestinationFileFormat string `json:"DestinationFileFormat,omitempty" xml:"DestinationFileFormat"`
    SplitNameRule string `json:"SplitNameRule,omitempty" xml:"SplitNameRule"`
    VerticalResolution int64 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
    HorizontalResolution int64 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
}


type TaskData struct {
 
    Tasks []TaskDescription `json:"Tasks,omitempty" xml:"Tasks"`
}


type TaskDescription struct {
 
    TaskType string `json:"TaskType,omitempty" xml:"TaskType"`
    TaskParameter *TaskParameter `json:"TaskParameter,omitempty" xml:"TaskParameter"`
}


type TaskParameter struct {
 
}


type TaskResultParameter struct {
     
 
}


type TaskRunResult struct {
 
    Description string `json:"Description,omitempty" xml:"Description"`
    OutFileList []DataSource `json:"OutFileList,omitempty" xml:"OutFileList"`
}


type ChartOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    ChartIndex int64 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
    ChartType string `json:"ChartType,omitempty" xml:"ChartType"`
    UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
    UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
    LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
    LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
    Area string `json:"Area,omitempty" xml:"Area"`
    IsVertical bool `json:"IsVertical,omitempty" xml:"IsVertical"`
    CategoryData string `json:"CategoryData,omitempty" xml:"CategoryData"`
    IsAutoGetSerialName bool `json:"IsAutoGetSerialName,omitempty" xml:"IsAutoGetSerialName"`
    Title string `json:"Title,omitempty" xml:"Title"`
}


type ListObjectOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    ListObject *ListObject `json:"ListObject,omitempty" xml:"ListObject"`
}


type OperateObject struct {
 
    OperateObjectType string `json:"OperateObjectType,omitempty" xml:"OperateObjectType"`
    Position *OperateObjectPosition `json:"Position,omitempty" xml:"Position"`
}


type OperateObjectPosition struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
    SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
    ChartIndex int64 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
    ShapeIndex int64 `json:"ShapeIndex,omitempty" xml:"ShapeIndex"`
    CellName string `json:"CellName,omitempty" xml:"CellName"`
    ListObjectIndex int64 `json:"ListObjectIndex,omitempty" xml:"ListObjectIndex"`
}


type OperateParameter struct {
 
    OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
}


type PageBreakOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    PageBreakType string `json:"PageBreakType,omitempty" xml:"PageBreakType"`
    Index int64 `json:"Index,omitempty" xml:"Index"`
    Row int64 `json:"Row,omitempty" xml:"Row"`
    Column int64 `json:"Column,omitempty" xml:"Column"`
    StartIndex int64 `json:"StartIndex,omitempty" xml:"StartIndex"`
    EndIndex int64 `json:"EndIndex,omitempty" xml:"EndIndex"`
}


type PageSetupOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    PageSetup *PageSetup `json:"PageSetup,omitempty" xml:"PageSetup"`
}


type PivotTableOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    SourceData string `json:"SourceData,omitempty" xml:"SourceData"`
    DestCellName string `json:"DestCellName,omitempty" xml:"DestCellName"`
    TableName string `json:"TableName,omitempty" xml:"TableName"`
    UseSameSource bool `json:"UseSameSource,omitempty" xml:"UseSameSource"`
    PivotTableIndex int64 `json:"PivotTableIndex,omitempty" xml:"PivotTableIndex"`
    PivotFieldRows []int64 `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
    PivotFieldColumns []int64 `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
    PivotFieldData []int64 `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
}


type ShapeOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    Shape *Shape `json:"Shape,omitempty" xml:"Shape"`
}


type WorkbookOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
}


type WorkbookSettingsOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    WorkbookSettings *WorkbookSettings `json:"WorkbookSettings,omitempty" xml:"WorkbookSettings"`
}


type WorksheetOperateParameter struct {
     
        OperateType string `json:"OperateType,omitempty" xml:"OperateType"`
 
    Name string `json:"Name,omitempty" xml:"Name"`
    SheetType string `json:"SheetType,omitempty" xml:"SheetType"`
    NewName string `json:"NewName,omitempty" xml:"NewName"`
    MovingRequest *WorksheetMovingRequest `json:"MovingRequest,omitempty" xml:"MovingRequest"`
}


type ListColumn struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Range_ *Range `json:"Range,omitempty" xml:"Range"`
    TotalsCalculation string `json:"TotalsCalculation,omitempty" xml:"TotalsCalculation"`
    Formula string `json:"Formula,omitempty" xml:"Formula"`
    TotalsRowLabel string `json:"TotalsRowLabel,omitempty" xml:"TotalsRowLabel"`
}


type ListObject struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
    DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
    StartColumn int64 `json:"StartColumn,omitempty" xml:"StartColumn"`
    StartRow int64 `json:"StartRow,omitempty" xml:"StartRow"`
    EndColumn int64 `json:"EndColumn,omitempty" xml:"EndColumn"`
    EndRow int64 `json:"EndRow,omitempty" xml:"EndRow"`
    ListColumns []ListColumn `json:"ListColumns,omitempty" xml:"ListColumns"`
    ShowHeaderRow bool `json:"ShowHeaderRow,omitempty" xml:"ShowHeaderRow"`
    ShowTableStyleColumnStripes bool `json:"ShowTableStyleColumnStripes,omitempty" xml:"ShowTableStyleColumnStripes"`
    ShowTableStyleFirstColumn bool `json:"ShowTableStyleFirstColumn,omitempty" xml:"ShowTableStyleFirstColumn"`
    ShowTableStyleLastColumn bool `json:"ShowTableStyleLastColumn,omitempty" xml:"ShowTableStyleLastColumn"`
    ShowTableStyleRowStripes bool `json:"ShowTableStyleRowStripes,omitempty" xml:"ShowTableStyleRowStripes"`
    ShowTotals bool `json:"ShowTotals,omitempty" xml:"ShowTotals"`
    TableStyleName string `json:"TableStyleName,omitempty" xml:"TableStyleName"`
    TableStyleType string `json:"TableStyleType,omitempty" xml:"TableStyleType"`
    DataRange *Range `json:"DataRange,omitempty" xml:"DataRange"`
    DataSourceType string `json:"DataSourceType,omitempty" xml:"DataSourceType"`
    Comment string `json:"Comment,omitempty" xml:"Comment"`
    XmlMap *XmlMap `json:"XmlMap,omitempty" xml:"XmlMap"`
    AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
    AlternativeDescription string `json:"AlternativeDescription,omitempty" xml:"AlternativeDescription"`
}


type ListObjects struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    ListObjectList []LinkElement `json:"ListObjectList,omitempty" xml:"ListObjectList"`
}


type PivotGlobalizationSettings struct {
 
}


type ArcShapeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *ArcShape `json:"Shape,omitempty" xml:"Shape"`
}


type AutoFilterResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
}


type AutoShapeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    AutoShape *AutoShape `json:"AutoShape,omitempty" xml:"AutoShape"`
}


type AutoShapesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    AutoShapes *AutoShapes `json:"AutoShapes,omitempty" xml:"AutoShapes"`
}


type AxisResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Axis *Axis `json:"Axis,omitempty" xml:"Axis"`
}


type BorderResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Border *Border `json:"Border,omitempty" xml:"Border"`
}


type ButtonResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *Button `json:"Shape,omitempty" xml:"Shape"`
}


type CalculateFormulaResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Result string `json:"Result,omitempty" xml:"Result"`
}


type CellResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Cell *Cell `json:"Cell,omitempty" xml:"Cell"`
}


type CellsDocumentPropertiesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    DocumentProperties *CellsDocumentProperties `json:"DocumentProperties,omitempty" xml:"DocumentProperties"`
}


type CellsDocumentPropertyResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    DocumentProperty *CellsDocumentProperty `json:"DocumentProperty,omitempty" xml:"DocumentProperty"`
}


type CellsDrawingResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *CellsDrawing `json:"Shape,omitempty" xml:"Shape"`
}


type CellsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Cells *Cells `json:"Cells,omitempty" xml:"Cells"`
}


type ChartAreaResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ChartArea *ChartArea `json:"ChartArea,omitempty" xml:"ChartArea"`
}


type ChartDataTableResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ChartDataTable *ChartDataTable `json:"ChartDataTable,omitempty" xml:"ChartDataTable"`
}


type ChartPointResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ChartPoint *ChartPoint `json:"ChartPoint,omitempty" xml:"ChartPoint"`
}


type ChartPointsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ChartPoints *ChartPoints `json:"ChartPoints,omitempty" xml:"ChartPoints"`
}


type ChartResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Chart *Chart `json:"Chart,omitempty" xml:"Chart"`
}


type ChartsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Charts *Charts `json:"Charts,omitempty" xml:"Charts"`
}


type CheckBoxResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *CheckBox `json:"Shape,omitempty" xml:"Shape"`
}


type CheckedExternalReferenceResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ReferenceOtherWorkbook bool `json:"ReferenceOtherWorkbook,omitempty" xml:"ReferenceOtherWorkbook"`
    ReferenceOtherWorksheet bool `json:"ReferenceOtherWorksheet,omitempty" xml:"ReferenceOtherWorksheet"`
    Formulas []string `json:"Formulas,omitempty" xml:"Formulas"`
}


type CheckedFormulaErrorsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    IsFormulasErrors bool `json:"IsFormulasErrors,omitempty" xml:"IsFormulasErrors"`
    FormulasErrors []string `json:"FormulasErrors,omitempty" xml:"FormulasErrors"`
}


type ColumnResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Column *Column `json:"Column,omitempty" xml:"Column"`
}


type ColumnsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Columns *Columns `json:"Columns,omitempty" xml:"Columns"`
}


type ComboBoxResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *ComboBox `json:"Shape,omitempty" xml:"Shape"`
}


type CommentResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Comment *Comment `json:"Comment,omitempty" xml:"Comment"`
}


type CommentShapeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Comment *CommentShape `json:"Comment,omitempty" xml:"Comment"`
}


type CommentsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Comments *Comments `json:"Comments,omitempty" xml:"Comments"`
}


type ConditionalFormattingResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ConditionalFormatting *ConditionalFormatting `json:"ConditionalFormatting,omitempty" xml:"ConditionalFormatting"`
}


type ConditionalFormattingsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ConditionalFormattings *ConditionalFormattings `json:"ConditionalFormattings,omitempty" xml:"ConditionalFormattings"`
}


type DataLabelsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
}


type DisplayUnitLabelResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    DisplayUnitLabel *DisplayUnitLabel `json:"DisplayUnitLabel,omitempty" xml:"DisplayUnitLabel"`
}


type DropBarsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    DropBars *DropBars `json:"DropBars,omitempty" xml:"DropBars"`
}


type ErrorBarResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ErrorBar *ErrorBar `json:"ErrorBar,omitempty" xml:"ErrorBar"`
}


type FillFormatResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    FillFormat *FillFormat `json:"FillFormat,omitempty" xml:"FillFormat"`
}


type FindResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Count int64 `json:"Count,omitempty" xml:"Count"`
    TextItems *TextItems `json:"TextItems,omitempty" xml:"TextItems"`
}


type FloorResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Floor *Floor `json:"Floor,omitempty" xml:"Floor"`
}


type FormResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Form *Form `json:"Form,omitempty" xml:"Form"`
}


type FormsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Forms *Forms `json:"Forms,omitempty" xml:"Forms"`
}


type GroupBoxResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *GroupBox `json:"Shape,omitempty" xml:"Shape"`
}


type HorizontalPageBreakResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    HorizontalPageBreak *HorizontalPageBreak `json:"HorizontalPageBreak,omitempty" xml:"HorizontalPageBreak"`
}


type HorizontalPageBreaksResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    HorizontalPageBreaks *HorizontalPageBreaks `json:"HorizontalPageBreaks,omitempty" xml:"HorizontalPageBreaks"`
}


type HyperlinkResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Hyperlink *Hyperlink `json:"Hyperlink,omitempty" xml:"Hyperlink"`
}


type HyperlinksResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Hyperlinks *Hyperlinks `json:"Hyperlinks,omitempty" xml:"Hyperlinks"`
}


type LabelResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *Label `json:"Shape,omitempty" xml:"Shape"`
}


type LegendEntriesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    LegendEntries *LegendEntries `json:"LegendEntries,omitempty" xml:"LegendEntries"`
}


type LegendEntryResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
}


type LegendResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Legend *Legend `json:"Legend,omitempty" xml:"Legend"`
}


type LineResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Line *Line `json:"Line,omitempty" xml:"Line"`
}


type LineShapeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *LineShape `json:"Shape,omitempty" xml:"Shape"`
}


type ListBoxResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *ListBox `json:"Shape,omitempty" xml:"Shape"`
}


type ListObjectResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ListObject *ListObject `json:"ListObject,omitempty" xml:"ListObject"`
}


type ListObjectsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    ListObjects *ListObjects `json:"ListObjects,omitempty" xml:"ListObjects"`
}


type MergedCellResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    MergedCell *MergedCell `json:"MergedCell,omitempty" xml:"MergedCell"`
}


type MergedCellsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    MergedCells *MergedCells `json:"MergedCells,omitempty" xml:"MergedCells"`
}


type NameResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Name *Name `json:"Name,omitempty" xml:"Name"`
}


type NamesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Names *Names `json:"Names,omitempty" xml:"Names"`
}


type OleObjectResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    OleObject *OleObject `json:"OleObject,omitempty" xml:"OleObject"`
}


type OleObjectsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    OleObjects *OleObjects `json:"OleObjects,omitempty" xml:"OleObjects"`
}


type OvalResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *Oval `json:"Shape,omitempty" xml:"Shape"`
}


type PageSectionsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PageSections []PageSection `json:"PageSections,omitempty" xml:"PageSections"`
}


type PageSetupResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PageSetup *PageSetup `json:"PageSetup,omitempty" xml:"PageSetup"`
}


type PictureResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Picture *Picture `json:"Picture,omitempty" xml:"Picture"`
}


type PicturesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Pictures *Pictures `json:"Pictures,omitempty" xml:"Pictures"`
}


type PivotFieldResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PivotField *PivotField `json:"PivotField,omitempty" xml:"PivotField"`
}


type PivotFilterResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PivotFilter *PivotFilter `json:"PivotFilter,omitempty" xml:"PivotFilter"`
}


type PivotFiltersResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PivotFilters []PivotFilter `json:"PivotFilters,omitempty" xml:"PivotFilters"`
}


type PivotTableResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PivotTable *PivotTable `json:"PivotTable,omitempty" xml:"PivotTable"`
}


type PivotTablesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PivotTables *PivotTables `json:"PivotTables,omitempty" xml:"PivotTables"`
}


type PlotAreaResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    PlotArea *PlotArea `json:"PlotArea,omitempty" xml:"PlotArea"`
}


type RadioButtonResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *RadioButton `json:"Shape,omitempty" xml:"Shape"`
}


type RangeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Range_ *Range `json:"Range,omitempty" xml:"Range"`
}


type RangesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Ranges *Ranges `json:"Ranges,omitempty" xml:"Ranges"`
}


type RangeValueResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    CellsList []Cell `json:"CellsList,omitempty" xml:"CellsList"`
}


type RectangleShapeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *RectangleShape `json:"Shape,omitempty" xml:"Shape"`
}


type RowResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Row *Row `json:"Row,omitempty" xml:"Row"`
}


type RowsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Rows *Rows `json:"Rows,omitempty" xml:"Rows"`
}


type SaveFilesToCloudResultResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    SaveFilesToCloudResult *SaveFilesToCloudResult `json:"SaveFilesToCloudResult,omitempty" xml:"SaveFilesToCloudResult"`
}


type ScrollBarResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *ScrollBar `json:"Shape,omitempty" xml:"Shape"`
}


type SeriesesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Serieses *SeriesItems `json:"Serieses,omitempty" xml:"Serieses"`
}


type SeriesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Series *Series `json:"Series,omitempty" xml:"Series"`
}


type ShapeResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *Shape `json:"Shape,omitempty" xml:"Shape"`
}


type ShapesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shapes *Shapes `json:"Shapes,omitempty" xml:"Shapes"`
}


type SingleValueResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Value *SingleValue `json:"Value,omitempty" xml:"Value"`
}


type SparklineGroupResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    SparklineGroup *SparklineGroup `json:"SparklineGroup,omitempty" xml:"SparklineGroup"`
}


type SparklineGroupsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    SparklineGroups *SparklineGroups `json:"SparklineGroups,omitempty" xml:"SparklineGroups"`
}


type SpinnerResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *Spinner `json:"Shape,omitempty" xml:"Shape"`
}


type SplitResultResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Result *SplitResult `json:"Result,omitempty" xml:"Result"`
}


type StyleResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Style *Style `json:"Style,omitempty" xml:"Style"`
}


type StylesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Styles *Styles `json:"Styles,omitempty" xml:"Styles"`
}


type TaskRunResultResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    TaskRunResult *TaskRunResult `json:"TaskRunResult,omitempty" xml:"TaskRunResult"`
}


type TextBoxResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Shape *TextBox `json:"Shape,omitempty" xml:"Shape"`
}


type TextItemResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    TextItem *TextItem `json:"TextItem,omitempty" xml:"TextItem"`
}


type TextItemsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    TextItems *TextItems `json:"TextItems,omitempty" xml:"TextItems"`
}


type TickLabelsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    TickLabels *TickLabels `json:"TickLabels,omitempty" xml:"TickLabels"`
}


type TitleResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Title *Title `json:"Title,omitempty" xml:"Title"`
}


type TrendlineResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Trendline *Trendline `json:"Trendline,omitempty" xml:"Trendline"`
}


type TrendlinesResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Trendlines *Trendlines `json:"Trendlines,omitempty" xml:"Trendlines"`
}


type ValidationResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Validation *Validation `json:"Validation,omitempty" xml:"Validation"`
}


type ValidationsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Validations *Validations `json:"Validations,omitempty" xml:"Validations"`
}


type VerticalPageBreakResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    VerticalPageBreak *VerticalPageBreak `json:"VerticalPageBreak,omitempty" xml:"VerticalPageBreak"`
}


type VerticalPageBreaksResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    VerticalPageBreaks *VerticalPageBreaks `json:"VerticalPageBreaks,omitempty" xml:"VerticalPageBreaks"`
}


type WallsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Walls *Walls `json:"Walls,omitempty" xml:"Walls"`
}


type WorkbookReplaceResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Matches int64 `json:"Matches,omitempty" xml:"Matches"`
    Workbook *LinkElement `json:"Workbook,omitempty" xml:"Workbook"`
}


type WorkbookResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Workbook *Workbook `json:"Workbook,omitempty" xml:"Workbook"`
}


type WorkbookSettingsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Settings *WorkbookSettings `json:"settings,omitempty" xml:"settings"`
}


type WorkbooksResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Workbooks []LinkElement `json:"Workbooks,omitempty" xml:"Workbooks"`
}


type WorksheetReplaceResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Matches int64 `json:"Matches,omitempty" xml:"Matches"`
    Worksheet *LinkElement `json:"Worksheet,omitempty" xml:"Worksheet"`
}


type WorksheetResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Worksheet *Worksheet `json:"Worksheet,omitempty" xml:"Worksheet"`
}


type WorksheetsResponse struct {
     
        Code int64 `json:"Code,omitempty" xml:"Code"`
        Status string `json:"Status,omitempty" xml:"Status"`
 
    Worksheets *Worksheets `json:"Worksheets,omitempty" xml:"Worksheets"`
}


type AnalyzeExcelRequest struct {
 
    Files []FileInfo `json:"Files,omitempty" xml:"Files"`
    NeedThumbnail bool `json:"NeedThumbnail,omitempty" xml:"NeedThumbnail"`
    BuildSuggestoinSheet bool `json:"BuildSuggestoinSheet,omitempty" xml:"BuildSuggestoinSheet"`
}


type BatchConvertRequest struct {
 
    SourceFolder string `json:"SourceFolder,omitempty" xml:"SourceFolder"`
    SourceStorage string `json:"SourceStorage,omitempty" xml:"SourceStorage"`
    MatchCondition *MatchConditionRequest `json:"MatchCondition,omitempty" xml:"MatchCondition"`
    Format string `json:"Format,omitempty" xml:"Format"`
    OutFolder string `json:"OutFolder,omitempty" xml:"OutFolder"`
    OutStorage string `json:"OutStorage,omitempty" xml:"OutStorage"`
    Region string `json:"Region,omitempty" xml:"Region"`
    PageWideFitOnPerSheet bool `json:"PageWideFitOnPerSheet,omitempty" xml:"PageWideFitOnPerSheet"`
    PageTallFitOnPerSheet bool `json:"PageTallFitOnPerSheet,omitempty" xml:"PageTallFitOnPerSheet"`
    SaveOptions *SaveOptions `json:"SaveOptions,omitempty" xml:"SaveOptions"`
}


type BatchLockRequest struct {
 
    SourceFolder string `json:"SourceFolder,omitempty" xml:"SourceFolder"`
    SourceStorage string `json:"SourceStorage,omitempty" xml:"SourceStorage"`
    MatchCondition *MatchConditionRequest `json:"MatchCondition,omitempty" xml:"MatchCondition"`
    Password string `json:"Password,omitempty" xml:"Password"`
    OutFolder string `json:"OutFolder,omitempty" xml:"OutFolder"`
    OutStorage string `json:"OutStorage,omitempty" xml:"OutStorage"`
}


type BatchProtectRequest struct {
 
    SourceFolder string `json:"SourceFolder,omitempty" xml:"SourceFolder"`
    SourceStorage string `json:"SourceStorage,omitempty" xml:"SourceStorage"`
    MatchCondition *MatchConditionRequest `json:"MatchCondition,omitempty" xml:"MatchCondition"`
    ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
    Password string `json:"Password,omitempty" xml:"Password"`
    OutFolder string `json:"OutFolder,omitempty" xml:"OutFolder"`
    OutStorage string `json:"OutStorage,omitempty" xml:"OutStorage"`
}


type BatchSplitRequest struct {
 
    SourceFolder string `json:"SourceFolder,omitempty" xml:"SourceFolder"`
    SourceStorage string `json:"SourceStorage,omitempty" xml:"SourceStorage"`
    MatchCondition *MatchConditionRequest `json:"MatchCondition,omitempty" xml:"MatchCondition"`
    Format string `json:"Format,omitempty" xml:"Format"`
    FromIndex int64 `json:"FromIndex,omitempty" xml:"FromIndex"`
    ToIndex int64 `json:"ToIndex,omitempty" xml:"ToIndex"`
    OutFolder string `json:"OutFolder,omitempty" xml:"OutFolder"`
    OutStorage string `json:"OutStorage,omitempty" xml:"OutStorage"`
    Region string `json:"Region,omitempty" xml:"Region"`
    SaveOptions *SaveOptions `json:"SaveOptions,omitempty" xml:"SaveOptions"`
}


type ColorFilterRequest struct {
 
    Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
    ForegroundColor *CellsColor `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
    BackgroundColor *CellsColor `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
}


type ConvertParameter struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Value string `json:"Value,omitempty" xml:"Value"`
}


type CreatePivotTableRequest struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    SourceData string `json:"SourceData,omitempty" xml:"SourceData"`
    DestCellName string `json:"DestCellName,omitempty" xml:"DestCellName"`
    UseSameSource bool `json:"UseSameSource,omitempty" xml:"UseSameSource"`
    PivotFieldRows []int64 `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
    PivotFieldColumns []int64 `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
    PivotFieldData []int64 `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
}


type DataCleansingRequest struct {
 
    File *FileInfo `json:"File,omitempty" xml:"File"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    Region string `json:"Region,omitempty" xml:"Region"`
    OutFileFormat string `json:"OutFileFormat,omitempty" xml:"OutFileFormat"`
    DataCleansing *DataCleansing `json:"DataCleansing,omitempty" xml:"DataCleansing"`
}


type DataDeduplicationRequest struct {
 
    File *FileInfo `json:"File,omitempty" xml:"File"`
    DeduplicationRegion *DeduplicationRegion `json:"DeduplicationRegion,omitempty" xml:"DeduplicationRegion"`
    OutFileFormat string `json:"OutFileFormat,omitempty" xml:"OutFileFormat"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    Region string `json:"Region,omitempty" xml:"Region"`
}


type DataFillRequest struct {
 
    File *FileInfo `json:"File,omitempty" xml:"File"`
    OutFileFormat string `json:"OutFileFormat,omitempty" xml:"OutFileFormat"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    Region string `json:"Region,omitempty" xml:"Region"`
    DataFill *DataFill `json:"DataFill,omitempty" xml:"DataFill"`
}


type DataTransformationRequest struct {
 
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    LoadData *LoadData `json:"LoadData,omitempty" xml:"LoadData"`
    AppliedSteps []AppliedStep `json:"AppliedSteps,omitempty" xml:"AppliedSteps"`
    Region string `json:"Region,omitempty" xml:"Region"`
    OutFormat string `json:"OutFormat,omitempty" xml:"OutFormat"`
}


type DeleteIncompleteRowsRequest struct {
 
    File *FileInfo `json:"File,omitempty" xml:"File"`
    OutFileFormat string `json:"OutFileFormat,omitempty" xml:"OutFileFormat"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    Region string `json:"Region,omitempty" xml:"Region"`
    Ranges []Range `json:"Ranges,omitempty" xml:"Ranges"`
}


type ImportJsonRequest struct {
 
    JsonFileSource *DataSource `json:"JsonFileSource,omitempty" xml:"JsonFileSource"`
    ImportPosition *ImportPosition `json:"ImportPosition,omitempty" xml:"ImportPosition"`
    JsonContent string `json:"JsonContent,omitempty" xml:"JsonContent"`
}


type ImportXMLRequest struct {
 
    XMLFileSource *DataSource `json:"XMLFileSource,omitempty" xml:"XMLFileSource"`
    ImportPosition *ImportPosition `json:"ImportPosition,omitempty" xml:"ImportPosition"`
    XMLContent string `json:"XMLContent,omitempty" xml:"XMLContent"`
}


type MatchConditionRequest struct {
 
    RegexPattern string `json:"RegexPattern,omitempty" xml:"RegexPattern"`
    FullMatchConditions []string `json:"FullMatchConditions,omitempty" xml:"FullMatchConditions"`
}


type PasswordRequest struct {
 
    Password string `json:"Password,omitempty" xml:"Password"`
}


type PivotTableFieldRequest struct {
 
    Data []int64 `json:"Data,omitempty" xml:"Data"`
}


type ProtectWorkbookRequest struct {
 
    AwaysOpenReadOnly bool `json:"AwaysOpenReadOnly,omitempty" xml:"AwaysOpenReadOnly"`
    EncryptWithPassword string `json:"EncryptWithPassword,omitempty" xml:"EncryptWithPassword"`
    ProtectCurrentSheet *Protection `json:"ProtectCurrentSheet,omitempty" xml:"ProtectCurrentSheet"`
    ProtectAllSheets *Protection `json:"ProtectAllSheets,omitempty" xml:"ProtectAllSheets"`
    ProtectWorkbookStructure string `json:"ProtectWorkbookStructure,omitempty" xml:"ProtectWorkbookStructure"`
    DigitalSignature *DigitalSignature `json:"DigitalSignature,omitempty" xml:"DigitalSignature"`
    MarkAsFinal bool `json:"MarkAsFinal,omitempty" xml:"MarkAsFinal"`
}


type RangeConvertRequest struct {
 
    Source *Range `json:"Source,omitempty" xml:"Source"`
    ImageType string `json:"ImageType,omitempty" xml:"ImageType"`
    ImageOrPrintOptions *ImageOrPrintOptions `json:"ImageOrPrintOptions,omitempty" xml:"ImageOrPrintOptions"`
    PageSetup *PageSetup `json:"PageSetup,omitempty" xml:"PageSetup"`
}


type RangeCopyRequest struct {
 
    Operate string `json:"Operate,omitempty" xml:"Operate"`
    Source *Range `json:"Source,omitempty" xml:"Source"`
    Target *Range `json:"Target,omitempty" xml:"Target"`
    TargetWorkbook string `json:"TargetWorkbook,omitempty" xml:"TargetWorkbook"`
    PasteOptions *PasteOptions `json:"PasteOptions,omitempty" xml:"PasteOptions"`
}


type RangeSetOutlineBorderRequest struct {
 
    Range_ *Range `json:"Range,omitempty" xml:"Range"`
    BorderEdge string `json:"borderEdge,omitempty" xml:"borderEdge"`
    BorderStyle string `json:"borderStyle,omitempty" xml:"borderStyle"`
    BorderColor *Color `json:"borderColor,omitempty" xml:"borderColor"`
}


type RangeSetStyleRequest struct {
 
    Range_ *Range `json:"Range,omitempty" xml:"Range"`
    Style *Style `json:"Style,omitempty" xml:"Style"`
}


type RangeSortRequest struct {
 
    DataSorter *DataSorter `json:"DataSorter,omitempty" xml:"DataSorter"`
    CellArea *Range `json:"CellArea,omitempty" xml:"CellArea"`
}


type TableTotalRequest struct {
 
    ListColumnIndex int64 `json:"ListColumnIndex,omitempty" xml:"ListColumnIndex"`
    TotalsCalculation string `json:"TotalsCalculation,omitempty" xml:"TotalsCalculation"`
    CustomFormula string `json:"CustomFormula,omitempty" xml:"CustomFormula"`
}


type TextWaterMarkerRequest struct {
 
    Text string `json:"Text,omitempty" xml:"Text"`
    FontName string `json:"FontName,omitempty" xml:"FontName"`
    FontSize int64 `json:"FontSize,omitempty" xml:"FontSize"`
    Height int64 `json:"Height,omitempty" xml:"Height"`
    Width int64 `json:"Width,omitempty" xml:"Width"`
    ImageAdaptOption string `json:"ImageAdaptOption,omitempty" xml:"ImageAdaptOption"`
}


type WorkbookEncryptionRequest struct {
 
    EncryptionType string `json:"EncryptionType,omitempty" xml:"EncryptionType"`
    KeyLength int64 `json:"KeyLength,omitempty" xml:"KeyLength"`
    Password string `json:"Password,omitempty" xml:"Password"`
}


type WorkbookProtectionRequest struct {
 
    ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
    Password string `json:"Password,omitempty" xml:"Password"`
}


type WorksheetMovingRequest struct {
 
    DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
    Position string `json:"Position,omitempty" xml:"Position"`
}


type AppliedOperate interface {
}


type AppliedStep struct {
 
    StepName string `json:"StepName,omitempty" xml:"StepName"`
    AppliedOperate AppliedOperate `json:"AppliedOperate,omitempty" xml:"AppliedOperate"`
}


type DataQuery struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    DataSourceDataType string `json:"DataSourceDataType,omitempty" xml:"DataSourceDataType"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    DataItem *DataItem `json:"DataItem,omitempty" xml:"DataItem"`
}


type LoadData struct {
 
    LoadTo *LoadTo `json:"LoadTo,omitempty" xml:"LoadTo"`
    DataQuery *DataQuery `json:"DataQuery,omitempty" xml:"DataQuery"`
}


type LoadTo struct {
 
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    BeginRowIndex int64 `json:"beginRowIndex,omitempty" xml:"beginRowIndex"`
    BeginColumnIndex int64 `json:"beginColumnIndex,omitempty" xml:"beginColumnIndex"`
}


type MergeQueries struct {
     
        AppliedOperateType string `json:"AppliedOperateType,omitempty" xml:"AppliedOperateType"`
 
    DataQueryNameA string `json:"DataQueryNameA,omitempty" xml:"DataQueryNameA"`
    DataAIndexField string `json:"DataAIndexField,omitempty" xml:"DataAIndexField"`
    DataQueryNameB string `json:"DataQueryNameB,omitempty" xml:"DataQueryNameB"`
    DataBIndexField string `json:"DataBIndexField,omitempty" xml:"DataBIndexField"`
    JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
}


type PivotColumn struct {
     
        AppliedOperateType string `json:"AppliedOperateType,omitempty" xml:"AppliedOperateType"`
 
    PivotColumnName string `json:"PivotColumnName,omitempty" xml:"PivotColumnName"`
    ValueColumnNames []string `json:"ValueColumnNames,omitempty" xml:"ValueColumnNames"`
}


type UnpivotColumn struct {
     
        AppliedOperateType string `json:"AppliedOperateType,omitempty" xml:"AppliedOperateType"`
 
    UnpivotColumnNames []string `json:"UnpivotColumnNames,omitempty" xml:"UnpivotColumnNames"`
    ColumnMapName string `json:"ColumnMapName,omitempty" xml:"ColumnMapName"`
    ValueMapName string `json:"ValueMapName,omitempty" xml:"ValueMapName"`
}


type PivotField struct {
 
    AutoShowCount int64 `json:"AutoShowCount,omitempty" xml:"AutoShowCount"`
    AutoShowField int64 `json:"AutoShowField,omitempty" xml:"AutoShowField"`
    AutoSortField int64 `json:"AutoSortField,omitempty" xml:"AutoSortField"`
    BaseField int64 `json:"BaseField,omitempty" xml:"BaseField"`
    BaseIndex int64 `json:"BaseIndex,omitempty" xml:"BaseIndex"`
    BaseItem int64 `json:"BaseItem,omitempty" xml:"BaseItem"`
    BaseItemPosition string `json:"BaseItemPosition,omitempty" xml:"BaseItemPosition"`
    CurrentPageItem int64 `json:"CurrentPageItem,omitempty" xml:"CurrentPageItem"`
    DataDisplayFormat string `json:"DataDisplayFormat,omitempty" xml:"DataDisplayFormat"`
    DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
    DragToColumn bool `json:"DragToColumn,omitempty" xml:"DragToColumn"`
    DragToData bool `json:"DragToData,omitempty" xml:"DragToData"`
    DragToHide bool `json:"DragToHide,omitempty" xml:"DragToHide"`
    DragToPage bool `json:"DragToPage,omitempty" xml:"DragToPage"`
    DragToRow bool `json:"DragToRow,omitempty" xml:"DragToRow"`
    Function string `json:"Function,omitempty" xml:"Function"`
    InsertBlankRow bool `json:"InsertBlankRow,omitempty" xml:"InsertBlankRow"`
    IsAscendShow bool `json:"IsAscendShow,omitempty" xml:"IsAscendShow"`
    IsAscendSort bool `json:"IsAscendSort,omitempty" xml:"IsAscendSort"`
    IsAutoShow bool `json:"IsAutoShow,omitempty" xml:"IsAutoShow"`
    IsAutoSort bool `json:"IsAutoSort,omitempty" xml:"IsAutoSort"`
    IsAutoSubtotals bool `json:"IsAutoSubtotals,omitempty" xml:"IsAutoSubtotals"`
    IsCalculatedField bool `json:"IsCalculatedField,omitempty" xml:"IsCalculatedField"`
    IsIncludeNewItemsInFilter bool `json:"IsIncludeNewItemsInFilter,omitempty" xml:"IsIncludeNewItemsInFilter"`
    IsInsertPageBreaksBetweenItems bool `json:"IsInsertPageBreaksBetweenItems,omitempty" xml:"IsInsertPageBreaksBetweenItems"`
    IsMultipleItemSelectionAllowed bool `json:"IsMultipleItemSelectionAllowed,omitempty" xml:"IsMultipleItemSelectionAllowed"`
    IsRepeatItemLabels bool `json:"IsRepeatItemLabels,omitempty" xml:"IsRepeatItemLabels"`
    ItemCount int64 `json:"ItemCount,omitempty" xml:"ItemCount"`
    Items []string `json:"Items,omitempty" xml:"Items"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Number int64 `json:"Number,omitempty" xml:"Number"`
    NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
    OriginalItems []string `json:"OriginalItems,omitempty" xml:"OriginalItems"`
    PivotItems []PivotItem `json:"PivotItems,omitempty" xml:"PivotItems"`
    Position int64 `json:"Position,omitempty" xml:"Position"`
    ShowAllItems bool `json:"ShowAllItems,omitempty" xml:"ShowAllItems"`
    ShowCompact bool `json:"ShowCompact,omitempty" xml:"ShowCompact"`
    ShowInOutlineForm bool `json:"ShowInOutlineForm,omitempty" xml:"ShowInOutlineForm"`
    ShowSubtotalAtTop bool `json:"ShowSubtotalAtTop,omitempty" xml:"ShowSubtotalAtTop"`
}


type PivotFilter struct {
 
    AutoFilter *AutoFilter `json:"AutoFilter,omitempty" xml:"AutoFilter"`
    EvaluationOrder int64 `json:"EvaluationOrder,omitempty" xml:"EvaluationOrder"`
    FieldIndex int64 `json:"FieldIndex,omitempty" xml:"FieldIndex"`
    FilterType string `json:"FilterType,omitempty" xml:"FilterType"`
    ValueFieldIndex int64 `json:"ValueFieldIndex,omitempty" xml:"ValueFieldIndex"`
    MemberPropertyFieldIndex int64 `json:"MemberPropertyFieldIndex,omitempty" xml:"MemberPropertyFieldIndex"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Value1 string `json:"Value1,omitempty" xml:"Value1"`
    Value2 string `json:"Value2,omitempty" xml:"Value2"`
    Top10Filter *Top10Filter `json:"Top10Filter,omitempty" xml:"Top10Filter"`
}


type PivotItem struct {
 
    Index int64 `json:"Index,omitempty" xml:"Index"`
    IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Value string `json:"Value,omitempty" xml:"Value"`
}


type PivotTable struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AltTextDescription string `json:"AltTextDescription,omitempty" xml:"AltTextDescription"`
    AltTextTitle string `json:"AltTextTitle,omitempty" xml:"AltTextTitle"`
    AutoFormatType string `json:"AutoFormatType,omitempty" xml:"AutoFormatType"`
    BaseFields []PivotField `json:"BaseFields,omitempty" xml:"BaseFields"`
    ColumnFields []PivotField `json:"ColumnFields,omitempty" xml:"ColumnFields"`
    ColumnGrand bool `json:"ColumnGrand,omitempty" xml:"ColumnGrand"`
    ColumnHeaderCaption string `json:"ColumnHeaderCaption,omitempty" xml:"ColumnHeaderCaption"`
    ColumnRange *CellArea `json:"ColumnRange,omitempty" xml:"ColumnRange"`
    CustomListSort bool `json:"CustomListSort,omitempty" xml:"CustomListSort"`
    DataBodyRange *CellArea `json:"DataBodyRange,omitempty" xml:"DataBodyRange"`
    DataField *PivotField `json:"DataField,omitempty" xml:"DataField"`
    DataFields []PivotField `json:"DataFields,omitempty" xml:"DataFields"`
    DataSource []string `json:"DataSource,omitempty" xml:"DataSource"`
    DisplayErrorString bool `json:"DisplayErrorString,omitempty" xml:"DisplayErrorString"`
    DisplayImmediateItems bool `json:"DisplayImmediateItems,omitempty" xml:"DisplayImmediateItems"`
    DisplayNullString bool `json:"DisplayNullString,omitempty" xml:"DisplayNullString"`
    EnableDataValueEditing bool `json:"EnableDataValueEditing,omitempty" xml:"EnableDataValueEditing"`
    EnableDrilldown bool `json:"EnableDrilldown,omitempty" xml:"EnableDrilldown"`
    EnableFieldDialog bool `json:"EnableFieldDialog,omitempty" xml:"EnableFieldDialog"`
    EnableFieldList bool `json:"EnableFieldList,omitempty" xml:"EnableFieldList"`
    EnableWizard bool `json:"EnableWizard,omitempty" xml:"EnableWizard"`
    ErrorString string `json:"ErrorString,omitempty" xml:"ErrorString"`
    FieldListSortAscending bool `json:"FieldListSortAscending,omitempty" xml:"FieldListSortAscending"`
    GrandTotalName string `json:"GrandTotalName,omitempty" xml:"GrandTotalName"`
    HasBlankRows bool `json:"HasBlankRows,omitempty" xml:"HasBlankRows"`
    Indent int64 `json:"Indent,omitempty" xml:"Indent"`
    IsAutoFormat bool `json:"IsAutoFormat,omitempty" xml:"IsAutoFormat"`
    IsGridDropZones bool `json:"IsGridDropZones,omitempty" xml:"IsGridDropZones"`
    IsMultipleFieldFilters bool `json:"IsMultipleFieldFilters,omitempty" xml:"IsMultipleFieldFilters"`
    IsSelected bool `json:"IsSelected,omitempty" xml:"IsSelected"`
    ItemPrintTitles bool `json:"ItemPrintTitles,omitempty" xml:"ItemPrintTitles"`
    ManualUpdate bool `json:"ManualUpdate,omitempty" xml:"ManualUpdate"`
    MergeLabels bool `json:"MergeLabels,omitempty" xml:"MergeLabels"`
    MissingItemsLimit string `json:"MissingItemsLimit,omitempty" xml:"MissingItemsLimit"`
    Name string `json:"Name,omitempty" xml:"Name"`
    NullString string `json:"NullString,omitempty" xml:"NullString"`
    PageFieldOrder string `json:"PageFieldOrder,omitempty" xml:"PageFieldOrder"`
    PageFields []PivotField `json:"PageFields,omitempty" xml:"PageFields"`
    PageFieldWrapCount int64 `json:"PageFieldWrapCount,omitempty" xml:"PageFieldWrapCount"`
    PivotFilters []PivotFilter `json:"PivotFilters,omitempty" xml:"PivotFilters"`
    PivotTableStyleName string `json:"PivotTableStyleName,omitempty" xml:"PivotTableStyleName"`
    PivotTableStyleType string `json:"PivotTableStyleType,omitempty" xml:"PivotTableStyleType"`
    PreserveFormatting bool `json:"PreserveFormatting,omitempty" xml:"PreserveFormatting"`
    PrintDrill bool `json:"PrintDrill,omitempty" xml:"PrintDrill"`
    PrintTitles bool `json:"PrintTitles,omitempty" xml:"PrintTitles"`
    RefreshDataFlag bool `json:"RefreshDataFlag,omitempty" xml:"RefreshDataFlag"`
    RefreshDataOnOpeningFile bool `json:"RefreshDataOnOpeningFile,omitempty" xml:"RefreshDataOnOpeningFile"`
    RowFields []PivotField `json:"RowFields,omitempty" xml:"RowFields"`
    RowGrand bool `json:"RowGrand,omitempty" xml:"RowGrand"`
    RowHeaderCaption string `json:"RowHeaderCaption,omitempty" xml:"RowHeaderCaption"`
    RowRange *CellArea `json:"RowRange,omitempty" xml:"RowRange"`
    SaveData bool `json:"SaveData,omitempty" xml:"SaveData"`
    ShowDataTips bool `json:"ShowDataTips,omitempty" xml:"ShowDataTips"`
    ShowDrill bool `json:"ShowDrill,omitempty" xml:"ShowDrill"`
    ShowEmptyCol bool `json:"ShowEmptyCol,omitempty" xml:"ShowEmptyCol"`
    ShowEmptyRow bool `json:"ShowEmptyRow,omitempty" xml:"ShowEmptyRow"`
    ShowMemberPropertyTips bool `json:"ShowMemberPropertyTips,omitempty" xml:"ShowMemberPropertyTips"`
    ShowPivotStyleColumnHeader bool `json:"ShowPivotStyleColumnHeader,omitempty" xml:"ShowPivotStyleColumnHeader"`
    ShowPivotStyleColumnStripes bool `json:"ShowPivotStyleColumnStripes,omitempty" xml:"ShowPivotStyleColumnStripes"`
    ShowPivotStyleLastColumn bool `json:"ShowPivotStyleLastColumn,omitempty" xml:"ShowPivotStyleLastColumn"`
    ShowPivotStyleRowHeader bool `json:"ShowPivotStyleRowHeader,omitempty" xml:"ShowPivotStyleRowHeader"`
    ShowPivotStyleRowStripes bool `json:"ShowPivotStyleRowStripes,omitempty" xml:"ShowPivotStyleRowStripes"`
    ShowRowHeaderCaption bool `json:"ShowRowHeaderCaption,omitempty" xml:"ShowRowHeaderCaption"`
    ShowValuesRow bool `json:"ShowValuesRow,omitempty" xml:"ShowValuesRow"`
    SubtotalHiddenPageItems bool `json:"SubtotalHiddenPageItems,omitempty" xml:"SubtotalHiddenPageItems"`
    TableRange1 *CellArea `json:"TableRange1,omitempty" xml:"TableRange1"`
    TableRange2 *CellArea `json:"TableRange2,omitempty" xml:"TableRange2"`
    Tag string `json:"Tag,omitempty" xml:"Tag"`
}


type PivotTables struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    PivotTableList []LinkElement `json:"PivotTableList,omitempty" xml:"PivotTableList"`
}


type AddTextOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
    Text string `json:"Text,omitempty" xml:"Text"`
    SelectPoistion string `json:"SelectPoistion,omitempty" xml:"SelectPoistion"`
    SelectText string `json:"SelectText,omitempty" xml:"SelectText"`
    SkipEmptyCells bool `json:"SkipEmptyCells,omitempty" xml:"SkipEmptyCells"`
}


type BaseOperateOptions interface {
}


type CharacterCountOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
}


type CheckExternalReferenceOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Ranged_Table string `json:"Ranged_Table,omitempty" xml:"Ranged_Table"`
    ChartIndex int64 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
}


type CheckFormulaErrorOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
    ChartIndex int64 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
    Names []string `json:"Names,omitempty" xml:"Names"`
}


type CombinationSourceData struct {
 
    Tag string `json:"Tag,omitempty" xml:"Tag"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    TableName string `json:"TableName,omitempty" xml:"TableName"`
    CellArea string `json:"CellArea,omitempty" xml:"CellArea"`
    HasHeader bool `json:"HasHeader,omitempty" xml:"HasHeader"`
}


type ConvertTextOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
    ConvertTextType string `json:"ConvertTextType,omitempty" xml:"ConvertTextType"`
    SourceCharacters string `json:"SourceCharacters,omitempty" xml:"SourceCharacters"`
    TargetCharacters string `json:"TargetCharacters,omitempty" xml:"TargetCharacters"`
}


type ConvertWorkbookOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    PageSetup *PageSetup `json:"PageSetup,omitempty" xml:"PageSetup"`
    SaveOptions *SaveOptions `json:"SaveOptions,omitempty" xml:"SaveOptions"`
    ConvertFormat string `json:"ConvertFormat,omitempty" xml:"ConvertFormat"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
}


type ConvertWorksheetOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    WorksheetName string `json:"WorksheetName,omitempty" xml:"WorksheetName"`
    PageSetup *PageSetup `json:"PageSetup,omitempty" xml:"PageSetup"`
    ImageOrPrintOptions *ImageOrPrintOptions `json:"ImageOrPrintOptions,omitempty" xml:"ImageOrPrintOptions"`
    ConvertFormat string `json:"ConvertFormat,omitempty" xml:"ConvertFormat"`
    CheckExcelRestriction bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
    Region string `json:"Region,omitempty" xml:"Region"`
}


type DataOutputLocation struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    BeginRowIndex int64 `json:"BeginRowIndex,omitempty" xml:"BeginRowIndex"`
    BeginColumnIndex int64 `json:"BeginColumnIndex,omitempty" xml:"BeginColumnIndex"`
}


type ExtractTextOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Range_ string `json:"Range,omitempty" xml:"Range"`
    ExtractTextType string `json:"ExtractTextType,omitempty" xml:"ExtractTextType"`
    BeforeText string `json:"BeforeText,omitempty" xml:"BeforeText"`
    AfterText string `json:"AfterText,omitempty" xml:"AfterText"`
    BeforePosition int64 `json:"BeforePosition,omitempty" xml:"BeforePosition"`
    AfterPosition int64 `json:"AfterPosition,omitempty" xml:"AfterPosition"`
    OutPositionRange string `json:"OutPositionRange,omitempty" xml:"OutPositionRange"`
}


type MergeTableOptions struct {
 
    MainTable *CombinationSourceData `json:"MainTable,omitempty" xml:"MainTable"`
    SecondaryTable *CombinationSourceData `json:"SecondaryTable,omitempty" xml:"SecondaryTable"`
    DataMergeType string `json:"DataMergeType,omitempty" xml:"DataMergeType"`
    OverwriteMainTable bool `json:"OverwriteMainTable,omitempty" xml:"OverwriteMainTable"`
    SyncDataToTargetWorkbook bool `json:"SyncDataToTargetWorkbook,omitempty" xml:"SyncDataToTargetWorkbook"`
    MergedDataToPosition *DataOutputLocation `json:"MergedDataToPosition,omitempty" xml:"MergedDataToPosition"`
}


type RemoveCharactersByCharacter struct {
 
    RemoveTextMethod string `json:"RemoveTextMethod,omitempty" xml:"RemoveTextMethod"`
    RemoveCharacters []string `json:"RemoveCharacters,omitempty" xml:"RemoveCharacters"`
    RemoveCharacterSetsType string `json:"RemoveCharacterSetsType,omitempty" xml:"RemoveCharacterSetsType"`
}


type RemoveCharactersOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
    RemoveCharactersByCharacter *RemoveCharactersByCharacter `json:"RemoveCharactersByCharacter,omitempty" xml:"RemoveCharactersByCharacter"`
    RemoveCharactersByPosition *RemoveCharactersByPosition `json:"RemoveCharactersByPosition,omitempty" xml:"RemoveCharactersByPosition"`
}


type RemoveDuplicatesOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
}


type ScopeItem struct {
 
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Ranges []string `json:"Ranges,omitempty" xml:"Ranges"`
}


type ScopeOptions struct {
 
    Scope string `json:"Scope,omitempty" xml:"Scope"`
    ScopeItems []ScopeItem `json:"ScopeItems,omitempty" xml:"ScopeItems"`
}


type SpecifyCellsObject struct {
 
    WorksheetName string `json:"WorksheetName,omitempty" xml:"WorksheetName"`
    PageIndex int64 `json:"PageIndex,omitempty" xml:"PageIndex"`
    Region string `json:"Region,omitempty" xml:"Region"`
}


type SpecifyWordsCountOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    SearchWord string `json:"SearchWord,omitempty" xml:"SearchWord"`
}


type SplitTextOptions struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    Worksheet string `json:"Worksheet,omitempty" xml:"Worksheet"`
    Range_ string `json:"Range,omitempty" xml:"Range"`
    SplitDelimitersType string `json:"SplitDelimitersType,omitempty" xml:"SplitDelimitersType"`
    CustomDelimiter string `json:"CustomDelimiter,omitempty" xml:"CustomDelimiter"`
    KeepDelimitersInResultingCells bool `json:"KeepDelimitersInResultingCells,omitempty" xml:"KeepDelimitersInResultingCells"`
    KeepDelimitersPosition string `json:"KeepDelimitersPosition,omitempty" xml:"KeepDelimitersPosition"`
    HowToSplit string `json:"HowToSplit,omitempty" xml:"HowToSplit"`
}


type TablePositionInfo struct {
 
}


type TrimContentOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    TrimContent string `json:"TrimContent,omitempty" xml:"TrimContent"`
    TrimLeading bool `json:"TrimLeading,omitempty" xml:"TrimLeading"`
    TrimTrailing bool `json:"TrimTrailing,omitempty" xml:"TrimTrailing"`
    TrimSpaceBetweenWordTo1 bool `json:"TrimSpaceBetweenWordTo1,omitempty" xml:"TrimSpaceBetweenWordTo1"`
    TrimNonBreakingSpaces bool `json:"TrimNonBreakingSpaces,omitempty" xml:"TrimNonBreakingSpaces"`
    RemoveExtraLineBreaks bool `json:"RemoveExtraLineBreaks,omitempty" xml:"RemoveExtraLineBreaks"`
    RemoveAllLineBreaks bool `json:"RemoveAllLineBreaks,omitempty" xml:"RemoveAllLineBreaks"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
}


type WordCaseOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
    WordCaseType string `json:"WordCaseType,omitempty" xml:"WordCaseType"`
    ScopeOptions *ScopeOptions `json:"ScopeOptions,omitempty" xml:"ScopeOptions"`
}


type WordsCountOptions struct {
 
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
}


type CellValue struct {
 
    RowIndex int64 `json:"rowIndex,omitempty" xml:"rowIndex"`
    ColumnIndex int64 `json:"columnIndex,omitempty" xml:"columnIndex"`
    Type_ string `json:"type,omitempty" xml:"type"`
    Value string `json:"value,omitempty" xml:"value"`
    Formula string `json:"formula,omitempty" xml:"formula"`
    Style *Style `json:"style,omitempty" xml:"style"`
}


type CustomParserConfig struct {
 
    ColumnIndex int64 `json:"ColumnIndex,omitempty" xml:"ColumnIndex"`
    ParseMethod string `json:"ParseMethod,omitempty" xml:"ParseMethod"`
    CustomStyle string `json:"CustomStyle,omitempty" xml:"CustomStyle"`
}


type Import2DimensionDoubleArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    Data []float64 `json:"Data,omitempty" xml:"Data"`
}


type Import2DimensionIntArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    Data []int64 `json:"Data,omitempty" xml:"Data"`
}


type Import2DimensionStringArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    Data []string `json:"Data,omitempty" xml:"Data"`
}


type ImportBatchDataOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    BatchData []CellValue `json:"BatchData,omitempty" xml:"BatchData"`
}


type ImportCSVDataOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    SeparatorString string `json:"SeparatorString,omitempty" xml:"SeparatorString"`
    ConvertNumericData bool `json:"ConvertNumericData,omitempty" xml:"ConvertNumericData"`
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    SourceFile string `json:"SourceFile,omitempty" xml:"SourceFile"`
    CustomParsers []CustomParserConfig `json:"CustomParsers,omitempty" xml:"CustomParsers"`
}


type ImportDoubleArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    IsVertical bool `json:"IsVertical,omitempty" xml:"IsVertical"`
    Data []float64 `json:"Data,omitempty" xml:"Data"`
}


type ImportIntArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    IsVertical bool `json:"IsVertical,omitempty" xml:"IsVertical"`
    Data []int64 `json:"Data,omitempty" xml:"Data"`
}


type ImportOption struct {
 
    DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
    IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
    ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
    DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
    Source *FileSource `json:"Source,omitempty" xml:"Source"`
}


type ImportPictureOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
    UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
    LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
    LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
    Filename string `json:"Filename,omitempty" xml:"Filename"`
    Data string `json:"Data,omitempty" xml:"Data"`
}


type ImportPosition struct {
 
    SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
    RowIndex int64 `json:"RowIndex,omitempty" xml:"RowIndex"`
    ColumnIndex int64 `json:"ColumnIndex,omitempty" xml:"ColumnIndex"`
}


type ImportStringArrayOption struct {
     
        DestinationWorksheet string `json:"DestinationWorksheet,omitempty" xml:"DestinationWorksheet"`
        IsInsert bool `json:"IsInsert,omitempty" xml:"IsInsert"`
        ImportDataType string `json:"ImportDataType,omitempty" xml:"ImportDataType"`
        DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
        Source *FileSource `json:"Source,omitempty" xml:"Source"`
 
    FirstRow int64 `json:"FirstRow,omitempty" xml:"FirstRow"`
    FirstColumn int64 `json:"FirstColumn,omitempty" xml:"FirstColumn"`
    IsVertical bool `json:"IsVertical,omitempty" xml:"IsVertical"`
    Data []string `json:"Data,omitempty" xml:"Data"`
}


type ArcShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
    BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
    BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
    EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
    EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
    EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
}


type Area struct {
 
    BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
    FillFormat *FillFormat `json:"FillFormat,omitempty" xml:"FillFormat"`
    ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
    Format string `json:"Format,omitempty" xml:"Format"`
    InvertIfNegative bool `json:"InvertIfNegative,omitempty" xml:"InvertIfNegative"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}


type AutoShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type AutoShapes struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoShapeList []LinkElement `json:"AutoShapeList,omitempty" xml:"AutoShapeList"`
}


type Button struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type CellsDrawing struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type CheckBox struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    CheckedValue string `json:"CheckedValue,omitempty" xml:"CheckedValue"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    Value bool `json:"Value,omitempty" xml:"Value"`
}


type ComboBox struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    DropDownLines int64 `json:"DropDownLines,omitempty" xml:"DropDownLines"`
    InputRange string `json:"InputRange,omitempty" xml:"InputRange"`
    SelectedCell *LinkElement `json:"SelectedCell,omitempty" xml:"SelectedCell"`
    SelectedIndex int64 `json:"SelectedIndex,omitempty" xml:"SelectedIndex"`
    SelectedValue string `json:"SelectedValue,omitempty" xml:"SelectedValue"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
}


type CommentShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    Comment *LinkElement `json:"Comment,omitempty" xml:"Comment"`
}


type FillFormat struct {
 
    Type_ string `json:"Type,omitempty" xml:"Type"`
    SolidFill *SolidFill `json:"SolidFill,omitempty" xml:"SolidFill"`
    PatternFill *PatternFill `json:"PatternFill,omitempty" xml:"PatternFill"`
    TextureFill *TextureFill `json:"TextureFill,omitempty" xml:"TextureFill"`
    GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
    ImageData string `json:"ImageData,omitempty" xml:"ImageData"`
}


type Form struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    FormType string `json:"FormType,omitempty" xml:"FormType"`
    CheckedValue string `json:"CheckedValue,omitempty" xml:"CheckedValue"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    InputRange string `json:"InputRange,omitempty" xml:"InputRange"`
    SelectedIndex int64 `json:"SelectedIndex,omitempty" xml:"SelectedIndex"`
    SelectedValue string `json:"SelectedValue,omitempty" xml:"SelectedValue"`
    SelectedCell *LinkElement `json:"SelectedCell,omitempty" xml:"SelectedCell"`
    DropDownLines int64 `json:"DropDownLines,omitempty" xml:"DropDownLines"`
    ItemCount int64 `json:"ItemCount,omitempty" xml:"ItemCount"`
    SelectedCells []LinkElement `json:"SelectedCells,omitempty" xml:"SelectedCells"`
    SelectionType string `json:"SelectionType,omitempty" xml:"SelectionType"`
    IsChecked bool `json:"IsChecked,omitempty" xml:"IsChecked"`
    CurrentValue int64 `json:"CurrentValue,omitempty" xml:"CurrentValue"`
    Min int64 `json:"Min,omitempty" xml:"Min"`
    Max int64 `json:"Max,omitempty" xml:"Max"`
    IncrementalChange int64 `json:"IncrementalChange,omitempty" xml:"IncrementalChange"`
    PageChange int64 `json:"PageChange,omitempty" xml:"PageChange"`
    IsHorizontal bool `json:"IsHorizontal,omitempty" xml:"IsHorizontal"`
}


type Forms struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    FormList []LinkElement `json:"FormList,omitempty" xml:"FormList"`
}


type GradientFill struct {
 
    FillType string `json:"FillType,omitempty" xml:"FillType"`
    DirectionType string `json:"DirectionType,omitempty" xml:"DirectionType"`
    Angle float64 `json:"Angle,omitempty" xml:"Angle"`
    GradientStops []GradientFillStop `json:"GradientStops,omitempty" xml:"GradientStops"`
}


type GradientFillStop struct {
 
    Color *Color `json:"Color,omitempty" xml:"Color"`
    Position float64 `json:"Position,omitempty" xml:"Position"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}


type GroupBox struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
}


type GroupShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type Label struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type Line struct {
 
    BeginArrowLength string `json:"BeginArrowLength,omitempty" xml:"BeginArrowLength"`
    BeginArrowWidth string `json:"BeginArrowWidth,omitempty" xml:"BeginArrowWidth"`
    BeginType string `json:"BeginType,omitempty" xml:"BeginType"`
    CapType string `json:"CapType,omitempty" xml:"CapType"`
    Color *Color `json:"Color,omitempty" xml:"Color"`
    CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
    DashType string `json:"DashType,omitempty" xml:"DashType"`
    EndArrowLength string `json:"EndArrowLength,omitempty" xml:"EndArrowLength"`
    EndArrowWidth string `json:"EndArrowWidth,omitempty" xml:"EndArrowWidth"`
    EndType string `json:"EndType,omitempty" xml:"EndType"`
    GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
    IsAuto bool `json:"IsAuto,omitempty" xml:"IsAuto"`
    IsAutomaticColor bool `json:"IsAutomaticColor,omitempty" xml:"IsAutomaticColor"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
    Style string `json:"Style,omitempty" xml:"Style"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
    Weight string `json:"Weight,omitempty" xml:"Weight"`
    WeightPt float64 `json:"WeightPt,omitempty" xml:"WeightPt"`
}


type LineFormat struct {
     
        Type_ string `json:"Type,omitempty" xml:"Type"`
        SolidFill *SolidFill `json:"SolidFill,omitempty" xml:"SolidFill"`
        PatternFill *PatternFill `json:"PatternFill,omitempty" xml:"PatternFill"`
        TextureFill *TextureFill `json:"TextureFill,omitempty" xml:"TextureFill"`
        GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
        ImageData string `json:"ImageData,omitempty" xml:"ImageData"`
 
    BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
    BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
    BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
    CapType string `json:"CapType,omitempty" xml:"CapType"`
    CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
    DashStyle string `json:"DashStyle,omitempty" xml:"DashStyle"`
    EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
    EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
    EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
    JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
    Weight float64 `json:"Weight,omitempty" xml:"Weight"`
}


type LineShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    BeginArrowheadLength string `json:"BeginArrowheadLength,omitempty" xml:"BeginArrowheadLength"`
    BeginArrowheadStyle string `json:"BeginArrowheadStyle,omitempty" xml:"BeginArrowheadStyle"`
    BeginArrowheadWidth string `json:"BeginArrowheadWidth,omitempty" xml:"BeginArrowheadWidth"`
    EndArrowheadLength string `json:"EndArrowheadLength,omitempty" xml:"EndArrowheadLength"`
    EndArrowheadStyle string `json:"EndArrowheadStyle,omitempty" xml:"EndArrowheadStyle"`
    EndArrowheadWidth string `json:"EndArrowheadWidth,omitempty" xml:"EndArrowheadWidth"`
}


type ListBox struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    InputRange string `json:"InputRange,omitempty" xml:"InputRange"`
    ItemCount int64 `json:"ItemCount,omitempty" xml:"ItemCount"`
    PageChange int64 `json:"PageChange,omitempty" xml:"PageChange"`
    SelectedCells []LinkElement `json:"SelectedCells,omitempty" xml:"SelectedCells"`
    SelectedIndex int64 `json:"SelectedIndex,omitempty" xml:"SelectedIndex"`
    SelectionType string `json:"SelectionType,omitempty" xml:"SelectionType"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
}


type OleObject struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    DisplayAsIcon bool `json:"DisplayAsIcon,omitempty" xml:"DisplayAsIcon"`
    FileFormatType string `json:"FileFormatType,omitempty" xml:"FileFormatType"`
    ImageSourceFullName string `json:"ImageSourceFullName,omitempty" xml:"ImageSourceFullName"`
    IsAutoSize bool `json:"IsAutoSize,omitempty" xml:"IsAutoSize"`
    IsLink bool `json:"IsLink,omitempty" xml:"IsLink"`
    ProgID string `json:"ProgID,omitempty" xml:"ProgID"`
    SourceFullName string `json:"SourceFullName,omitempty" xml:"SourceFullName"`
}


type OleObjects struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    OleObjectList []LinkElement `json:"OleObjectList,omitempty" xml:"OleObjectList"`
}


type Oval struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type PatternFill struct {
 
    Pattern string `json:"Pattern,omitempty" xml:"Pattern"`
    BackgroundCellsColor *CellsColor `json:"BackgroundCellsColor,omitempty" xml:"BackgroundCellsColor"`
    ForegroundCellsColor *CellsColor `json:"ForegroundCellsColor,omitempty" xml:"ForegroundCellsColor"`
    ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
    BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
    BackTransparency float64 `json:"BackTransparency,omitempty" xml:"BackTransparency"`
    ForeTransparency float64 `json:"ForeTransparency,omitempty" xml:"ForeTransparency"`
}


type PicFormatOption struct {
 
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Scale float64 `json:"Scale,omitempty" xml:"Scale"`
    Left float64 `json:"Left,omitempty" xml:"Left"`
    Right float64 `json:"Right,omitempty" xml:"Right"`
    Top float64 `json:"Top,omitempty" xml:"Top"`
    Bottom float64 `json:"Bottom,omitempty" xml:"Bottom"`
}


type Picture struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    BorderLineColor *Color `json:"BorderLineColor,omitempty" xml:"BorderLineColor"`
    BorderWeight float64 `json:"BorderWeight,omitempty" xml:"BorderWeight"`
    OriginalHeight int64 `json:"OriginalHeight,omitempty" xml:"OriginalHeight"`
    OriginalWidth int64 `json:"OriginalWidth,omitempty" xml:"OriginalWidth"`
    ImageFormat string `json:"ImageFormat,omitempty" xml:"ImageFormat"`
    SourceFullName string `json:"SourceFullName,omitempty" xml:"SourceFullName"`
}


type Pictures struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    PictureList []LinkElement `json:"PictureList,omitempty" xml:"PictureList"`
}


type RadioButton struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    GroupBox *GroupBox `json:"GroupBox,omitempty" xml:"GroupBox"`
    IsChecked bool `json:"IsChecked,omitempty" xml:"IsChecked"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
}


type RectangleShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type ScrollBar struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    CurrentValue int64 `json:"CurrentValue,omitempty" xml:"CurrentValue"`
    IncrementalChange int64 `json:"IncrementalChange,omitempty" xml:"IncrementalChange"`
    IsHorizontal bool `json:"IsHorizontal,omitempty" xml:"IsHorizontal"`
    Max int64 `json:"Max,omitempty" xml:"Max"`
    Min int64 `json:"Min,omitempty" xml:"Min"`
    PageChange int64 `json:"PageChange,omitempty" xml:"PageChange"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
}


type ShadowEffect struct {
 
    Angle float64 `json:"Angle,omitempty" xml:"Angle"`
    Blur float64 `json:"Blur,omitempty" xml:"Blur"`
    Color *CellsColor `json:"Color,omitempty" xml:"Color"`
    Distance float64 `json:"Distance,omitempty" xml:"Distance"`
    PresetType string `json:"PresetType,omitempty" xml:"PresetType"`
    Size float64 `json:"Size,omitempty" xml:"Size"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}


type Shape struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Name string `json:"Name,omitempty" xml:"Name"`
    MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
    AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
    Placement string `json:"Placement,omitempty" xml:"Placement"`
    UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
    Top int64 `json:"Top,omitempty" xml:"Top"`
    UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
    Left int64 `json:"Left,omitempty" xml:"Left"`
    LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
    Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
    LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
    Right int64 `json:"Right,omitempty" xml:"Right"`
    Width int64 `json:"Width,omitempty" xml:"Width"`
    Height int64 `json:"Height,omitempty" xml:"Height"`
    X int64 `json:"X,omitempty" xml:"X"`
    Y int64 `json:"Y,omitempty" xml:"Y"`
    RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
    Text string `json:"Text,omitempty" xml:"Text"`
    AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
    TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
    TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
    IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
    IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
    IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
    IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
    IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
    IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
    IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
    LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
    ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
    Font *Font `json:"Font,omitempty" xml:"Font"`
    Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
}


type Shapes struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    ShapeList []LinkElement `json:"ShapeList,omitempty" xml:"ShapeList"`
}


type SolidFill struct {
 
    Color *Color `json:"Color,omitempty" xml:"Color"`
    CellsColor *CellsColor `json:"CellsColor,omitempty" xml:"CellsColor"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
}


type Spinner struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
    CurrentValue int64 `json:"CurrentValue,omitempty" xml:"CurrentValue"`
    IncrementalChange int64 `json:"IncrementalChange,omitempty" xml:"IncrementalChange"`
    Max int64 `json:"Max,omitempty" xml:"Max"`
    Min int64 `json:"Min,omitempty" xml:"Min"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
}


type TextBox struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type TextureFill struct {
 
    Type_ string `json:"Type,omitempty" xml:"Type"`
    Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
    Scale float64 `json:"Scale,omitempty" xml:"Scale"`
    TilePicOption *TilePicOption `json:"TilePicOption,omitempty" xml:"TilePicOption"`
    PicFormatOption *PicFormatOption `json:"PicFormatOption,omitempty" xml:"PicFormatOption"`
    Image *LinkElement `json:"Image,omitempty" xml:"Image"`
}


type TilePicOption struct {
 
    OffsetX float64 `json:"OffsetX,omitempty" xml:"OffsetX"`
    OffsetY float64 `json:"OffsetY,omitempty" xml:"OffsetY"`
    ScaleX float64 `json:"ScaleX,omitempty" xml:"ScaleX"`
    ScaleY float64 `json:"ScaleY,omitempty" xml:"ScaleY"`
    AlignmentType string `json:"AlignmentType,omitempty" xml:"AlignmentType"`
    MirrorType string `json:"MirrorType,omitempty" xml:"MirrorType"`
}


type DigitalSignature struct {
 
    Comments string `json:"Comments,omitempty" xml:"Comments"`
    SignTime string `json:"SignTime,omitempty" xml:"SignTime"`
    Id string `json:"Id,omitempty" xml:"Id"`
    Password string `json:"Password,omitempty" xml:"Password"`
    Image []int64 `json:"Image,omitempty" xml:"Image"`
    ProviderId string `json:"ProviderId,omitempty" xml:"ProviderId"`
    IsValid bool `json:"IsValid,omitempty" xml:"IsValid"`
    XAdESType string `json:"XAdESType,omitempty" xml:"XAdESType"`
}


type Axis struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    AxisBetweenCategories bool `json:"AxisBetweenCategories,omitempty" xml:"AxisBetweenCategories"`
    AxisLine *Line `json:"AxisLine,omitempty" xml:"AxisLine"`
    BaseUnitScale string `json:"BaseUnitScale,omitempty" xml:"BaseUnitScale"`
    CategoryType string `json:"CategoryType,omitempty" xml:"CategoryType"`
    CrossAt float64 `json:"CrossAt,omitempty" xml:"CrossAt"`
    CrossType string `json:"CrossType,omitempty" xml:"CrossType"`
    DisplayUnit string `json:"DisplayUnit,omitempty" xml:"DisplayUnit"`
    DisplayUnitLabel *DisplayUnitLabel `json:"DisplayUnitLabel,omitempty" xml:"DisplayUnitLabel"`
    HasMultiLevelLabels bool `json:"HasMultiLevelLabels,omitempty" xml:"HasMultiLevelLabels"`
    IsAutomaticMajorUnit bool `json:"IsAutomaticMajorUnit,omitempty" xml:"IsAutomaticMajorUnit"`
    IsAutomaticMaxValue bool `json:"IsAutomaticMaxValue,omitempty" xml:"IsAutomaticMaxValue"`
    IsAutomaticMinorUnit bool `json:"IsAutomaticMinorUnit,omitempty" xml:"IsAutomaticMinorUnit"`
    IsAutomaticMinValue bool `json:"IsAutomaticMinValue,omitempty" xml:"IsAutomaticMinValue"`
    IsDisplayUnitLabelShown bool `json:"IsDisplayUnitLabelShown,omitempty" xml:"IsDisplayUnitLabelShown"`
    IsLogarithmic bool `json:"IsLogarithmic,omitempty" xml:"IsLogarithmic"`
    IsPlotOrderReversed bool `json:"IsPlotOrderReversed,omitempty" xml:"IsPlotOrderReversed"`
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    LogBase float64 `json:"LogBase,omitempty" xml:"LogBase"`
    MajorGridLines *Line `json:"MajorGridLines,omitempty" xml:"MajorGridLines"`
    MajorTickMark string `json:"MajorTickMark,omitempty" xml:"MajorTickMark"`
    MajorUnit float64 `json:"MajorUnit,omitempty" xml:"MajorUnit"`
    MajorUnitScale string `json:"MajorUnitScale,omitempty" xml:"MajorUnitScale"`
    MaxValue float64 `json:"MaxValue,omitempty" xml:"MaxValue"`
    MinorGridLines *Line `json:"MinorGridLines,omitempty" xml:"MinorGridLines"`
    MinorTickMark string `json:"MinorTickMark,omitempty" xml:"MinorTickMark"`
    MinorUnit float64 `json:"MinorUnit,omitempty" xml:"MinorUnit"`
    MinorUnitScale string `json:"MinorUnitScale,omitempty" xml:"MinorUnitScale"`
    MinValue float64 `json:"MinValue,omitempty" xml:"MinValue"`
    TickLabelPosition string `json:"TickLabelPosition,omitempty" xml:"TickLabelPosition"`
    TickLabels *TickLabels `json:"TickLabels,omitempty" xml:"TickLabels"`
    TickLabelSpacing int64 `json:"TickLabelSpacing,omitempty" xml:"TickLabelSpacing"`
    TickMarkSpacing int64 `json:"TickMarkSpacing,omitempty" xml:"TickMarkSpacing"`
    Title *Title `json:"Title,omitempty" xml:"Title"`
}


type Chart struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoScaling bool `json:"AutoScaling,omitempty" xml:"AutoScaling"`
    BackWall *Walls `json:"BackWall,omitempty" xml:"BackWall"`
    CategoryAxis *Axis `json:"CategoryAxis,omitempty" xml:"CategoryAxis"`
    ChartArea *ChartArea `json:"ChartArea,omitempty" xml:"ChartArea"`
    ChartDataTable *ChartDataTable `json:"ChartDataTable,omitempty" xml:"ChartDataTable"`
    ChartObject *LinkElement `json:"ChartObject,omitempty" xml:"ChartObject"`
    DepthPercent int64 `json:"DepthPercent,omitempty" xml:"DepthPercent"`
    Elevation int64 `json:"Elevation,omitempty" xml:"Elevation"`
    FirstSliceAngle int64 `json:"FirstSliceAngle,omitempty" xml:"FirstSliceAngle"`
    Floor *Floor `json:"Floor,omitempty" xml:"Floor"`
    GapDepth int64 `json:"GapDepth,omitempty" xml:"GapDepth"`
    GapWidth int64 `json:"GapWidth,omitempty" xml:"GapWidth"`
    HeightPercent int64 `json:"HeightPercent,omitempty" xml:"HeightPercent"`
    HidePivotFieldButtons bool `json:"HidePivotFieldButtons,omitempty" xml:"HidePivotFieldButtons"`
    Is3D bool `json:"Is3D,omitempty" xml:"Is3D"`
    IsRectangularCornered bool `json:"IsRectangularCornered,omitempty" xml:"IsRectangularCornered"`
    Legend *Legend `json:"Legend,omitempty" xml:"Legend"`
    Name string `json:"Name,omitempty" xml:"Name"`
    NSeries *SeriesItems `json:"NSeries,omitempty" xml:"NSeries"`
    PageSetup *LinkElement `json:"PageSetup,omitempty" xml:"PageSetup"`
    Perspective int64 `json:"Perspective,omitempty" xml:"Perspective"`
    PivotSource string `json:"PivotSource,omitempty" xml:"PivotSource"`
    Placement string `json:"Placement,omitempty" xml:"Placement"`
    PlotArea *PlotArea `json:"PlotArea,omitempty" xml:"PlotArea"`
    PlotEmptyCellsType string `json:"PlotEmptyCellsType,omitempty" xml:"PlotEmptyCellsType"`
    PlotVisibleCells bool `json:"PlotVisibleCells,omitempty" xml:"PlotVisibleCells"`
    PrintSize string `json:"PrintSize,omitempty" xml:"PrintSize"`
    RightAngleAxes bool `json:"RightAngleAxes,omitempty" xml:"RightAngleAxes"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    SecondCategoryAxis *LinkElement `json:"SecondCategoryAxis,omitempty" xml:"SecondCategoryAxis"`
    SecondValueAxis *LinkElement `json:"SecondValueAxis,omitempty" xml:"SecondValueAxis"`
    SeriesAxis *LinkElement `json:"SeriesAxis,omitempty" xml:"SeriesAxis"`
    Shapes *LinkElement `json:"Shapes,omitempty" xml:"Shapes"`
    ShowDataTable bool `json:"ShowDataTable,omitempty" xml:"ShowDataTable"`
    ShowLegend bool `json:"ShowLegend,omitempty" xml:"ShowLegend"`
    SideWall *LinkElement `json:"SideWall,omitempty" xml:"SideWall"`
    SizeWithWindow bool `json:"SizeWithWindow,omitempty" xml:"SizeWithWindow"`
    Style int64 `json:"Style,omitempty" xml:"Style"`
    Title *LinkElement `json:"Title,omitempty" xml:"Title"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    ValueAxis *Axis `json:"ValueAxis,omitempty" xml:"ValueAxis"`
    Walls *LinkElement `json:"Walls,omitempty" xml:"Walls"`
    WallsAndGridlines2D bool `json:"WallsAndGridlines2D,omitempty" xml:"WallsAndGridlines2D"`
}


type ChartArea struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
}


type ChartDataTable struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
    BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
    Font *Font `json:"Font,omitempty" xml:"Font"`
    HasBorderHorizontal bool `json:"HasBorderHorizontal,omitempty" xml:"HasBorderHorizontal"`
    HasBorderOutline bool `json:"HasBorderOutline,omitempty" xml:"HasBorderOutline"`
    HasBorderVertical bool `json:"HasBorderVertical,omitempty" xml:"HasBorderVertical"`
    ShowLegendKey bool `json:"ShowLegendKey,omitempty" xml:"ShowLegendKey"`
}


type ChartFrame struct {
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
    BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
    Font *Font `json:"Font,omitempty" xml:"Font"`
    IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
    IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    Width int64 `json:"Width,omitempty" xml:"Width"`
    Height int64 `json:"Height,omitempty" xml:"Height"`
    X int64 `json:"X,omitempty" xml:"X"`
    Y int64 `json:"Y,omitempty" xml:"Y"`
}


type ChartGlobalizationSettings struct {
 
}


type ChartPoint struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
    Explosion int64 `json:"Explosion,omitempty" xml:"Explosion"`
    Marker *Marker `json:"Marker,omitempty" xml:"Marker"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    XValue *interface{} `json:"XValue,omitempty" xml:"XValue"`
    YValue *interface{} `json:"YValue,omitempty" xml:"YValue"`
    IsInSecondaryPlot bool `json:"IsInSecondaryPlot,omitempty" xml:"IsInSecondaryPlot"`
}


type ChartPoints struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    ChartPointList []ChartPoint `json:"ChartPointList,omitempty" xml:"ChartPointList"`
}


type Charts struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    ChartList []LinkElement `json:"ChartList,omitempty" xml:"ChartList"`
}


type ChartShape struct {
     
        Name string `json:"Name,omitempty" xml:"Name"`
        MsoDrawingType string `json:"MsoDrawingType,omitempty" xml:"MsoDrawingType"`
        AutoShapeType string `json:"AutoShapeType,omitempty" xml:"AutoShapeType"`
        Placement string `json:"Placement,omitempty" xml:"Placement"`
        UpperLeftRow int64 `json:"UpperLeftRow,omitempty" xml:"UpperLeftRow"`
        Top int64 `json:"Top,omitempty" xml:"Top"`
        UpperLeftColumn int64 `json:"UpperLeftColumn,omitempty" xml:"UpperLeftColumn"`
        Left int64 `json:"Left,omitempty" xml:"Left"`
        LowerRightRow int64 `json:"LowerRightRow,omitempty" xml:"LowerRightRow"`
        Bottom int64 `json:"Bottom,omitempty" xml:"Bottom"`
        LowerRightColumn int64 `json:"LowerRightColumn,omitempty" xml:"LowerRightColumn"`
        Right int64 `json:"Right,omitempty" xml:"Right"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
        RotationAngle float64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
        HtmlText string `json:"HtmlText,omitempty" xml:"HtmlText"`
        Text string `json:"Text,omitempty" xml:"Text"`
        AlternativeText string `json:"AlternativeText,omitempty" xml:"AlternativeText"`
        TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
        TextHorizontalOverflow string `json:"TextHorizontalOverflow,omitempty" xml:"TextHorizontalOverflow"`
        TextOrientationType string `json:"TextOrientationType,omitempty" xml:"TextOrientationType"`
        TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
        TextVerticalOverflow string `json:"TextVerticalOverflow,omitempty" xml:"TextVerticalOverflow"`
        IsGroup bool `json:"IsGroup,omitempty" xml:"IsGroup"`
        IsHidden bool `json:"IsHidden,omitempty" xml:"IsHidden"`
        IsLockAspectRatio bool `json:"IsLockAspectRatio,omitempty" xml:"IsLockAspectRatio"`
        IsLocked bool `json:"IsLocked,omitempty" xml:"IsLocked"`
        IsPrintable bool `json:"IsPrintable,omitempty" xml:"IsPrintable"`
        IsTextWrapped bool `json:"IsTextWrapped,omitempty" xml:"IsTextWrapped"`
        IsWordArt bool `json:"IsWordArt,omitempty" xml:"IsWordArt"`
        LinkedCell string `json:"LinkedCell,omitempty" xml:"LinkedCell"`
        ZOrderPosition int64 `json:"ZOrderPosition,omitempty" xml:"ZOrderPosition"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        Hyperlink string `json:"Hyperlink,omitempty" xml:"Hyperlink"`
 
}


type DataLabels struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    IsAutoText bool `json:"IsAutoText,omitempty" xml:"IsAutoText"`
    IsDeleted bool `json:"IsDeleted,omitempty" xml:"IsDeleted"`
    LinkedSource string `json:"LinkedSource,omitempty" xml:"LinkedSource"`
    Number int64 `json:"Number,omitempty" xml:"Number"`
    NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
    NumberFormatLinked bool `json:"NumberFormatLinked,omitempty" xml:"NumberFormatLinked"`
    Position string `json:"Position,omitempty" xml:"Position"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    Separator string `json:"Separator,omitempty" xml:"Separator"`
    ShowBubbleSize bool `json:"ShowBubbleSize,omitempty" xml:"ShowBubbleSize"`
    ShowCategoryName bool `json:"ShowCategoryName,omitempty" xml:"ShowCategoryName"`
    ShowLegendKey bool `json:"ShowLegendKey,omitempty" xml:"ShowLegendKey"`
    ShowPercentage bool `json:"ShowPercentage,omitempty" xml:"ShowPercentage"`
    ShowSeriesName bool `json:"ShowSeriesName,omitempty" xml:"ShowSeriesName"`
    ShowValue bool `json:"ShowValue,omitempty" xml:"ShowValue"`
    Text string `json:"Text,omitempty" xml:"Text"`
    TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}


type DisplayUnitLabel struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    LinkedSource string `json:"LinkedSource,omitempty" xml:"LinkedSource"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    Text string `json:"Text,omitempty" xml:"Text"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}


type DropBars struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
}


type ErrorBar struct {
     
        BeginArrowLength string `json:"BeginArrowLength,omitempty" xml:"BeginArrowLength"`
        BeginArrowWidth string `json:"BeginArrowWidth,omitempty" xml:"BeginArrowWidth"`
        BeginType string `json:"BeginType,omitempty" xml:"BeginType"`
        CapType string `json:"CapType,omitempty" xml:"CapType"`
        Color *Color `json:"Color,omitempty" xml:"Color"`
        CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
        DashType string `json:"DashType,omitempty" xml:"DashType"`
        EndArrowLength string `json:"EndArrowLength,omitempty" xml:"EndArrowLength"`
        EndArrowWidth string `json:"EndArrowWidth,omitempty" xml:"EndArrowWidth"`
        EndType string `json:"EndType,omitempty" xml:"EndType"`
        GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
        IsAuto bool `json:"IsAuto,omitempty" xml:"IsAuto"`
        IsAutomaticColor bool `json:"IsAutomaticColor,omitempty" xml:"IsAutomaticColor"`
        IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
        JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
        Style string `json:"Style,omitempty" xml:"Style"`
        Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
        Weight string `json:"Weight,omitempty" xml:"Weight"`
        WeightPt float64 `json:"WeightPt,omitempty" xml:"WeightPt"`
 
    Link *Link `json:"Link,omitempty" xml:"Link"`
    Amount float64 `json:"Amount,omitempty" xml:"Amount"`
    DisplayType string `json:"DisplayType,omitempty" xml:"DisplayType"`
    MinusValue string `json:"MinusValue,omitempty" xml:"MinusValue"`
    PlusValue string `json:"PlusValue,omitempty" xml:"PlusValue"`
    ShowMarkerTTop bool `json:"ShowMarkerTTop,omitempty" xml:"ShowMarkerTTop"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type Floor struct {
     
        BackgroundColor *Color `json:"BackgroundColor,omitempty" xml:"BackgroundColor"`
        FillFormat *FillFormat `json:"FillFormat,omitempty" xml:"FillFormat"`
        ForegroundColor *Color `json:"ForegroundColor,omitempty" xml:"ForegroundColor"`
        Format string `json:"Format,omitempty" xml:"Format"`
        InvertIfNegative bool `json:"InvertIfNegative,omitempty" xml:"InvertIfNegative"`
        Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
 
    Border *Line `json:"Border,omitempty" xml:"Border"`
}


type Legend struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    Position string `json:"Position,omitempty" xml:"Position"`
    LegendEntries *LinkElement `json:"LegendEntries,omitempty" xml:"LegendEntries"`
}


type LegendEntries struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    LegendEntryList []LinkElement `json:"legendEntryList,omitempty" xml:"legendEntryList"`
}


type LegendEntry struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
    BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
    Font *Font `json:"Font,omitempty" xml:"Font"`
    IsDeleted bool `json:"IsDeleted,omitempty" xml:"IsDeleted"`
}


type Marker struct {
 
    Border *Line `json:"Border,omitempty" xml:"Border"`
    Area *Area `json:"Area,omitempty" xml:"Area"`
    MarkerSize int64 `json:"MarkerSize,omitempty" xml:"MarkerSize"`
    MarkerStyle string `json:"MarkerStyle,omitempty" xml:"MarkerStyle"`
}


type PlotArea struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    InnerHeight int64 `json:"InnerHeight,omitempty" xml:"InnerHeight"`
    InnerWidth int64 `json:"InnerWidth,omitempty" xml:"InnerWidth"`
    InnerX int64 `json:"InnerX,omitempty" xml:"InnerX"`
    InnerY int64 `json:"InnerY,omitempty" xml:"InnerY"`
}


type Series struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    Area *Area `json:"Area,omitempty" xml:"Area"`
    Bar3DShapeType string `json:"Bar3DShapeType,omitempty" xml:"Bar3DShapeType"`
    Border *Line `json:"Border,omitempty" xml:"Border"`
    BubbleScale int64 `json:"BubbleScale,omitempty" xml:"BubbleScale"`
    BubbleSizes string `json:"BubbleSizes,omitempty" xml:"BubbleSizes"`
    CountOfDataValues int64 `json:"CountOfDataValues,omitempty" xml:"CountOfDataValues"`
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
    DisplayName string `json:"DisplayName,omitempty" xml:"DisplayName"`
    DoughnutHoleSize int64 `json:"DoughnutHoleSize,omitempty" xml:"DoughnutHoleSize"`
    DownBars *DropBars `json:"DownBars,omitempty" xml:"DownBars"`
    DropLines *Line `json:"DropLines,omitempty" xml:"DropLines"`
    Explosion int64 `json:"Explosion,omitempty" xml:"Explosion"`
    FirstSliceAngle int64 `json:"FirstSliceAngle,omitempty" xml:"FirstSliceAngle"`
    GapWidth int64 `json:"GapWidth,omitempty" xml:"GapWidth"`
    Has3DEffect bool `json:"Has3DEffect,omitempty" xml:"Has3DEffect"`
    HasDropLines bool `json:"HasDropLines,omitempty" xml:"HasDropLines"`
    HasHiLoLines bool `json:"HasHiLoLines,omitempty" xml:"HasHiLoLines"`
    HasLeaderLines bool `json:"HasLeaderLines,omitempty" xml:"HasLeaderLines"`
    HasRadarAxisLabels bool `json:"HasRadarAxisLabels,omitempty" xml:"HasRadarAxisLabels"`
    HasSeriesLines bool `json:"HasSeriesLines,omitempty" xml:"HasSeriesLines"`
    HasUpDownBars bool `json:"HasUpDownBars,omitempty" xml:"HasUpDownBars"`
    HiLoLines *Line `json:"HiLoLines,omitempty" xml:"HiLoLines"`
    IsAutoSplit bool `json:"IsAutoSplit,omitempty" xml:"IsAutoSplit"`
    IsColorVaried bool `json:"IsColorVaried,omitempty" xml:"IsColorVaried"`
    LeaderLines *Line `json:"LeaderLines,omitempty" xml:"LeaderLines"`
    LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
    Marker *Marker `json:"Marker,omitempty" xml:"Marker"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Overlap int64 `json:"Overlap,omitempty" xml:"Overlap"`
    PlotOnSecondAxis bool `json:"PlotOnSecondAxis,omitempty" xml:"PlotOnSecondAxis"`
    Points *LinkElement `json:"Points,omitempty" xml:"Points"`
    SecondPlotSize int64 `json:"SecondPlotSize,omitempty" xml:"SecondPlotSize"`
    SeriesLines *Line `json:"SeriesLines,omitempty" xml:"SeriesLines"`
    Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
    ShowNegativeBubbles bool `json:"ShowNegativeBubbles,omitempty" xml:"ShowNegativeBubbles"`
    SizeRepresents string `json:"SizeRepresents,omitempty" xml:"SizeRepresents"`
    Smooth bool `json:"Smooth,omitempty" xml:"Smooth"`
    SplitType string `json:"SplitType,omitempty" xml:"SplitType"`
    SplitValue float64 `json:"SplitValue,omitempty" xml:"SplitValue"`
    TrendLines *Trendlines `json:"TrendLines,omitempty" xml:"TrendLines"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    UpBars *DropBars `json:"UpBars,omitempty" xml:"UpBars"`
    Values string `json:"Values,omitempty" xml:"Values"`
    XErrorBar *ErrorBar `json:"XErrorBar,omitempty" xml:"XErrorBar"`
    XValues string `json:"XValues,omitempty" xml:"XValues"`
    YErrorBar *ErrorBar `json:"YErrorBar,omitempty" xml:"YErrorBar"`
}


type SeriesItems struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    CategoryData string `json:"CategoryData,omitempty" xml:"CategoryData"`
    IsColorVaried bool `json:"IsColorVaried,omitempty" xml:"IsColorVaried"`
    SecondCatergoryData string `json:"SecondCatergoryData,omitempty" xml:"SecondCatergoryData"`
    SeriesList []Series `json:"SeriesList,omitempty" xml:"SeriesList"`
}


type Sparkline struct {
 
    Column int64 `json:"Column,omitempty" xml:"Column"`
    DataRange string `json:"DataRange,omitempty" xml:"DataRange"`
    Row int64 `json:"Row,omitempty" xml:"Row"`
}


type SparklineGroup struct {
 
    DisplayHidden bool `json:"DisplayHidden,omitempty" xml:"DisplayHidden"`
    FirstPointColor *CellsColor `json:"FirstPointColor,omitempty" xml:"FirstPointColor"`
    HighPointColor *CellsColor `json:"HighPointColor,omitempty" xml:"HighPointColor"`
    HorizontalAxisColor *CellsColor `json:"HorizontalAxisColor,omitempty" xml:"HorizontalAxisColor"`
    HorizontalAxisDateRange string `json:"HorizontalAxisDateRange,omitempty" xml:"HorizontalAxisDateRange"`
    LastPointColor *CellsColor `json:"LastPointColor,omitempty" xml:"LastPointColor"`
    LineWeight float64 `json:"LineWeight,omitempty" xml:"LineWeight"`
    LowPointColor *CellsColor `json:"LowPointColor,omitempty" xml:"LowPointColor"`
    MarkersColor *CellsColor `json:"MarkersColor,omitempty" xml:"MarkersColor"`
    NegativePointsColor *CellsColor `json:"NegativePointsColor,omitempty" xml:"NegativePointsColor"`
    PlotEmptyCellsType string `json:"PlotEmptyCellsType,omitempty" xml:"PlotEmptyCellsType"`
    PlotRightToLeft bool `json:"PlotRightToLeft,omitempty" xml:"PlotRightToLeft"`
    PresetStyle string `json:"PresetStyle,omitempty" xml:"PresetStyle"`
    SeriesColor *CellsColor `json:"SeriesColor,omitempty" xml:"SeriesColor"`
    ShowFirstPoint bool `json:"ShowFirstPoint,omitempty" xml:"ShowFirstPoint"`
    ShowHighPoint bool `json:"ShowHighPoint,omitempty" xml:"ShowHighPoint"`
    ShowHorizontalAxis bool `json:"ShowHorizontalAxis,omitempty" xml:"ShowHorizontalAxis"`
    ShowLastPoint bool `json:"ShowLastPoint,omitempty" xml:"ShowLastPoint"`
    ShowLowPoint bool `json:"ShowLowPoint,omitempty" xml:"ShowLowPoint"`
    ShowMarkers bool `json:"ShowMarkers,omitempty" xml:"ShowMarkers"`
    ShowNegativePoints bool `json:"ShowNegativePoints,omitempty" xml:"ShowNegativePoints"`
    SparklineCollection []Sparkline `json:"SparklineCollection,omitempty" xml:"SparklineCollection"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    VerticalAxisMaxValue float64 `json:"VerticalAxisMaxValue,omitempty" xml:"VerticalAxisMaxValue"`
    VerticalAxisMaxValueType string `json:"VerticalAxisMaxValueType,omitempty" xml:"VerticalAxisMaxValueType"`
    VerticalAxisMinValue float64 `json:"VerticalAxisMinValue,omitempty" xml:"VerticalAxisMinValue"`
    VerticalAxisMinValueType string `json:"VerticalAxisMinValueType,omitempty" xml:"VerticalAxisMinValueType"`
}


type SparklineGroups struct {
 
    SparklineGroupList []SparklineGroup `json:"SparklineGroupList,omitempty" xml:"SparklineGroupList"`
}


type TickLabels struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
    BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
    Font *Font `json:"Font,omitempty" xml:"Font"`
    Number int64 `json:"Number,omitempty" xml:"Number"`
    NumberFormat string `json:"NumberFormat,omitempty" xml:"NumberFormat"`
    NumberFormatLinked bool `json:"NumberFormatLinked,omitempty" xml:"NumberFormatLinked"`
    Offset int64 `json:"Offset,omitempty" xml:"Offset"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
    ReadingOrder string `json:"ReadingOrder,omitempty" xml:"ReadingOrder"`
    DirectionType string `json:"DirectionType,omitempty" xml:"DirectionType"`
}


type Title struct {
     
        Area *Area `json:"Area,omitempty" xml:"Area"`
        AutoScaleFont bool `json:"AutoScaleFont,omitempty" xml:"AutoScaleFont"`
        BackgroundMode string `json:"BackgroundMode,omitempty" xml:"BackgroundMode"`
        Border *Line `json:"Border,omitempty" xml:"Border"`
        Font *Font `json:"Font,omitempty" xml:"Font"`
        IsAutomaticSize bool `json:"IsAutomaticSize,omitempty" xml:"IsAutomaticSize"`
        IsInnerMode bool `json:"IsInnerMode,omitempty" xml:"IsInnerMode"`
        Shadow bool `json:"Shadow,omitempty" xml:"Shadow"`
        Width int64 `json:"Width,omitempty" xml:"Width"`
        Height int64 `json:"Height,omitempty" xml:"Height"`
        X int64 `json:"X,omitempty" xml:"X"`
        Y int64 `json:"Y,omitempty" xml:"Y"`
 
    IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
    LinkedSource string `json:"LinkedSource,omitempty" xml:"LinkedSource"`
    RotationAngle int64 `json:"RotationAngle,omitempty" xml:"RotationAngle"`
    Text string `json:"Text,omitempty" xml:"Text"`
    TextDirection string `json:"TextDirection,omitempty" xml:"TextDirection"`
    TextHorizontalAlignment string `json:"TextHorizontalAlignment,omitempty" xml:"TextHorizontalAlignment"`
    TextVerticalAlignment string `json:"TextVerticalAlignment,omitempty" xml:"TextVerticalAlignment"`
}


type Trendline struct {
     
        BeginArrowLength string `json:"BeginArrowLength,omitempty" xml:"BeginArrowLength"`
        BeginArrowWidth string `json:"BeginArrowWidth,omitempty" xml:"BeginArrowWidth"`
        BeginType string `json:"BeginType,omitempty" xml:"BeginType"`
        CapType string `json:"CapType,omitempty" xml:"CapType"`
        Color *Color `json:"Color,omitempty" xml:"Color"`
        CompoundType string `json:"CompoundType,omitempty" xml:"CompoundType"`
        DashType string `json:"DashType,omitempty" xml:"DashType"`
        EndArrowLength string `json:"EndArrowLength,omitempty" xml:"EndArrowLength"`
        EndArrowWidth string `json:"EndArrowWidth,omitempty" xml:"EndArrowWidth"`
        EndType string `json:"EndType,omitempty" xml:"EndType"`
        GradientFill *GradientFill `json:"GradientFill,omitempty" xml:"GradientFill"`
        IsAuto bool `json:"IsAuto,omitempty" xml:"IsAuto"`
        IsAutomaticColor bool `json:"IsAutomaticColor,omitempty" xml:"IsAutomaticColor"`
        IsVisible bool `json:"IsVisible,omitempty" xml:"IsVisible"`
        JoinType string `json:"JoinType,omitempty" xml:"JoinType"`
        Style string `json:"Style,omitempty" xml:"Style"`
        Transparency float64 `json:"Transparency,omitempty" xml:"Transparency"`
        Weight string `json:"Weight,omitempty" xml:"Weight"`
        WeightPt float64 `json:"WeightPt,omitempty" xml:"WeightPt"`
 
    Link *Link `json:"link,omitempty" xml:"link"`
    Backward float64 `json:"Backward,omitempty" xml:"Backward"`
    DataLabels *DataLabels `json:"DataLabels,omitempty" xml:"DataLabels"`
    DisplayEquation bool `json:"DisplayEquation,omitempty" xml:"DisplayEquation"`
    DisplayRSquared bool `json:"DisplayRSquared,omitempty" xml:"DisplayRSquared"`
    Forward float64 `json:"Forward,omitempty" xml:"Forward"`
    Intercept float64 `json:"Intercept,omitempty" xml:"Intercept"`
    IsNameAuto bool `json:"IsNameAuto,omitempty" xml:"IsNameAuto"`
    LegendEntry *LegendEntry `json:"LegendEntry,omitempty" xml:"LegendEntry"`
    Name string `json:"Name,omitempty" xml:"Name"`
    Order int64 `json:"Order,omitempty" xml:"Order"`
    Period int64 `json:"Period,omitempty" xml:"Period"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
}


type Trendlines struct {
     
        Link *Link `json:"link,omitempty" xml:"link"`
 
    TrendlineList []Trendline `json:"TrendlineList,omitempty" xml:"TrendlineList"`
}


type Walls struct {
     
        Border *Line `json:"Border,omitempty" xml:"Border"`
 
    CenterX int64 `json:"CenterX,omitempty" xml:"CenterX"`
    CenterY int64 `json:"CenterY,omitempty" xml:"CenterY"`
    Depth int64 `json:"Depth,omitempty" xml:"Depth"`
    Height int64 `json:"Height,omitempty" xml:"Height"`
    Width int64 `json:"Width,omitempty" xml:"Width"`
}


type AnalyzedColumnDescription struct {
 
    Index int64 `json:"Index,omitempty" xml:"Index"`
    ColumnIndex int64 `json:"ColumnIndex,omitempty" xml:"ColumnIndex"`
    Title string `json:"Title,omitempty" xml:"Title"`
    RepetitionRate float64 `json:"RepetitionRate,omitempty" xml:"RepetitionRate"`
    ColumnDataDataType string `json:"ColumnDataDataType,omitempty" xml:"ColumnDataDataType"`
    NumberCategoryType string `json:"NumberCategoryType,omitempty" xml:"NumberCategoryType"`
    TextCategoryType string `json:"TextCategoryType,omitempty" xml:"TextCategoryType"`
    StyleNumber int64 `json:"StyleNumber,omitempty" xml:"StyleNumber"`
    ColumnDataExceptionDescription string `json:"columnDataExceptionDescription,omitempty" xml:"columnDataExceptionDescription"`
}


type AnalyzedResult struct {
 
    Filename string `json:"Filename,omitempty" xml:"Filename"`
    Description string `json:"Description,omitempty" xml:"Description"`
    BasicStatistics *ExcelDataStatistics `json:"BasicStatistics,omitempty" xml:"BasicStatistics"`
    Results []AnalyzedTableDescription `json:"Results,omitempty" xml:"Results"`
    SuggestedFile string `json:"SuggestedFile,omitempty" xml:"SuggestedFile"`
}


type AnalyzedTableDescription struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
    Columns []AnalyzedColumnDescription `json:"Columns,omitempty" xml:"Columns"`
    DateColumns []int64 `json:"DateColumns,omitempty" xml:"DateColumns"`
    NumberColumns []int64 `json:"NumberColumns,omitempty" xml:"NumberColumns"`
    TextColumns []int64 `json:"TextColumns,omitempty" xml:"TextColumns"`
    ExceptionColumns []int64 `json:"ExceptionColumns,omitempty" xml:"ExceptionColumns"`
    HasTableHeaderRow bool `json:"HasTableHeaderRow,omitempty" xml:"HasTableHeaderRow"`
    HasTableTotalRow bool `json:"HasTableTotalRow,omitempty" xml:"HasTableTotalRow"`
    StartDataColumnIndex int64 `json:"StartDataColumnIndex,omitempty" xml:"StartDataColumnIndex"`
    EndDataColumnIndex int64 `json:"EndDataColumnIndex,omitempty" xml:"EndDataColumnIndex"`
    StartDataRowIndex int64 `json:"StartDataRowIndex,omitempty" xml:"StartDataRowIndex"`
    EndDataRowIndex int64 `json:"EndDataRowIndex,omitempty" xml:"EndDataRowIndex"`
    Thumbnail string `json:"Thumbnail,omitempty" xml:"Thumbnail"`
    DiscoverCharts []DiscoverChart `json:"DiscoverCharts,omitempty" xml:"DiscoverCharts"`
    DiscoverPivotTables []DiscoverPivotTable `json:"DiscoverPivotTables,omitempty" xml:"DiscoverPivotTables"`
}


type DiscoverChart struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
    Title string `json:"Title,omitempty" xml:"Title"`
    Type_ string `json:"Type,omitempty" xml:"Type"`
    DataRange string `json:"DataRange,omitempty" xml:"DataRange"`
    Thumbnail string `json:"Thumbnail,omitempty" xml:"Thumbnail"`
}


type DiscoverPivotTable struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    Title string `json:"Title,omitempty" xml:"Title"`
    DataRange string `json:"DataRange,omitempty" xml:"DataRange"`
    PivotFieldRows []int64 `json:"PivotFieldRows,omitempty" xml:"PivotFieldRows"`
    PivotFieldColumns []int64 `json:"PivotFieldColumns,omitempty" xml:"PivotFieldColumns"`
    PivotFieldData []int64 `json:"PivotFieldData,omitempty" xml:"PivotFieldData"`
    Thumbnail string `json:"Thumbnail,omitempty" xml:"Thumbnail"`
}


type ExcelDataStatistics struct {
 
    WorksheetDataStatistics []WorksheetDataStatistics `json:"WorksheetDataStatistics,omitempty" xml:"WorksheetDataStatistics"`
}


type WorksheetDataStatistics struct {
 
    Name string `json:"Name,omitempty" xml:"Name"`
    ChartsCount int64 `json:"ChartsCount,omitempty" xml:"ChartsCount"`
    TablesCount int64 `json:"TablesCount,omitempty" xml:"TablesCount"`
    PivotTablesCount int64 `json:"PivotTablesCount,omitempty" xml:"PivotTablesCount"`
    ShapesCount int64 `json:"ShapesCount,omitempty" xml:"ShapesCount"`
    HyperlinksCount int64 `json:"HyperlinksCount,omitempty" xml:"HyperlinksCount"`
    QueryTablesCount int64 `json:"QueryTablesCount,omitempty" xml:"QueryTablesCount"`
    CellsCount int64 `json:"CellsCount,omitempty" xml:"CellsCount"`
    CellsCountInTable int64 `json:"CellsCountInTable,omitempty" xml:"CellsCountInTable"`
    CellsCountIsFormula int64 `json:"CellsCountIsFormula,omitempty" xml:"CellsCountIsFormula"`
}


type Point struct {
	IsEmpty bool  `json:"IsEmpty,omitempty" xml:"IsEmpty"`
	X       int64 `json:"X,omitempty" xml:"X"`
	Y       int64 `json:"Y,omitempty" xml:"Y"`
}
