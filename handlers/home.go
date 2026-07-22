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

	// Fetch all artists from the API.
	artists, err := api.FetchArtists(r.Context())
	if err != nil {
		http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}

	// Filter the artists based on the user's search.
	// If the search query is empty, all artists are returned.
	filteredArtists := helpers.SearchArtists(artists, searchQuery)

	// Prepare the data that will be sent to the template.
	pageData := models.PageData{
		Title:        "Groupie Tracker",
		Artists:      filteredArtists,
		Search:       searchQuery,
		ResultCount:  len(filteredArtists),
		SearchActive: searchQuery != "",
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
