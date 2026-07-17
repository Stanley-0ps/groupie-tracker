package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
)

const relationURL = "https://groupietrackers.herokuapp.com/api/relation"

func FetchRelations() ([]models.Relation, error) {
	response, err := http.Get(relationURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", response.StatusCode)
	}
	var resp models.RelationsResponse

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return resp.Index, nil
}
