package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestChartAreaController_GetChartArea(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetChartAreaRequest(
		0,
		remoteName,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestChartAreaController_GetChartArea \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartAreaController_GetChartAreaFillFormat(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetChartAreaFillFormatRequest(
		0,
		remoteName,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestChartAreaController_GetChartAreaFillFormat \n", GetBaseTest().GetTestNumber())
	}
}

func TestChartAreaController_GetChartAreaBorder(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetChartAreaBorderRequest(
		0,
		remoteName,
		"Sheet4",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestChartAreaController_GetChartAreaBorder \n", GetBaseTest().GetTestNumber())
	}
}
