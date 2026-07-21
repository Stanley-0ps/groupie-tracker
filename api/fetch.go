package api

import (
	"groupie-tracker/models"
)

const artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

// FetchArtists fetches all artists from the API.
func FetchArtists() ([]models.Artist, error) {

	var artists []models.Artist

	err := FetchJSON(artistsURL, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
