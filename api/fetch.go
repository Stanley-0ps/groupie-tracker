package api

import (
	"context"
	"groupie-tracker/models"
)

const artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

// FetchArtists fetches all artists from the API.
func FetchArtists(ctx context.Context) ([]models.Artist, error) {

	var artists []models.Artist

	err := FetchJSON(ctx, artistsURL, &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
