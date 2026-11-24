/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AutoFilter">
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

// AutoFilter Represents autofiltering for the specified worksheet.
type AutoFilter struct {
	LinkElement
	// Gets the collection of the filter columns.
	FilterColumns []FilterColumn `json:"FilterColumns,omitempty" xml:"FilterColumns"`
	// Represents the range to which the specified AutoFilter applies.
	Range string `json:"Range,omitempty" xml:"Range"`
	// Gets the data sorter.
	Sorter *DataSorter `json:"Sorter,omitempty" xml:"Sorter"`
	// Indicates whether the AutoFilter button for this column is visible.
	ShowFilterButton *bool `json:"ShowFilterButton,omitempty" xml:"ShowFilterButton"`
}
