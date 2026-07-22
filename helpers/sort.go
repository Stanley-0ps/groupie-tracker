package helpers

import (
	"fmt"
	"groupie-tracker/models"
	"sort"
	"strings"
)

// SortArtists returns a copied artist slice in the requested order so sorting
// never changes the source data used by other parts of the page.
func SortArtists(artists []models.Artist, sortBy string) ([]models.Artist, error) {
	sortedArtists := make([]models.Artist, len(artists))
	copy(sortedArtists, artists)

	switch sortBy {
	case "":
		return sortedArtists, nil
	case "name-asc":
		sort.SliceStable(sortedArtists, func(i, j int) bool {
			return strings.ToLower(sortedArtists[i].Name) < strings.ToLower(sortedArtists[j].Name)
		})
	case "name-desc":
		sort.SliceStable(sortedArtists, func(i, j int) bool {
			return strings.ToLower(sortedArtists[i].Name) > strings.ToLower(sortedArtists[j].Name)
		})
	case "creation-asc":
		sort.SliceStable(sortedArtists, func(i, j int) bool {
			return sortedArtists[i].CreationDate < sortedArtists[j].CreationDate
		})
	case "creation-desc":
		sort.SliceStable(sortedArtists, func(i, j int) bool {
			return sortedArtists[i].CreationDate > sortedArtists[j].CreationDate
		})
	default:
		return nil, fmt.Errorf("invalid sort option")
	}

	return sortedArtists, nil
}
