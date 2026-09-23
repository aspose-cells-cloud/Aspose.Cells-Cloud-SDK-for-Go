package unittest_test

import (
	"encoding/json"
	"testing"

	asposecellscloud "asposecellscloud"
	"asposecellscloud/models"
)

// The specification declares models.Color's A/R/G/B as `Byte`, and the generator
// used to map that identifier to Go's []byte. encoding/json renders []byte as a
// base64 STRING, so a colour filter was sent as {"R":"MA=="} where the service
// expects a number. `Byte` now maps to *int32, like `Integer`.
//
// These tests pin the numeric encoding, and pin that a component explicitly set
// to 0 survives serialization -- every model field carries an unconditional
// `omitempty`, so a value-typed int would drop a fully transparent A or a black
// R/G/B on the wire.

func TestColorEncodesComponentsAsNumbers(t *testing.T) {
	c := &models.Color{
		A: asposecellscloud.Int32Ptr(255),
		R: asposecellscloud.Int32Ptr(48),
		G: asposecellscloud.Int32Ptr(0),
		B: asposecellscloud.Int32Ptr(0),
	}

	encoded, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("marshal Color: %v", err)
	}

	const want = `{"A":255,"R":48,"G":0,"B":0}`
	if string(encoded) != want {
		t.Fatalf("Color JSON = %s, want %s", encoded, want)
	}
}

func TestColorOmitsUnsetComponents(t *testing.T) {
	encoded, err := json.Marshal(&models.Color{})
	if err != nil {
		t.Fatalf("marshal empty Color: %v", err)
	}
	if string(encoded) != "{}" {
		t.Fatalf("empty Color JSON = %s, want {}", encoded)
	}

	// A nil component is absent; the ones set alongside it are unaffected.
	partial, err := json.Marshal(&models.Color{R: asposecellscloud.Int32Ptr(1)})
	if err != nil {
		t.Fatalf("marshal partial Color: %v", err)
	}
	if string(partial) != `{"R":1}` {
		t.Fatalf("partial Color JSON = %s, want {\"R\":1}", partial)
	}
}

func TestColorRoundTripPreservesZeroComponents(t *testing.T) {
	// Fully transparent (A=0) and black (R=G=B=0) are legitimate values, not
	// "unset" -- they must come back out of a round trip intact.
	const wire = `{"A":0,"R":0,"G":0,"B":0}`

	var c models.Color
	if err := json.Unmarshal([]byte(wire), &c); err != nil {
		t.Fatalf("unmarshal Color: %v", err)
	}
	for name, got := range map[string]*int32{"A": c.A, "R": c.R, "G": c.G, "B": c.B} {
		if got == nil {
			t.Fatalf("round trip: component %s is nil, want pointer to 0", name)
		}
		if *got != 0 {
			t.Fatalf("round trip: component %s = %d, want 0", name, *got)
		}
	}

	encoded, err := json.Marshal(&c)
	if err != nil {
		t.Fatalf("re-marshal Color: %v", err)
	}
	if string(encoded) != wire {
		t.Fatalf("re-marshalled Color = %s, want %s", encoded, wire)
	}
}

// TestColorFilterRequestBodyIsNumeric covers the path the live tests actually
// exercise: ColorFilterRequest -> CellsColor -> Color. It reads the components
// back out of generic JSON so it fails on a base64 string, which is the wire
// form the []byte mapping produced.
func TestColorFilterRequestBodyIsNumeric(t *testing.T) {
	request := &models.ColorFilterRequest{
		Pattern: "Solid",
		ForegroundColor: &models.CellsColor{
			Type: "Automatic",
			Color: &models.Color{
				R: asposecellscloud.Int32Ptr(48),
				G: asposecellscloud.Int32Ptr(48),
				B: asposecellscloud.Int32Ptr(48),
			},
		},
	}

	body, err := json.Marshal(request)
	if err != nil {
		t.Fatalf("marshal ColorFilterRequest: %v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("unmarshal ColorFilterRequest: %v", err)
	}

	color := nestedObject(t, decoded, "ForegroundColor", "Color")
	for _, name := range []string{"R", "G", "B"} {
		component, ok := color[name]
		if !ok {
			t.Fatalf("ColorFilterRequest body %s has no %q: %s", body, name, body)
		}
		number, ok := component.(float64)
		if !ok {
			t.Fatalf("ColorFilterRequest body component %s = %#v (%T), want a JSON number", name, component, component)
		}
		if number != 48 {
			t.Fatalf("ColorFilterRequest body component %s = %v, want 48", name, number)
		}
	}
}

// nestedObject walks a decoded JSON document down the named objects.
func nestedObject(t *testing.T, doc map[string]interface{}, path ...string) map[string]interface{} {
	t.Helper()
	for _, key := range path {
		child, ok := doc[key].(map[string]interface{})
		if !ok {
			t.Fatalf("expected %q to be a JSON object, got %#v", key, doc[key])
		}
		doc = child
	}
	return doc
}
