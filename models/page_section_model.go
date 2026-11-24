/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="PageSection">
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

// PageSection Class summary: Includes interactive discussions, group projects, guest speakers, and practical applications of course material.
type PageSection struct {
	// 0,1,2  left , middle ,right
	Section *int32 `json:"Section,omitempty" xml:"Section"`
	// page context script
	Context string `json:"Context,omitempty" xml:"Context"`
	// A property named "Picture" with both getter and setter methods of type string is defined.
	Picture string `json:"Picture,omitempty" xml:"Picture"`
	// fisrt page context script
	FisrtPageContext string `json:"FisrtPageContext,omitempty" xml:"FisrtPageContext"`
	// Even page context script
	EvenPageContext string `json:"EvenPageContext,omitempty" xml:"EvenPageContext"`
}
