package models

// SearchFilters stores optional exact-value filters selected on the home page.
// A zero value means that the corresponding filter is not active.
type SearchFilters struct {
	CreationYear   int
	FirstAlbumYear int
	MemberCount    int
}

// FilterOptions contains the values available for each home-page filter.
type FilterOptions struct {
	CreationYears   []int
	FirstAlbumYears []int
	MemberCounts    []int
}
