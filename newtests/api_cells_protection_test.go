package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestProtection_ProtectSpreadsheet(t *testing.T) {
	ctx := context.Background()
	request := requests.NewProtectSpreadsheetRequest(
		"123456",
		"123456",
		"TestData/EmployeeSalesSummary.xlsx",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestProtection_ProtectSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestProtection_UnprotectSpreadsheet(t *testing.T) {
	ctx := context.Background()
	request := requests.NewUnprotectSpreadsheetRequest(
		"123456",
		"123456",
		"TestData/EmployeeSalesSummary_Locked.xlsx",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestProtection_UnprotectSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}
