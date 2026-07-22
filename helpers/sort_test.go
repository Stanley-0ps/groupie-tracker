package helpers

import (
	"groupie-tracker/models"
	"reflect"
	"testing"
)

func TestSortArtists(t *testing.T) {
	t.Parallel()

	artists := []models.Artist{{Name: "Queen", CreationDate: 1970}, {Name: "ABBA", CreationDate: 1972}, {Name: "Coldplay", CreationDate: 1996}}
	got, err := SortArtists(artists, "name-asc")
	if err != nil {
		t.Fatalf("SortArtists() error = %v", err)
	}
	want := []models.Artist{artists[1], artists[2], artists[0]}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SortArtists() = %#v, want %#v", got, want)
	}
	if !reflect.DeepEqual(artists, []models.Artist{{Name: "Queen", CreationDate: 1970}, {Name: "ABBA", CreationDate: 1972}, {Name: "Coldplay", CreationDate: 1996}}) {
		t.Fatalf("SortArtists() changed the input slice: %#v", artists)
	}
}
