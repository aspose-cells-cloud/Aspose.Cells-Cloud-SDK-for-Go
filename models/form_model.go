/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="Form">
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

// Form I'm happy to help! Please provide the features you would like me to summarize into one sentence for the class.
type Form struct {
	Shape
	// A property named "FormType" of type string which can be both accessed and modified.
	FormType          string        `json:"FormType,omitempty" xml:"FormType"`
	CheckedValue      string        `json:"CheckedValue,omitempty" xml:"CheckedValue"`
	Shadow            *bool         `json:"Shadow,omitempty" xml:"Shadow"`
	InputRange        string        `json:"InputRange,omitempty" xml:"InputRange"`
	SelectedIndex     *int32        `json:"SelectedIndex,omitempty" xml:"SelectedIndex"`
	SelectedValue     string        `json:"SelectedValue,omitempty" xml:"SelectedValue"`
	SelectedCell      *LinkElement  `json:"SelectedCell,omitempty" xml:"SelectedCell"`
	DropDownLines     *int32        `json:"DropDownLines,omitempty" xml:"DropDownLines"`
	ItemCount         *int32        `json:"ItemCount,omitempty" xml:"ItemCount"`
	SelectedCells     []LinkElement `json:"SelectedCells,omitempty" xml:"SelectedCells"`
	SelectionType     string        `json:"SelectionType,omitempty" xml:"SelectionType"`
	IsChecked         *bool         `json:"IsChecked,omitempty" xml:"IsChecked"`
	CurrentValue      *int32        `json:"CurrentValue,omitempty" xml:"CurrentValue"`
	Min               *int32        `json:"Min,omitempty" xml:"Min"`
	Max               *int32        `json:"Max,omitempty" xml:"Max"`
	IncrementalChange *int32        `json:"IncrementalChange,omitempty" xml:"IncrementalChange"`
	PageChange        *int32        `json:"PageChange,omitempty" xml:"PageChange"`
	IsHorizontal      *bool         `json:"IsHorizontal,omitempty" xml:"IsHorizontal"`
	Link              *Link         `json:"link,omitempty" xml:"link"`
}
