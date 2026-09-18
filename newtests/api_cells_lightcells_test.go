package newtests

import (
	"context"
	"fmt"
	"testing"

	"asposecellscloud"
	"asposecellscloud/models"
	"asposecellscloud/requests"
)

func TestLightCells_PostSplit_csv(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "csv"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_html(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "html"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_html \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_ods(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "ods"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_ods \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_pdf(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "pdf"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_xps(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "xps"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_xps \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_md(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "md"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_md \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_svg(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "svg"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_svg \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_docx(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "docx"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_docx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_pptx(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "pptx"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_pptx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSplit_json(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	outFormat := "json"
	request := requests.NewPostSplitRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		outFormat,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSplit_json \n", GetBaseTest().GetTestNumber())
	}
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_csv(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_html(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_ods(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_pdf(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_md(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_svg(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_docx(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_pptx(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostAssemble_json(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

func TestLightCells_PostExport_csv_workbook(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "csv"
	objectType := "workbook"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_csv_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_html_workbook(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "html"
	objectType := "workbook"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_html_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_workbook(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "pdf"
	objectType := "workbook"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pptx_workbook(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "pptx"
	objectType := "workbook"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pptx_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_json_workbook(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "json"
	objectType := "workbook"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_json_workbook \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_html_worksheet(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "html"
	objectType := "worksheet"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_html_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_docx_worksheet(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "docx"
	objectType := "worksheet"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_docx_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_json_worksheet(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "json"
	objectType := "worksheet"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_json_worksheet \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_chart(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "pdf"
	objectType := "chart"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_chart(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "png"
	objectType := "chart"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_png_picture(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "png"
	objectType := "picture"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_png_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_csv_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "csv"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_csv_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_ods_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "ods"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_ods_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pdf_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "pdf"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pdf_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_md_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "md"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_md_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_docx_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "docx"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_docx_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_pptx_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "pptx"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_pptx_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostExport_json_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "json"
	objectType := "listobject"
	request := requests.NewPostExportRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("objectType", objectType),
		requests.WithCommonParameter("format", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostExport_json_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostCompress_50(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	compressLevel := 50
	request := requests.NewPostCompressRequest(
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		requests.WithCommonParameter("CompressLevel", intPtr(compressLevel)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostCompress_50 \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostCompress_90(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	compressLevel := 90
	request := requests.NewPostCompressRequest(
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		requests.WithCommonParameter("CompressLevel", intPtr(compressLevel)),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostCompress_90 \n", GetBaseTest().GetTestNumber())
	}
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostMerge_html_true(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostMerge_pdf_true(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostMerge_xlsx_true(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

// PostAssemble and PostMerge upload several workbooks at once: the legacy
// request took File map[string]string and sent every entry as its own
// multipart part. The v4.0 request exposes a single File string, so the multi-
// file assembly cannot be expressed.
func TestLightCells_PostMerge_json_false(t *testing.T) {
	t.Skip("multi-file upload is not expressible in the v4.0 request model")
}

func TestLightCells_PostUnlock(t *testing.T) {
	ctx := context.Background()
	needUnlockXlsx := "needUnlock.xlsx"
	request := requests.NewPostUnlockRequest(
		GetBaseTest().localTestDataFolder+needUnlockXlsx,
		"123456",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostUnlock \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostLock(t *testing.T) {
	ctx := context.Background()
	needlockXlsx := "needlock.xlsx"
	request := requests.NewPostLockRequest(
		GetBaseTest().localTestDataFolder+needlockXlsx,
		"123456",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostLock \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostProtect(t *testing.T) {
	ctx := context.Background()
	assemblyTestXlsx := "assemblytest.xlsx"
	protectWorkbookRequest := &models.ProtectWorkbookRequest{}
	protectWorkbookRequest.AwaysOpenReadOnly = asposecellscloud.BoolPtr(true)
	protectWorkbookRequest.EncryptWithPassword = "123456"
	request := requests.NewPostProtectRequest(
		GetBaseTest().localTestDataFolder+assemblyTestXlsx,
		protectWorkbookRequest,
		requests.WithCommonParameter("password", "123456"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostProtect \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostProtect_ProtectWorkbookRequest(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	protectWorkbookRequest := &models.ProtectWorkbookRequest{}
	protectWorkbookRequest.AwaysOpenReadOnly = asposecellscloud.BoolPtr(true)
	protectWorkbookRequest.EncryptWithPassword = "123456"
	request := requests.NewPostProtectRequest(
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		protectWorkbookRequest,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostProtect_ProtectWorkbookRequest \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostSearch(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	request := requests.NewPostSearchRequest(
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		"12",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostSearch \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReplace(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	request := requests.NewPostReplaceRequest(
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		"newtext",
		"12",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReplace \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReplaceOnlySheetname(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	request := requests.NewPostReplaceRequest(
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		"newtext",
		"12",
		requests.WithCommonParameter("sheetname", "Sheet1"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReplaceOnlySheetname \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostWatermark(t *testing.T) {
	ctx := context.Background()
	dataSourceXlsx := "datasource.xlsx"
	request := requests.NewPostWatermarkRequest(
		"#773322",
		GetBaseTest().localTestDataFolder+dataSourceXlsx,
		"aspose.cells cloud sdk",
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostWatermark \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_chart(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "chart"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_chart \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_comment(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "comment"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_comment \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_picture(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "picture"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_picture \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_shape(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "shape"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_shape \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_listobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "listobject"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_listobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_hyperlink(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "hyperlink"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_hyperlink \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_oleobject(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "oleobject"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_oleobject \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_pivottable(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "pivottable"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_pivottable \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_validation(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "validation"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_validation \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostClearObjects_Background(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	objecttype := "Background"
	request := requests.NewPostClearObjectsRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		objecttype,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostClearObjects_Background \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostRepair_xlsx(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "xlsx"
	request := requests.NewPostRepairRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostRepair_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostRepair_pdf(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	format := "pdf"
	request := requests.NewPostRepairRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostRepair_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_rows_pdf(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	rotateType := "rows"
	format := "pdf"
	request := requests.NewPostReverseRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		rotateType,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_rows_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_cols_pdf(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	rotateType := "cols"
	format := "pdf"
	request := requests.NewPostReverseRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		rotateType,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_cols_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_both_pdf(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	rotateType := "both"
	format := "pdf"
	request := requests.NewPostReverseRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		rotateType,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_both_pdf \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_rows_csv(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	rotateType := "rows"
	format := "csv"
	request := requests.NewPostReverseRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		rotateType,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_rows_csv \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_cols_png(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	rotateType := "cols"
	format := "png"
	request := requests.NewPostReverseRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		rotateType,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_cols_png \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostReverse_both_xlsx(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	rotateType := "both"
	format := "xlsx"
	request := requests.NewPostReverseRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		rotateType,
		requests.WithCommonParameter("outFormat", format),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostReverse_both_xlsx \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_GetMetadata(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	request := requests.NewGetMetadataRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("type", "all"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_GetMetadata \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_DeleteMetadata(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	request := requests.NewDeleteMetadataRequest(
		GetBaseTest().localTestDataFolder+book1Xlsx,
		requests.WithCommonParameter("type", "all"),
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_DeleteMetadata \n", GetBaseTest().GetTestNumber())
	}
}

func TestLightCells_PostMetadata(t *testing.T) {
	ctx := context.Background()
	book1Xlsx := "Book1.xlsx"
	cellsDocumentscellsDocument0 := &models.CellsDocumentProperty{}
	cellsDocumentscellsDocument0.Name = "Author"
	cellsDocumentscellsDocument0.Value = "roy.wang"
	var cellsDocuments = []models.CellsDocumentProperty{*cellsDocumentscellsDocument0}
	request := requests.NewPostMetadataRequest(
		cellsDocuments,
		GetBaseTest().localTestDataFolder+book1Xlsx,
	)

	_, err := asposecellscloud.DoChecked(ctx, GetBaseTest().Client, request)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("%d\tTestLightCells_PostMetadata \n", GetBaseTest().GetTestNumber())
	}
}
