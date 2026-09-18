package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestMerger_MergeRemoteSpreadsheet(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	book1Xlsx := "Book1.xlsx"
	bookTextXlsx := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookTextXlsx, GetBaseTest().localTestDataFolder+bookTextXlsx, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/"+book1Xlsx, GetBaseTest().localTestDataFolder+book1Xlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewMergeRemoteSpreadsheetRequest(
		remoteFolder+"/"+book1Xlsx,
		bookTextXlsx,
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestMerger_MergeRemoteSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}
