//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
	"asposecellscloud/converter"
	"asposecellscloud/datasource"
)

func main() {
	// If no environment variables are configured, obtain the ClientId and
	// ClientSecret from https://dashboard.aspose.cloud/#/applications and pass
	// them to NewAsposeCellsCloudClient explicitly:
	//   asposecellscloud.NewAsposeCellsCloudClient("YourClientId", "YourClientSecret", "")
	client := asposecellscloud.NewAsposeCellsCloudClient(
		os.Getenv("ProductClientId"),
		os.Getenv("ProductClientSecret"),
		os.Getenv("CellsCloudApiBaseUrl"),
	)

	// Convert a local spreadsheet to PDF and write the result to a local file.
	err := converter.Convert(
		context.Background(),
		client,
		datasource.FilePathSource("EmployeeSalesSummary.xlsx"),
		datasource.FilePathSink("EmployeeSalesSummary.pdf"),
		converter.FormatPDF,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("T")
}
