/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="status.go">
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

import (
	"encoding/json"
	"fmt"

	"asposecellscloud/models"
)

// CheckResponseStatus verifies that an HTTP response is successful (2xx).
// client.Do only reports transport/retry errors and does not inspect the HTTP
// status code, so the high-level functional APIs must call this after every
// request. It returns a *SDKError wrapping the status code and, when
// available, the "Status" field from the API error body.
func CheckResponseStatus(resp *RichResponse) error {
	if resp == nil {
		return fmt.Errorf("%w: nil response", ErrRequestFailed)
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	msg := fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(resp.Body))
	var cc models.CellsCloudResponse
	_ = json.Unmarshal(resp.Body, &cc)
	if cc.Status != "" {
		msg = fmt.Sprintf("HTTP %d: %s", resp.StatusCode, cc.Status)
	}
	return &SDKError{Code: resp.StatusCode, Message: "request failed", Err: fmt.Errorf("%w: %s", ErrRequestFailed, msg)}
}
