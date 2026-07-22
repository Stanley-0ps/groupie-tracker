package helpers

import (
	"groupie-tracker/models"
	"reflect"
	"testing"
)

func TestAttachLocations(t *testing.T) {
	t.Parallel()

	artists := []models.Artist{{ID: 1, Name: "Queen"}, {ID: 2, Name: "ABBA"}}
	locations := []models.Location{{ID: 2, Locations: []string{"stockholm-sweden"}}}

	got := AttachLocations(artists, locations)
	if want := []string{"stockholm-sweden"}; !reflect.DeepEqual(got[1].Locations, want) {
		t.Fatalf("AttachLocations() locations = %#v, want %#v", got[1].Locations, want)
	}
	if artists[1].Locations != nil {
		t.Fatalf("AttachLocations() changed the input slice: %#v", artists[1].Locations)
	}
}
