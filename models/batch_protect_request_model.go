/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="BatchProtectRequest">
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

// BatchProtectRequest Indicates batch protect file request
type BatchProtectRequest struct {
	// The directory stores files that need to format conversion.
	SourceFolder string `json:"SourceFolder,omitempty" xml:"SourceFolder"`
	// Aspose Cloud storage name
	SourceStorage string `json:"SourceStorage,omitempty" xml:"SourceStorage"`
	// Indicates the match condition that needs to be processed for the file name.
	MatchCondition *MatchConditionRequest `json:"MatchCondition,omitempty" xml:"MatchCondition"`
	// This property defines the protection type of an object.
	ProtectionType string `json:"ProtectionType,omitempty" xml:"ProtectionType"`
	//
	Password string `json:"Password,omitempty" xml:"Password"`
	// The directory that stores files whose format conversion was successful.
	OutFolder string `json:"OutFolder,omitempty" xml:"OutFolder"`
	// Aspose Cloud storage name.
	OutStorage string `json:"OutStorage,omitempty" xml:"OutStorage"`
}
