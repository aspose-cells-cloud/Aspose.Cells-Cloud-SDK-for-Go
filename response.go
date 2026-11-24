/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="response.go">
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

package asposecellscloud

import "encoding/json"

// RichResponse unified API response structure
type RichResponse struct {
	StatusCode int                 // HTTP status code
	Headers    map[string][]string // Response Header
	Body       []byte              // Raw response body
}

// GetJSON Parse the response body into the specified struct
func (r *RichResponse) GetJSON(target interface{}) error {
	if len(r.Body) == 0 || target == nil {
		return nil
	}
	return json.Unmarshal(r.Body, target)
}

func (r *RichResponse) ToBytes() []byte {
	return r.Body
}

// ToString returns the response body as a string.
func (r *RichResponse) ToString() string {
	return string(r.Body)
}
