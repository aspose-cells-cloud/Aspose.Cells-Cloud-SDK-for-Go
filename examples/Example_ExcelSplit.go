//go:build ignore

package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"

	"asposecellscloud"
	"asposecellscloud/dataprocessing"
	"asposecellscloud/datasource"
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

	// Split a local spreadsheet into one file per worksheet. The v4.0 endpoint
	// returns a zip archive containing the generated files.
	const zipPath = "splittedFiles.zip"
	err := dataprocessing.SplitSpreadsheet(ctx, client,
		datasource.FilePathSource("EmployeeSalesSummary.xlsx"),
		datasource.FilePathSink(zipPath),
		dataprocessing.WithOutFormat("xlsx"),
		dataprocessing.WithFrom(0),
		dataprocessing.WithTo(2),
	)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Successful split Excel files.")

	if err := unzip(zipPath); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

// unzip extracts every entry of the archive at zipPath into the current
// directory.
func unzip(zipPath string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}

		content, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return err
		}

		if err := os.WriteFile(file.Name, content, file.Mode()); err != nil {
			return err
		}
	}
	return nil
}
