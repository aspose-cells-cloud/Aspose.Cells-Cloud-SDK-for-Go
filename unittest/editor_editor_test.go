package unittest_test

import (
	"context"
	"errors"
	"testing"

	"asposecellscloud"
	"asposecellscloud/datasource"
	"asposecellscloud/editor"
	"asposecellscloud/internal/testutil"
)

// TestAddWorksheet tests the v4.0 AddWorksheet function.
// This is the equivalent of the v3.0 PutAddNewWorksheet for local files.
func TestAddWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	err := editor.AddWorksheet(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		"NewSheet",
		editor.WithPosition(3))
	if err != nil {
		t.Fatalf("editor.AddWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/add/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/add/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("sheetName"); got != "NewSheet" {
		t.Errorf("sheetName = %q, want NewSheet", got)
	}
	if got := c.Query.Get("position"); got != "3" {
		t.Errorf("position = %q, want 3", got)
	}
	if got := string(c.Files["Spreadsheet"]); got != "source-xlsx" {
		t.Errorf("Spreadsheet part = %q, want source-xlsx", got)
	}
}

// TestDeleteWorksheet tests the v4.0 DeleteWorksheet function.
// This is the equivalent of the v3.0 DeleteWorksheet for local files.
func TestDeleteWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	err := editor.DeleteWorksheet(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		"Sheet1")
	if err != nil {
		t.Fatalf("editor.DeleteWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/delete/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/delete/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("sheetName"); got != "Sheet1" {
		t.Errorf("sheetName = %q, want Sheet1", got)
	}
}

// TestRenameWorksheet tests the v4.0 RenameWorksheet function.
// This is the equivalent of the v3.0 PostRenameWorksheet for local files.
func TestRenameWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	err := editor.RenameWorksheet(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		"OldName",
		"NewName")
	if err != nil {
		t.Fatalf("editor.RenameWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/rename/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/rename/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("sourceName"); got != "OldName" {
		t.Errorf("sourceName = %q, want OldName", got)
	}
	if got := c.Query.Get("targetName"); got != "NewName" {
		t.Errorf("targetName = %q, want NewName", got)
	}
}

// TestMoveWorksheet tests the v4.0 MoveWorksheet function.
// This is the equivalent of the v3.0 PostMoveWorksheet for local files.
func TestMoveWorksheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "updated")
	sink := &datasource.BytesSink{}

	err := editor.MoveWorksheet(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")),
		sink,
		"Sheet1",
		2)
	if err != nil {
		t.Fatalf("editor.MoveWorksheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/move/worksheet" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/move/worksheet", c.Method, c.Path)
	}
	if got := c.Query.Get("worksheet"); got != "Sheet1" {
		t.Errorf("worksheet = %q, want Sheet1", got)
	}
	if got := c.Query.Get("position"); got != "2" {
		t.Errorf("position = %q, want 2", got)
	}
}

// TestListWorksheets tests the v4.0 ListWorksheets function.
// This is the equivalent of the v3.0 GetWorksheets for local files.
func TestListWorksheets(t *testing.T) {
	client, capture := testutil.NewServer(t, `[{"WorksheetName":"Sheet1","SheetType":"Worksheet"},{"WorksheetName":"Sheet2","SheetType":"Chart"}]`)

	worksheets, err := editor.ListWorksheets(context.Background(), client,
		datasource.BytesSource([]byte("source-xlsx")))
	if err != nil {
		t.Fatalf("editor.ListWorksheets failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/worksheets" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/worksheets", c.Method, c.Path)
	}
	if len(worksheets) != 2 {
		t.Fatalf("got %d worksheets, want 2", len(worksheets))
	}
	if worksheets[0].Name != "Sheet1" || worksheets[0].Type != "Worksheet" {
		t.Errorf("worksheets[0] = %+v, want {Sheet1 Worksheet}", worksheets[0])
	}
	if worksheets[1].Name != "Sheet2" || worksheets[1].Type != "Chart" {
		t.Errorf("worksheets[1] = %+v, want {Sheet2 Chart}", worksheets[1])
	}
}

// TestCreateSpreadsheet tests the v4.0 CreateSpreadsheet function.
// This creates a new empty spreadsheet.
func TestCreateSpreadsheet(t *testing.T) {
	client, capture := testutil.NewServer(t, "new-xlsx")
	sink := &datasource.BytesSink{}

	err := editor.CreateSpreadsheet(context.Background(), client, sink)
	if err != nil {
		t.Fatalf("editor.CreateSpreadsheet failed: %v", err)
	}

	c := capture()
	if c.Method != "PUT" || c.Path != "/v4.0/cells/spreadsheet/create" {
		t.Errorf("request = %s %s, want PUT /v4.0/cells/spreadsheet/create", c.Method, c.Path)
	}
	if len(c.Files) != 0 {
		t.Errorf("expected no file parts, got %v", c.Files)
	}
	if got := string(sink.Bytes()); got != "new-xlsx" {
		t.Errorf("sink = %q, want new-xlsx", got)
	}
}

// TestEditor_Validation tests validation of editor parameters.
func TestEditor_Validation(t *testing.T) {
	client, _ := testutil.NewServer(t, "")
	ctx := context.Background()
	src := datasource.BytesSource([]byte("x"))
	sink := &datasource.BytesSink{}

	tests := []struct {
		name string
		fn   func() error
	}{
		{"AddWorksheet nil source", func() error {
			return editor.AddWorksheet(ctx, client, nil, sink, "S")
		}},
		{"AddWorksheet nil sink", func() error {
			return editor.AddWorksheet(ctx, client, src, nil, "S")
		}},
		{"AddWorksheet empty sheet", func() error {
			return editor.AddWorksheet(ctx, client, src, sink, "")
		}},
		{"DeleteWorksheet nil sink", func() error {
			return editor.DeleteWorksheet(ctx, client, src, nil, "S")
		}},
		{"DeleteWorksheet empty sheet", func() error {
			return editor.DeleteWorksheet(ctx, client, src, sink, "")
		}},
		{"RenameWorksheet empty old", func() error {
			return editor.RenameWorksheet(ctx, client, src, sink, "", "New")
		}},
		{"RenameWorksheet empty new", func() error {
			return editor.RenameWorksheet(ctx, client, src, sink, "Old", "")
		}},
		{"MoveWorksheet empty sheet", func() error {
			return editor.MoveWorksheet(ctx, client, src, sink, "", 0)
		}},
		{"ListWorksheets nil source", func() error {
			_, err := editor.ListWorksheets(ctx, client, nil)
			return err
		}},
		{"CreateSpreadsheet nil sink", func() error {
			return editor.CreateSpreadsheet(ctx, client, nil)
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
