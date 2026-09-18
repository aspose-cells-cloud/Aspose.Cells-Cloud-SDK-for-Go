package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestBatchController_PostBatchConvert(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localBook1 := "Book1.xlsx"
	remoteBook1 := "Book1.xlsx"
	localMyDoc := "myDocument.xlsx"
	remoteMyDoc := "myDocument.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteBook1, GetBaseTest().localTestDataFolder+localBook1, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/"+remoteMyDoc, GetBaseTest().localTestDataFolder+localMyDoc, ""); err != nil {
		t.Fatal(err)
	}

	batchConvertRequestMatchCondition := &models.MatchConditionRequest{}
	batchConvertRequestMatchCondition.RegexPattern = "(^Book)(.+)(xlsx$)"
	batchConvertRequest := &models.BatchConvertRequest{}
	batchConvertRequest.SourceFolder = remoteFolder
	batchConvertRequest.Format = "pdf"
	batchConvertRequest.OutFolder = "OutResult"
	batchConvertRequest.MatchCondition = batchConvertRequestMatchCondition
	request := requests.NewPostBatchConvertRequest(
		batchConvertRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchConvert \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchProtect(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localBook1 := "Book1.xlsx"
	remoteBook1 := "Book1.xlsx"
	localMyDoc := "myDocument.xlsx"
	remoteMyDoc := "myDocument.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteBook1, GetBaseTest().localTestDataFolder+localBook1, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/"+remoteMyDoc, GetBaseTest().localTestDataFolder+localMyDoc, ""); err != nil {
		t.Fatal(err)
	}

	batchProtectRequestMatchCondition := &models.MatchConditionRequest{}
	batchProtectRequestMatchCondition.RegexPattern = "(^Book)(.+)(xlsx$)"
	batchProtectRequest := &models.BatchProtectRequest{}
	batchProtectRequest.SourceFolder = remoteFolder
	batchProtectRequest.ProtectionType = "All"
	batchProtectRequest.Password = "123456"
	batchProtectRequest.OutFolder = "OutResult"
	batchProtectRequest.MatchCondition = batchProtectRequestMatchCondition
	request := requests.NewPostBatchProtectRequest(
		batchProtectRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchProtect \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchLock(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localBook1 := "Book1.xlsx"
	remoteBook1 := "Book1.xlsx"
	localMyDoc := "myDocument.xlsx"
	remoteMyDoc := "myDocument.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteBook1, GetBaseTest().localTestDataFolder+localBook1, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/"+remoteMyDoc, GetBaseTest().localTestDataFolder+localMyDoc, ""); err != nil {
		t.Fatal(err)
	}

	batchLockRequestMatchCondition := &models.MatchConditionRequest{}
	batchLockRequestMatchCondition.RegexPattern = "(^Book)(.+)(xlsx$)"
	batchLockRequest := &models.BatchLockRequest{}
	batchLockRequest.SourceFolder = remoteFolder
	batchLockRequest.Password = "123456"
	batchLockRequest.OutFolder = "OutResult"
	batchLockRequest.MatchCondition = batchLockRequestMatchCondition
	request := requests.NewPostBatchLockRequest(
		batchLockRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchLock \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchUnlock(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localBook1 := "Book1.xlsx"
	remoteBook1 := "Book1.xlsx"
	localMyDoc := "myDocument.xlsx"
	remoteMyDoc := "myDocument.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteBook1, GetBaseTest().localTestDataFolder+localBook1, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/"+remoteMyDoc, GetBaseTest().localTestDataFolder+localMyDoc, ""); err != nil {
		t.Fatal(err)
	}

	batchLockRequestMatchCondition := &models.MatchConditionRequest{}
	batchLockRequestMatchCondition.RegexPattern = "(^Book)(.+)(xlsx$)"
	batchLockRequest := &models.BatchLockRequest{}
	batchLockRequest.SourceFolder = remoteFolder
	batchLockRequest.Password = "123456"
	batchLockRequest.OutFolder = "OutResult"
	batchLockRequest.MatchCondition = batchLockRequestMatchCondition
	request := requests.NewPostBatchUnlockRequest(
		batchLockRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchUnlock \n", GetBaseTest().GetTestNumber())
	}
}

func TestBatchController_PostBatchSplit(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localBook1 := "Book1.xlsx"
	remoteBook1 := "Book1.xlsx"
	localMyDoc := "myDocument.xlsx"
	remoteMyDoc := "myDocument.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteBook1, GetBaseTest().localTestDataFolder+localBook1, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/"+remoteMyDoc, GetBaseTest().localTestDataFolder+localMyDoc, ""); err != nil {
		t.Fatal(err)
	}

	batchSplitRequestMatchCondition := &models.MatchConditionRequest{}
	batchSplitRequestMatchCondition.RegexPattern = "(^Book)(.+)(xlsx$)"
	batchSplitRequest := &models.BatchSplitRequest{}
	batchSplitRequest.SourceFolder = remoteFolder
	batchSplitRequest.Format = "Pdf"
	batchSplitRequest.OutFolder = "OutResult"
	batchSplitRequest.MatchCondition = batchSplitRequestMatchCondition
	request := requests.NewPostBatchSplitRequest(
		batchSplitRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestBatchController_PostBatchSplit \n", GetBaseTest().GetTestNumber())
	}
}
