package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestAddText_AddText(t *testing.T) {
	ctx := context.Background()
	request := requests.NewAddTextRequest(
		"AtTheBeginning",
		"TestData/BookText.xlsx",
		"New",
		requests.WithCommonParameter("selectText", "text"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAddText_AddText \n", GetBaseTest().GetTestNumber())
	}
}

func TestAddText_AddTextSkipEmptyCells(t *testing.T) {
	ctx := context.Background()
	request := requests.NewAddTextRequest(
		"AtTheBeginning",
		"TestData/BookText.xlsx",
		"New",
		requests.WithCommonParameter("selectText", ""),
		requests.WithCommonParameter("skipEmptyCells", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("worksheet", "Bikes"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAddText_AddTextSkipEmptyCells \n", GetBaseTest().GetTestNumber())
	}
}

func TestAddText_AddTextInRange(t *testing.T) {
	ctx := context.Background()
	request := requests.NewAddTextRequest(
		"AtTheBeginning",
		"TestData/BookText.xlsx",
		"New",
		requests.WithCommonParameter("selectText", ""),
		requests.WithCommonParameter("skipEmptyCells", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("worksheet", "Bikes"),
		requests.WithCommonParameter("range", "A1:B15"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAddText_AddTextInRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestAddText_AddTextBeforeText(t *testing.T) {
	ctx := context.Background()
	request := requests.NewAddTextRequest(
		"BeforeText",
		"TestData/BookText.xlsx",
		"New",
		requests.WithCommonParameter("selectText", "bike"),
		requests.WithCommonParameter("worksheet", "Bikes"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAddText_AddTextBeforeText \n", GetBaseTest().GetTestNumber())
	}
}
