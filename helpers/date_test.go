package helpers

import "testing"

func TestFormatDate(t *testing.T) {
	t.Parallel()

	if got, want := FormatDate("14-12-1973"), "December 14, 1973"; got != want {
		t.Fatalf("FormatDate() = %q, want %q", got, want)
	}
	if got, want := FormatDate("unknown"), "unknown"; got != want {
		t.Fatalf("FormatDate() with invalid date = %q, want %q", got, want)
	}
}
