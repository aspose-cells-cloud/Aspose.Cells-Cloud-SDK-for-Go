package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestConversion30_WorkbookSaveAs_csv_OutResultPostExcelSaveAscsv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "csv"
	newfilename := "OutResult/PostExcelSaveAs.csv"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_csv_OutResultPostExcelSaveAscsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_html_OutResultPostExcelSaveAshtml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "html"
	newfilename := "OutResult/PostExcelSaveAs.html"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_html_OutResultPostExcelSaveAshtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_mhtml_OutResultPostExcelSaveAsmhtml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "mhtml"
	newfilename := "OutResult/PostExcelSaveAs.mhtml"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_mhtml_OutResultPostExcelSaveAsmhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_ods_OutResultPostExcelSaveAsods(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "ods"
	newfilename := "OutResult/PostExcelSaveAs.ods"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_ods_OutResultPostExcelSaveAsods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_pdf_OutResultPostExcelSaveAspdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pdf"
	newfilename := "OutResult/PostExcelSaveAs.pdf"
	// PdfSaveOptions is a member of the SaveOptions family, two levels below the base
	// through PaginatedSaveOptions; the parameter accepts any member of the family.
	saveOptions := &models.PdfSaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_pdf_OutResultPostExcelSaveAspdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_xml_OutResultPostExcelSaveAsxml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "xml"
	newfilename := "OutResult/PostExcelSaveAs.xml"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_xml_OutResultPostExcelSaveAsxml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_txt_OutResultPostExcelSaveAstxt(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "txt"
	newfilename := "OutResult/PostExcelSaveAs.txt"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_txt_OutResultPostExcelSaveAstxt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_tif_OutResultPostExcelSaveAstif(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "tif"
	newfilename := "OutResult/PostExcelSaveAs.tif"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_tif_OutResultPostExcelSaveAstif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_xlsb_OutResultPostExcelSaveAsxlsb(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "xlsb"
	newfilename := "OutResult/PostExcelSaveAs.xlsb"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_xlsb_OutResultPostExcelSaveAsxlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_xps_OutResultPostExcelSaveAsxps(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "xps"
	newfilename := "OutResult/PostExcelSaveAs.xps"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_xps_OutResultPostExcelSaveAsxps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_png_OutResultPostExcelSaveAspng(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "png"
	newfilename := "OutResult/PostExcelSaveAs.png"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_png_OutResultPostExcelSaveAspng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_md_OutResultPostExcelSaveAsmd(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "md"
	newfilename := "OutResult/PostExcelSaveAs.md"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_md_OutResultPostExcelSaveAsmd \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_svg_OutResultPostExcelSaveAssvg(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "svg"
	newfilename := "OutResult/PostExcelSaveAs.svg"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_svg_OutResultPostExcelSaveAssvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_docx_OutResultPostExcelSaveAsdocx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "docx"
	newfilename := "OutResult/PostExcelSaveAs.docx"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_docx_OutResultPostExcelSaveAsdocx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_pptx_OutResultPostExcelSaveAspptx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pptx"
	newfilename := "OutResult/PostExcelSaveAs.pptx"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_pptx_OutResultPostExcelSaveAspptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_json_OutResultPostExcelSaveAsjson(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "json"
	newfilename := "OutResult/PostExcelSaveAs.json"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_json_OutResultPostExcelSaveAsjson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_WorkbookSaveAs_sql_OutResultPostExcelSaveAssql(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "sql"
	newfilename := "OutResult/PostExcelSaveAs.sql"
	saveOptions := &models.SaveOptions{}
	saveOptions.SaveFormat = format
	request := requests.NewPostWorkbookSaveAsRequest(
		remoteName,
		newfilename,
		requests.WithCommonParameter("saveOptions", saveOptions),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_WorkbookSaveAs_sql_OutResultPostExcelSaveAssql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_csv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "csv"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_html(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "html"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_mhtml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "mhtml"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_ods(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "ods"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_pdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pdf"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_xml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "xml"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_txt(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "txt"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_tif(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "tif"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_xps(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "xps"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_png(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "png"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_md(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "md"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_svg(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "svg"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_docx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "docx"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_pptx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "pptx"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_json(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "json"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_GetWorkbookFormat_sql(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	remoteName := "Book1.xlsx"
	remoteFolder := "TestData/In"
	if err := mustUploadFile(t, remoteFolder+"/"+remoteName, GetBaseTest().localTestDataFolder+localName, ""); err != nil {
		t.Fatal(err)
	}

	format := "sql"
	request := requests.NewGetWorkbookRequest(
		remoteName,
		requests.WithCommonParameter("format", format),
		requests.WithCommonParameter("folder", remoteFolder),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_GetWorkbookFormat_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_csv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "csv"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_xls(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xls"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_xls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_html(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "html"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_mhtml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "mhtml"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_mhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_ods(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "ods"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_pdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "pdf"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_xml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xml"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_xml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_txt(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "txt"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_txt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_tif(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "tif"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_tif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_xlsb(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xlsb"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_xlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_xps(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xps"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_png(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "png"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_md(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "md"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_svg(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "svg"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_docx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "docx"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_pptx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "pptx"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_json(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "json"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_json \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbook_sql(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "sql"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbook_sql \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_csv_OutResultConvertWorkbookcsv(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "csv"
	outPath := "OutResult/ConvertWorkbook.csv"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_csv_OutResultConvertWorkbookcsv \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_xls_OutResultConvertWorkbookxls(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xls"
	outPath := "OutResult/ConvertWorkbook.xls"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_xls_OutResultConvertWorkbookxls \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_html_OutResultConvertWorkbookhtml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "html"
	outPath := "OutResult/ConvertWorkbook.html"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_html_OutResultConvertWorkbookhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_mhtml_OutResultConvertWorkbookmhtml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "mhtml"
	outPath := "OutResult/ConvertWorkbook.mhtml"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_mhtml_OutResultConvertWorkbookmhtml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_ods_OutResultConvertWorkbookods(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "ods"
	outPath := "OutResult/ConvertWorkbook.ods"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_ods_OutResultConvertWorkbookods \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_pdf_OutResultConvertWorkbookpdf(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "pdf"
	outPath := "OutResult/ConvertWorkbook.pdf"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_pdf_OutResultConvertWorkbookpdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_xml_OutResultConvertWorkbookxml(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xml"
	outPath := "OutResult/ConvertWorkbook.xml"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_xml_OutResultConvertWorkbookxml \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_txt_OutResultConvertWorkbooktxt(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "txt"
	outPath := "OutResult/ConvertWorkbook.txt"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_txt_OutResultConvertWorkbooktxt \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_tif_OutResultConvertWorkbooktif(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "tif"
	outPath := "OutResult/ConvertWorkbook.tif"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_tif_OutResultConvertWorkbooktif \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_xlsb_OutResultConvertWorkbookxlsb(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xlsb"
	outPath := "OutResult/ConvertWorkbook.xlsb"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_xlsb_OutResultConvertWorkbookxlsb \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_xltm_OutResultConvertWorkbookxltm(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xltm"
	outPath := "OutResult/ConvertWorkbook.xltm"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_xltm_OutResultConvertWorkbookxltm \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_xps_OutResultConvertWorkbookxps(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "xps"
	outPath := "OutResult/ConvertWorkbook.xps"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_xps_OutResultConvertWorkbookxps \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_png_OutResultConvertWorkbookpng(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "png"
	outPath := "OutResult/ConvertWorkbook.png"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_png_OutResultConvertWorkbookpng \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_md_OutResultConvertWorkbookmd(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "md"
	outPath := "OutResult/ConvertWorkbook.md"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_md_OutResultConvertWorkbookmd \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_svg_OutResultConvertWorkbooksvg(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "svg"
	outPath := "OutResult/ConvertWorkbook.svg"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_svg_OutResultConvertWorkbooksvg \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_docx_OutResultConvertWorkbookdocx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "docx"
	outPath := "OutResult/ConvertWorkbook.docx"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_docx_OutResultConvertWorkbookdocx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_pptx_OutResultConvertWorkbookpptx(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "pptx"
	outPath := "OutResult/ConvertWorkbook.pptx"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_pptx_OutResultConvertWorkbookpptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_json_OutResultConvertWorkbookjson(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "json"
	outPath := "OutResult/ConvertWorkbook.json"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_json_OutResultConvertWorkbookjson \n", GetBaseTest().GetTestNumber())
	}
}

func TestConversion30_ConvertWorkbookSaveCloud_sql_OutResultConvertWorkbooksql(t *testing.T) {
	ctx := context.Background()
	localName := "Book1.xlsx"
	format := "sql"
	outPath := "OutResult/ConvertWorkbook.sql"
	request := requests.NewPutConvertWorkbookRequest(
		GetBaseTest().localTestDataFolder+localName,
		format,
		requests.WithCommonParameter("outPath", outPath),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestConversion30_ConvertWorkbookSaveCloud_sql_OutResultConvertWorkbooksql \n", GetBaseTest().GetTestNumber())
	}
}
