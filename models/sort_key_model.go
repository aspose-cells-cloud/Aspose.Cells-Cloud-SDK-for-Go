/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="SortKey">
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

// SortKey Represents sort key.
type SortKey struct {
	// Represents the key of sorting.
	Key *int32 `json:"Key,omitempty" xml:"Key"`
	// Represents the order of sorting.
	SortOrder string `json:"SortOrder,omitempty" xml:"SortOrder"`
	// This class includes a property named CustomList that is an array of strings with both getter and setter methods.
	CustomList []interface{} `json:"CustomList,omitempty" xml:"CustomList"`
	// Indicates the order of sorting.
	Order string `json:"Order,omitempty" xml:"Order"`
	// Gets the sorted column index(absolute position, column A is 0, B is 1, ...).
	Index *int32 `json:"Index,omitempty" xml:"Index"`
	// Represents the type of sorting.
	Type string `json:"Type,omitempty" xml:"Type"`
}
