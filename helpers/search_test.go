package helpers

import (
	"groupie-tracker/models"
	"reflect"
	"testing"
)

func TestSearchArtists(t *testing.T) {
	t.Parallel()

	artists := []models.Artist{
		{ID: 1, Name: "Queen", CreationDate: 1970, FirstAlbum: "14-12-1973", Members: []string{"Freddie Mercury", "Brian May", "John Deacon"}, Locations: []string{"london-uk"}},
		{ID: 2, Name: "Coldplay", CreationDate: 1996, FirstAlbum: "26-06-2000", Members: []string{"Chris Martin", "Jonny Buckland", "Will Champion", "Guy Berryman"}, Locations: []string{"paris-france"}},
		{ID: 3, Name: "The Beatles", CreationDate: 1960, FirstAlbum: "22-03-1963", Members: []string{"John Lennon", "Paul McCartney", "George Harrison", "Ringo Starr"}, Locations: []string{"liverpool-uk", "london-uk"}},
	}

	got := SearchArtists(artists, "  QUEEN ")
	want := artists[:1]
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() = %#v, want %#v", got, want)
	}

	if got := SearchArtists(artists, ""); !reflect.DeepEqual(got, artists) {
		t.Fatalf("SearchArtists() with empty query = %#v, want all artists", got)
	}

	if got, want := SearchArtists(artists, "fReDdIe"), artists[:1]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by member = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "martin"), artists[1:2]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by partial member name = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "john"), []models.Artist{artists[0], artists[2]}; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by shared member name = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "1970"), artists[:1]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by creation year = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "19"), artists; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by partial creation year = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "26-06-2000"), artists[1:2]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by first album date = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "-03-1963"), artists[2:]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by partial first album date = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "PARIS"), artists[1:2]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by concert location = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "uk"), []models.Artist{artists[0], artists[2]}; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by shared concert location = %#v, want %#v", got, want)
	}

	if got, want := SearchArtists(artists, "queen freddie 1970"), artists[:1]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by combined fields = %#v, want %#v", got, want)
	}

	if got := SearchArtists(artists, "queen 1996"); len(got) != 0 {
		t.Fatalf("SearchArtists() with unmatched combined fields = %#v, want no artists", got)
	}

	if got, want := SearchArtists(artists, "john 1960"), artists[2:]; !reflect.DeepEqual(got, want) {
		t.Fatalf("SearchArtists() by member and creation year = %#v, want %#v", got, want)
	}
}
