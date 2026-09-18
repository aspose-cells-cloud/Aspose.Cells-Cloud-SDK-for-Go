package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestCalculate_AggregateCellsByColor(t *testing.T) {
	ctx := context.Background()
	request := requests.NewAggregateCellsByColorRequest(
		"TestData/AggregateCellsByColor.xlsx",
		requests.WithCommonParameter("worksheet", "Sheet1"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCalculate_AggregateCellsByColor \n", GetBaseTest().GetTestNumber())
	}
}

func TestCalculate_MathCalculate(t *testing.T) {
	ctx := context.Background()
	request := requests.NewMathCalculateRequest(
		"add",
		"TestData/EmployeeSalesSummary-BlankWorksheet.xlsx",
		"12.3",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCalculate_MathCalculate \n", GetBaseTest().GetTestNumber())
	}
}
