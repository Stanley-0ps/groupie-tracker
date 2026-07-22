package helpers

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

// SearchArtists returns artists for which every search term matches at least
// one searchable field: name, member, creation year, first album, or location.
func SearchArtists(artists []models.Artist, query string) []models.Artist {
	// Fields normalizes whitespace and case, then supports combined word searches.
	terms := strings.Fields(strings.ToLower(strings.TrimSpace(query)))
	if len(terms) == 0 {
		return artists
	}

	var filtered []models.Artist
	for _, artist := range artists {
		if matchesAllTerms(artist, terms) {
			// The artist is appended once after all terms match.
			filtered = append(filtered, artist)
		}
	}

	return filtered
}

// matchesAllTerms checks whether each term occurs in any searchable artist field.
func matchesAllTerms(artist models.Artist, terms []string) bool {
	for _, term := range terms {
		if !matchesTerm(artist, term) {
			// One missing term means this artist does not satisfy the combined search.
			return false
		}
	}

	return true
}

// matchesTerm checks one normalized term against every field available for search.
func matchesTerm(artist models.Artist, term string) bool {
	if strings.Contains(strings.ToLower(artist.Name), term) ||
		strings.Contains(strconv.Itoa(artist.CreationDate), term) ||
		strings.Contains(strings.ToLower(artist.FirstAlbum), term) {
		return true
	}

	for _, location := range artist.Locations {
		if strings.Contains(strings.ToLower(location), term) {
			return true
		}
	}

	for _, member := range artist.Members {
		if strings.Contains(strings.ToLower(member), term) {
			return true
		}
	}

	return false
}
