package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

const artistURL = "https://groupietrackers.herokuapp.com/api/artists"

// FetchArtists fetches all artists from the API.
func FetchArtists() ([]models.Artist, error) {
	//Send GET request
	response, err := http.Get(artistURL)
	if err != nil {
		return nil, err
	}

	//Close response body when the function exits
	defer response.Body.Close()

	// Check if the request was successful
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", response.StatusCode)
	}

	// Variable to hold the decoded artists
	var artists []models.Artist

	// Decode JSON into the slice
	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}
