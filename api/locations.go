package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

const locationsURL = "https://groupietrackers.herokuapp.com/api/locations"

func FetchLocations() ([]models.Location, error) {
	resp, err := http.Get(locationsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result models.LocationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Index, nil
}
