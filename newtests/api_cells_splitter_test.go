package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestSplitter_SplitLocalFile(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "PDF"
	request := requests.NewSplitSpreadsheetRequest(
		"TestData/"+book1Xlsx,
		requests.WithCommonParameter("outFormat", outFormat),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSplitter_SplitLocalFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestSplitter_SplitLocalFileToRemoteFolder(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "PDF"
	request := requests.NewSplitSpreadsheetRequest(
		"TestData/"+book1Xlsx,
		requests.WithCommonParameter("outFormat", outFormat),
		requests.WithCommonParameter("outPath", "TestData/Out"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSplitter_SplitLocalFileToRemoteFolder \n", GetBaseTest().GetTestNumber())
	}
}

func TestSplitter_SplitFileInRemote(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	book1Xlsx := "Book1.xlsx"
	outFormat := "PDF"
	request := requests.NewSplitRemoteSpreadsheetRequest(
		book1Xlsx,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("outFormat", outFormat),
		requests.WithCommonParameter("outPath", "TestData/Out"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSplitter_SplitFileInRemote \n", GetBaseTest().GetTestNumber())
	}
}
