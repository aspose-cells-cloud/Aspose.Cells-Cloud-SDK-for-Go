//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	"asposecellscloud"
	"asposecellscloud/converter"
	"asposecellscloud/datasource"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

const remoteFolder = "GoSDK"

func main() {
	ctx := context.Background()
	input := "EmployeeSalesSummary.xlsx"

	// If no environment variables are configured, obtain the ClientId and
	// ClientSecret from https://dashboard.aspose.cloud/#/applications and pass
	// them to NewAsposeCellsCloudClient explicitly:
	//   asposecellscloud.NewAsposeCellsCloudClient("YourClientId", "YourClientSecret", "")
	client := asposecellscloud.NewAsposeCellsCloudClient(
		os.Getenv("ProductClientId"),
		os.Getenv("ProductClientSecret"),
		os.Getenv("CellsCloudApiBaseUrl"),
	)

	// -------------------------------------------------------------------------
	// Local file processing
	// -------------------------------------------------------------------------

	// Convert a local Excel file to another format file directly and save it
	// to a local file. This completes in one step.
	must(converter.Convert(ctx, client,
		datasource.FilePathSource(input),
		datasource.FilePathSink("EmployeeSalesSummary.pdf"),
		converter.FormatPDF,
	))

	// The same conversion completed in two steps: convert a local Excel file
	// to byte data, then save the bytes yourself.
	sink := &datasource.BytesSink{}
	must(converter.Convert(ctx, client,
		datasource.FilePathSource(input),
		sink,
		converter.FormatPDF,
	))
	must(os.WriteFile("EmployeeSalesSummary1.pdf", sink.Bytes(), 0644))

	// -------------------------------------------------------------------------
	// Cloud file processing
	// -------------------------------------------------------------------------

	// 1. Upload a local Excel file to Cells Cloud Storage.
	uploadFile(ctx, client, input, remoteFolder+"/"+input)

	// 2.1 Save an Excel file of Cells Cloud as another format file of Cells
	// Cloud.
	_, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewSaveSpreadsheetAsRequest("pdf", input,
			requests.WithCommonParameter("folder", remoteFolder),
			requests.WithCommonParameter("saveOptionsData", &models.SaveOptionsData{Filename: "EmployeeSalesSummaryXlsx.pdf"}),
		))
	must(err)

	wf := &asposecellscloud.WorkbookRef{Name: input, Folder: remoteFolder}

	// 2.2 Export a worksheet of a cloud Excel file to another format file
	// directly.
	must(converter.Worksheet(ctx, client, wf, "Sales",
		datasource.FilePathSink("EmployeeSalesSummary_Sales.pdf"),
		converter.FormatPDF,
	))

	// 2.3 Export a worksheet of a cloud Excel file as PNG.
	must(converter.Worksheet(ctx, client, wf, "Sales",
		datasource.FilePathSink("EmployeeSalesSummary_Sale_1.png"),
		converter.FormatPNG,
	))

	// 2.4 Export a specified cells area (B5:L36) of a worksheet as PNG.
	resp, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewExportRangeAsFormatRequest("png", input, "B5:L36", "Sales",
			requests.WithCommonParameter("folder", remoteFolder),
		))
	must(err)
	must(os.WriteFile("EmployeeSalesSummary_Sale_Area.png", resp.Body, 0644))

	// Get the conversion result: export a cloud workbook as PDF.
	must(converter.Workbook(ctx, client, wf,
		datasource.FilePathSink("EmployeeSalesSummary2.pdf"),
		converter.FormatPDF,
	))

	// Convert a local Excel file directly to the other supported formats. The
	// result is written to a local file in each case.
	localConversions := []struct {
		outPath string
		format  string
	}{
		{"EmployeeSalesSummary3.pdf", converter.FormatPDF},
		{"EmployeeSalesSummary.json", converter.FormatJSON},
		{"EmployeeSalesSummary.docx", converter.FormatDocx},
		{"EmployeeSalesSummary.png", converter.FormatPNG},
		{"EmployeeSalesSummary.pptx", converter.FormatPptx},
		{"EmployeeSalesSummary.html", converter.FormatHTML},
	}
	for _, c := range localConversions {
		must(converter.Convert(ctx, client,
			datasource.FilePathSource(input),
			datasource.FilePathSink(c.outPath),
			c.format,
		))
	}
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
