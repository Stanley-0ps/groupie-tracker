package helpers

import (
	"fmt"
	"groupie-tracker/models"
	"sort"
	"strconv"
	"strings"
)

// ParseSearchFilters validates query-string values before they are used to
// filter artists. Empty values intentionally mean that a filter is disabled.
func ParseSearchFilters(creationYear, firstAlbumYear, memberCount string) (models.SearchFilters, error) {
	creation, err := parseOptionalPositiveInt(creationYear, "creation year")
	if err != nil {
		return models.SearchFilters{}, err
	}
	album, err := parseOptionalPositiveInt(firstAlbumYear, "first album year")
	if err != nil {
		return models.SearchFilters{}, err
	}
	members, err := parseOptionalPositiveInt(memberCount, "member count")
	if err != nil {
		return models.SearchFilters{}, err
	}

	return models.SearchFilters{CreationYear: creation, FirstAlbumYear: album, MemberCount: members}, nil
}

// FilterArtists applies every active filter. An artist must meet all selected
// filters, which makes filtering predictable when combined with search.
func FilterArtists(artists []models.Artist, filters models.SearchFilters) []models.Artist {
	if filters == (models.SearchFilters{}) {
		return artists
	}

	var filtered []models.Artist
	for _, artist := range artists {
		if filters.CreationYear != 0 && artist.CreationDate != filters.CreationYear {
			continue
		}
		if filters.FirstAlbumYear != 0 && firstAlbumYear(artist.FirstAlbum) != filters.FirstAlbumYear {
			continue
		}
		if filters.MemberCount != 0 && len(artist.Members) != filters.MemberCount {
			continue
		}
		filtered = append(filtered, artist)
	}

	return filtered
}

// BuildFilterOptions derives sorted, unique filter values from the API data.
func BuildFilterOptions(artists []models.Artist) models.FilterOptions {
	creationYears := make(map[int]struct{})
	firstAlbumYears := make(map[int]struct{})
	memberCounts := make(map[int]struct{})

	for _, artist := range artists {
		creationYears[artist.CreationDate] = struct{}{}
		if year := firstAlbumYear(artist.FirstAlbum); year != 0 {
			firstAlbumYears[year] = struct{}{}
		}
		memberCounts[len(artist.Members)] = struct{}{}
	}

	return models.FilterOptions{
		CreationYears:   sortedKeys(creationYears),
		FirstAlbumYears: sortedKeys(firstAlbumYears),
		MemberCounts:    sortedKeys(memberCounts),
	}
}

// parseOptionalPositiveInt keeps malformed filter values from silently changing
// the result set when a user edits the URL manually.
func parseOptionalPositiveInt(value, fieldName string) (int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return 0, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed < 1 {
		return 0, fmt.Errorf("invalid %s", fieldName)
	}
	return parsed, nil
}

// firstAlbumYear extracts the final four-digit year from the API's album date.
func firstAlbumYear(firstAlbum string) int {
	if len(firstAlbum) < 4 {
		return 0
	}
	year, err := strconv.Atoi(firstAlbum[len(firstAlbum)-4:])
	if err != nil {
		return 0
	}
	return year
}

// sortedKeys converts a set to a stable order for predictable filter controls.
func sortedKeys(values map[int]struct{}) []int {
	keys := make([]int, 0, len(values))
	for value := range values {
		keys = append(keys, value)
	}
	sort.Ints(keys)
	return keys
}
