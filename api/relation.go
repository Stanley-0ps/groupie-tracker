package api

import (
	"groupie-tracker/models"
)

const relationURL = "https://groupietrackers.herokuapp.com/api/relation"

func FetchRelations() ([]models.Relation, error) {

	var response models.RelationsResponse

	err := FetchJSON(relationURL, &response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}
