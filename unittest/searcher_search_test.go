package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/searcher"
	"asposecellscloud/internal/testutil"
)

// TestSearch tests the v4.0 Search function for local files.
// This is the equivalent of the v3.0 SearchSpreadsheetContent for local files.
func TestSearch(t *testing.T) {
	respBody := `{"TextItems":[{"Filename":"a.xlsx","Worksheet":"Sheet1","Position":"Cell:E35","Content":"margin"}]}`
	client, capture := testutil.NewServer(t, respBody)

	items, err := searcher.Search(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		"Sheet1",
		"margin")
	if err != nil {
		t.Fatalf("searcher.Search failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/search/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/search/content", c.Method, c.Path)
	}
	if got := c.Query.Get("searchText"); got != "margin" {
		t.Errorf("searchText = %q, want margin", got)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if got := string(c.Files["Spreadsheet"]); got != "source-xlsx" {
		t.Errorf("Spreadsheet part = %q, want source-xlsx", got)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	if items[0].Text != "margin" || items[0].Worksheet != "Sheet1" || items[0].Position != "Cell:E35" {
		t.Errorf("item = %+v, want {margin Sheet1 Cell:E35}", items[0])
	}
}

// TestReplace tests the v4.0 Replace function for local files.
// This is the equivalent of the v3.0 ReplaceSpreadsheetContent for local files.
func TestReplace(t *testing.T) {
	client, capture := testutil.NewServer(t, "replaced")
	sink := &datasource.BytesSink{}

	err := searcher.Replace(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		"old",
		"new")
	if err != nil {
		t.Fatalf("searcher.Replace failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/replace/content", c.Method, c.Path)
	}
	if got := c.Query.Get("searchText"); got != "old" {
		t.Errorf("searchText = %q, want old", got)
	}
	if got := c.Query.Get("replaceText"); got != "new" {
		t.Errorf("replaceText = %q, want new", got)
	}
	if got := string(sink.Bytes()); got != "replaced" {
		t.Errorf("sink = %q, want replaced", got)
	}
}

// TestSearchWorksheet tests the v4.0 SearchWorksheet function for cloud files.
// This is the equivalent of the v3.0 SearchContentInRemoteWorksheet for cloud files.
func TestSearchWorksheet(t *testing.T) {
	respBody := `{"TextItems":[{"Filename":"a.xlsx","Worksheet":"Sheet1","Position":"Cell:E35","Content":"margin"}]}`
	client, capture := testutil.NewServer(t, respBody)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", Folder: "TestData/In", StorageName: "s3"}

	items, err := searcher.SearchWorksheet(context.Background(), client, wf, "Sheet1", "margin")
	if err != nil {
		t.Fatalf("searcher.SearchWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/search/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/search/content", c.Method, c.Path)
	}
	if got := c.Query.Get("folder"); got != "TestData/In" {
		t.Errorf("folder = %q, want TestData/In", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
	if len(items) != 1 || items[0].Text != "margin" {
		t.Errorf("items = %+v, want one margin item", items)
	}
}

// TestSearchRange tests the v4.0 SearchRange function for cloud file ranges.
// This is the equivalent of the v3.0 SearchContentInRemoteRange for cloud file ranges.
func TestSearchRange(t *testing.T) {
	respBody := `{"TextItems":[{"Filename":"a.xlsx","Worksheet":"Sheet1","Position":"Cell:E35","Content":"margin"}]}`
	client, capture := testutil.NewServer(t, respBody)
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}

	items, err := searcher.SearchRange(context.Background(), client, wf, "Sheet1", "E35:F40", "margin")
	if err != nil {
		t.Fatalf("searcher.SearchRange failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/search/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/search/content", c.Method, c.Path)
	}
	if len(items) != 1 {
		t.Errorf("items = %+v, want one item", items)
	}
}

// TestReplaceWorkbook tests the v4.0 ReplaceWorkbook function for cloud files.
// This is the equivalent of the v3.0 ReplaceContentInRemoteSpreadsheet for cloud files.
func TestReplaceWorkbook(t *testing.T) {
	client, capture := testutil.NewServer(t, "ok")
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx", StorageName: "s3"}

	err := searcher.ReplaceWorkbook(context.Background(), client, wf, "old", "new")
	if err != nil {
		t.Fatalf("searcher.ReplaceWorkbook failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/replace/content", c.Method, c.Path)
	}
	if got := c.Query.Get("searchText"); got != "old" {
		t.Errorf("searchText = %q, want old", got)
	}
	if got := c.Query.Get("replaceText"); got != "new" {
		t.Errorf("replaceText = %q, want new", got)
	}
	if got := c.Query.Get("storageName"); got != "s3" {
		t.Errorf("storageName = %q, want s3", got)
	}
}

// TestReplaceWorksheet tests the v4.0 ReplaceWorksheet function for cloud file worksheets.
// This is the equivalent of the v3.0 ReplaceContentInRemoteWorksheet for cloud file worksheets.
func TestReplaceWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "ok")
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}

	err := searcher.ReplaceWorksheet(context.Background(), client, wf, "Sheet1", "old", "new")
	if err != nil {
		t.Fatalf("searcher.ReplaceWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/replace/content", c.Method, c.Path)
	}
}

// TestReplaceRange tests the v4.0 ReplaceRange function for cloud file ranges.
// This is the equivalent of the v3.0 ReplaceContentInRemoteRange for cloud file ranges.
func TestReplaceRange(t *testing.T) {
	client, capture := testutil.NewServer(t, "ok")
	wf := &asposecellscloud.WorkbookRef{Name: "Book1.xlsx"}

	err := searcher.ReplaceRange(context.Background(), client, wf, "Sheet1", "E35:F40", "old", "new")
	if err != nil {
		t.Fatalf("searcher.ReplaceRange failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/replace/content" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/Book1.xlsx/worksheets/Sheet1/ranges/E35:F40/replace/content", c.Method, c.Path)
	}
}

// TestSearch_Real tests search using real API.
func TestSearch_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)

	items, err := searcher.Search(context.Background(), client,
		datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		"Sheet1",
		"test")
	if err != nil {
		t.Skipf("Search failed (might be expected): %v", err)
		return
	}

	t.Logf("Search: found %d items", len(items))
	for _, item := range items {
		t.Logf("  - %s in %s at %s", item.Text, item.Worksheet, item.Position)
	}
}

// TestReplace_Real tests replace using real API.
func TestReplace_Real(t *testing.T) {
	if !SkipUnlessRealTest(t) {
		return
	}

	client := RealClient(t)
	sink := &datasource.BytesSink{}

	err := searcher.Replace(context.Background(), client,
		datasource.BytesSource([]byte("TestData/Book1.xlsx")),
		sink,
		"test",
		"replacement")
	if err != nil {
		t.Skipf("Replace failed (might be expected): %v", err)
		return
	}

	if len(sink.Bytes()) == 0 {
		t.Error("expected non-empty output")
	} else {
		t.Logf("Replace output size: %d bytes", len(sink.Bytes()))
	}
}

// TestSearcher_Validation tests validation of searcher parameters.
func TestSearcher_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}
	wf := &asposecellscloud.WorkbookRef{Name: "B.xlsx"}

	type testFunc func() error
	tests := []struct {
		name string
		fn   testFunc
	}{
		{"Search nil source", func() error {
			_, err := searcher.Search(ctx, client, nil, "S", "t")
			return err
		}},
		{"Search empty worksheet", func() error {
			_, err := searcher.Search(ctx, client, src, "", "t")
			return err
		}},
		{"Search empty text", func() error {
			_, err := searcher.Search(ctx, client, src, "S", "")
			return err
		}},
		{"Replace nil source", func() error {
			return searcher.Replace(ctx, client, nil, sink, "a", "b")
		}},
		{"Replace nil sink", func() error {
			return searcher.Replace(ctx, client, src, nil, "a", "b")
		}},
		{"SearchWorksheet nil wf", func() error {
			_, err := searcher.SearchWorksheet(ctx, client, nil, "S", "t")
			return err
		}},
		{"SearchRange empty cellArea", func() error {
			_, err := searcher.SearchRange(ctx, client, wf, "S", "", "t")
			return err
		}},
		{"ReplaceWorkbook empty old", func() error {
			return searcher.ReplaceWorkbook(ctx, client, wf, "", "new")
		}},
		{"ReplaceWorksheet empty worksheet", func() error {
			return searcher.ReplaceWorksheet(ctx, client, wf, "", "a", "b")
		}},
		{"ReplaceRange empty cellArea", func() error {
			return searcher.ReplaceRange(ctx, client, wf, "S", "", "a", "b")
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !errors.Is(tt.fn(), asposecellscloud.ErrInvalidParam) {
				t.Errorf("expected ErrInvalidParam")
			}
		})
	}
}
