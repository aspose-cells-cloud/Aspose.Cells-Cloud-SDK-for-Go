package unittest_test

import (
	"os"
	"testing"

	asposecellscloud "asposecellscloud"
)

// RealClient returns a client configured with credentials from environment variables.
// It fails the test if credentials are not available.
func RealClient(t *testing.T) *asposecellscloud.AsposeCellsCloudClient {
	t.Helper()

	clientID := os.Getenv("CellsCloudClientId")
	clientSecret := os.Getenv("CellsCloudClientSecret")
	baseURL := os.Getenv("CellsCloudApiBaseUrl")

	if clientID == "" {
		t.Fatal("CellsCloudClientId environment variable is required")
	}
	if clientSecret == "" {
		t.Fatal("CellsCloudClientSecret environment variable is required")
	}

	if baseURL == "" {
		baseURL = "https://api.aspose.cloud"
	}

	return asposecellscloud.NewAsposeCellsCloudClient(clientID, clientSecret, baseURL)
}

// SkipUnlessRealTest returns true if real API tests should run.
// Set CellsCloudRealTests=1 to enable real API tests.
func SkipUnlessRealTest(t *testing.T) bool {
	t.Helper()
	if os.Getenv("CellsCloudRealTests") != "1" {
		t.Skip("Skipping real API test. Set CellsCloudRealTests=1 to run")
		return false
	}
	return true
}
