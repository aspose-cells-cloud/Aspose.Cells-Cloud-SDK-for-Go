package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestShapesController_GetWorksheetShapes(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetShapesRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_GetWorksheetShapes \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_GetWorksheetShape(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorksheetShapeRequest(
		remoteName,
		0,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_GetWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PutWorksheetShape(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	// ArcShape is one of the 22 subclasses of Shape; the parameter takes the family
	// (the generated models.ShapeLike). The shape itself is described by the query
	// parameters below, so the DTO only has to be a member of the family.
	shapeDTO := &models.ArcShape{}
	request := requests.NewPutWorksheetShapeRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("shapeDTO", shapeDTO),
		requests.WithCommonParameter("DrawingType", "arc"),
		requests.WithCommonParameter("upperLeftRow", intPtr(1)),
		requests.WithCommonParameter("upperLeftColumn", intPtr(1)),
		requests.WithCommonParameter("top", intPtr(10)),
		requests.WithCommonParameter("left", intPtr(10)),
		requests.WithCommonParameter("width", intPtr(100)),
		requests.WithCommonParameter("height", intPtr(100)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_PutWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_DeleteWorksheetShapes(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetShapesRequest(
		remoteName,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_DeleteWorksheetShapes \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_DeleteWorksheetShape(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorksheetShapeRequest(
		remoteName,
		0,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_DeleteWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PostWorksheetShape(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	dto := &models.Shape{}
	dto.LowerRightColumn = asposecellscloud.Int32Ptr(10)
	request := requests.NewPostWorksheetShapeRequest(
		dto,
		remoteName,
		0,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_PostWorksheetShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PostWorksheetGroupShape(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	var listShape = []interface{}{int64(0),
		int64(1)}
	request := requests.NewPostWorksheetGroupShapeRequest(
		listShape,
		remoteName,
		"Sheet6",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_PostWorksheetGroupShape \n", GetBaseTest().GetTestNumber())
	}
}

func TestShapesController_PostWorksheetUngroupShape(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorksheetUngroupShapeRequest(
		remoteName,
		0,
		"Sheet1",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestShapesController_PostWorksheetUngroupShape \n", GetBaseTest().GetTestNumber())
	}
}
