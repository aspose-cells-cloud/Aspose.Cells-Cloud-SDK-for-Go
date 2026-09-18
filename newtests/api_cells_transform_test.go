package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestTransform_RemoveSpreadsheetBlankRows(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewRemoveSpreadsheetBlankRowsRequest(
		"TestData/" + localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTransform_RemoveSpreadsheetBlankRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestTransform_RemoveSpreadsheetBlankColumns(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewRemoveSpreadsheetBlankColumnsRequest(
		"TestData/" + localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTransform_RemoveSpreadsheetBlankColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestTransform_RemoveSpreadsheetBlankWorksheets(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary-BlankWorksheet.xlsx"
	request := requests.NewRemoveSpreadsheetBlankWorksheetsRequest(
		"TestData/" + localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTransform_RemoveSpreadsheetBlankWorksheets \n", GetBaseTest().GetTestNumber())
	}
}
