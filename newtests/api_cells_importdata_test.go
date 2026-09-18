package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/requests"
)

func TestImportData_ImportDataIntoSpreadsheet(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	csVDataFile := "BookCsvDuplicateData.csv"
	request := requests.NewImportDataIntoSpreadsheetRequest(
		"TestData/"+csVDataFile,
		"TestData/"+book1Xlsx,
		"E3",
		"Sheet1",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestImportData_ImportDataIntoSpreadsheet \n", GetBaseTest().GetTestNumber())
	}
}
