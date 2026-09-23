package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestWorkbookController_PostDigitalSignature(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	roywangPFX := "roywang.pfx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/roywang.pfx", GetBaseTest().localTestDataFolder+roywangPFX, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostDigitalSignatureRequest(
		remoteFolder+"/roywang.pfx",
		remoteName,
		"123456",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostDigitalSignature \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostEncryptWorkbook(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	encryption := &models.WorkbookEncryptionRequest{}
	encryption.Password = "123456"
	encryption.EncryptionType = "XOR"
	encryption.KeyLength = asposecellscloud.Int32Ptr(128)
	request := requests.NewPostEncryptWorkbookRequest(
		encryption,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostEncryptWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteDecryptWorkbook(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	encryption := &models.WorkbookEncryptionRequest{}
	encryption.Password = "123456"
	encryption.EncryptionType = "XOR"
	encryption.KeyLength = asposecellscloud.Int32Ptr(128)
	request := requests.NewDeleteDecryptWorkbookRequest(
		encryption,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteDecryptWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostProtectWorkbook(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	protectWorkbookRequest := &models.ProtectWorkbookRequest{}
	protectWorkbookRequest.EncryptWithPassword = "123456"
	protectWorkbookRequest.ProtectWorkbookStructure = "ALL"
	request := requests.NewPostProtectWorkbookRequest(
		remoteName,
		protectWorkbookRequest,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostProtectWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteUnProtectWorkbook(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteUnProtectWorkbookRequest(
		remoteName,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteUnProtectWorkbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookDefaultStyle(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorkbookDefaultStyleRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookDefaultStyle \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookTextItems(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorkbookTextItemsRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookTextItems \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookNames(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorkbookNamesRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookNames \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookName(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	newName := &models.Name{}
	newName.Text = "name_1804"
	newName.Comment = "KeepSourceFormatting"
	newName.RefersTo = "=Sheet1!$I$4"
	request := requests.NewPutWorkbookNameRequest(
		remoteName,
		newName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookName(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorkbookNameRequest(
		remoteName,
		"Name_2",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookName(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	newName := &models.Name{}
	newName.Text = "name_1804"
	newName.Comment = "KeepSourceFormatting"
	newName.RefersTo = "=Sheet1!$I$4"
	request := requests.NewPostWorkbookNameRequest(
		remoteName,
		"Name_2",
		newName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookNameValue(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorkbookNameValueRequest(
		remoteName,
		"Name_2",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookNameValue \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteWorkbookNames(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorkbookNamesRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteWorkbookNames \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteWorkbookName(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorkbookNameRequest(
		remoteName,
		"Name_2",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteWorkbookName \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutDocumentProtectFromChanges(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	password := &models.PasswordRequest{}
	password.Password = "123456"
	request := requests.NewPutDocumentProtectFromChangesRequest(
		remoteName,
		password,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutDocumentProtectFromChanges \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteDocumentUnProtectFromChanges(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteDocumentUnProtectFromChangesRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteDocumentUnProtectFromChanges \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbooksMerge(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	myDocumentXLSX := "myDocument.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/myDocument.xlsx", GetBaseTest().localTestDataFolder+myDocumentXLSX, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorkbooksMergeRequest(
		remoteFolder+"/myDocument.xlsx",
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
		requests.WithCommonParameter("mergedStorageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbooksMerge \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbooksTextSearch(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorkbooksTextSearchRequest(
		remoteName,
		"1234",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbooksTextSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookTextReplace(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorkbookTextReplaceRequest(
		remoteName,
		"5678",
		"1234",
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookTextReplace \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookGetSmartMarkerResult(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	reportDataXML := "ReportData.xml"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/ReportData.xml", GetBaseTest().localTestDataFolder+reportDataXML, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorkbookGetSmartMarkerResultRequest(
		remoteName,
		requests.WithCommonParameter("xmlFile", remoteFolder+"/ReportData.xml"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("outPath", "OutResult/SmartMarkerResult.xlsx"),
		requests.WithCommonParameter("storageName", ""),
		requests.WithCommonParameter("outStorageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookGetSmartMarkerResult \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookCreate(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	reportDataXML := "ReportData.xml"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	if err := mustUploadFile(t, remoteFolder+"/ReportData.xml", GetBaseTest().localTestDataFolder+reportDataXML, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPutWorkbookCreateRequest(
		"PutWorkbookCreate.xlsx",
		requests.WithCommonParameter("templateFile", remoteFolder+"/"+remoteName),
		requests.WithCommonParameter("dataFile", remoteFolder+"/ReportData.xml"),
		requests.WithCommonParameter("isWriteOver", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
		requests.WithCommonParameter("checkExcelRestriction", asposecellscloud.BoolPtr(true)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookCreate \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookSplit(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostWorkbookSplitRequest(
		remoteName,
		requests.WithCommonParameter("format", "png"),
		requests.WithCommonParameter("outFolder", "OutResult"),
		requests.WithCommonParameter("from", intPtr(1)),
		requests.WithCommonParameter("to", intPtr(5)),
		requests.WithCommonParameter("horizontalResolution", intPtr(96)),
		requests.WithCommonParameter("verticalResolution", intPtr(96)),
		requests.WithCommonParameter("splitNameRule", "sheetname"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
		requests.WithCommonParameter("outStorageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookSplit \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostImportData(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	var importOptionData = []interface{}{int64(1),
		int64(2),
		int64(3),
		int64(4)}
	// ImportIntArrayOption is one of the nine subclasses ImportOption has in the
	// specification, and the request parameter accepts the whole family (the
	// generated models.ImportOptionLike), so the concrete option is handed over
	// as-is. ImportDataType names the payload here: unlike the abstract
	// AppliedOperate family, the service has to be told which subclass this is.
	importOption := &models.ImportIntArrayOption{}
	importOption.DestinationWorksheet = "Sheet1"
	importOption.FirstColumn = asposecellscloud.Int32Ptr(1)
	importOption.FirstRow = asposecellscloud.Int32Ptr(3)
	importOption.ImportDataType = "IntArray"
	importOption.IsInsert = asposecellscloud.BoolPtr(true)
	importOption.IsVertical = asposecellscloud.BoolPtr(true)
	importOption.Data = importOptionData
	request := requests.NewPostImportDataRequest(
		remoteName,
		requests.WithCommonParameter("importOption", importOption),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostImportData \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookCalculateFormula(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	options := &models.CalculationOptions{}
	options.IgnoreError = asposecellscloud.BoolPtr(true)
	options.Recursive = asposecellscloud.BoolPtr(true)
	request := requests.NewPostWorkbookCalculateFormulaRequest(
		remoteName,
		requests.WithCommonParameter("options", options),
		requests.WithCommonParameter("ignoreError", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookCalculateFormula \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostAutofitWorkbookRows(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostAutofitWorkbookRowsRequest(
		remoteName,
		requests.WithCommonParameter("startRow", intPtr(1)),
		requests.WithCommonParameter("endRow", intPtr(100)),
		requests.WithCommonParameter("onlyAuto", asposecellscloud.BoolPtr(true)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostAutofitWorkbookRows \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostAutofitWorkbookColumns(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewPostAutofitWorkbookColumnsRequest(
		remoteName,
		requests.WithCommonParameter("startColumn", intPtr(1)),
		requests.WithCommonParameter("endColumn", intPtr(20)),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostAutofitWorkbookColumns \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetWorkbookSettings(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetWorkbookSettingsRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetWorkbookSettings \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PostWorkbookSettings(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	settings := &models.WorkbookSettings{}
	settings.AutoCompressPictures = asposecellscloud.BoolPtr(true)
	settings.HidePivotFieldList = asposecellscloud.BoolPtr(true)
	request := requests.NewPostWorkbookSettingsRequest(
		remoteName,
		settings,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PostWorkbookSettings \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookBackground(t *testing.T) {
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

	request := requests.NewPutWorkbookBackgroundRequest(
		remoteName,
		requests.WithCommonParameter("picPath", remoteFolder+"/WaterMark.png"),
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookBackground \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_DeleteWorkbookBackground(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewDeleteWorkbookBackgroundRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_DeleteWorkbookBackground \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_PutWorkbookWaterMarker(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	textWaterMarkerRequest := &models.TextWaterMarkerRequest{}
	textWaterMarkerRequest.Text = "Aspose Cells Cloud"
	textWaterMarkerRequest.FontSize = asposecellscloud.Int32Ptr(12)
	request := requests.NewPutWorkbookWaterMarkerRequest(
		remoteName,
		textWaterMarkerRequest,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_PutWorkbookWaterMarker \n", GetBaseTest().GetTestNumber())
	}
}

func TestWorkbookController_GetPageCount(t *testing.T) {
	ctx := context.Background()
	remoteFolder := "TestData/In"
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewGetPageCountRequest(
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
		requests.WithCommonParameter("storageName", ""),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestWorkbookController_GetPageCount \n", GetBaseTest().GetTestNumber())
	}
}
