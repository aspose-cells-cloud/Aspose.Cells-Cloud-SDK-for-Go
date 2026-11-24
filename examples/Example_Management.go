//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/editor"
	"asposecellscloud/requests"
)

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
		inputPath  = "EmployeeSalesSummary.xlsx"
		outputPath = "EmployeeSalesSummary_Output.xlsx"
	)

	// Add a new worksheet to a local Excel file.
	must(editor.AddWorksheet(ctx, client,
		datasource.FilePathSource(inputPath),
		datasource.FilePathSink(outputPath),
		"Sheet1",
		editor.WithPosition(0),
	))

	// Rename a worksheet in the Excel file.
	must(editor.RenameWorksheet(ctx, client,
		datasource.FilePathSource(outputPath),
		datasource.FilePathSink(outputPath),
		"Sheet1",
		"NewSheet",
	))

	// Move a worksheet to a new position in the Excel file.
	must(editor.MoveWorksheet(ctx, client,
		datasource.FilePathSource(outputPath),
		datasource.FilePathSink(outputPath),
		"NewSheet",
		2,
	))

	// Delete a worksheet from the Excel file.
	must(editor.DeleteWorksheet(ctx, client,
		datasource.FilePathSource(outputPath),
		datasource.FilePathSink(outputPath),
		"NewSheet",
	))

	// Compress the Excel file.
	data, err := os.ReadFile(outputPath)
	must(err)

	compressRequest := requests.NewCompressSpreadsheetRequest(9, "")
	compressRequest.SetSpreadsheetBytes(data, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, compressRequest)
	must(err)
	must(os.WriteFile(outputPath, resp.Body, 0644))
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
