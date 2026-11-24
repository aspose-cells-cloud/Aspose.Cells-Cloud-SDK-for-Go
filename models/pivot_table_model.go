/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PivotTable">
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

// PivotTable Summary description for PivotTable.
type PivotTable struct {
	LinkElement
	// Gets the description of the alt text
	AltTextDescription     string        `json:"AltTextDescription,omitempty" xml:"AltTextDescription"`
	AltTextTitle           string        `json:"AltTextTitle,omitempty" xml:"AltTextTitle"`
	AutoFormatType         string        `json:"AutoFormatType,omitempty" xml:"AutoFormatType"`
	BaseFields             []PivotField  `json:"BaseFields,omitempty" xml:"BaseFields"`
	ColumnFields           []PivotField  `json:"ColumnFields,omitempty" xml:"ColumnFields"`
	ColumnGrand            *bool         `json:"ColumnGrand,omitempty" xml:"ColumnGrand"`
	ColumnHeaderCaption    string        `json:"ColumnHeaderCaption,omitempty" xml:"ColumnHeaderCaption"`
	ColumnRange            *CellArea     `json:"ColumnRange,omitempty" xml:"ColumnRange"`
	CustomListSort         *bool         `json:"CustomListSort,omitempty" xml:"CustomListSort"`
	DataBodyRange          *CellArea     `json:"DataBodyRange,omitempty" xml:"DataBodyRange"`
	DataField              *PivotField   `json:"DataField,omitempty" xml:"DataField"`
	DataFields             []PivotField  `json:"DataFields,omitempty" xml:"DataFields"`
	DataSource             []interface{} `json:"DataSource,omitempty" xml:"DataSource"`
	DisplayErrorString     *bool         `json:"DisplayErrorString,omitempty" xml:"DisplayErrorString"`
	DisplayImmediateItems  *bool         `json:"DisplayImmediateItems,omitempty" xml:"DisplayImmediateItems"`
	DisplayNullString      *bool         `json:"DisplayNullString,omitempty" xml:"DisplayNullString"`
	EnableDataValueEditing *bool         `json:"EnableDataValueEditing,omitempty" xml:"EnableDataValueEditing"`
	EnableDrilldown        *bool         `json:"EnableDrilldown,omitempty" xml:"EnableDrilldown"`
	EnableFieldDialog      *bool         `json:"EnableFieldDialog,omitempty" xml:"EnableFieldDialog"`
	EnableFieldList        *bool         `json:"EnableFieldList,omitempty" xml:"EnableFieldList"`
	EnableWizard           *bool         `json:"EnableWizard,omitempty" xml:"EnableWizard"`
	ErrorString            string        `json:"ErrorString,omitempty" xml:"ErrorString"`
	FieldListSortAscending *bool         `json:"FieldListSortAscending,omitempty" xml:"FieldListSortAscending"`
	GrandTotalName         string        `json:"GrandTotalName,omitempty" xml:"GrandTotalName"`
	HasBlankRows           *bool         `json:"HasBlankRows,omitempty" xml:"HasBlankRows"`
	Indent                 *int32        `json:"Indent,omitempty" xml:"Indent"`
	IsAutoFormat           *bool         `json:"IsAutoFormat,omitempty" xml:"IsAutoFormat"`
	IsGridDropZones        *bool         `json:"IsGridDropZones,omitempty" xml:"IsGridDropZones"`
	IsMultipleFieldFilters *bool         `json:"IsMultipleFieldFilters,omitempty" xml:"IsMultipleFieldFilters"`
	IsSelected             *bool         `json:"IsSelected,omitempty" xml:"IsSelected"`
	ItemPrintTitles        *bool         `json:"ItemPrintTitles,omitempty" xml:"ItemPrintTitles"`
	ManualUpdate           *bool         `json:"ManualUpdate,omitempty" xml:"ManualUpdate"`
	MergeLabels            *bool         `json:"MergeLabels,omitempty" xml:"MergeLabels"`
	MissingItemsLimit      string        `json:"MissingItemsLimit,omitempty" xml:"MissingItemsLimit"`
	// Represents pivot table name.
	Name                        string        `json:"Name,omitempty" xml:"Name"`
	NullString                  string        `json:"NullString,omitempty" xml:"NullString"`
	PageFieldOrder              string        `json:"PageFieldOrder,omitempty" xml:"PageFieldOrder"`
	PageFields                  []PivotField  `json:"PageFields,omitempty" xml:"PageFields"`
	PageFieldWrapCount          *int32        `json:"PageFieldWrapCount,omitempty" xml:"PageFieldWrapCount"`
	PivotFilters                []PivotFilter `json:"PivotFilters,omitempty" xml:"PivotFilters"`
	PivotTableStyleName         string        `json:"PivotTableStyleName,omitempty" xml:"PivotTableStyleName"`
	PivotTableStyleType         string        `json:"PivotTableStyleType,omitempty" xml:"PivotTableStyleType"`
	PreserveFormatting          *bool         `json:"PreserveFormatting,omitempty" xml:"PreserveFormatting"`
	PrintDrill                  *bool         `json:"PrintDrill,omitempty" xml:"PrintDrill"`
	PrintTitles                 *bool         `json:"PrintTitles,omitempty" xml:"PrintTitles"`
	RefreshDataFlag             *bool         `json:"RefreshDataFlag,omitempty" xml:"RefreshDataFlag"`
	RefreshDataOnOpeningFile    *bool         `json:"RefreshDataOnOpeningFile,omitempty" xml:"RefreshDataOnOpeningFile"`
	RowFields                   []PivotField  `json:"RowFields,omitempty" xml:"RowFields"`
	RowGrand                    *bool         `json:"RowGrand,omitempty" xml:"RowGrand"`
	RowHeaderCaption            string        `json:"RowHeaderCaption,omitempty" xml:"RowHeaderCaption"`
	RowRange                    *CellArea     `json:"RowRange,omitempty" xml:"RowRange"`
	SaveData                    *bool         `json:"SaveData,omitempty" xml:"SaveData"`
	ShowDataTips                *bool         `json:"ShowDataTips,omitempty" xml:"ShowDataTips"`
	ShowDrill                   *bool         `json:"ShowDrill,omitempty" xml:"ShowDrill"`
	ShowEmptyCol                *bool         `json:"ShowEmptyCol,omitempty" xml:"ShowEmptyCol"`
	ShowEmptyRow                *bool         `json:"ShowEmptyRow,omitempty" xml:"ShowEmptyRow"`
	ShowMemberPropertyTips      *bool         `json:"ShowMemberPropertyTips,omitempty" xml:"ShowMemberPropertyTips"`
	ShowPivotStyleColumnHeader  *bool         `json:"ShowPivotStyleColumnHeader,omitempty" xml:"ShowPivotStyleColumnHeader"`
	ShowPivotStyleColumnStripes *bool         `json:"ShowPivotStyleColumnStripes,omitempty" xml:"ShowPivotStyleColumnStripes"`
	ShowPivotStyleLastColumn    *bool         `json:"ShowPivotStyleLastColumn,omitempty" xml:"ShowPivotStyleLastColumn"`
	ShowPivotStyleRowHeader     *bool         `json:"ShowPivotStyleRowHeader,omitempty" xml:"ShowPivotStyleRowHeader"`
	ShowPivotStyleRowStripes    *bool         `json:"ShowPivotStyleRowStripes,omitempty" xml:"ShowPivotStyleRowStripes"`
	ShowRowHeaderCaption        *bool         `json:"ShowRowHeaderCaption,omitempty" xml:"ShowRowHeaderCaption"`
	ShowValuesRow               *bool         `json:"ShowValuesRow,omitempty" xml:"ShowValuesRow"`
	SubtotalHiddenPageItems     *bool         `json:"SubtotalHiddenPageItems,omitempty" xml:"SubtotalHiddenPageItems"`
	TableRange1                 *CellArea     `json:"TableRange1,omitempty" xml:"TableRange1"`
	TableRange2                 *CellArea     `json:"TableRange2,omitempty" xml:"TableRange2"`
	Tag                         string        `json:"Tag,omitempty" xml:"Tag"`
}
