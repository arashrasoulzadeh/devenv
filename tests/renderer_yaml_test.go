package tests

import (
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/renderer"
)

func TestParseYAML_EmptyMap(t *testing.T) {
	input := map[string]any{}

	got := renderer.ParseYAML(input)
	want := ""

	if got != want {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestParseYAML_BasicTypes(t *testing.T) {
	input := map[string]any{
		"foo":   "bar",
		"num":   123,
		"pi":    3.14,
		"debug": true,
	}

	got := renderer.ParseYAML(input)

	// sorted order: a-z
	want := `debug: true
foo: "bar"
num: 123
pi: 3.14`

	if got != want {
		t.Errorf("unexpected output:\nGOT:\n%s\nWANT:\n%s", got, want)
	}
}

func TestParseYAML_StringWithQuotes(t *testing.T) {
	input := map[string]any{
		"name": `John "Johnny"`,
	}

	got := renderer.ParseYAML(input)
	want := `name: "John \"Johnny\""`

	if got != want {
		t.Errorf("unexpected escaping:\nGOT: %q\nWANT: %q", got, want)
	}
}

func TestParseYAML_NilValuesSkipped(t *testing.T) {
	input := map[string]any{
		"skip_me": nil,
		"x":       "y",
	}

	got := renderer.ParseYAML(input)
	want := `x: "y"`

	if got != want {
		t.Errorf("expected nil values to be skipped, got %q", got)
	}
}

func TestParseYAML_MixedTypesOrdering(t *testing.T) {
	input := map[string]any{
		"z": 1,
		"a": false,
		"m": "middle",
	}

	got := renderer.ParseYAML(input)

	// sorted keys: a, m, z
	want := `a: false
m: "middle"
z: 1`

	if got != want {
		t.Errorf("unexpected ordering:\nGOT:\n%s\nWANT:\n%s", got, want)
	}
}
