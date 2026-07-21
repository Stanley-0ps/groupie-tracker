package api

import (
	"groupie-tracker/models"
)

const locationsURL = "https://groupietrackers.herokuapp.com/api/locations"

func FetchLocations() ([]models.Location, error) {

	var response models.LocationsResponse

	err := FetchJSON(locationsURL, &response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}
