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

const (
	remoteFolder       = "GoSDK"
	remoteOutputFolder = "GoSDK/Output"
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

	uploadFile(ctx, client, "CompanySales.xlsx", remoteFolder+"/CompanySales.xlsx")
	uploadFile(ctx, client, "EmployeeSalesSummary.xlsx", remoteFolder+"/EmployeeSalesSummary.xlsx")

	// Batch convert Excel files.
	_, err := asposecellscloud.DoChecked(ctx, client,
		requests.NewPostBatchConvertRequest(&models.BatchConvertRequest{
			SourceFolder:   remoteFolder,
			Format:         "Pdf",
			OutFolder:      remoteOutputFolder,
			MatchCondition: matchXlsx(),
		}))
	report("Successful batch conversion of Excel files.", err)

	// Batch lock Excel files.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPostBatchLockRequest(&models.BatchLockRequest{
			SourceFolder:   remoteFolder,
			Password:       "123456",
			OutFolder:      remoteOutputFolder,
			MatchCondition: matchXlsx(),
		}))
	report("Successful batch lock of Excel files.", err)

	// Batch unlock Excel files.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPostBatchUnlockRequest(&models.BatchLockRequest{
			SourceFolder:   remoteFolder,
			Password:       "123456",
			OutFolder:      remoteOutputFolder,
			MatchCondition: matchXlsx(),
		}))
	report("Successful batch unlock of Excel files.", err)

	// Batch split Excel files.
	_, err = asposecellscloud.DoChecked(ctx, client,
		requests.NewPostBatchSplitRequest(&models.BatchSplitRequest{
			SourceFolder:   remoteFolder,
			Format:         "pdf",
			OutFolder:      remoteOutputFolder,
			MatchCondition: matchXlsx(),
		}))
	report("Successful batch split of Excel files.", err)
}

func matchXlsx() *models.MatchConditionRequest {
	return &models.MatchConditionRequest{RegexPattern: "xlsx"}
}

// uploadFile uploads a local file to the given path in cloud storage.
func uploadFile(ctx context.Context, client *asposecellscloud.AsposeCellsCloudClient, localPath, remotePath string) {
	data, err := os.ReadFile(localPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	req := requests.NewUploadFileRequest(remotePath, "")
	req.SetUploadFilesBytes(data, "UploadFiles")

	if _, err = asposecellscloud.DoChecked(ctx, client, req); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func report(message string, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(message)
}
