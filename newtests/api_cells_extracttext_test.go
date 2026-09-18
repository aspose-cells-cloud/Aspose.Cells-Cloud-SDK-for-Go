package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestExtractText_ExtractText(t *testing.T) {
	ctx := context.Background()
	request := requests.NewExtractTextRequest(
		"ExtractFirstCharacter",
		"F1:F10",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("beforeText", ""),
		requests.WithCommonParameter("afterText", ""),
		requests.WithCommonParameter("beforePosition", intPtr(10)),
		requests.WithCommonParameter("afterPosition", intPtr(0)),
		requests.WithCommonParameter("worksheet", "Bikes"),
		requests.WithCommonParameter("range", "A1:A10"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestExtractText_ExtractText \n", GetBaseTest().GetTestNumber())
	}
}

func TestExtractText_ExtractTextWithExtractLastCharacter(t *testing.T) {
	ctx := context.Background()
	request := requests.NewExtractTextRequest(
		"ExtractLastCharacter",
		"F1:F10",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("beforeText", ""),
		requests.WithCommonParameter("afterText", ""),
		requests.WithCommonParameter("beforePosition", intPtr(0)),
		requests.WithCommonParameter("afterPosition", intPtr(10)),
		requests.WithCommonParameter("worksheet", "Bikes"),
		requests.WithCommonParameter("range", "A1:A10"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestExtractText_ExtractTextWithExtractLastCharacter \n", GetBaseTest().GetTestNumber())
	}
}

func TestExtractText_ExtractTextWithExtractTextAfter(t *testing.T) {
	ctx := context.Background()
	request := requests.NewExtractTextRequest(
		"ExtractTextAfter",
		"F1:F10",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("beforeText", ""),
		requests.WithCommonParameter("afterText", "bikes"),
		requests.WithCommonParameter("beforePosition", intPtr(0)),
		requests.WithCommonParameter("afterPosition", intPtr(0)),
		requests.WithCommonParameter("worksheet", "Bikes"),
		requests.WithCommonParameter("range", "A1:A10"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestExtractText_ExtractTextWithExtractTextAfter \n", GetBaseTest().GetTestNumber())
	}
}
