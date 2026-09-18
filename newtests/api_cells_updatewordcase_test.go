package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestUpdateWordCase_UpdateWordCase(t *testing.T) {
	ctx := context.Background()
	request := requests.NewUpdateWordCaseRequest(
		"TestData/BookText.xlsx",
		"ProperCase",
		requests.WithCommonParameter("worksheet", "Bikes"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestUpdateWordCase_UpdateWordCase \n", GetBaseTest().GetTestNumber())
	}
}

func TestUpdateWordCase_UpdateWordCaseInRange(t *testing.T) {
	ctx := context.Background()
	request := requests.NewUpdateWordCaseRequest(
		"TestData/BookText.xlsx",
		"ProperCase",
		requests.WithCommonParameter("worksheet", "Bikes"),
		requests.WithCommonParameter("range", "A1:B15"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestUpdateWordCase_UpdateWordCaseInRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestUpdateWordCase_UpdateWordCaseInSpreadsheet(t *testing.T) {
	ctx := context.Background()
	request := requests.NewUpdateWordCaseRequest(
		"TestData/BookText.xlsx",
		"ProperCase",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestUpdateWordCase_UpdateWordCaseInSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}
