/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Protection">
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

// Protection Represents the various types of protection options available for a worksheet.
type Protection struct {
	// Represents if the deletion of columns is allowed on a protected worksheet.
	AllowDeletingColumn *bool `json:"AllowDeletingColumn,omitempty" xml:"AllowDeletingColumn"`
	// Represents if the deletion of rows is allowed on a protected worksheet.
	AllowDeletingRow *bool `json:"AllowDeletingRow,omitempty" xml:"AllowDeletingRow"`
	// Represents if the user is allowed to make use of an AutoFilter that was created before the sheet was protected.
	AllowFiltering *bool `json:"AllowFiltering,omitempty" xml:"AllowFiltering"`
	// Represents if the formatting of cells is allowed on a protected worksheet.
	AllowFormattingCell *bool `json:"AllowFormattingCell,omitempty" xml:"AllowFormattingCell"`
	// Represents if the formatting of columns is allowed on a protected worksheet
	AllowFormattingColumn *bool `json:"AllowFormattingColumn,omitempty" xml:"AllowFormattingColumn"`
	// Represents if the formatting of rows is allowed on a protected worksheet
	AllowFormattingRow *bool `json:"AllowFormattingRow,omitempty" xml:"AllowFormattingRow"`
	// Represents if the insertion of columns is allowed on a protected worksheet
	AllowInsertingColumn *bool `json:"AllowInsertingColumn,omitempty" xml:"AllowInsertingColumn"`
	// Represents if the insertion of hyperlinks is allowed on a protected worksheet
	AllowInsertingHyperlink *bool `json:"AllowInsertingHyperlink,omitempty" xml:"AllowInsertingHyperlink"`
	// Represents if the insertion of rows is allowed on a protected worksheet
	AllowInsertingRow *bool `json:"AllowInsertingRow,omitempty" xml:"AllowInsertingRow"`
	// Represents if the sorting option is allowed on a protected worksheet.
	AllowSorting *bool `json:"AllowSorting,omitempty" xml:"AllowSorting"`
	// Represents if the user is allowed to manipulate pivot tables on a protected worksheet.
	AllowUsingPivotTable *bool `json:"AllowUsingPivotTable,omitempty" xml:"AllowUsingPivotTable"`
	// Represents if the user is allowed to edit contents of locked cells on a protected worksheet.
	AllowEditingContent *bool `json:"AllowEditingContent,omitempty" xml:"AllowEditingContent"`
	// Represents if the user is allowed to manipulate drawing objects on a protected worksheet.
	AllowEditingObject *bool `json:"AllowEditingObject,omitempty" xml:"AllowEditingObject"`
	// Represents if the user is allowed to edit scenarios on a protected worksheet.
	AllowEditingScenario *bool `json:"AllowEditingScenario,omitempty" xml:"AllowEditingScenario"`
	// Represents the password to protect the worksheet.
	Password string `json:"Password,omitempty" xml:"Password"`
	// Represents if the user is allowed to select locked cells on a protected worksheet.
	AllowSelectingLockedCell *bool `json:"AllowSelectingLockedCell,omitempty" xml:"AllowSelectingLockedCell"`
	// Represents if the user is allowed to select unlocked cells on a protected worksheet.
	AllowSelectingUnlockedCell *bool `json:"AllowSelectingUnlockedCell,omitempty" xml:"AllowSelectingUnlockedCell"`
}
