package helpers

import "groupie-tracker/models"

// AttachLocations returns artists enriched with the concert locations that
// share their ID. The API sends these resources separately, so search needs
// this step before it can match location names.
func AttachLocations(artists []models.Artist, locations []models.Location) []models.Artist {
	locationsByArtist := make(map[int][]string, len(locations))
	for _, location := range locations {
		// The map makes locating an artist's concert locations efficient.
		locationsByArtist[location.ID] = location.Locations
	}

	// Copy the slice so callers keep their original artist data unchanged.
	enrichedArtists := make([]models.Artist, len(artists))
	copy(enrichedArtists, artists)
	for index := range enrichedArtists {
		enrichedArtists[index].Locations = locationsByArtist[enrichedArtists[index].ID]
	}

	return enrichedArtists
}
