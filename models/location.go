package models

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations`
}

type LocationsResponse struct {
	Index []Location `json:"index"`
}
