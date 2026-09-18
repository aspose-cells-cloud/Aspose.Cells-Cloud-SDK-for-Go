package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestRemoveCharacters_RemoveCharacters(t *testing.T) {
	ctx := context.Background()
	request := requests.NewRemoveCharactersRequest(
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("removeTextMethod", "RemoveCharacterSets"),
		requests.WithCommonParameter("characterSets", "NonPrintingCharacters"),
		requests.WithCommonParameter("removeCustomValue", ""),
		requests.WithCommonParameter("worksheet", "Text"),
		requests.WithCommonParameter("range", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRemoveCharacters_RemoveCharacters \n", GetBaseTest().GetTestNumber())
	}
}

func TestRemoveCharacters_RemoveDuplicateSubstrings(t *testing.T) {
	ctx := context.Background()
	request := requests.NewRemoveDuplicateSubstringsRequest(
		"Space",
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("worksheet", "Text"),
		requests.WithCommonParameter("range", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRemoveCharacters_RemoveDuplicateSubstrings \n", GetBaseTest().GetTestNumber())
	}
}

func TestRemoveCharacters_RemoveCharactersWithFirstNCharacters(t *testing.T) {
	ctx := context.Background()
	request := requests.NewRemoveCharactersByPositionRequest(
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("theFirstNCharacters", intPtr(5)),
		requests.WithCommonParameter("theLastNCharacters", intPtr(3)),
		requests.WithCommonParameter("allCharactersBeforeText", ""),
		requests.WithCommonParameter("allCharactersAfterText", ""),
		requests.WithCommonParameter("worksheet", "Text"),
		requests.WithCommonParameter("range", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRemoveCharacters_RemoveCharactersWithFirstNCharacters \n", GetBaseTest().GetTestNumber())
	}
}

func TestRemoveCharacters_RemoveCharactersWithAllCharactersBeforeText(t *testing.T) {
	ctx := context.Background()
	request := requests.NewRemoveCharactersByPositionRequest(
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("theFirstNCharacters", intPtr(0)),
		requests.WithCommonParameter("theLastNCharacters", intPtr(0)),
		requests.WithCommonParameter("allCharactersBeforeText", "Designed"),
		requests.WithCommonParameter("allCharactersAfterText", "distance"),
		requests.WithCommonParameter("worksheet", "Text"),
		requests.WithCommonParameter("range", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRemoveCharacters_RemoveCharactersWithAllCharactersBeforeText \n", GetBaseTest().GetTestNumber())
	}
}

func TestRemoveCharacters_RemoveCharactersByPosition(t *testing.T) {
	ctx := context.Background()
	request := requests.NewRemoveCharactersByPositionRequest(
		"TestData/BookText.xlsx",
		requests.WithCommonParameter("theFirstNCharacters", intPtr(5)),
		requests.WithCommonParameter("theLastNCharacters", intPtr(3)),
		requests.WithCommonParameter("allCharactersBeforeText", ""),
		requests.WithCommonParameter("allCharactersAfterText", ""),
		requests.WithCommonParameter("worksheet", ""),
		requests.WithCommonParameter("range", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRemoveCharacters_RemoveCharactersByPosition \n", GetBaseTest().GetTestNumber())
	}
}
