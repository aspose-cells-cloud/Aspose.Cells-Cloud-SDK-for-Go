/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ProtectWorkbookRequest">
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

// ProtectWorkbookRequest Indicates protect workbook request
type ProtectWorkbookRequest struct {
	// Indicates aways open read-only.
	AwaysOpenReadOnly *bool `json:"AwaysOpenReadOnly,omitempty" xml:"AwaysOpenReadOnly"`
	// Indicates encrypt with password.
	EncryptWithPassword string `json:"EncryptWithPassword,omitempty" xml:"EncryptWithPassword"`
	// Represents the various types of protection options available for a worksheet.
	ProtectCurrentSheet *Protection `json:"ProtectCurrentSheet,omitempty" xml:"ProtectCurrentSheet"`
	// Represents the various types of protection options available for all worksheets.
	ProtectAllSheets *Protection `json:"ProtectAllSheets,omitempty" xml:"ProtectAllSheets"`
	// Indicates protect workbook structure. All, Contents, Objects, Scenarios, Structure, Windows, and None.
	ProtectWorkbookStructure string `json:"ProtectWorkbookStructure,omitempty" xml:"ProtectWorkbookStructure"`
	// Indicates signature in file.
	DigitalSignature *DigitalSignature `json:"DigitalSignature,omitempty" xml:"DigitalSignature"`
	// Indicates mark as final.
	MarkAsFinal *bool `json:"MarkAsFinal,omitempty" xml:"MarkAsFinal"`
}
