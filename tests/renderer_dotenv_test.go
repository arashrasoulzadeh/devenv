package tests

import (
	"testing"

	"github.com/arashrasoulzadeh/devenv/src/renderer"
)

func TestParseDotEnv_EmptyMap(t *testing.T) {
	input := map[string]any{}

	got := renderer.ParseDotEnv(input)
	want := ""

	if got != want {
		t.Errorf("expected empty string, got %q", got)
	}
}

func TestParseDotEnv_SingleString(t *testing.T) {
	input := map[string]any{
		"APP_NAME": "my app",
	}

	got := renderer.ParseDotEnv(input)
	want := `APP_NAME="my app"`

	if got != want {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestParseDotEnv_MultipleTypes(t *testing.T) {
	input := map[string]any{
		"DEBUG": true,
		"PORT":  8080,
		"NAME":  "api",
	}

	got := renderer.ParseDotEnv(input)

	// sorted keys: DEBUG, NAME, PORT
	want := `DEBUG=true
NAME="api"
PORT=8080`

	if got != want {
		t.Errorf("unexpected output:\nGOT:\n%q\nWANT:\n%q", got, want)
	}
}

func TestParseDotEnv_NilValueSkipped(t *testing.T) {
	input := map[string]any{
		"A": nil,
		"B": "value",
	}

	got := renderer.ParseDotEnv(input)
	want := `B="value"`

	if got != want {
		t.Errorf("expected nil to be skipped, got %q", got)
	}
}

func TestParseDotEnv_StringEscaping(t *testing.T) {
	input := map[string]any{
		"KEY": `he said "hello"`,
	}

	got := renderer.ParseDotEnv(input)
	want := `KEY="he said \"hello\""`

	if got != want {
		t.Errorf("unexpected escaping:\nGOT: %q\nWANT: %q", got, want)
	}
}
