/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="WorkbookSettings">
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

// WorkbookSettings            Represents all settings of the workbook.
type WorkbookSettings struct {
	// Specifies a boolean value that indicates the application automatically compressed pictures in the workbook.
	AutoCompressPictures *bool `json:"AutoCompressPictures,omitempty" xml:"AutoCompressPictures"`
	// Indicates whether the file is mark for auto-recovery.
	AutoRecover *bool `json:"AutoRecover,omitempty" xml:"AutoRecover"`
	// Specifies the incremental public release of the application.
	BuildVersion string `json:"BuildVersion,omitempty" xml:"BuildVersion"`
	// It specifies whether to calculate formulas manually,             automatically or automatically except for multiple table operations.
	CalcMode string `json:"CalcMode,omitempty" xml:"CalcMode"`
	// Specifies the version of the calculation engine used to calculate values in the workbook.
	CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
	// Indicates whether check comptiliblity when saving workbook.                           Remarks:  The default value is true.
	CheckComptiliblity *bool `json:"CheckComptiliblity,omitempty" xml:"CheckComptiliblity"`
	// Whether check restriction of excel file when user modify cells related objects.             For example, excel does not allow inputting string value longer than 32K.             When you input a value longer than 32K such as by Cell.PutValue(string), if this property is true, you will get an Exception.             If this property is false, we will accept your input string value as the cell's value so that later             you can output the complete string value for other file formats such as CSV.             However, if you have set such kind of value that is invalid for excel file format,             you should not save the workbook as excel file format later. Otherwise there may be unexpected error for the generated excel file.
	CheckExcelRestriction *bool `json:"CheckExcelRestriction,omitempty" xml:"CheckExcelRestriction"`
	// indicates whether the application last saved the workbook file after a crash.
	CrashSave *bool `json:"CrashSave,omitempty" xml:"CrashSave"`
	// Whether creates calculated formulas chain. Default is false.
	CreateCalcChain *bool `json:"CreateCalcChain,omitempty" xml:"CreateCalcChain"`
	// indicates whether the application last opened the workbook for data recovery.
	DataExtractLoad *bool `json:"DataExtractLoad,omitempty" xml:"DataExtractLoad"`
	// Gets or sets a value which represents if the workbook uses the 1904 date system.
	Date1904 *bool `json:"Date1904,omitempty" xml:"Date1904"`
	// Indicates whether and how to show objects in the workbook.
	DisplayDrawingObjects string `json:"DisplayDrawingObjects,omitempty" xml:"DisplayDrawingObjects"`
	// Enable macros;
	EnableMacros *bool `json:"EnableMacros,omitempty" xml:"EnableMacros"`
	// Gets or sets the first visible worksheet tab.
	FirstVisibleTab *int32 `json:"FirstVisibleTab,omitempty" xml:"FirstVisibleTab"`
	// Gets and sets whether hide the field list for the PivotTable.
	HidePivotFieldList *bool `json:"HidePivotFieldList,omitempty" xml:"HidePivotFieldList"`
	// Indicates whether encrypting the workbook with default password if Structure and Windows of the workbook are locked.
	IsDefaultEncrypted *bool `json:"IsDefaultEncrypted,omitempty" xml:"IsDefaultEncrypted"`
	// Indicates whether this workbook is hidden.
	IsHidden *bool `json:"IsHidden,omitempty" xml:"IsHidden"`
	// Gets or sets a value indicating whether the generated spreadsheet will contain a horizontal scroll bar.
	IsHScrollBarVisible *bool `json:"IsHScrollBarVisible,omitempty" xml:"IsHScrollBarVisible"`
	// Represents whether the generated spreadsheet will be opened Minimized.
	IsMinimized *bool `json:"IsMinimized,omitempty" xml:"IsMinimized"`
	// Gets or sets a value indicating whether the generated spreadsheet will contain a vertical scroll bar.
	IsVScrollBarVisible *bool `json:"IsVScrollBarVisible,omitempty" xml:"IsVScrollBarVisible"`
	// Indicates whether enable iterative calculation to resolve circular references.
	Iteration *bool `json:"Iteration,omitempty" xml:"Iteration"`
	// Gets or sets the user interface language of the Workbook version based on CountryCode that has saved the file.
	LanguageCode string `json:"LanguageCode,omitempty" xml:"LanguageCode"`
	// Returns or sets the maximum number of change to resolve a circular reference.
	MaxChange *float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
	// Returns or sets the maximum number of iterations to resolve a circular reference.
	MaxIteration *int32 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
	// Gets or sets the memory usage options. The new option will be taken as the default option for newly created worksheets but does not take effect for existing worksheets.
	MemorySetting string `json:"MemorySetting,omitempty" xml:"MemorySetting"`
	// Gets or sets the decimal separator for formatting/parsing numeric values. Default is the decimal separator of current Region.
	NumberDecimalSeparator string `json:"NumberDecimalSeparator,omitempty" xml:"NumberDecimalSeparator"`
	// Gets or sets the character that separates groups of digits to the left of the decimal in numeric values. Default is the group separator of current Region.
	NumberGroupSeparator string `json:"NumberGroupSeparator,omitempty" xml:"NumberGroupSeparator"`
	// Indicates whether parsing the formula when reading the file.
	ParsingFormulaOnOpen *bool `json:"ParsingFormulaOnOpen,omitempty" xml:"ParsingFormulaOnOpen"`
	// True if calculations in this workbook will be done using only the precision of the numbers as they're displayed
	PrecisionAsDisplayed *bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
	// Indicates whether to recalculate before saving the document.
	RecalculateBeforeSave *bool `json:"RecalculateBeforeSave,omitempty" xml:"RecalculateBeforeSave"`
	// Indicates whether re-calculate all formulas on opening file.
	ReCalculateOnOpen *bool `json:"ReCalculateOnOpen,omitempty" xml:"ReCalculateOnOpen"`
	// Indicates if the Read Only Recommended option is selected.
	RecommendReadOnly *bool `json:"RecommendReadOnly,omitempty" xml:"RecommendReadOnly"`
	// Gets or sets the regional settings for workbook.
	Region string `json:"Region,omitempty" xml:"Region"`
	// True if personal information can be removed from the specified workbook.
	RemovePersonalInformation *bool `json:"RemovePersonalInformation,omitempty" xml:"RemovePersonalInformation"`
	// Indicates whether the application last opened the workbook in safe or repair mode.
	RepairLoad *bool `json:"RepairLoad,omitempty" xml:"RepairLoad"`
	// Gets or sets a value that indicates whether the Workbook is shared.
	Shared *bool `json:"Shared,omitempty" xml:"Shared"`
	// Width of worksheet tab bar (in 1/1000 of window width).
	SheetTabBarWidth *int32 `json:"SheetTabBarWidth,omitempty" xml:"SheetTabBarWidth"`
	// Get or sets a value whether the Workbook tabs are displayed.
	ShowTabs *bool `json:"ShowTabs,omitempty" xml:"ShowTabs"`
	// Indicates whether update adjacent cells' border.
	UpdateAdjacentCellsBorder *bool `json:"UpdateAdjacentCellsBorder,omitempty" xml:"UpdateAdjacentCellsBorder"`
	// Gets and sets how updates external links when the workbook is opened.
	UpdateLinksType string `json:"UpdateLinksType,omitempty" xml:"UpdateLinksType"`
	// The height of the window, in unit of point.
	WindowHeight *float64 `json:"WindowHeight,omitempty" xml:"WindowHeight"`
	// The distance from the left edge of the client area to the left edge of the window, in unit of point.
	WindowLeft *float64 `json:"WindowLeft,omitempty" xml:"WindowLeft"`
	// The distance from the top edge of the client area to the top edge of the window, in unit of point.
	WindowTop *float64 `json:"WindowTop,omitempty" xml:"WindowTop"`
	// The width of the window, in unit of point.
	WindowWidth *float64 `json:"WindowWidth,omitempty" xml:"WindowWidth"`
	// Gets and sets the author of the file.
	Author string `json:"Author,omitempty" xml:"Author"`
	// Indicates whether checking custom number format when setting Style.Custom.
	CheckCustomNumberFormat *bool `json:"CheckCustomNumberFormat,omitempty" xml:"CheckCustomNumberFormat"`
	// Gets the protection type of the workbook.
	ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
	// Gets and sets the globalization settings.
	GlobalizationSettings *GlobalizationSettings `json:"GlobalizationSettings,omitempty" xml:"GlobalizationSettings"`
	// Represents Workbook file encryption password.
	Password string `json:"Password,omitempty" xml:"Password"`
	// Provides access to the workbook write protection options.
	WriteProtection *WriteProtection `json:"WriteProtection,omitempty" xml:"WriteProtection"`
	// Gets a value that indicates whether a password is required to open this workbook.
	IsEncrypted *bool `json:"IsEncrypted,omitempty" xml:"IsEncrypted"`
	// Gets a value that indicates whether the structure or window of the Workbook is protected.
	IsProtected *bool `json:"IsProtected,omitempty" xml:"IsProtected"`
	// Gets the max row index, zero-based.
	MaxRow *int32 `json:"MaxRow,omitempty" xml:"MaxRow"`
	// Gets the max column index, zero-based.
	MaxColumn *int32 `json:"MaxColumn,omitempty" xml:"MaxColumn"`
	// Gets and sets the number of significant digits.             The default value is .
	SignificantDigits *int32 `json:"SignificantDigits,omitempty" xml:"SignificantDigits"`
	// Indicates whether check compatibility with earlier versions when saving workbook.
	CheckCompatibility *bool `json:"CheckCompatibility,omitempty" xml:"CheckCompatibility"`
	// Gets and sets the default print paper size.
	PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
	// Gets and sets the max row number of shared formula.
	MaxRowsOfSharedFormula *int32 `json:"MaxRowsOfSharedFormula,omitempty" xml:"MaxRowsOfSharedFormula"`
	// Specifies the OOXML version for the output document. The default value is Ecma376_2006.
	Compliance string `json:"Compliance,omitempty" xml:"Compliance"`
	// Indicates whether setting  property when entering the string value(which starts  with single quote mark ) to the cell
	QuotePrefixToStyle *bool `json:"QuotePrefixToStyle,omitempty" xml:"QuotePrefixToStyle"`
	// Gets the settings for formula-related features.
	FormulaSettings *FormulaSettings `json:"FormulaSettings,omitempty" xml:"FormulaSettings"`
	// Fully calculates every time when a calculation is triggered.
	ForceFullCalculate *bool `json:"ForceFullCalculate,omitempty" xml:"ForceFullCalculate"`
}
