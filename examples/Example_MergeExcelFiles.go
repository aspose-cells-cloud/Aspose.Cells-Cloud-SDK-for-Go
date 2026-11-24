//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
	"asposecellscloud/dataprocessing"
	"asposecellscloud/datasource"
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
		employeeSalesSummary = "EmployeeSalesSummary.xlsx"
		companySales         = "CompanySales.xlsx"
	)

	uploadFile(ctx, client, companySales, remoteFolder+"/"+companySales)
	uploadFile(ctx, client, employeeSalesSummary, remoteFolder+"/"+employeeSalesSummary)

	// Merge a spreadsheet file into another cloud workbook in place.
	_, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewPostWorkbooksMergeRequest(
			remoteFolder+"/"+employeeSalesSummary,
			companySales,
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	fmt.Println("Successful merge Excel files.")

	// Merge the worksheets of a local spreadsheet into a single consolidated
	// workbook. The v4.0 local merge endpoint merges the worksheets of one
	// workbook; the legacy map-of-local-files form is no longer supported.
	must(dataprocessing.MergeSpreadsheet(ctx, client,
		datasource.FilePathSource(companySales),
		datasource.FilePathSink("MergedExcel.xlsx"),
		dataprocessing.WithOutFormat("xlsx"),
		dataprocessing.WithMergeInOneSheet(true),
	))
}

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
