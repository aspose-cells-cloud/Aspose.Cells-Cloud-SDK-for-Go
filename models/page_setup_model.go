/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PageSetup">
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

// PageSetup excel print page setting
type PageSetup struct {
	// Represents if elements of the document will be printed in black and white.
	BlackAndWhite *bool `json:"BlackAndWhite,omitempty" xml:"BlackAndWhite"`
	// Represents the size of the bottom margin, in unit of centimeters.
	BottomMargin *float64 `json:"BottomMargin,omitempty" xml:"BottomMargin"`
	// Represent if the sheet is printed centered horizontally.
	CenterHorizontally *bool `json:"CenterHorizontally,omitempty" xml:"CenterHorizontally"`
	// Represent if the sheet is printed centered vertically.
	CenterVertically *bool `json:"CenterVertically,omitempty" xml:"CenterVertically"`
	// Represents the first page number that will be used when this sheet is printed.
	FirstPageNumber *int32 `json:"FirstPageNumber,omitempty" xml:"FirstPageNumber"`
	// Represents  the number of pages tall the worksheet will be scaled to when it's printed.             The default value is 1.
	FitToPagesTall *int32 `json:"FitToPagesTall,omitempty" xml:"FitToPagesTall"`
	// Represents the number of pages wide the worksheet will be scaled to when it's printed.             The default value is 1.
	FitToPagesWide *int32 `json:"FitToPagesWide,omitempty" xml:"FitToPagesWide"`
	// Represents the distance from the bottom of the page to the footer, in unit of centimeters.
	FooterMargin *float64 `json:"FooterMargin,omitempty" xml:"FooterMargin"`
	// Represents the distance from the top of the page to the header, in unit of centimeters.
	HeaderMargin *float64 `json:"HeaderMargin,omitempty" xml:"HeaderMargin"`
	// Indicates whether the first the page number is automatically assigned.
	IsAutoFirstPageNumber *bool `json:"IsAutoFirstPageNumber,omitempty" xml:"IsAutoFirstPageNumber"`
	// Indicates whether header and footer margins are aligned with the page margins.             If this property is true, the left header and footer will be aligned with the left margin,             and the right header and footer will be aligned with the right margin.             This option is enabled by default.
	IsHFAlignMargins *bool `json:"IsHFAlignMargins,omitempty" xml:"IsHFAlignMargins"`
	// True means that the header/footer of the first page is different with other pages.
	IsHFDiffFirst *bool `json:"IsHFDiffFirst,omitempty" xml:"IsHFDiffFirst"`
	// True means that the header/footer of the odd pages is different with odd pages.
	IsHFDiffOddEven *bool `json:"IsHFDiffOddEven,omitempty" xml:"IsHFDiffOddEven"`
	// Indicates whether header and footer are scaled with document scaling.             Only applies for Excel 2007.
	IsHFScaleWithDoc *bool `json:"IsHFScaleWithDoc,omitempty" xml:"IsHFScaleWithDoc"`
	// If this property is False, the FitToPagesWide and FitToPagesTall properties control how the worksheet is scaled.
	IsPercentScale *bool `json:"IsPercentScale,omitempty" xml:"IsPercentScale"`
	// Represents the size of the left margin, in unit of centimeters.
	LeftMargin *float64 `json:"LeftMargin,omitempty" xml:"LeftMargin"`
	// Represents the order that Microsoft Excel uses to number pages when printing a large worksheet.
	Order string `json:"Order,omitempty" xml:"Order"`
	// Represents page print orientation.
	Orientation string `json:"Orientation,omitempty" xml:"Orientation"`
	// Represents the size of the paper.
	PaperSize string `json:"PaperSize,omitempty" xml:"PaperSize"`
	// Represents the range to be printed.
	PrintArea string `json:"PrintArea,omitempty" xml:"PrintArea"`
	// Represents the way comments are printed with the sheet.
	PrintComments string `json:"PrintComments,omitempty" xml:"PrintComments"`
	// Get and sets number of copies to print.
	PrintCopies *int32 `json:"PrintCopies,omitempty" xml:"PrintCopies"`
	// Represents if the sheet will be printed without graphics.
	PrintDraft *bool `json:"PrintDraft,omitempty" xml:"PrintDraft"`
	// Specifies the type of print error displayed.
	PrintErrors string `json:"PrintErrors,omitempty" xml:"PrintErrors"`
	// Represents if cell gridlines are printed on the page.
	PrintGridlines *bool `json:"PrintGridlines,omitempty" xml:"PrintGridlines"`
	// Represents if row and column headings are printed with this page.
	PrintHeadings *bool `json:"PrintHeadings,omitempty" xml:"PrintHeadings"`
	// Represents the print quality.
	PrintQuality *int32 `json:"PrintQuality,omitempty" xml:"PrintQuality"`
	// Represents the columns that contain the cells to be repeated on the left side of each page.
	PrintTitleColumns string `json:"PrintTitleColumns,omitempty" xml:"PrintTitleColumns"`
	// Represents the rows that contain the cells to be repeated at the top of each page.
	PrintTitleRows string `json:"PrintTitleRows,omitempty" xml:"PrintTitleRows"`
	// Represents the size of the right margin, in unit of centimeters.
	RightMargin *float64 `json:"RightMargin,omitempty" xml:"RightMargin"`
	// Represents the size of the top margin, in unit of centimeters.
	TopMargin *float64 `json:"TopMargin,omitempty" xml:"TopMargin"`
	// Represents the scaling factor in percent. It should be between 10 and 400.
	Zoom *int32 `json:"Zoom,omitempty" xml:"Zoom"`
	// Represents the page header.
	Header []PageSection `json:"Header,omitempty" xml:"Header"`
	// Represents the page footor.
	Footer []PageSection `json:"Footer,omitempty" xml:"Footer"`
}
