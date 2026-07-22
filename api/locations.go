package api

import (
	"context"
	"groupie-tracker/models"
)

const locationsURL = "https://groupietrackers.herokuapp.com/api/locations"

func FetchLocations(ctx context.Context) ([]models.Location, error) {

	var response models.LocationsResponse

	err := FetchJSON(ctx, locationsURL, &response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}
