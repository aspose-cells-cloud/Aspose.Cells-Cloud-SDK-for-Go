/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="SplitWorkbookTaskParameter">
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

// SplitWorkbookTaskParameter Represents split workbook task parameter.
type SplitWorkbookTaskParameter struct {
	TaskParameter
	// Represents data source of task object.
	Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
	// Represents data source of task object.
	DestinationFilePosition *FileSource `json:"DestinationFilePosition,omitempty" xml:"DestinationFilePosition"`
	// Represents destination data source.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents destination data source.
	TargetDataSource *DataSource `json:"TargetDataSource,omitempty" xml:"TargetDataSource"`
	// Represents destination data file format.
	DestinationFileFormat string `json:"DestinationFileFormat,omitempty" xml:"DestinationFileFormat"`
	// SheetName /NewGuid
	SplitNameRule string `json:"SplitNameRule,omitempty" xml:"SplitNameRule"`
	// When destination file format is image , vertical resolution can not be null.
	VerticalResolution *int32 `json:"VerticalResolution,omitempty" xml:"VerticalResolution"`
	// When destination file format is image , horizontal resolution can not be null.
	HorizontalResolution *int32 `json:"HorizontalResolution,omitempty" xml:"HorizontalResolution"`
}
