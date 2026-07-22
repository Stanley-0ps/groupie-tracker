package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"sort"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/artist" {
		http.NotFound(w, r)
		return
	}

	artistID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || artistID < 1 {
		http.Error(w, "invalid artist ID", http.StatusBadRequest)
		return
	}

	artists, err := api.FetchArtists(r.Context())
	if err != nil {
		http.Error(w, "unable to fetch artists", http.StatusInternalServerError)
		return
	}
	locations, err := api.FetchLocations(r.Context())
	if err != nil {
		http.Error(w, "unable to fetch locations", http.StatusInternalServerError)
		return
	}
	relations, err := api.FetchRelations(r.Context())
	if err != nil {
		http.Error(w, "unable to fetch relations", http.StatusInternalServerError)
		return
	}

	var artist models.Artist
	found := false
	for _, candidate := range artists {
		if candidate.ID == artistID {
			artist, found = candidate, true
			break
		}
	}
	if !found {
		http.NotFound(w, r)
		return
	}

	var artistLocations []string
	for _, location := range locations {
		if location.ID == artistID {
			artistLocations = location.Locations
			break
		}
	}

	var schedule []models.ScheduleEntry
	for _, relation := range relations {
		if relation.ID != artistID {
			continue
		}
		for location, dates := range relation.DatesLocations {
			schedule = append(schedule, models.ScheduleEntry{Location: location, Dates: dates})
		}
		break
	}
	sort.Slice(schedule, func(i, j int) bool {
		return schedule[i].Location < schedule[j].Location
	})

	tmpl, err := template.New("artist.html").Funcs(template.FuncMap{
		"formatLocation": helpers.FormatLocation,
	}).ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "unable to load template", http.StatusInternalServerError)
		return
	}

	pageData := models.ArtistPageData{
		Artist:    artist,
		Locations: artistLocations,
		Schedule:  schedule,
	}
	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, "unable to render template", http.StatusInternalServerError)
	}
}
