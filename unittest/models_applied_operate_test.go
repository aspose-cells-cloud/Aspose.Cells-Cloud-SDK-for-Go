package unittest_test

import (
	"encoding/json"
	"testing"

	asposecellscloud "asposecellscloud"
	"asposecellscloud/models"
)

// The specification models polymorphism with an abstract parent plus a
// discriminator property: AppliedOperate is abstract, MergeQueries / PivotColumn /
// UnpivotColumn inherit from it, and each carries an inherited AppliedOperateType
// whose spec description reads "Utilizes a custom JSON converter to serialize and
// deserialize an enum property". That converter is on the service side, and it
// dispatches on that property. Measured against the live API:
//
//   - no operate, or an operate without the discriminator -> HTTP 500
//     ("Object reference not set to an instance of an object");
//   - a valid member on the wrong payload -> HTTP 500 ("Value cannot be null.
//     (Parameter 'key')"), i.e. it really routes by the discriminator;
//   - a non-member -> HTTP 500 ("Could not convert 'whatever' to
//     AppliedOperateType"), so it is an enum of the concrete type names;
//   - the member -> HTTP 200. Casing is irrelevant (the converter is
//     case-insensitive).
//
// Two generator defects used to make the model unusable here:
//
//   - a Class-typed property referencing an abstract model became `*AbstractType`
//     -- a pointer to an interface, which no concrete model can be assigned to, so
//     AppliedStep.AppliedOperate could never be populated. It is now the interface
//     itself (`AppliedOperate`), so a *UnpivotColumn assigns directly.
//   - inherited properties were skipped unconditionally on the assumption that
//     embedding the parent would supply them. Since an abstract parent becomes an
//     interface it is *not* embedded, so the discriminator was lost outright. It is
//     now emitted into the child whenever the parent is not embedded.
//
// The value of the discriminator is implied by the concrete Go type, so each child
// now fills it in from its own name on serialization -- the caller only supplies
// the data. These tests pin that, and the wire shape the live
// PostDataTransformation test depends on.

func TestAppliedStepAcceptsConcreteOperate(t *testing.T) {
	step := &models.AppliedStep{StepName: "UnpivotColumn"}
	// The line that used to be impossible: a concrete operate assigned to the
	// parent-typed field. The discriminator is not set -- it is derived.
	step.AppliedOperate = &models.UnpivotColumn{
		UnpivotColumnNames: []string{"2017", "2018", "2019"},
		ColumnMapName:      "Date",
		ValueMapName:       "Count",
	}

	body, err := json.Marshal(step)
	if err != nil {
		t.Fatalf("marshal AppliedStep: %v", err)
	}

	const want = `{"StepName":"UnpivotColumn","AppliedOperate":{"UnpivotColumnNames":["2017","2018","2019"],` +
		`"ColumnMapName":"Date","ValueMapName":"Count","AppliedOperateType":"UnpivotColumn"}}`
	if string(body) != want {
		t.Fatalf("AppliedStep JSON = %s, want %s", body, want)
	}
}

func TestAppliedOperateDiscriminatorComesFromTheConcreteType(t *testing.T) {
	// Every child of an abstract parent must name itself, without being asked.
	operates := []models.AppliedOperate{
		&models.UnpivotColumn{},
		&models.PivotColumn{},
		&models.MergeQueries{},
	}
	wantMembers := []string{"UnpivotColumn", "PivotColumn", "MergeQueries"}

	for i, operate := range operates {
		body, err := json.Marshal(operate)
		if err != nil {
			t.Fatalf("marshal %T: %v", operate, err)
		}

		var decoded map[string]interface{}
		if err := json.Unmarshal(body, &decoded); err != nil {
			t.Fatalf("unmarshal %T: %v", operate, err)
		}

		got, ok := decoded["AppliedOperateType"]
		if !ok {
			t.Fatalf("%T JSON = %s, want a derived AppliedOperateType", operate, body)
		}
		if got != wantMembers[i] {
			t.Fatalf("%T AppliedOperateType = %v, want %q", operate, got, wantMembers[i])
		}
	}
}

func TestAppliedOperateExplicitDiscriminatorIsKept(t *testing.T) {
	// The derived value is a default, not a clamp: an explicit member survives,
	// which is what lets a caller pin a casing the service accepts.
	body, err := json.Marshal(&models.UnpivotColumn{
		UnpivotColumnNames: []string{"2017"},
		AppliedOperateType: "unpivotColumn",
	})
	if err != nil {
		t.Fatalf("marshal UnpivotColumn: %v", err)
	}
	if want := `{"UnpivotColumnNames":["2017"],"AppliedOperateType":"unpivotColumn"}`; string(body) != want {
		t.Fatalf("UnpivotColumn JSON = %s, want %s", body, want)
	}
}

func TestAppliedOperateDiscriminatorSurvivesARoundTrip(t *testing.T) {
	// A payload read back from the service (or written without the member) still
	// names its own type when it goes out again.
	var operate models.UnpivotColumn
	if err := json.Unmarshal([]byte(`{"ValueMapName":"Count"}`), &operate); err != nil {
		t.Fatalf("unmarshal UnpivotColumn: %v", err)
	}

	body, err := json.Marshal(operate)
	if err != nil {
		t.Fatalf("marshal UnpivotColumn: %v", err)
	}
	if want := `{"ValueMapName":"Count","AppliedOperateType":"UnpivotColumn"}`; string(body) != want {
		t.Fatalf("round-tripped UnpivotColumn JSON = %s, want %s", body, want)
	}
}

func TestAppliedStepOmitsUnsetOperate(t *testing.T) {
	// A nil interface must be dropped, not rendered as `"AppliedOperate":null` --
	// the service rejects a present-but-null operate the same way it rejects an
	// absent one.
	body, err := json.Marshal(&models.AppliedStep{StepName: "UnpivotColumn"})
	if err != nil {
		t.Fatalf("marshal AppliedStep: %v", err)
	}
	if string(body) != `{"StepName":"UnpivotColumn"}` {
		t.Fatalf("AppliedStep JSON = %s, want {\"StepName\":\"UnpivotColumn\"}", body)
	}
}

// TestDataTransformationRequestBodyShape mirrors the request built by the live
// TestDataProcessingController_PostDataTransformation, which is the only
// end-to-end check of this model. It asserts the operate is nested under the step
// carrying its own discriminator, which is what the service's converter needs.
func TestDataTransformationRequestBodyShape(t *testing.T) {
	request := &models.DataTransformationRequest{
		LoadData: &models.LoadData{
			LoadTo: &models.LoadTo{
				Worksheet:        "L2W",
				BeginRowIndex:    asposecellscloud.Int32Ptr(3),
				BeginColumnIndex: asposecellscloud.Int32Ptr(2),
			},
			DataQuery: &models.DataQuery{
				Name:               "DataQuery",
				DataSourceDataType: "ListObject",
				DataSource: &models.DataSource{
					DataSourceType: "CloudFileSystem",
					DataPath:       "TestData/In/BookTableL2W.xlsx",
				},
				DataItem: &models.DataItem{DataItemType: "Table", Value: "Table1"},
			},
		},
		AppliedSteps: []models.AppliedStep{{
			StepName: "UnpivotColumn",
			AppliedOperate: &models.UnpivotColumn{
				UnpivotColumnNames: []string{"2017", "2018", "2019"},
				ColumnMapName:      "Date",
				ValueMapName:       "Count",
			},
		}},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal DataTransformationRequest: %v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal DataTransformationRequest: %v", err)
	}

	steps, ok := decoded["AppliedSteps"].([]interface{})
	if !ok || len(steps) != 1 {
		t.Fatalf("AppliedSteps = %#v, want a one-element array", decoded["AppliedSteps"])
	}
	step, ok := steps[0].(map[string]interface{})
	if !ok {
		t.Fatalf("AppliedSteps[0] = %#v, want a JSON object", steps[0])
	}
	if step["StepName"] != "UnpivotColumn" {
		t.Fatalf("AppliedSteps[0].StepName = %v, want UnpivotColumn", step["StepName"])
	}

	operate, ok := step["AppliedOperate"].(map[string]interface{})
	if !ok {
		t.Fatalf("AppliedSteps[0].AppliedOperate = %#v, want the concrete operate nested under the step", step["AppliedOperate"])
	}
	if operate["AppliedOperateType"] != "UnpivotColumn" {
		t.Fatalf("AppliedSteps[0].AppliedOperate.AppliedOperateType = %v, want UnpivotColumn", operate["AppliedOperateType"])
	}
	if operate["ValueMapName"] != "Count" || operate["ColumnMapName"] != "Date" {
		t.Fatalf("AppliedSteps[0].AppliedOperate lost its concrete data: %s", body)
	}
}
