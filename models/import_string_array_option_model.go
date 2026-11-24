/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="ImportStringArrayOption">
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

// ImportStringArrayOption I'm happy to help! Could you please provide me with the features you would like me to summarize for the class?
type ImportStringArrayOption struct {
	ImportOption
	// A property named "FirstRow" of type integer with both a getter and a setter is defined in the class.
	FirstRow    *int32        `json:"FirstRow,omitempty" xml:"FirstRow"`
	FirstColumn *int32        `json:"FirstColumn,omitempty" xml:"FirstColumn"`
	IsVertical  *bool         `json:"IsVertical,omitempty" xml:"IsVertical"`
	Data        []interface{} `json:"Data,omitempty" xml:"Data"`
}
