package helpers

import (
	"groupie-tracker/models"
	"reflect"
	"testing"
)

func TestFilterArtists(t *testing.T) {
	t.Parallel()

	artists := []models.Artist{
		{ID: 1, CreationDate: 1970, FirstAlbum: "14-12-1973", Members: []string{"A", "B", "C", "D"}},
		{ID: 2, CreationDate: 1996, FirstAlbum: "26-06-2000", Members: []string{"A", "B", "C", "D"}},
		{ID: 3, CreationDate: 1970, FirstAlbum: "22-03-1963", Members: []string{"A", "B"}},
	}

	filters := models.SearchFilters{CreationYear: 1970, MemberCount: 4}
	if got, want := FilterArtists(artists, filters), artists[:1]; !reflect.DeepEqual(got, want) {
		t.Fatalf("FilterArtists() = %#v, want %#v", got, want)
	}
}

func TestParseSearchFilters(t *testing.T) {
	t.Parallel()

	got, err := ParseSearchFilters("1970", "1973", "4")
	if err != nil {
		t.Fatalf("ParseSearchFilters() error = %v", err)
	}
	want := models.SearchFilters{CreationYear: 1970, FirstAlbumYear: 1973, MemberCount: 4}
	if got != want {
		t.Fatalf("ParseSearchFilters() = %#v, want %#v", got, want)
	}

	if _, err := ParseSearchFilters("invalid", "", ""); err == nil {
		t.Fatal("ParseSearchFilters() error = nil, want validation error")
	}
}
