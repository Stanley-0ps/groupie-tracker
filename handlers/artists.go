package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow the /artist route
	if r.URL.Path != "/artist" {
		http.NotFound(w, r)
		return
	}

	// Get the id from the URL
	id := r.URL.Query().Get("id")

	// Convert it to an integer
	artistID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch all artists
	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}

	// Fetch all artist locations
	locations, err := api.FetchLocations()
	if err != nil {
		http.Error(w, "Unable to fetch locations", http.StatusInternalServerError)
		return
	}

	//fetch artist relation
	relations, err := api.FetchRelations()
	if err != nil {
		http.Error(w, "Unable to fetch relations", http.StatusInternalServerError)
		return
	}

	// Find the requested artist
	var selectedArtist models.Artist
	found := false

	for _, artist := range artists {
		if artist.ID == artistID {
			selectedArtist = artist
			found = true
			break
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	// Find artist locations
	var artistLocations []string

	for _, location := range locations {
		if location.ID == artistID {
			artistLocations = location.Locations
			break
		}
	}

	//Find artist relations
	var artistRelation map[string][]string

	for _, relation := range relations {
		if relation.ID == artistID {
			artistRelation = relation.DatesLocations
			break
		}
	}

	// Register custom template functions.
	funcMap := template.FuncMap{
		"formatLocation": helpers.FormatLocation,
	}

	tmpl, err := template.New("artist.html").
		Funcs(funcMap).
		ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	pageData := models.ArtistPageData{
		Artist:    selectedArtist,
		Locations: artistLocations,
		Relation:  artistRelation,
	}
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
