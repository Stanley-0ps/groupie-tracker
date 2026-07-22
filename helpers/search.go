package helpers

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func SearchArtists(artists []models.Artist, query string) []models.Artist {

	// Remove extra spaces and convert to lowercase.
	query = strings.TrimSpace(strings.ToLower(query))

	// If the search box is empty,
	// return every artist.
	if query == "" {
		return artists
	}

	var filtered []models.Artist

	// Check every searchable artist field. Each artist is appended at most once.
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), query) {
			filtered = append(filtered, artist)
			continue
		}

		// Convert the numeric year so it can use the same partial-match search.
		if strings.Contains(strconv.Itoa(artist.CreationDate), query) {
			filtered = append(filtered, artist)
			continue
		}

		// FirstAlbum is an API date string, so it can be searched directly.
		if strings.Contains(strings.ToLower(artist.FirstAlbum), query) {
			filtered = append(filtered, artist)
			continue
		}

		locationMatched := false
		for _, location := range artist.Locations {
			if strings.Contains(strings.ToLower(location), query) {
				filtered = append(filtered, artist)
				locationMatched = true
				break
			}
		}

		// A location match has already added the artist, so do not duplicate it.
		if locationMatched {
			continue
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), query) {
				// Add each matching artist only once, even if more than one
				// member matches the query.
				filtered = append(filtered, artist)
				break
			}
		}
	}

	return filtered
}
