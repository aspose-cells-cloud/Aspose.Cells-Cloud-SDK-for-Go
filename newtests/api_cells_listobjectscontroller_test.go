package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestListObjectsController_GetWorksheetListObjects(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetListObjectsRequest(
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_GetWorksheetListObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_GetWorksheetListObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetListObjectRequest(
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("format", "pdf"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_GetWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PutWorksheetListObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetListObjectRequest(
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("startRow", intPtr(1)),
		requests.WithCommonParameter("startColumn", intPtr(1)),
		requests.WithCommonParameter("endRow", intPtr(6)),
		requests.WithCommonParameter("endColumn", intPtr(6)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("hasHeaders", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("displayName", "true"),
		requests.WithCommonParameter("showTotals", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PutWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_DeleteWorksheetListObjects(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetListObjectsRequest(
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_DeleteWorksheetListObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_DeleteWorksheetListObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetListObjectRequest(
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_DeleteWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	listObject := &models.ListObject{}
	listObject.ShowHeaderRow = asposecellscloud.BoolPtr(true)
	request := requests.NewPostWorksheetListObjectRequest(
		listObject,
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectConvertToRange(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetListObjectConvertToRangeRequest(
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectConvertToRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectSummarizeWithPivotTable(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	var createPivotTableRequestPivotFieldColumns = []interface{}{int64(2)}
	var createPivotTableRequestPivotFieldData = []interface{}{int64(1)}
	var createPivotTableRequestPivotFieldRows = []interface{}{int64(0)}
	createPivotTableRequest := &models.CreatePivotTableRequest{}
	createPivotTableRequest.DestCellName = "C1"
	createPivotTableRequest.Name = "testp"
	createPivotTableRequest.SourceData = "=Sheet2!A1:E8"
	createPivotTableRequest.UseSameSource = asposecellscloud.BoolPtr(true)
	createPivotTableRequest.PivotFieldColumns = createPivotTableRequestPivotFieldColumns
	createPivotTableRequest.PivotFieldData = createPivotTableRequestPivotFieldData
	createPivotTableRequest.PivotFieldRows = createPivotTableRequestPivotFieldRows
	request := requests.NewPostWorksheetListObjectSummarizeWithPivotTableRequest(
		createPivotTableRequest,
		"Sheet2",
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectSummarizeWithPivotTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectSortTable(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	dataSorter := &models.DataSorter{}
	dataSorter.CaseSensitive = asposecellscloud.BoolPtr(true)
	request := requests.NewPostWorksheetListObjectSortTableRequest(
		dataSorter,
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectSortTable \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListColumn(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	listColumn := &models.ListColumn{}
	listColumn.Name = "test cloumn"
	request := requests.NewPostWorksheetListColumnRequest(
		0,
		listColumn,
		0,
		remoteName,
		"Sheet7",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListColumn \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListColumnsTotal(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	tableTotalRequeststableTotalRequest0 := &models.TableTotalRequest{}
	tableTotalRequeststableTotalRequest0.ListColumnIndex = asposecellscloud.Int32Ptr(1)
	tableTotalRequeststableTotalRequest0.TotalsCalculation = "Average"
	var tableTotalRequests = []models.TableTotalRequest{*tableTotalRequeststableTotalRequest0}
	request := requests.NewPostWorksheetListColumnsTotalRequest(
		0,
		remoteName,
		"Sheet7",
		tableTotalRequests,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListColumnsTotal \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectRemoveDuplicates(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestTables.xlsx"
	remoteName := "TestTables.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetListObjectRemoveDuplicatesRequest(
		0,
		remoteName,
		"Sheet2",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectRemoveDuplicates \n", GetBaseTest().GetTestNumber())
	}
}

func TestListObjectsController_PostWorksheetListObjectInsertSlicer(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestTables.xlsx"
	remoteName := "TestTables.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetListObjectInsertSlicerRequest(
		2,
		"j9",
		0,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestListObjectsController_PostWorksheetListObjectInsertSlicer \n", GetBaseTest().GetTestNumber())
	}
}
