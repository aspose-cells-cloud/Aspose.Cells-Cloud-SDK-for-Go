package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestPicturesController_GetWorksheetPictures(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetPicturesRequest(
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPicturesController_GetWorksheetPictures \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_GetWorksheetPictureWithFormat(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetPictureWithFormatRequest(
		"png",
		remoteName,
		0,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPicturesController_GetWorksheetPictureWithFormat \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_PutWorksheetAddPicture(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	waterMarkPNG := "WaterMark.png"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/WaterMark.png", GetBaseTest().localTestDataFolder+waterMarkPNG, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetAddPictureRequest(
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("upperLeftRow", intPtr(1)),
		requests.WithCommonParameter("upperLeftColumn", intPtr(1)),
		requests.WithCommonParameter("lowerRightRow", intPtr(10)),
		requests.WithCommonParameter("lowerRightColumn", intPtr(10)),
		requests.WithCommonParameter("picturePath", remoteFolder+"/WaterMark.png"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPicturesController_PutWorksheetAddPicture \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_PostWorksheetPicture(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	picture := &models.Picture{}
	picture.Left = asposecellscloud.Int32Ptr(10)
	picture.Bottom = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetPictureRequest(
		remoteName,
		picture,
		0,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPicturesController_PostWorksheetPicture \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_DeleteWorksheetPicture(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetPictureRequest(
		remoteName,
		0,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPicturesController_DeleteWorksheetPicture \n", GetBaseTest().GetTestNumber())
	}
}

func TestPicturesController_DeleteWorksheetPictures(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetPicturesRequest(
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestPicturesController_DeleteWorksheetPictures \n", GetBaseTest().GetTestNumber())
	}
}
