package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestSparklineGroupsController_GetWorksheetSparklineGroups(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetSparklineGroupsRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_GetWorksheetSparklineGroups \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_GetWorksheetSparklineGroup(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetSparklineGroupRequest(
		remoteName,
		"Sheet1",
		0,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_GetWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_DeleteWorksheetSparklineGroups(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetSparklineGroupsRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_DeleteWorksheetSparklineGroups \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_DeleteWorksheetSparklineGroup(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetSparklineGroupRequest(
		remoteName,
		"Sheet1",
		0,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_DeleteWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_PutWorksheetSparklineGroup(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetSparklineGroupRequest(
		"C6:E13",
		false,
		"G6:G13",
		remoteName,
		"Sheet1",
		"Line",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_PutWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}

func TestSparklineGroupsController_PostWorksheetSparklineGroup(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "TestCase.xlsx"
	remoteName := "TestCase.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	sparklineGroup := &models.SparklineGroup{}
	sparklineGroup.DisplayHidden = asposecellscloud.BoolPtr(true)
	sparklineGroup.PlotRightToLeft = asposecellscloud.BoolPtr(true)
	request := requests.NewPostWorksheetSparklineGroupRequest(
		remoteName,
		"Sheet1",
		sparklineGroup,
		0,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestSparklineGroupsController_PostWorksheetSparklineGroup \n", GetBaseTest().GetTestNumber())
	}
}
