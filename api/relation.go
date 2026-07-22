package api

import (
	"context"
	"groupie-tracker/models"
)

const relationURL = "https://groupietrackers.herokuapp.com/api/relation"

func FetchRelations(ctx context.Context) ([]models.Relation, error) {

	var response models.RelationsResponse

	err := FetchJSON(ctx, relationURL, &response)
	if err != nil {
		return nil, err
	}

	return response.Index, nil
}
