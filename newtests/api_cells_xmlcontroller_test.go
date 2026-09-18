package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestXmlController_PostWorkbookExportXML(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Template.xlsx"
	remoteName := "Template.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorkbookExportXMLRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestXmlController_PostWorkbookExportXML \n", GetBaseTest().GetTestNumber())
	}
}

func TestXmlController_PostWorkbookImportXML(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Template.xlsx"
	dataXML := "data.xml"
	remoteName := "Template.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/data.xml", GetBaseTest().localTestDataFolder+dataXML, ""); err != nil {
		t.Fatal(err)
	}

	importXMLRequestXMLFileSource := &models.DataSource{}
	importXMLRequestXMLFileSource.DataSourceType = "CloudFileSystem"
	importXMLRequestXMLFileSource.DataPath = remoteFolder + "/data.xml"
	importXMLRequestImportPosition := &models.ImportPosition{}
	importXMLRequestImportPosition.SheetName = "Sheet1"
	importXMLRequestImportPosition.RowIndex = asposecellscloud.Int32Ptr(3)
	importXMLRequestImportPosition.ColumnIndex = asposecellscloud.Int32Ptr(4)
	importXMLRequest := &models.ImportXMLRequest{}
	importXMLRequest.XMLFileSource = importXMLRequestXMLFileSource
	importXMLRequest.ImportPosition = importXMLRequestImportPosition
	request := requests.NewPostWorkbookImportXMLRequest(
		importXMLRequest,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestXmlController_PostWorkbookImportXML \n", GetBaseTest().GetTestNumber())
	}
}
