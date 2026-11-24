/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="DataTransformationRequest">
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

// DataTransformationRequest Data Transformation Request
type DataTransformationRequest struct {
	// Indicates the source of the mount data.
	FileInfo *FileInfo `json:"FileInfo,omitempty" xml:"FileInfo"`
	// Indicates the source of the mount data.
	DataSource *DataSource `json:"DataSource,omitempty" xml:"DataSource"`
	// Indicates load data.
	LoadData *LoadData `json:"LoadData,omitempty" xml:"LoadData"`
	// Indicates applied step list.
	AppliedSteps []AppliedStep `json:"AppliedSteps,omitempty" xml:"AppliedSteps"`
	// This class has a property named "Region" of type string with both a getter and a setter.
	Region string `json:"Region,omitempty" xml:"Region"`
	// Indicates output format
	OutFormat string `json:"OutFormat,omitempty" xml:"OutFormat"`
}
