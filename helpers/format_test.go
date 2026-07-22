package helpers

import "testing"

func TestFormatLocation(t *testing.T) {
	t.Parallel()

	if got, want := FormatLocation("são_paulo-usa"), "São Paulo, USA"; got != want {
		t.Fatalf("FormatLocation() = %q, want %q", got, want)
	}
}

func TestCapitalizeWordHandlesUTF8(t *testing.T) {
	t.Parallel()

	if got, want := CapitalizeWord("évreux"), "Évreux"; got != want {
		t.Fatalf("CapitalizeWord() = %q, want %q", got, want)
	}
}
