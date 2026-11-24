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

const remoteFolder = "SDKGo"

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
		companySales         = "CompanySales.xlsx"
		employeeSalesSummary = "EmployeeSalesSummary.xlsx"
	)

	// Check if the folder exists, and create it when it does not.
	resp, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewObjectExistsRequest(remoteFolder))
	must(err)

	var exists models.ObjectExist
	must(resp.GetJSON(&exists))
	if exists.Exists == nil || !*exists.Exists {
		_, err = asposecellscloud.DoChecked(ctx, client,
			requests.NewCreateFolderRequest(remoteFolder))
		must(err)
	}

	// Upload source files.
	uploadFile(ctx, client, companySales, remoteFolder+"/"+companySales)
	uploadFile(ctx, client, employeeSalesSummary, remoteFolder+"/"+employeeSalesSummary)

	// Create a folder on cloud storage.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewCreateFolderRequest(remoteFolder+"/CellsCloud"))
	must(err)

	// Move a file to another folder on cloud storage.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewMoveFileRequest(remoteFolder+"/CellsCloud/"+companySales, remoteFolder+"/"+companySales))
	must(err)

	// Copy a file to another file on cloud storage.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewCopyFileRequest(remoteFolder+"/CellsCloud/"+employeeSalesSummary, remoteFolder+"/"+employeeSalesSummary))
	must(err)

	// Move a folder on cloud storage.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewMoveFolderRequest(remoteFolder+"/CellsCloud2", remoteFolder+"/CellsCloud"))
	must(err)

	// Get the files list in a folder.
	resp, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewGetFilesListRequest(
			requests.WithCommonParameter("path", remoteFolder+"/CellsCloud2"),
		))
	must(err)

	var fileList models.FilesList
	must(resp.GetJSON(&fileList))
	for index, storageFile := range fileList.Value {
		fmt.Printf("Get File %d %s\n", index, storageFile.Name)
	}

	// Delete the folder on cloud storage.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewDeleteFolderRequest(remoteFolder+"/CellsCloud2"))
	must(err)

	// Delete the file on cloud storage.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewDeleteFileRequest(remoteFolder+"/"+companySales))
	must(err)
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
