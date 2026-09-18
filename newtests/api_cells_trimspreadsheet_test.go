package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestTrimSpreadsheet_TrimCharacter(t *testing.T) {
	ctx := context.Background()
	request := requests.NewTrimCharacterRequest(
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("trimLeading", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("trimTrailing", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("trimSpaceBetweenWordTo1", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("trimNonBreakingSpaces", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("removeExtraLineBreaks", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("removeAllLineBreaks", asposecellscloud.BoolPtr(true)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTrimSpreadsheet_TrimCharacter \n", GetBaseTest().GetTestNumber())
	}
}

func TestTrimSpreadsheet_StartTrimCharacter(t *testing.T) {
	ctx := context.Background()
	request := requests.NewTrimCharacterRequest(
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("trimLeading", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("trimTrailing", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("trimSpaceBetweenWordTo1", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("trimNonBreakingSpaces", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("removeExtraLineBreaks", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("removeAllLineBreaks", asposecellscloud.BoolPtr(true)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTrimSpreadsheet_StartTrimCharacter \n", GetBaseTest().GetTestNumber())
	}
}
