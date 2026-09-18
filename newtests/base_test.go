/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="base_test.go">
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

package newtests

import (
	"context"
	"os"
	"strconv"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

// These tests are the new-model (v4.0) port of the legacy integration tests in
// ../integrationtests. Test names, file names and test data are kept identical
// so the two suites can be compared one to one.

var BaseTestInstance *BaseTest

type BaseTest struct {
	remoteFolder        string
	localTestDataFolder string
	Client              *asposecellscloud.AsposeCellsCloudClient
	TestNumber          int
}

func (bt *BaseTest) FromBealoonToString(v bool) string {
	return strconv.FormatBool(v)
}

func (bt *BaseTest) FromIntToString(v int) string {
	return strconv.FormatInt(int64(v), 10)
}

func (bt *BaseTest) GetTestNumber() int {
	bt.TestNumber++
	return bt.TestNumber
}

func NewBaseTest() *BaseTest {
	bt := &BaseTest{
		remoteFolder:        "GoTest",
		localTestDataFolder: "TestData/",
		TestNumber:          0,
		// Get Client Secret and Client Id from https://aspose.cloud
		Client: asposecellscloud.NewAsposeCellsCloudClient(
			os.Getenv("CellsCloudClientId"),
			os.Getenv("CellsCloudClientSecret"),
			os.Getenv("CellsCloudApiBaseUrl"),
		),
	}
	return bt
}

func GetBaseTest() *BaseTest {
	if BaseTestInstance == nil {
		BaseTestInstance = NewBaseTest()
	}
	return BaseTestInstance
}

// mustUploadFile uploads a local file to the remote storage. The legacy tests
// ignored the outcome of this call; here a failure aborts the test, because a
// missing remote file makes every assertion that follows meaningless.
func mustUploadFile(t *testing.T, path string, uploadFiles string, storageName string) error {
	t.Helper()
	if uploadFiles == "" {
		return nil
	}
	req := requests.NewUploadFileRequest(path, uploadFiles,
		requests.WithCommonParameter("storageName", storageName))
	_, err := asposecellscloud.DoChecked(context.Background(), GetBaseTest().Client, req)
	return err
}

// intPtr covers the *int option parameters of the generated requests package.
// The SDK's own ptr.go only provides Int32Ptr/Int64Ptr/Float64Ptr/BoolPtr.
func intPtr(v int) *int {
	return &v
}
