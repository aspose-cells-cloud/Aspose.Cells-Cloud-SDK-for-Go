/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="offline_test.go">
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

package asposecellscloud_test

import (
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

// intPtr is a small local helper for optional integer parameters, which the
// request contract models as pointers so that "unset" differs from zero.
func intPtr(v int) *int { return &v }

// TestOfflineStatusRequestContract verifies the generated request contract
// without contacting the live API. It is the smoke test that CI can run with
// no credentials.
func TestOfflineStatusRequestContract(t *testing.T) {
	req := requests.NewGetAsposeCellsCloudStatusRequest()

	if got := req.GetMethod(); got != "GET" {
		t.Fatalf("GetMethod() = %q, want GET", got)
	}
	if got := req.GetPath(); got != "/v4.0/cells" {
		t.Fatalf("GetPath() = %q, want /v4.0/cells", got)
	}
	if err := req.Validate(); err != nil {
		t.Fatalf("Validate() = %v, want nil", err)
	}
}

// TestOfflinePathAndOptionContract covers path escaping and optional pointer
// parameters, both of which are easy to break in code generation.
func TestOfflinePathAndOptionContract(t *testing.T) {
	req := requests.NewPostClearContentsRequest(
		"Book 1.xlsx",
		"Sheet/1",
		requests.WithCommonParameter("range", "A1:C10"),
		requests.WithCommonParameter("startRow", intPtr(1)),
		requests.WithCommonParameter("endRow", intPtr(3)),
		requests.WithCommonParameter("folder", "TestData/In"),
	)

	wantPath := "/v3.0/cells/Book%201.xlsx/worksheets/Sheet%2F1/cells/clearcontents"
	if got := req.GetPath(); got != wantPath {
		t.Fatalf("GetPath() = %q, want %q", got, wantPath)
	}

	query := req.GetQueryParameters()
	if got := query.Get("range"); got != "A1:C10" {
		t.Errorf("query range = %q, want A1:C10", got)
	}
	if got := query.Get("startRow"); got != "1" {
		t.Errorf("query startRow = %q, want 1", got)
	}
	if got := query.Get("endRow"); got != "3" {
		t.Errorf("query endRow = %q, want 3", got)
	}
	if got := query.Get("folder"); got != "TestData/In" {
		t.Errorf("query folder = %q, want TestData/In", got)
	}
}

// TestOfflineExtraQueryAndValidation covers the implicit query extension and
// required-parameter validation paths.
func TestOfflineExtraQueryAndValidation(t *testing.T) {
	req := requests.NewGetAsposeCellsCloudStatusRequest(
		requests.WithQueryParameter("traceId", "abc-123"),
	)
	if got := req.GetQueryParameters().Get("traceId"); got != "abc-123" {
		t.Errorf("traceId = %q, want abc-123", got)
	}

	missing := requests.NewPostClearContentsRequest("", "Sheet1")
	if err := missing.Validate(); err == nil {
		t.Fatal("Validate() = nil, want an error for missing name")
	}
}

// TestOfflineResponseStatus verifies the shared status-code helper.
func TestOfflineResponseStatus(t *testing.T) {
	ok := &asposecellscloud.RichResponse{StatusCode: 200, Body: []byte(`{"Status":"OK"}`)}
	if err := asposecellscloud.CheckResponseStatus(ok); err != nil {
		t.Fatalf("CheckResponseStatus(200) = %v, want nil", err)
	}

	bad := &asposecellscloud.RichResponse{StatusCode: 401, Body: []byte(`{"Status":"Unauthorized"}`)}
	if err := asposecellscloud.CheckResponseStatus(bad); err == nil {
		t.Fatal("CheckResponseStatus(401) = nil, want an error")
	}
}
