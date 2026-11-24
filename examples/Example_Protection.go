//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
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
		lockedPath = "EmployeeSalesSummary_Lock.xlsx"
		openPath   = "EmployeeSalesSummary_unlock.xlsx"
	)

	// Protect an Excel file.
	input, err := os.ReadFile(inputPath)
	must(err)

	protectRequest := requests.NewProtectSpreadsheetRequest("123456", "123456", "")
	protectRequest.SetSpreadsheetBytes(input, "Spreadsheet")

	resp, err := asposecellscloud.DoChecked(ctx, client, protectRequest)
	must(err)
	must(os.WriteFile(lockedPath, resp.Body, 0644))

	// Unprotect an Excel file.
	locked, err := os.ReadFile(lockedPath)
	must(err)

	unprotectRequest := requests.NewUnprotectSpreadsheetRequest("123456", "123456", "")
	unprotectRequest.SetSpreadsheetBytes(locked, "Spreadsheet")

	resp, err = asposecellscloud.DoChecked(ctx, client, unprotectRequest)
	must(err)
	must(os.WriteFile(openPath, resp.Body, 0644))
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
