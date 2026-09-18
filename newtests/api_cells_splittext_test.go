package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestSplitText_SplitText(t *testing.T) {
	ctx := context.Background()
	request := requests.NewSplitTextRequest(
		"Comma",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("keepDelimitersInResultingCells", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("keepDelimitersPosition", "BeforeText"),
		requests.WithCommonParameter("HowToSplit", "SplitToColumns"),
		requests.WithCommonParameter("worksheet", "Bikes"),
		requests.WithCommonParameter("range", "A1:A10"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSplitText_SplitText \n", GetBaseTest().GetTestNumber())
	}
}
