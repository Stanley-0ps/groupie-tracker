package models

type PageData struct {
	Title         string
	Artists       []Artist
	Search        string
	Filters       SearchFilters
	FilterOptions FilterOptions
	SortBy        string
	ResultCount   int
	SearchActive  bool
	FiltersActive bool
}

// ArtistPageData holds everything needed by artist.html.
type ArtistPageData struct {
	Artist    Artist
	Locations []string
	Schedule  []ScheduleEntry
}

// ScheduleEntry represents the dates for one concert location.
type ScheduleEntry struct {
	Location string
	Dates    []string
}
