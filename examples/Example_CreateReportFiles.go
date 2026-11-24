//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

const remoteFolder = "GoSDK"

func main() {
	ctx := context.Background()

	// If no environment variables are configured, obtain the ClientId and
	// ClientSecret from https://dashboard.aspose.cloud/#/applications and pass
	// them to NewAsposeCellsCloudClient explicitly:
	//   asposecellscloud.NewAsposeCellsCloudClient("YourClientId", "YourClientSecret", "")
	client := asposecellscloud.NewAsposeCellsCloudClient(
		os.Getenv("ProductClientId"),
		os.Getenv("ProductClientSecret"),
		os.Getenv("CellsCloudApiBaseUrl"),
	)

	const (
		reportTemplate  = "ProductSalteReportTemplate.xlsx"
		reportName      = "ProductSaleReport.xlsx"
		sheetName       = "SalesData"
		reportSheetName = "Sales Report"
		reportDataFile  = "ReportData.xml"
	)

	// Create the template file.
	_, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewPutWorkbookCreateRequest(reportTemplate,
			requests.WithCommonParameter("folder", remoteFolder),
			requests.WithCommonParameter("isWriteOver", asposecellscloud.BoolPtr(true)),
		))
	must(err)
	fmt.Println("Create template file.")

	// Add Sheet1.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPutAddNewWorksheetRequest(reportTemplate, "Sheet1",
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	fmt.Println("Add Sheet1.")

	// Upload the report data file.
	uploadFile(ctx, client, reportDataFile, remoteFolder+"/"+reportDataFile)
	fmt.Println("Update data file.")

	// Fill the smart marker cells.
	setCell := func(cellName, value string) {
		_, err := asposecellscloud.DoChecked(ctx, client,
			requests.NewPostWorksheetCellSetValueRequest(cellName, reportTemplate, "Sheet1",
				requests.WithCommonParameter("folder", remoteFolder),
				requests.WithCommonParameter("value", value),
				requests.WithCommonParameter("type", "string"),
			))
		must(err)
		fmt.Printf("Set %s Value.\n", cellName)
	}
	setCell("A1", "Product Category")
	setCell("B1", "Sales")
	setCell("A2", "&=ProductSaleRecord.ProductCategory")
	setCell("B2", "&=ProductSaleRecord.Sales")

	// Generate the report from the smart markers.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPostWorkbookGetSmartMarkerResultRequest(reportTemplate,
			requests.WithCommonParameter("folder", remoteFolder),
			requests.WithCommonParameter("xmlFile", remoteFolder+"/"+reportDataFile),
			requests.WithCommonParameter("outPath", remoteFolder+"/"+reportName),
		))
	must(err)
	fmt.Println("Smart Marker Successful.")

	// Apply a number format to the data range.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPostWorksheetCellsRangeStyleRequest(reportName,
			&models.RangeSetStyleRequest{
				Range: &models.Range{
					FirstColumn: asposecellscloud.Int32Ptr(1),
					ColumnCount: asposecellscloud.Int32Ptr(1),
					FirstRow:    asposecellscloud.Int32Ptr(1),
					RowCount:    asposecellscloud.Int32Ptr(4),
				},
				Style: &models.Style{Number: asposecellscloud.Int32Ptr(8)},
			},
			"Sheet1",
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	fmt.Println("Set range style.")

	// Add a list object for the data range.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPutWorksheetListObjectRequest(reportName, "Sheet1",
			requests.WithCommonParameter("folder", remoteFolder),
			requests.WithCommonParameter("startRow", intPtr(0)),
			requests.WithCommonParameter("startColumn", intPtr(0)),
			requests.WithCommonParameter("endColumn", intPtr(1)),
			requests.WithCommonParameter("endRow", intPtr(4)),
			requests.WithCommonParameter("hasHeaders", asposecellscloud.BoolPtr(true)),
		))
	must(err)
	fmt.Println("Add a list object.")

	// Rename Sheet1 to SalesData.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPostRenameWorksheetRequest(reportName, sheetName, "Sheet1",
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	fmt.Println("Create simple report.")

	// Add the report sheet.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPutAddNewWorksheetRequest(reportName, reportSheetName,
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	fmt.Println("Add Sale Report Sheet.")

	// Hide the data sheet.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPutChangeVisibilityWorksheetRequest(false, reportName, sheetName,
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	fmt.Println("Set Sheet visibility.")

	// Add a chart in the report sheet.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPutWorksheetChartRequest("Pie", reportName, reportSheetName,
			requests.WithCommonParameter("folder", remoteFolder),
			requests.WithCommonParameter("isVertical", asposecellscloud.BoolPtr(true)),
			requests.WithCommonParameter("title", "Aspose Chart"),
			requests.WithCommonParameter("upperLeftRow", intPtr(5)),
			requests.WithCommonParameter("upperLeftColumn", intPtr(3)),
			requests.WithCommonParameter("lowerRightColumn", intPtr(9)),
			requests.WithCommonParameter("lowerRightRow", intPtr(15)),
			requests.WithCommonParameter("area", sheetName+"!A1:B5"),
		))
	must(err)
	fmt.Println("Add Chart in Sheet.")
}

// intPtr returns a pointer to v. Generated request parameters model optional
// integers as *int so that "unset" differs from zero.
func intPtr(v int) *int { return &v }

// uploadFile uploads a local file to the given path in cloud storage.
func uploadFile(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, localPath, remotePath string) {
	data, err := os.ReadFile(localPath)
	must(err)

	req := requests.NewUploadFileRequest(remotePath, "")
	req.SetUploadFilesBytes(data, "UploadFiles")

	_, err = asposecellscloud.DoChecked(ctx, client, req)
	must(err)
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
