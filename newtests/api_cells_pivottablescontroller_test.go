package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestPivotTablesController_GetWorksheetPivotTables(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetPivotTablesRequest(
		remoteName,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTables \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetWorksheetPivotTable(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetPivotTableRequest(
		remoteName,
		0,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetPivotTableField(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetPivotTableFieldRequest(
		remoteName,
		0,
		"Row",
		0,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetPivotTableField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_GetWorksheetPivotTableFilters(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetPivotTableFiltersRequest(
		remoteName,
		0,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_GetWorksheetPivotTableFilters \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PutWorksheetPivotTable(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetPivotTableRequest(
		remoteName,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("sourceData", "=Sheet1!C6:E13"),
		requests.WithCommonParameter("destCellName", "C1"),
		requests.WithCommonParameter("tableName", "TestPivot"),
		requests.WithCommonParameter("useSameSource", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PutWorksheetPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PutPivotTableField(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	var pivotTableFieldRequestData = []interface{}{int64(0)}
	pivotTableFieldRequest := &models.PivotTableFieldRequest{}
	pivotTableFieldRequest.Data = pivotTableFieldRequestData
	request := requests.NewPutPivotTableFieldRequest(
		remoteName,
		"Row",
		pivotTableFieldRequest,
		0,
		"Sheet4",
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PutPivotTableField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableFieldHideItem(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostPivotTableFieldHideItemRequest(
		0,
		true,
		1,
		remoteName,
		"Row",
		0,
		"Sheet4",
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableFieldHideItem \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableFieldMoveTo(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostPivotTableFieldMoveToRequest(
		0,
		"Row",
		remoteName,
		0,
		"Sheet4",
		"Column",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableFieldMoveTo \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableCellStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	styleFont := &models.Font{}
	styleFont.Size = asposecellscloud.Int32Ptr(16)
	style := &models.Style{}
	style.Font = styleFont
	request := requests.NewPostPivotTableCellStyleRequest(
		1,
		remoteName,
		0,
		1,
		"Sheet4",
		style,
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableCellStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	styleFont := &models.Font{}
	styleFont.Size = asposecellscloud.Int32Ptr(16)
	style := &models.Style{}
	style.Font = styleFont
	request := requests.NewPostPivotTableStyleRequest(
		remoteName,
		0,
		"Sheet4",
		style,
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableUpdatePivotFields(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	pivotField := &models.PivotField{}
	pivotField.ShowCompact = asposecellscloud.BoolPtr(true)
	request := requests.NewPostPivotTableUpdatePivotFieldsRequest(
		remoteName,
		pivotField,
		"Row",
		0,
		"Sheet4",
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableUpdatePivotFields \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostPivotTableUpdatePivotField(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	pivotField := &models.PivotField{}
	pivotField.ShowCompact = asposecellscloud.BoolPtr(true)
	request := requests.NewPostPivotTableUpdatePivotFieldRequest(
		remoteName,
		pivotField,
		0,
		"Row",
		0,
		"Sheet4",
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostPivotTableUpdatePivotField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostWorksheetPivotTableCalculate(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetPivotTableCalculateRequest(
		remoteName,
		0,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostWorksheetPivotTableCalculate \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_PostWorksheetPivotTableMove(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetPivotTableMoveRequest(
		remoteName,
		0,
		"Sheet4",
		requests.WithCommonParameter("row", intPtr(1)),
		requests.WithCommonParameter("column", intPtr(1)),
		requests.WithCommonParameter("destCellName", "C10"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_PostWorksheetPivotTableMove \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTables(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetPivotTablesRequest(
		remoteName,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTables \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTable(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetPivotTableRequest(
		remoteName,
		0,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeletePivotTableField(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	var pivotTableFieldRequestData = []interface{}{int64(0)}
	pivotTableFieldRequest := &models.PivotTableFieldRequest{}
	pivotTableFieldRequest.Data = pivotTableFieldRequestData
	request := requests.NewDeletePivotTableFieldRequest(
		remoteName,
		"Row",
		pivotTableFieldRequest,
		0,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeletePivotTableField \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTableFilters(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetPivotTableFiltersRequest(
		remoteName,
		0,
		"Sheet3",
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTableFilters \n", GetBaseTest().GetTestNumber())
	}
}

func TestPivotTablesController_DeleteWorksheetPivotTableFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetPivotTableFilterRequest(
		0,
		remoteName,
		0,
		"Sheet3",
		requests.WithCommonParameter("needReCalculate", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPivotTablesController_DeleteWorksheetPivotTableFilter \n", GetBaseTest().GetTestNumber())
	}
}
