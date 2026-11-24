/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PivotField">
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

// PivotField            Represents a field in a PivotTable report.
type PivotField struct {
	// Represent the number of top or bottom items            that are automatically shown in the specified PivotTable field.
	AutoShowCount                  *int32        `json:"AutoShowCount,omitempty" xml:"AutoShowCount"`
	AutoShowField                  *int32        `json:"AutoShowField,omitempty" xml:"AutoShowField"`
	AutoSortField                  *int32        `json:"AutoSortField,omitempty" xml:"AutoSortField"`
	BaseField                      *int32        `json:"BaseField,omitempty" xml:"BaseField"`
	BaseIndex                      *int32        `json:"BaseIndex,omitempty" xml:"BaseIndex"`
	BaseItem                       *int32        `json:"BaseItem,omitempty" xml:"BaseItem"`
	BaseItemPosition               string        `json:"BaseItemPosition,omitempty" xml:"BaseItemPosition"`
	CurrentPageItem                *int32        `json:"CurrentPageItem,omitempty" xml:"CurrentPageItem"`
	DataDisplayFormat              string        `json:"DataDisplayFormat,omitempty" xml:"DataDisplayFormat"`
	DisplayName                    string        `json:"DisplayName,omitempty" xml:"DisplayName"`
	DragToColumn                   *bool         `json:"DragToColumn,omitempty" xml:"DragToColumn"`
	DragToData                     *bool         `json:"DragToData,omitempty" xml:"DragToData"`
	DragToHide                     *bool         `json:"DragToHide,omitempty" xml:"DragToHide"`
	DragToPage                     *bool         `json:"DragToPage,omitempty" xml:"DragToPage"`
	DragToRow                      *bool         `json:"DragToRow,omitempty" xml:"DragToRow"`
	Function                       string        `json:"Function,omitempty" xml:"Function"`
	InsertBlankRow                 *bool         `json:"InsertBlankRow,omitempty" xml:"InsertBlankRow"`
	IsAscendShow                   *bool         `json:"IsAscendShow,omitempty" xml:"IsAscendShow"`
	IsAscendSort                   *bool         `json:"IsAscendSort,omitempty" xml:"IsAscendSort"`
	IsAutoShow                     *bool         `json:"IsAutoShow,omitempty" xml:"IsAutoShow"`
	IsAutoSort                     *bool         `json:"IsAutoSort,omitempty" xml:"IsAutoSort"`
	IsAutoSubtotals                *bool         `json:"IsAutoSubtotals,omitempty" xml:"IsAutoSubtotals"`
	IsCalculatedField              *bool         `json:"IsCalculatedField,omitempty" xml:"IsCalculatedField"`
	IsIncludeNewItemsInFilter      *bool         `json:"IsIncludeNewItemsInFilter,omitempty" xml:"IsIncludeNewItemsInFilter"`
	IsInsertPageBreaksBetweenItems *bool         `json:"IsInsertPageBreaksBetweenItems,omitempty" xml:"IsInsertPageBreaksBetweenItems"`
	IsMultipleItemSelectionAllowed *bool         `json:"IsMultipleItemSelectionAllowed,omitempty" xml:"IsMultipleItemSelectionAllowed"`
	IsRepeatItemLabels             *bool         `json:"IsRepeatItemLabels,omitempty" xml:"IsRepeatItemLabels"`
	ItemCount                      *int32        `json:"ItemCount,omitempty" xml:"ItemCount"`
	Items                          []interface{} `json:"Items,omitempty" xml:"Items"`
	Name                           string        `json:"Name,omitempty" xml:"Name"`
	Number                         *int32        `json:"Number,omitempty" xml:"Number"`
	NumberFormat                   string        `json:"NumberFormat,omitempty" xml:"NumberFormat"`
	OriginalItems                  []interface{} `json:"OriginalItems,omitempty" xml:"OriginalItems"`
	PivotItems                     []PivotItem   `json:"PivotItems,omitempty" xml:"PivotItems"`
	Position                       *int32        `json:"Position,omitempty" xml:"Position"`
	ShowAllItems                   *bool         `json:"ShowAllItems,omitempty" xml:"ShowAllItems"`
	ShowCompact                    *bool         `json:"ShowCompact,omitempty" xml:"ShowCompact"`
	ShowInOutlineForm              *bool         `json:"ShowInOutlineForm,omitempty" xml:"ShowInOutlineForm"`
	ShowSubtotalAtTop              *bool         `json:"ShowSubtotalAtTop,omitempty" xml:"ShowSubtotalAtTop"`
}
