package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"groupie-tracker/api"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Only allow requests to the home page.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	searchQuery := strings.TrimSpace(r.URL.Query().Get("search"))
	filters, err := helpers.ParseSearchFilters(
		r.URL.Query().Get("creationYear"),
		r.URL.Query().Get("firstAlbumYear"),
		r.URL.Query().Get("memberCount"),
	)
	if err != nil {
		http.Error(w, "invalid filter", http.StatusBadRequest)
		return
	}

	// Fetch all artists from the API.
	artists, err := api.FetchArtists(r.Context())
	if err != nil {
		http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}
	if searchQuery != "" {
		// Location data is a separate API resource and is needed only for searches.
		locations, err := api.FetchLocations(r.Context())
		if err != nil {
			http.Error(w, "Unable to fetch artist locations", http.StatusInternalServerError)
			return
		}
		artists = helpers.AttachLocations(artists, locations)
	}

	// Filter the artists based on the user's search.
	// If the search query is empty, all artists are returned.
	filteredArtists := helpers.FilterArtists(helpers.SearchArtists(artists, searchQuery), filters)
	filteredArtists, err = helpers.SortArtists(filteredArtists, r.URL.Query().Get("sort"))
	if err != nil {
		http.Error(w, "invalid sort option", http.StatusBadRequest)
		return
	}

	// Prepare the data that will be sent to the template.
	pageData := models.PageData{
		Title:         "Groupie Tracker",
		Artists:       filteredArtists,
		Search:        searchQuery,
		Filters:       filters,
		FilterOptions: helpers.BuildFilterOptions(artists),
		SortBy:        r.URL.Query().Get("sort"),
		ResultCount:   len(filteredArtists),
		SearchActive:  searchQuery != "",
		FiltersActive: filters != (models.SearchFilters{}),
	}

	// Parse the HTML template.
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Render the template with the page data.
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
