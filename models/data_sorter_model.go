/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataSorter">
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

// DataSorter Summary description for DataSorter.
type DataSorter struct {
	// Gets and sets whether case sensitive when comparing string.
	CaseSensitive *bool `json:"CaseSensitive,omitempty" xml:"CaseSensitive"`
	// Represents whether the range has headers.
	HasHeaders *bool `json:"HasHeaders,omitempty" xml:"HasHeaders"`
	// Gets the key list of data sorter.
	KeyList []SortKey `json:"KeyList,omitempty" xml:"KeyList"`
	// True means that sorting orientation is from left to right.             False means that sorting orientation is from top to bottom.             The default value is false.
	SortLeftToRight *bool `json:"SortLeftToRight,omitempty" xml:"SortLeftToRight"`
	// Indicates whether sorting anything that looks like a number.
	SortAsNumber *bool `json:"SortAsNumber,omitempty" xml:"SortAsNumber"`
	// Gets the key list of data sorter.
	Keys []DataSorterKey `json:"Keys,omitempty" xml:"Keys"`
}
