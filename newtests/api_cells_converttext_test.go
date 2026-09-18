package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestConvertText_ConvertText(t *testing.T) {
	ctx := context.Background()
	request := requests.NewConvertTextRequest(
		"ConvertNumberToText",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("sourceCharacters", ""),
		requests.WithCommonParameter("targetCharacters", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConvertText_ConvertText \n", GetBaseTest().GetTestNumber())
	}
}

func TestConvertText_ConvertTextWithConvertCharacters(t *testing.T) {
	ctx := context.Background()
	request := requests.NewConvertTextRequest(
		"ConvertCharacters",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("sourceCharacters", "Bikes"),
		requests.WithCommonParameter("targetCharacters", "MOTO"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConvertText_ConvertTextWithConvertCharacters \n", GetBaseTest().GetTestNumber())
	}
}

func TestConvertText_ConvertTextWithConvertWriteSpace(t *testing.T) {
	ctx := context.Background()
	request := requests.NewConvertTextRequest(
		"ConvertWriteSpace",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("sourceCharacters", ""),
		requests.WithCommonParameter("targetCharacters", "MOTO"),
		requests.WithCommonParameter("worksheet", "Bikes"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConvertText_ConvertTextWithConvertWriteSpace \n", GetBaseTest().GetTestNumber())
	}
}
