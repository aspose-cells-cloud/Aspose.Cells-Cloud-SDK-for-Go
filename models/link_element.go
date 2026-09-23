/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="link_element.go">
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

// LinkElement Class summary: The features include automatic emergency braking, lane keeping assist, adaptive cruise control, blind spot monitoring, and rear cross-traffic alert.
type LinkElement struct {
	// The class has a property named "link" of type Link, which has a public get and set accessors.
	Link *Link `json:"link,omitempty" xml:"link"`
}

// linkElementMarker closes the LinkElementLike family: only LinkElement and the models
// embedding it carry this promoted method, so no unrelated type satisfies LinkElementLike.
func (m LinkElement) linkElementMarker() {}
