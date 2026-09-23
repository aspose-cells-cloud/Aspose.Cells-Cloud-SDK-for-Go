package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestDataProcessingController_PostWorkbookDataCleansing(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookCsvDuplicateData.csv"
	remoteName := "BookCsvDuplicateData.csv"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	dataCleansingDataFillDataFillDefaultValue := &models.DataFillValue{}
	dataCleansingDataFillDataFillDefaultValue.DefaultDate = "2024-01-01"
	dataCleansingDataFillDataFillDefaultValue.DefaultNumber = asposecellscloud.Int32Ptr(0)
	dataCleansingDataFillDataFillDefaultValue.DefaultBoolean = asposecellscloud.BoolPtr(false)
	dataCleansingDataFill := &models.DataFill{}
	dataCleansingDataFill.DataFillDefaultValue = dataCleansingDataFillDataFillDefaultValue
	dataCleansing := &models.DataCleansing{}
	dataCleansing.NeedFillData = asposecellscloud.BoolPtr(true)
	dataCleansing.DataFill = dataCleansingDataFill
	request := requests.NewPostWorkbookDataCleansingRequest(
		dataCleansing,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataCleansing \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostWorkbookDataDeduplication(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookCsvDuplicateData.csv"
	remoteName := "BookCsvDuplicateData.csv"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	var deduplicationRegionRanges = []models.Range{}
	deduplicationRegion := &models.DeduplicationRegion{}
	deduplicationRegion.Ranges = deduplicationRegionRanges
	request := requests.NewPostWorkbookDataDeduplicationRequest(
		deduplicationRegion,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataDeduplication \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostWorkbookDataFill(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookCsvDuplicateData.csv"
	remoteName := "BookCsvDuplicateData.csv"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	dataFillDataFillDefaultValue := &models.DataFillValue{}
	dataFillDataFillDefaultValue.DefaultDate = "2024-01-01"
	dataFillDataFillDefaultValue.DefaultNumber = asposecellscloud.Int32Ptr(0)
	dataFillDataFillDefaultValue.DefaultBoolean = asposecellscloud.BoolPtr(false)
	dataFill := &models.DataFill{}
	dataFill.DataFillDefaultValue = dataFillDataFillDefaultValue
	request := requests.NewPostWorkbookDataFillRequest(
		dataFill,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostWorkbookDataFill \n", GetBaseTest().GetTestNumber())
	}
}

func TestDataProcessingController_PostDataTransformation(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "BookTableL2W.xlsx"
	remoteName := "BookTableL2W.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	dataTransformationRequestLoadDataLoadTo := &models.LoadTo{}
	dataTransformationRequestLoadDataLoadTo.BeginColumnIndex = asposecellscloud.Int32Ptr(2)
	dataTransformationRequestLoadDataLoadTo.BeginRowIndex = asposecellscloud.Int32Ptr(3)
	dataTransformationRequestLoadDataLoadTo.Worksheet = "L2W"
	dataTransformationRequestLoadDataDataQueryDataItem := &models.DataItem{}
	dataTransformationRequestLoadDataDataQueryDataItem.DataItemType = "Table"
	dataTransformationRequestLoadDataDataQueryDataItem.Value = "Table1"
	dataTransformationRequestLoadDataDataQueryDataSource := &models.DataSource{}
	dataTransformationRequestLoadDataDataQueryDataSource.DataSourceType = "CloudFileSystem"
	dataTransformationRequestLoadDataDataQueryDataSource.DataPath = "TestData/In/BookTableL2W.xlsx"
	dataTransformationRequestLoadDataDataQuery := &models.DataQuery{}
	dataTransformationRequestLoadDataDataQuery.Name = "DataQuery"
	dataTransformationRequestLoadDataDataQuery.DataItem = dataTransformationRequestLoadDataDataQueryDataItem
	dataTransformationRequestLoadDataDataQuery.DataSource = dataTransformationRequestLoadDataDataQueryDataSource
	dataTransformationRequestLoadDataDataQuery.DataSourceDataType = "ListObject"
	dataTransformationRequestLoadData := &models.LoadData{}
	dataTransformationRequestLoadData.LoadTo = dataTransformationRequestLoadDataLoadTo
	dataTransformationRequestLoadData.DataQuery = dataTransformationRequestLoadDataDataQuery
	var dataTransformationRequestAppliedStepsAppliedStep0AppliedOperateUnpivotColumnNames = []string{"2017",
		"2018",
		"2019"}
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate := &models.UnpivotColumn{}
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.ValueMapName = "Count"
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.ColumnMapName = "Date"
	dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate.UnpivotColumnNames = dataTransformationRequestAppliedStepsAppliedStep0AppliedOperateUnpivotColumnNames
	dataTransformationRequestAppliedStepsAppliedStep0 := &models.AppliedStep{}
	dataTransformationRequestAppliedStepsAppliedStep0.StepName = "UnpivotColumn"
	// UnpivotColumn satisfies the AppliedOperate interface, so the concrete operation
	// is handed to the step as-is. Its AppliedOperateType discriminator -- the member
	// the service's JSON converter dispatches on -- is filled in from the concrete
	// type when the request is serialized.
	dataTransformationRequestAppliedStepsAppliedStep0.AppliedOperate = dataTransformationRequestAppliedStepsAppliedStep0AppliedOperate
	var dataTransformationRequestAppliedSteps = []models.AppliedStep{*dataTransformationRequestAppliedStepsAppliedStep0}
	dataTransformationRequest := &models.DataTransformationRequest{}
	dataTransformationRequest.LoadData = dataTransformationRequestLoadData
	dataTransformationRequest.AppliedSteps = dataTransformationRequestAppliedSteps
	request := requests.NewPostDataTransformationRequest(
		dataTransformationRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestDataProcessingController_PostDataTransformation \n", GetBaseTest().GetTestNumber())
	}
}
