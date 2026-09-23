package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestAutoFilterController_GetWorksheetAutoFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetAutoFilterRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)
	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_GetWorksheetAutoFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetDateFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetDateFilterRequest(
		"Year",
		0,
		remoteName,
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("year", intPtr(1920)),
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetDateFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetFilterRequest(
		"Year",
		0,
		remoteName,
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetIconFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetIconFilterRequest(
		0,
		1,
		"ArrowsGray3",
		remoteName,
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetIconFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetCustomFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetCustomFilterRequest(
		"1",
		0,
		remoteName,
		"LessOrEqual",
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetCustomFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetDynamicFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetDynamicFilterRequest(
		"BelowAverage",
		0,
		remoteName,
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetDynamicFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetFilterTop10(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetFilterTop10Request(
		0,
		true,
		true,
		1,
		remoteName,
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(false)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetFilterTop10 \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PutWorksheetColorFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	colorFilterForegroundColorColor := &models.Color{}
	colorFilterForegroundColorColor.R = asposecellscloud.Int32Ptr(48)
	colorFilterForegroundColorColor.G = asposecellscloud.Int32Ptr(48)
	colorFilterForegroundColorColor.B = asposecellscloud.Int32Ptr(48)
	colorFilterForegroundColor := &models.CellsColor{}
	colorFilterForegroundColor.Type = "Automatic"
	colorFilterForegroundColor.Color = colorFilterForegroundColorColor
	colorFilter := &models.ColorFilterRequest{}
	colorFilter.Pattern = "Solid"
	colorFilter.ForegroundColor = colorFilterForegroundColor
	request := requests.NewPutWorksheetColorFilterRequest(
		colorFilter,
		0,
		remoteName,
		"A1:B1",
		"Sheet1",
		requests.WithCommonParameter("matchBlanks", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("refresh", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PutWorksheetColorFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PostWorksheetMatchBlanks(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetMatchBlanksRequest(
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
		fmt.Printf("%d\tTestAutoFilterController_PostWorksheetMatchBlanks \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PostWorksheetMatchNonBlanks(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetMatchNonBlanksRequest(
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
		fmt.Printf("%d\tTestAutoFilterController_PostWorksheetMatchNonBlanks \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_PostWorksheetAutoFilterRefresh(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetAutoFilterRefreshRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_PostWorksheetAutoFilterRefresh \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_DeleteWorksheetDateFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetDateFilterRequest(
		"Year",
		0,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("year", intPtr(1920)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_DeleteWorksheetDateFilter \n", GetBaseTest().GetTestNumber())
	}
}

func TestAutoFilterController_DeleteWorksheetFilter(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetFilterRequest(
		0,
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("criteria", "year"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestAutoFilterController_DeleteWorksheetFilter \n", GetBaseTest().GetTestNumber())
	}
}
