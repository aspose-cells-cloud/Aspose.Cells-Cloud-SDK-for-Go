package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestOleObjectsController_GetWorksheetOleObjects(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetOleObjectsRequest(
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestOleObjectsController_GetWorksheetOleObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_GetWorksheetOleObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetOleObjectRequest(
		remoteName,
		0,
		"Sheet6",
		requests.WithCommonParameter("format", "png"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestOleObjectsController_GetWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_DeleteWorksheetOleObjects(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetOleObjectsRequest(
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestOleObjectsController_DeleteWorksheetOleObjects \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_DeleteWorksheetOleObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetOleObjectRequest(
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
		fmt.Printf("%d\tTestOleObjectsController_DeleteWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_PostUpdateWorksheetOleObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	ole := &models.OleObject{}
	ole.Left = asposecellscloud.Int32Ptr(10)
	ole.Right = asposecellscloud.Int32Ptr(10)
	ole.Height = asposecellscloud.Int32Ptr(90)
	ole.Width = asposecellscloud.Int32Ptr(78)
	request := requests.NewPostUpdateWorksheetOleObjectRequest(
		remoteName,
		ole,
		0,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestOleObjectsController_PostUpdateWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}

func TestOleObjectsController_PutWorksheetOleObject(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	oLEDoc := "OLEDoc.docx"
	wordJPG := "word.jpg"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, "OLEDoc.docx", GetBaseTest().localTestDataFolder+oLEDoc, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, "word.jpg", GetBaseTest().localTestDataFolder+wordJPG, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorksheetOleObjectRequest(
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("upperLeftRow", intPtr(1)),
		requests.WithCommonParameter("upperLeftColumn", intPtr(1)),
		requests.WithCommonParameter("height", intPtr(100)),
		requests.WithCommonParameter("width", intPtr(80)),
		requests.WithCommonParameter("oleFile", "OLEDoc.docx"),
		requests.WithCommonParameter("imageFile", "word.jpg"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestOleObjectsController_PutWorksheetOleObject \n", GetBaseTest().GetTestNumber())
	}
}
