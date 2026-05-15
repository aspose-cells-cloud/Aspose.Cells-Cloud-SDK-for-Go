package main

import (
	"fmt"
	"os"

	. "github.com/aspose-cells-cloud/aspose-cells-cloud-go/v26"
)

func main() {
	// If no environment variables are configured, please obtain the ClientId and ClientSecret from https://dashboard.aspose.cloud/#/applications and replace the following values:
	// instance  := NewCellsApiService('YourClientId','YourClientSecret')
	instance := NewCellsApiService(os.Getenv("ProductClientId"), os.Getenv("ProductClientSecret"))
	reportTemplate := "ProductSalteReportTemplate.xlsx"
	reportName := "ProductSaleReport.xlsx"
	sheetName := "SalesData"
	reportSheetName := "Sales Report"
	reportDataFile := "ReportData.xml"
	remoteFolder := "GoSDK"
	_, httpResponse, err := instance.PutWorkbookCreate(&PutWorkbookCreateRequest{Name: reportTemplate, Folder: remoteFolder, IsWriteOver: true})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Create template file.\n")
	}
	_, httpResponse, err = instance.PutAddNewWorksheet(&PutAddNewWorksheetRequest{Name: reportTemplate, Folder: remoteFolder, SheetName: "Sheet1"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Add Sheet1.\n")
	}
	_, httpResponse, err = instance.UploadFile(&UploadFileRequest{UploadFiles: reportDataFile, Path: remoteFolder + "/" + reportDataFile})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Update data file.\n")
	}
	_, httpResponse, err = instance.PostWorksheetCellSetValue(&PostWorksheetCellSetValueRequest{Name: reportTemplate, Folder: remoteFolder, SheetName: "Sheet1", CellName: "A1", Value: "Product Category", Type_: "string"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Set A1 Value.\n")
	}
	_, httpResponse, err = instance.PostWorksheetCellSetValue(&PostWorksheetCellSetValueRequest{Name: reportTemplate, Folder: remoteFolder, SheetName: "Sheet1", CellName: "B1", Value: "Sales", Type_: "string"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Set B1 Value.\n")
	}
	_, httpResponse, err = instance.PostWorksheetCellSetValue(&PostWorksheetCellSetValueRequest{Name: reportTemplate, Folder: remoteFolder, SheetName: "Sheet1", CellName: "A2", Value: "&=ProductSaleRecord.ProductCategory", Type_: "string"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Set A2 Value.\n")
	}
	_, httpResponse, err = instance.PostWorksheetCellSetValue(&PostWorksheetCellSetValueRequest{Name: reportTemplate, Folder: remoteFolder, SheetName: "Sheet1", CellName: "B2", Value: "&=ProductSaleRecord.Sales", Type_: "string"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Set B2 Value.\n")
	}
	_, httpResponse, err = instance.PostWorkbookGetSmartMarkerResult(&PostWorkbookGetSmartMarkerResultRequest{Name: reportTemplate, Folder: remoteFolder, XmlFile: remoteFolder + "/" + reportDataFile, OutPath: remoteFolder + "/" + reportName})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Smart Marker Successful.\n")
	}
	_, httpResponse, err = instance.PostWorksheetCellsRangeStyle(&PostWorksheetCellsRangeStyleRequest{Name: reportName, Folder: remoteFolder, SheetName: "Sheet1", RangeOperate: &RangeSetStyleRequest{Range_: &Range{FirstColumn: 1, ColumnCount: 1, FirstRow: 1, RowCount: 4}, Style: &Style{Number: 8}}})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Set range style.\n")
	}
	_, httpResponse, err = instance.PutWorksheetListObject(&PutWorksheetListObjectRequest{Name: reportName, Folder: remoteFolder, SheetName: "Sheet1", StartRow: 0, StartColumn: 0, EndColumn: 1, EndRow: 4, HasHeaders: true})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Add a list object.\n")
	}
	_, httpResponse, err = instance.PostRenameWorksheet(&PostRenameWorksheetRequest{Name: reportName, Folder: remoteFolder, SheetName: "Sheet1", Newname: sheetName})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Create simple report.\n")
	}
	_, httpResponse, err = instance.PutAddNewWorksheet(&PutAddNewWorksheetRequest{Name: reportName, Folder: remoteFolder, SheetName: reportSheetName})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Add Sale Report Sheet.\n")
	}
	_, httpResponse, err = instance.PutChangeVisibilityWorksheet(&PutChangeVisibilityWorksheetRequest{Name: reportName, Folder: remoteFolder, SheetName: sheetName, IsVisible: false})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Set Sheet visibility.\n")
	}

	_, httpResponse, err = instance.PutWorksheetChart(&PutWorksheetChartRequest{
		Name: reportName, Folder: remoteFolder, SheetName: reportSheetName, ChartType: "Pie", IsVertical: true,
		Title:        "Aspose Chart",
		UpperLeftRow: 5, UpperLeftColumn: 3, LowerRightColumn: 9, LowerRightRow: 15, Area: sheetName + "!A1:B5"})
	if err != nil {
		fmt.Print(err)
	} else if httpResponse.StatusCode < 200 || httpResponse.StatusCode > 299 {
		fmt.Print(httpResponse.StatusCode)
	} else {
		print("Add Chart in Sheet.\n")
	}
}
