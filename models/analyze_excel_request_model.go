/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="AnalyzeExcelRequest">
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

// AnalyzeExcelRequest Can you please provide the features that you would like me to summarize for the class?
type AnalyzeExcelRequest struct {
	// The class has a property called "Files" that is a list of FileInfo objects with a default value of an empty list.
	Files                []FileInfo `json:"Files,omitempty" xml:"Files"`
	NeedThumbnail        *bool      `json:"NeedThumbnail,omitempty" xml:"NeedThumbnail"`
	BuildSuggestoinSheet *bool      `json:"BuildSuggestoinSheet,omitempty" xml:"BuildSuggestoinSheet"`
}
