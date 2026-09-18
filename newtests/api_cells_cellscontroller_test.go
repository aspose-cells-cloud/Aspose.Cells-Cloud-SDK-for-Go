package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestCellsController_PostClearContents(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostClearContentsRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("range", "A1:C10"),
		requests.WithCommonParameter("startRow", intPtr(1)),
		requests.WithCommonParameter("startColumn", intPtr(1)),
		requests.WithCommonParameter("endRow", intPtr(3)),
		requests.WithCommonParameter("endColumn", intPtr(3)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostClearContents \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostClearFormats(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostClearFormatsRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("range", "A1:C10"),
		requests.WithCommonParameter("startRow", intPtr(1)),
		requests.WithCommonParameter("startColumn", intPtr(1)),
		requests.WithCommonParameter("endRow", intPtr(3)),
		requests.WithCommonParameter("endColumn", intPtr(3)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostClearFormats \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUpdateWorksheetRangeStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	styleFont := &models.Font{}
	styleFont.Size = asposecellscloud.Int32Ptr(16)
	style := &models.Style{}
	style.Font = styleFont
	request := requests.NewPostUpdateWorksheetRangeStyleRequest(
		remoteName,
		"A1:C10",
		"Sheet1",
		style,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUpdateWorksheetRangeStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostWorksheetMerge(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetMergeRequest(
		remoteName,
		"Sheet1",
		1,
		1,
		4,
		4,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostWorksheetMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostWorksheetUnmerge(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetUnmergeRequest(
		remoteName,
		"Sheet1",
		1,
		1,
		4,
		4,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostWorksheetUnmerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetCells(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetCellsRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("offest", intPtr(1)),
		requests.WithCommonParameter("count", intPtr(10)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetCells \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetCell(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetCellRequest(
		"A1",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetCell \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetCellStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetCellStyleRequest(
		"A1",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetCellStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostWorksheetCellSetValue(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetCellSetValueRequest(
		"A1",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("value", "1"),
		requests.WithCommonParameter("type", "int"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostWorksheetCellSetValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUpdateWorksheetCellStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	styleFont := &models.Font{}
	styleFont.Size = asposecellscloud.Int32Ptr(16)
	style := &models.Style{}
	style.Font = styleFont
	request := requests.NewPostUpdateWorksheetCellStyleRequest(
		"A1",
		remoteName,
		"Sheet1",
		style,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUpdateWorksheetCellStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostSetCellRangeValue(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostSetCellRangeValueRequest(
		"A1:C10",
		remoteName,
		"Sheet1",
		"string",
		"Test",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostSetCellRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCopyCellIntoCell(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostCopyCellIntoCellRequest(
		"C1",
		remoteName,
		"Sheet1",
		"Sheet2",
		requests.WithCommonParameter("cellname", "A1"),
		requests.WithCommonParameter("row", intPtr(1)),
		requests.WithCommonParameter("column", intPtr(1)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostCopyCellIntoCell \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetCellHtmlString(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetCellHtmlStringRequest(
		"A1",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetCellHtmlString \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostSetCellHtmlString(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostSetCellHtmlStringRequest(
		"A1",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostSetCellHtmlString \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCellCalculate(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	options := &models.CalculationOptions{}
	options.Recursive = asposecellscloud.BoolPtr(true)
	options.IgnoreError = asposecellscloud.BoolPtr(true)
	request := requests.NewPostCellCalculateRequest(
		"A1",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("options", options),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostCellCalculate \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCellCharacters(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	optionsvalue0Font := &models.Font{}
	optionsvalue0Font.IsBold = asposecellscloud.BoolPtr(true)
	optionsvalue0Font.Size = asposecellscloud.Int32Ptr(16)
	optionsvalue0 := &models.FontSetting{}
	optionsvalue0.Length = asposecellscloud.Int32Ptr(5)
	optionsvalue0.StartIndex = asposecellscloud.Int32Ptr(0)
	optionsvalue0.Font = optionsvalue0Font
	var options = []models.FontSetting{*optionsvalue0}
	request := requests.NewPostCellCharactersRequest(
		"E36",
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("options", options),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostCellCharacters \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetColumnsRequest(
		requests.WithCommonParameter("name", remoteName),
		requests.WithCommonParameter("sheetName", "Sheet1"),
		requests.WithCommonParameter("offset", intPtr(1)),
		requests.WithCommonParameter("count", intPtr(10)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostSetWorksheetColumnWidth(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostSetWorksheetColumnWidthRequest(
		1,
		remoteName,
		"Sheet1",
		10.9,
		requests.WithCommonParameter("count", intPtr(10)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostSetWorksheetColumnWidth \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetColumn(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetColumnRequest(
		1,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetColumn \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PutInsertWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutInsertWorksheetColumnsRequest(
		1,
		10,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("updateReference", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PutInsertWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_DeleteWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetColumnsRequest(
		1,
		10,
		remoteName,
		"Sheet1",
		true,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_DeleteWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostHideWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostHideWorksheetColumnsRequest(
		remoteName,
		"Sheet1",
		1,
		10,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostHideWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUnhideWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostUnhideWorksheetColumnsRequest(
		remoteName,
		"Sheet1",
		1,
		10,
		requests.WithCommonParameter("width", asposecellscloud.Float64Ptr(10.9)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUnhideWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostGroupWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostGroupWorksheetColumnsRequest(
		1,
		9,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("hide", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostGroupWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUngroupWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostUngroupWorksheetColumnsRequest(
		1,
		9,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUngroupWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCopyWorksheetColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostCopyWorksheetColumnsRequest(
		8,
		19,
		remoteName,
		"Sheet1",
		1,
		requests.WithCommonParameter("worksheet", "Sheet2"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostCopyWorksheetColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostColumnStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	styleFont := &models.Font{}
	styleFont.Size = asposecellscloud.Int32Ptr(16)
	style := &models.Style{}
	style.Font = styleFont
	request := requests.NewPostColumnStyleRequest(
		1,
		remoteName,
		"Sheet1",
		style,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostColumnStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetRowsRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("offset", intPtr(1)),
		requests.WithCommonParameter("count", intPtr(10)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_GetWorksheetRow(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetRowRequest(
		remoteName,
		1,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_GetWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_DeleteWorksheetRow(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetRowRequest(
		remoteName,
		1,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_DeleteWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_DeleteWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetRowsRequest(
		remoteName,
		"Sheet1",
		1,
		requests.WithCommonParameter("totalRows", intPtr(10)),
		requests.WithCommonParameter("updateReference", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_DeleteWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PutInsertWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutInsertWorksheetRowsRequest(
		remoteName,
		"Sheet1",
		1,
		requests.WithCommonParameter("totalRows", intPtr(10)),
		requests.WithCommonParameter("updateReference", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PutInsertWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PutInsertWorksheetRow(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutInsertWorksheetRowRequest(
		remoteName,
		1,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PutInsertWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUpdateWorksheetRow(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostUpdateWorksheetRowRequest(
		remoteName,
		1,
		"Sheet1",
		requests.WithCommonParameter("height", asposecellscloud.Float64Ptr(10.8)),
		requests.WithCommonParameter("count", intPtr(9)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUpdateWorksheetRow \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostHideWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostHideWorksheetRowsRequest(
		remoteName,
		"Sheet1",
		1,
		6,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostHideWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUnhideWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostUnhideWorksheetRowsRequest(
		remoteName,
		"Sheet1",
		1,
		8,
		requests.WithCommonParameter("height", asposecellscloud.Float64Ptr(10.9)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUnhideWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostGroupWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostGroupWorksheetRowsRequest(
		1,
		9,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("hide", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostGroupWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostUngroupWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostUngroupWorksheetRowsRequest(
		1,
		9,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("isAll", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostUngroupWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostCopyWorksheetRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostCopyWorksheetRowsRequest(
		12,
		remoteName,
		5,
		"Sheet1",
		1,
		requests.WithCommonParameter("worksheet", "Sheet2"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostCopyWorksheetRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestCellsController_PostRowStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	styleFont := &models.Font{}
	styleFont.Size = asposecellscloud.Int32Ptr(16)
	style := &models.Style{}
	style.Font = styleFont
	request := requests.NewPostRowStyleRequest(
		remoteName,
		1,
		"Sheet1",
		style,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestCellsController_PostRowStyle \n", GetBaseTest().GetTestNumber())
	}
}
