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

	// Split a spreadsheet in cloud storage and write the generated files to a
	// cloud output folder.
	uploadFile(ctx, client, employeeSalesSummary, remoteFolder+"/"+employeeSalesSummary)

	_, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewCreateFolderRequest(remoteFolder+"/Output"))
	must(err)

	wf := &asposecellscloud.WorkbookRef{Name: employeeSalesSummary, Folder: remoteFolder}
	_, err = dataprocessing.SplitRemoteSpreadsheet(ctx, client, wf, remoteFolder+"/Output",
		dataprocessing.WithOutFormat("pdf"),
		dataprocessing.WithTo(3),
	)
	must(err)
	fmt.Println("Successful split Excel files.")

	// Split a local Excel file to PNG files.
	splitLocal(ctx, client, companySales, "png", "CompanySales_png.zip", 3)
	// Split a local Excel file to PDF files.
	splitLocal(ctx, client, companySales, "pdf", "CompanySales_pdf.zip", 3)
}

// splitLocal splits a local spreadsheet into one file per worksheet, saves the
// returned zip archive and extracts every generated file.
func splitLocal(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, input, format, zipPath string, to int) {
	must(dataprocessing.SplitSpreadsheet(ctx, client,
		datasource.FilePathSource(input),
		datasource.FilePathSink(zipPath),
		dataprocessing.WithOutFormat(format),
		dataprocessing.WithTo(to),
	))
	fmt.Printf("Successful split %s. Begin to write every file.\n", input)

	must(unzip(zipPath))
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
