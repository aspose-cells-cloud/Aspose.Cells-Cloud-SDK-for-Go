/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="OperateObjectPosition">
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

// OperateObjectPosition Represents operate object position.
type OperateObjectPosition struct {
	// Represents data source of operate object.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Represents data source of operate object.
	Workbook *FileSource `json:"Workbook,omitempty" xml:"Workbook"`
	// Represents worksheet name of operate object.
	SheetName string `json:"SheetName,omitempty" xml:"SheetName"`
	// Represents chart index of operate object.
	ChartIndex *int32 `json:"ChartIndex,omitempty" xml:"ChartIndex"`
	// Represents shape index of operate object.
	ShapeIndex *int32 `json:"ShapeIndex,omitempty" xml:"ShapeIndex"`
	// Represents cell name of operate object.
	CellName string `json:"CellName,omitempty" xml:"CellName"`
	// Represents list object index of operate object.
	ListObjectIndex *int32 `json:"ListObjectIndex,omitempty" xml:"ListObjectIndex"`
}
