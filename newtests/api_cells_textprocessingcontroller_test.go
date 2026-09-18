package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestTextProcessingController_PostAddTextContent(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookText.xlsx"
	remoteName := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	addTextOptionsDataSource := &models.DataSource{}
	addTextOptionsDataSource.DataSourceType = "CloudFileSystem"
	addTextOptionsDataSource.DataPath = "TestData/In/BookText.xlsx"
	addTextOptionsScopeOptions := &models.ScopeOptions{}
	addTextOptionsScopeOptions.Scope = "Workbook"
	addTextOptions := &models.AddTextOptions{}
	addTextOptions.DataSource = addTextOptionsDataSource
	addTextOptions.Text = "Aspose.Cells Cloud is an excellent product."
	addTextOptions.ScopeOptions = addTextOptionsScopeOptions
	addTextOptions.SelectPoistion = "AtTheBeginning"
	addTextOptions.SkipEmptyCells = asposecellscloud.BoolPtr(true)
	request := requests.NewPostAddTextContentRequest(
		addTextOptions,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTextProcessingController_PostAddTextContent \n", GetBaseTest().GetTestNumber())
	}
}

func TestTextProcessingController_PostTrimContent(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookText.xlsx"
	remoteName := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	trimContentOptionsDataSource := &models.DataSource{}
	trimContentOptionsDataSource.DataSourceType = "CloudFileSystem"
	trimContentOptionsDataSource.DataPath = "TestData/In/BookText.xlsx"
	trimContentOptionsScopeOptions := &models.ScopeOptions{}
	trimContentOptionsScopeOptions.Scope = "EntireWorkbook"
	trimContentOptions := &models.TrimContentOptions{}
	trimContentOptions.DataSource = trimContentOptionsDataSource
	trimContentOptions.TrimLeading = asposecellscloud.BoolPtr(true)
	trimContentOptions.TrimTrailing = asposecellscloud.BoolPtr(true)
	trimContentOptions.TrimSpaceBetweenWordTo1 = asposecellscloud.BoolPtr(true)
	trimContentOptions.RemoveAllLineBreaks = asposecellscloud.BoolPtr(true)
	trimContentOptions.ScopeOptions = trimContentOptionsScopeOptions
	request := requests.NewPostTrimContentRequest(
		trimContentOptions,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTextProcessingController_PostTrimContent \n", GetBaseTest().GetTestNumber())
	}
}

func TestTextProcessingController_PostUpdateWordCase(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookText.xlsx"
	remoteName := "BookText.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	wordCaseOptionsDataSource := &models.DataSource{}
	wordCaseOptionsDataSource.DataSourceType = "CloudFileSystem"
	wordCaseOptionsDataSource.DataPath = "TestData/In/BookText.xlsx"
	wordCaseOptionsScopeOptions := &models.ScopeOptions{}
	wordCaseOptionsScopeOptions.Scope = "EntireWorkbook"
	wordCaseOptions := &models.WordCaseOptions{}
	wordCaseOptions.DataSource = wordCaseOptionsDataSource
	wordCaseOptions.WordCaseType = "None"
	wordCaseOptions.ScopeOptions = wordCaseOptionsScopeOptions
	request := requests.NewPostUpdateWordCaseRequest(
		wordCaseOptions,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestTextProcessingController_PostUpdateWordCase \n", GetBaseTest().GetTestNumber())
	}
}
