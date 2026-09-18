package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestSearcher_SearchTextInLocalFile(t *testing.T) {
	ctx := context.Background()
	bookTextXlsx := "BookText.xlsx"
	request := requests.NewSearchSpreadsheetContentRequest(
		"Bike",
		"TestData/"+bookTextXlsx,
		requests.WithCommonParameter("ignoringCase", asposecellscloud.BoolPtr(false)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchTextInLocalFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchTextFromWorksheetInLocalFile(t *testing.T) {
	ctx := context.Background()
	bookTextXlsx := "BookText.xlsx"
	request := requests.NewSearchSpreadsheetContentRequest(
		"Bike",
		"TestData/"+bookTextXlsx,
		requests.WithCommonParameter("ignoringCase", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("worksheet", "Sales"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchTextFromWorksheetInLocalFile \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchTextInRemoteSpreadsheet(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	bookTextXlsx := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookTextXlsx, GetBaseTest().localTestDataFolder+bookTextXlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewSearchContentInRemoteSpreadsheetRequest(
		bookTextXlsx,
		"Bike",
		requests.WithCommonParameter("ignoringCase", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchTextInRemoteSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchTextInRemoteRange(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	bookTextXlsx := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookTextXlsx, GetBaseTest().localTestDataFolder+bookTextXlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewSearchContentInRemoteRangeRequest(
		"A1:A10",
		bookTextXlsx,
		"Bike",
		"Sales",
		requests.WithCommonParameter("ignoringCase", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchTextInRemoteRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchSpreadsheetBrokenLinks(t *testing.T) {
	ctx := context.Background()
	bookFormulaXlsx := "BookFormula.xlsx"
	request := requests.NewSearchSpreadsheetBrokenLinksRequest(
		"TestData/" + bookFormulaXlsx,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchSpreadsheetBrokenLinks \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchBrokenLinksInRemoteSpreadsheet(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	bookFormulaXlsx := "BookFormula.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookFormulaXlsx, GetBaseTest().localTestDataFolder+bookFormulaXlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewSearchBrokenLinksInRemoteSpreadsheetRequest(
		bookFormulaXlsx,
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchBrokenLinksInRemoteSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchBrokenLinksInRemoteWorksheet(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	bookFormulaXlsx := "BookFormula.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookFormulaXlsx, GetBaseTest().localTestDataFolder+bookFormulaXlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewSearchBrokenLinksInRemoteWorksheetRequest(
		bookFormulaXlsx,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchBrokenLinksInRemoteWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestSearcher_SearchBrokenLinksInRemoteRange(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	bookFormulaXlsx := "BookFormula.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+bookFormulaXlsx, GetBaseTest().localTestDataFolder+bookFormulaXlsx, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewSearchBrokenLinksInRemoteRangeRequest(
		"A1:F40",
		bookFormulaXlsx,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSearcher_SearchBrokenLinksInRemoteRange \n", GetBaseTest().GetTestNumber())
	}
}
