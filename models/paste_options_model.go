/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PasteOptions">
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

// PasteOptions Represents the paste special options.
type PasteOptions struct {
	// True means only copying visible cells.
	OnlyVisibleCells *bool `json:"OnlyVisibleCells,omitempty" xml:"OnlyVisibleCells"`
	// The paste special type.
	PasteType string `json:"PasteType,omitempty" xml:"PasteType"`
	// Indicates whether skips blank cells.
	SkipBlanks *bool `json:"SkipBlanks,omitempty" xml:"SkipBlanks"`
	// True to transpose rows and columns when the range is pasted. The default value is False.
	Transpose *bool `json:"Transpose,omitempty" xml:"Transpose"`
}
