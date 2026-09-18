package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestReplacer_ReplaceTextInLocalFile(t *testing.T) {
	ctx := context.Background()
	bookTextXlsx := "BookText.xlsx"
	request := requests.NewReplaceSpreadsheetContentRequest(
		"****",
		"Bike",
		"TestData/"+bookTextXlsx,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestReplacer_ReplaceTextInLocalFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestReplacer_ReplaceTextFromWorksheetInLocalFile(t *testing.T) {
	ctx := context.Background()
	bookTextXlsx := "BookText.xlsx"
	request := requests.NewReplaceSpreadsheetContentRequest(
		"****",
		"Bike",
		"TestData/"+bookTextXlsx,
		requests.WithCommonParameter("worksheet", "Sales"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestReplacer_ReplaceTextFromWorksheetInLocalFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestReplacer_ReplaceTextInRemoteRange(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	bookTextXlsx := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookTextXlsx, GetBaseTest().localTestDataFolder+bookTextXlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewReplaceContentInRemoteRangeRequest(
		"A1:A10",
		bookTextXlsx,
		"****",
		"Bike",
		"Sales",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestReplacer_ReplaceTextInRemoteRange \n", GetBaseTest().GetTestNumber())
	}
}
