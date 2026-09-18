package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestConversion_WorkbookSaveAs_csv_OutResultPostExcelSaveAscsv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "csv"
	newfilename := "OutResult/PostExcelSaveAs.csv"
	saveOptionsData := &models.SaveOptionsData{}
	saveOptionsData.Filename = newfilename
	request := requests.NewSaveSpreadsheetAsRequest(
		format,
		remoteName,
		requests.WithCommonParameter("saveOptionsData", saveOptionsData),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_csv_OutResultPostExcelSaveAscsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_WorkbookSaveAs_pdf_OutResultPostExcelSaveAspdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pdf"
	newfilename := "OutResult/PostExcelSaveAs.pdf"
	saveOptionsData := &models.SaveOptionsData{}
	saveOptionsData.Filename = newfilename
	request := requests.NewSaveSpreadsheetAsRequest(
		format,
		remoteName,
		requests.WithCommonParameter("saveOptionsData", saveOptionsData),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_WorkbookSaveAs_pdf_OutResultPostExcelSaveAspdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_png(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "png"
	request := requests.NewConvertSpreadsheetRequest(
		format,
		"TestData/"+localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbook_sql(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "sql"
	request := requests.NewConvertSpreadsheetRequest(
		format,
		"TestData/"+localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbook_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertSpreadsheetToPdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	request := requests.NewConvertSpreadsheetToPdfRequest(
		"TestData/" + localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertSpreadsheetToPdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertSpreadsheetToCsv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	request := requests.NewConvertSpreadsheetToCsvRequest(
		"TestData/" + localName,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertSpreadsheetToCsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorkbookSaveCloud_csv_OutResultConvertWorkbookcsv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "csv"
	outPath := "OutResult/ConvertWorkbook.csv"
	request := requests.NewConvertSpreadsheetRequest(
		format,
		"TestData/"+localName,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorkbookSaveCloud_csv_OutResultConvertWorkbookcsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorksheetToSvg(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertWorksheetToImageRequest(
		"svg",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorksheetToSvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorksheetToPng(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertWorksheetToImageRequest(
		"png",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorksheetToPng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorksheetToTiff(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertWorksheetToImageRequest(
		"tiff",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorksheetToTiff \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertWorksheetToPdf(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertWorksheetToPdfRequest(
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertWorksheetToPdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertTableToSvg(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertTableToImageRequest(
		"svg",
		"TestData/"+localName,
		"Table13",
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertTableToSvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertTableToPng(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertTableToImageRequest(
		"png",
		"TestData/"+localName,
		"Table13",
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertTableToPng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertTableToPdf(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertTableToPdfRequest(
		"TestData/"+localName,
		"Table13",
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertTableToPdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertTableToCsv(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertTableToCsvRequest(
		"TestData/"+localName,
		"Table13",
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertTableToCsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertTableToHtml(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertTableToHtmlRequest(
		"TestData/"+localName,
		"Table13",
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertTableToHtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertTableToJson(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertTableToJsonRequest(
		"TestData/"+localName,
		"Table13",
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertTableToJson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertRangeToImage(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertRangeToImageRequest(
		"svg",
		"B2:F10",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertRangeToImage \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertRangeToPdf(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertRangeToPdfRequest(
		"A1:F10",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertRangeToPdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertRangeToCsv(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertRangeToCsvRequest(
		"A1:F10",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertRangeToCsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertRangeToHtml(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertRangeToHtmlRequest(
		"A1:F10",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertRangeToHtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertRangeToJson(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	request := requests.NewConvertRangeToJsonRequest(
		"A1:F10",
		"TestData/"+localName,
		"Sheet2",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertRangeToJson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertChartToImage_svg(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	format := "svg"
	request := requests.NewConvertChartToImageRequest(
		0,
		format,
		"TestData/"+localName,
		"Sales",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertChartToImage_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ConvertChartToPdf(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	request := requests.NewConvertChartToPdfRequest(
		0,
		"TestData/"+localName,
		"Sales",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ConvertChartToPdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportSpreadsheetAsFormat_pdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pdf"
	request := requests.NewExportSpreadsheetAsFormatRequest(
		format,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportSpreadsheetAsFormat_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportSpreadsheetAsFormat_pptx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pptx"
	request := requests.NewExportSpreadsheetAsFormatRequest(
		format,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportSpreadsheetAsFormat_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportSpreadsheetAsFormat_json(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "json"
	request := requests.NewExportSpreadsheetAsFormatRequest(
		format,
		remoteName,
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportSpreadsheetAsFormat_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportWorksheetAsFormat(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	remoteName := "EmployeeSalesSummary.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	request := requests.NewExportWorksheetAsFormatRequest(
		"svg",
		localName,
		"Sales",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportWorksheetAsFormat \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportChartAsFormat_svg(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	remoteName := "EmployeeSalesSummary.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "svg"
	request := requests.NewExportChartAsFormatRequest(
		0,
		format,
		localName,
		"Sales",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportChartAsFormat_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportTableAsFormat_svg(t *testing.T) {
	ctx := context.Background()
	localName := "TestTables.xlsx"
	remoteName := "TestTables.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "svg"
	request := requests.NewExportTableAsFormatRequest(
		format,
		localName,
		"Table13",
		"Sheet2",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportTableAsFormat_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion_ExportRangeAsFormat_svg(t *testing.T) {
	ctx := context.Background()
	localName := "EmployeeSalesSummary.xlsx"
	remoteName := "EmployeeSalesSummary.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "svg"
	request := requests.NewExportRangeAsFormatRequest(
		format,
		localName,
		"A1:F16",
		"Sales",
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion_ExportRangeAsFormat_svg \n", GetBaseTest().GetTestNumber())
	}
}
