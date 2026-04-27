package tests

import (
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/renderer"
)

func TestParseTOML_EmptyMap(t *testing.T) {
	input := map[string]any{}

	got := renderer.ParseTOML(input)
	want := ""

	if got != want {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestParseTOML_SingleString(t *testing.T) {
	input := map[string]any{
		"title": "my app",
	}

	got := renderer.ParseTOML(input)
	want := `title = "my app"`

	if got != want {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestParseTOML_BasicTypes(t *testing.T) {
	input := map[string]any{
		"debug": true,
		"port":  8080,
		"name":  "api",
	}

	got := renderer.ParseTOML(input)

	// sorted keys: debug, name, port
	want := `debug = true
name = "api"
port = 8080`

	if got != want {
		t.Errorf("unexpected output:\nGOT:\n%s\nWANT:\n%s", got, want)
	}
}

func TestParseTOML_NilValuesSkipped(t *testing.T) {
	input := map[string]any{
		"a": nil,
		"b": "value",
	}

	got := renderer.ParseTOML(input)
	want := `b = "value"`

	if got != want {
		t.Errorf("expected nil values to be skipped, got %q", got)
	}
}

func TestParseTOML_StringEscaping(t *testing.T) {
	input := map[string]any{
		"key": `he said "hello"`,
	}

	got := renderer.ParseTOML(input)
	want := `key = "he said \"hello\""`

	if got != want {
		t.Errorf("unexpected escaping:\nGOT: %q\nWANT: %q", got, want)
	}
}

func TestParseTOML_MixedTypesOrdering(t *testing.T) {
	input := map[string]any{
		"z": 1,
		"a": false,
		"m": "middle",
	}

	got := renderer.ParseTOML(input)

	// sorted keys: a, m, z
	want := `a = false
m = "middle"
z = 1`

	if got != want {
		t.Errorf("unexpected ordering:\nGOT:\n%s\nWANT:\n%s", got, want)
	}
}
