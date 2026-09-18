package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestRangesController_PostWorksheetCellsRangesCopy(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	rangeOperateSource := &models.Range{}
	rangeOperateSource.ColumnCount = asposecellscloud.Int32Ptr(3)
	rangeOperateSource.FirstColumn = asposecellscloud.Int32Ptr(8)
	rangeOperateSource.FirstRow = asposecellscloud.Int32Ptr(3)
	rangeOperateSource.RowCount = asposecellscloud.Int32Ptr(2)
	rangeOperateTarget := &models.Range{}
	rangeOperateTarget.ColumnCount = asposecellscloud.Int32Ptr(3)
	rangeOperateTarget.FirstColumn = asposecellscloud.Int32Ptr(8)
	rangeOperateTarget.FirstRow = asposecellscloud.Int32Ptr(13)
	rangeOperateTarget.RowCount = asposecellscloud.Int32Ptr(2)
	rangeOperate := &models.RangeCopyRequest{}
	rangeOperate.Operate = "copydata"
	rangeOperate.Source = rangeOperateSource
	rangeOperate.Target = rangeOperateTarget
	request := requests.NewPostWorksheetCellsRangesCopyRequest(
		remoteName,
		rangeOperate,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangesCopy \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeMerge(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	range_ := &models.Range{}
	range_.ColumnCount = asposecellscloud.Int32Ptr(1)
	range_.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	range_.FirstRow = asposecellscloud.Int32Ptr(1)
	range_.RowCount = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetCellsRangeMergeRequest(
		remoteName,
		range_,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeUnMerge(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	range_ := &models.Range{}
	range_.ColumnCount = asposecellscloud.Int32Ptr(1)
	range_.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	range_.FirstRow = asposecellscloud.Int32Ptr(1)
	range_.RowCount = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetCellsRangeUnMergeRequest(
		remoteName,
		range_,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeUnMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	rangeOperateStyleFont := &models.Font{}
	rangeOperateStyleFont.Size = asposecellscloud.Int32Ptr(16)
	rangeOperateStyle := &models.Style{}
	rangeOperateStyle.Font = rangeOperateStyleFont
	rangeOperateRange := &models.Range{}
	rangeOperateRange.ColumnCount = asposecellscloud.Int32Ptr(1)
	rangeOperateRange.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	rangeOperateRange.FirstRow = asposecellscloud.Int32Ptr(1)
	rangeOperateRange.RowCount = asposecellscloud.Int32Ptr(10)
	rangeOperate := &models.RangeSetStyleRequest{}
	rangeOperate.Style = rangeOperateStyle
	rangeOperate.Range = rangeOperateRange
	request := requests.NewPostWorksheetCellsRangeStyleRequest(
		remoteName,
		rangeOperate,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_GetWorksheetCellsRangeValue(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetCellsRangeValueRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("namerange", "Name_2"),
		requests.WithCommonParameter("firstRow", intPtr(0)),
		requests.WithCommonParameter("firstColumn", intPtr(0)),
		requests.WithCommonParameter("rowCount", intPtr(3)),
		requests.WithCommonParameter("columnCount", intPtr(2)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_GetWorksheetCellsRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeValue(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	range_ := &models.Range{}
	range_.ColumnCount = asposecellscloud.Int32Ptr(1)
	range_.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	range_.FirstRow = asposecellscloud.Int32Ptr(1)
	range_.RowCount = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetCellsRangeValueRequest(
		remoteName,
		range_,
		"Sheet1",
		"100",
		requests.WithCommonParameter("isConverted", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("setStyle", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeMoveTo(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	range_ := &models.Range{}
	range_.ColumnCount = asposecellscloud.Int32Ptr(1)
	range_.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	range_.FirstRow = asposecellscloud.Int32Ptr(1)
	range_.RowCount = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetCellsRangeMoveToRequest(
		10,
		10,
		remoteName,
		range_,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeMoveTo \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeOutlineBorder(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	rangeOperateborderColor := &models.Color{}
	rangeOperateborderColor.R = []byte{48}
	rangeOperateborderColor.G = []byte{48}
	rangeOperateborderColor.B = []byte{48}
	rangeOperateRange := &models.Range{}
	rangeOperateRange.ColumnCount = asposecellscloud.Int32Ptr(1)
	rangeOperateRange.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	rangeOperateRange.FirstRow = asposecellscloud.Int32Ptr(1)
	rangeOperateRange.RowCount = asposecellscloud.Int32Ptr(10)
	rangeOperate := &models.RangeSetOutlineBorderRequest{}
	rangeOperate.BorderEdge = "LeftBorder"
	rangeOperate.BorderStyle = "Dotted"
	rangeOperate.BorderColor = rangeOperateborderColor
	rangeOperate.Range = rangeOperateRange
	request := requests.NewPostWorksheetCellsRangeOutlineBorderRequest(
		remoteName,
		rangeOperate,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeOutlineBorder \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeColumnWidth(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	range_ := &models.Range{}
	range_.ColumnCount = asposecellscloud.Int32Ptr(1)
	range_.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	range_.FirstRow = asposecellscloud.Int32Ptr(1)
	range_.RowCount = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetCellsRangeColumnWidthRequest(
		remoteName,
		range_,
		"Sheet1",
		10.7,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeColumnWidth \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeRowHeight(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	range_ := &models.Range{}
	range_.ColumnCount = asposecellscloud.Int32Ptr(1)
	range_.ColumnWidth = asposecellscloud.Float64Ptr(10.0)
	range_.FirstRow = asposecellscloud.Int32Ptr(1)
	range_.RowCount = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetCellsRangeRowHeightRequest(
		remoteName,
		range_,
		"Sheet1",
		10.9,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeRowHeight \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PutWorksheetCellsRange(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetCellsRangeRequest(
		remoteName,
		"A1:C6",
		"Sheet1",
		"Down",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PutWorksheetCellsRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_DeleteWorksheetCellsRange(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetCellsRangeRequest(
		remoteName,
		"A1:C6",
		"Sheet1",
		"Up",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_DeleteWorksheetCellsRange \n", GetBaseTest().GetTestNumber())
	}
}

func TestRangesController_PostWorksheetCellsRangeSort(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Group.xlsx"
	remoteName := "Group.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	rangeSortRequestDataSorter := &models.DataSorter{}
	rangeSortRequestDataSorter.CaseSensitive = asposecellscloud.BoolPtr(true)
	rangeSortRequestCellArea := &models.Range{}
	rangeSortRequestCellArea.ColumnCount = asposecellscloud.Int32Ptr(3)
	rangeSortRequestCellArea.FirstColumn = asposecellscloud.Int32Ptr(0)
	rangeSortRequestCellArea.FirstRow = asposecellscloud.Int32Ptr(0)
	rangeSortRequestCellArea.RowCount = asposecellscloud.Int32Ptr(15)
	rangeSortRequest := &models.RangeSortRequest{}
	rangeSortRequest.DataSorter = rangeSortRequestDataSorter
	rangeSortRequest.CellArea = rangeSortRequestCellArea
	request := requests.NewPostWorksheetCellsRangeSortRequest(
		remoteName,
		rangeSortRequest,
		"book1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestRangesController_PostWorksheetCellsRangeSort \n", GetBaseTest().GetTestNumber())
	}
}
