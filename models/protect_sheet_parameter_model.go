/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ProtectSheetParameter">
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

// ProtectSheetParameter Class Summary: The features discussed include natural language processing, image recognition, data visualization, clustering, regression, classification, and recommendation systems.
type ProtectSheetParameter struct {
	// Property Summary: Defines a public string property for the ProtectionType.
	ProtectionType             string   `json:"ProtectionType,omitempty" xml:"ProtectionType"`
	Password                   string   `json:"Password,omitempty" xml:"Password"`
	AllowEditArea              []string `json:"AllowEditArea,omitempty" xml:"AllowEditArea"`
	AllowDeletingColumn        string   `json:"AllowDeletingColumn,omitempty" xml:"AllowDeletingColumn"`
	AllowDeletingRow           string   `json:"AllowDeletingRow,omitempty" xml:"AllowDeletingRow"`
	AllowFiltering             string   `json:"AllowFiltering,omitempty" xml:"AllowFiltering"`
	AllowFormattingCell        string   `json:"AllowFormattingCell,omitempty" xml:"AllowFormattingCell"`
	AllowFormattingColumn      string   `json:"AllowFormattingColumn,omitempty" xml:"AllowFormattingColumn"`
	AllowFormattingRow         string   `json:"AllowFormattingRow,omitempty" xml:"AllowFormattingRow"`
	AllowInsertingColumn       string   `json:"AllowInsertingColumn,omitempty" xml:"AllowInsertingColumn"`
	AllowInsertingHyperlink    string   `json:"AllowInsertingHyperlink,omitempty" xml:"AllowInsertingHyperlink"`
	AllowInsertingRow          string   `json:"AllowInsertingRow,omitempty" xml:"AllowInsertingRow"`
	AllowSelectingLockedCell   string   `json:"AllowSelectingLockedCell,omitempty" xml:"AllowSelectingLockedCell"`
	AllowSelectingUnlockedCell string   `json:"AllowSelectingUnlockedCell,omitempty" xml:"AllowSelectingUnlockedCell"`
	AllowSorting               string   `json:"AllowSorting,omitempty" xml:"AllowSorting"`
	AllowUsingPivotTable       string   `json:"AllowUsingPivotTable,omitempty" xml:"AllowUsingPivotTable"`
}
