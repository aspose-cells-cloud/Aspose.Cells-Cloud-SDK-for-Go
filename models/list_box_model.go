/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ListBox">
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

// ListBox Represents a list box object.
type ListBox struct {
	Shape
	// This property stores the input range for the class.
	InputRange string `json:"InputRange,omitempty" xml:"InputRange"`
	// Gets the number of items in the list box.
	ItemCount *int32 `json:"ItemCount,omitempty" xml:"ItemCount"`
	// Specifies the amount by which the control's value is changed                          when the user clicks on the scrollbar's page up or page down region.
	PageChange *int32 `json:"PageChange,omitempty" xml:"PageChange"`
	// Gets the selected cells.                         Returns null if the input range is not set or no item is selected
	SelectedCells []LinkElement `json:"SelectedCells,omitempty" xml:"SelectedCells"`
	// Gets or sets the index number of the currently selected item in a list box or combo box.                         Zero-based.
	SelectedIndex *int32 `json:"SelectedIndex,omitempty" xml:"SelectedIndex"`
	// Gets or sets the selection mode of the specified list box.
	SelectionType string `json:"SelectionType,omitempty" xml:"SelectionType"`
	// Indicates whether the combobox has 3-D shading.
	Shadow *bool `json:"Shadow,omitempty" xml:"Shadow"`
	Link   *Link `json:"link,omitempty" xml:"link"`
}
