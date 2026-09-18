package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestManagement_CreateSpreadsheet(t *testing.T) {
	ctx := context.Background()
	request := requests.NewCreateSpreadsheetRequest(
		requests.WithCommonParameter("format", "xlsx"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_CreateSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestManagement_CreateSpreadsheetWithTemplate(t *testing.T) {
	ctx := context.Background()
	request := requests.NewCreateSpreadsheetRequest(
		requests.WithCommonParameter("format", "pdf"),
		requests.WithCommonParameter("template", "SalesDataComparisonXLSX"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_CreateSpreadsheetWithTemplate \n", GetBaseTest().GetTestNumber())
	}
}

func TestManagement_AddWorksheet(t *testing.T) {
	ctx := context.Background()
	request := requests.NewAddWorksheetToSpreadsheetRequest(
		"TestData/AggregateCellsByColor.xlsx",
		requests.WithCommonParameter("sheetType", "Worksheet"),
		requests.WithCommonParameter("position", intPtr(1)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_AddWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestManagement_DeleteWorksheet(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewDeleteWorksheetFromSpreadsheetRequest(
		"Sales",
		"TestData/"+localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_DeleteWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestManagement_RenameWorksheet(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewRenameWorksheetInSpreadsheetRequest(
		"Sales",
		"TestData/"+localName,
		"SalesData",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_RenameWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestManagement_MoveWorksheet(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewMoveWorksheetInSpreadsheetRequest(
		1,
		"TestData/"+localName,
		"Sales",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_MoveWorksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestManagement_CompressSpreadsheet(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewCompressSpreadsheetRequest(
		9,
		"TestData/"+localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestManagement_CompressSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}
