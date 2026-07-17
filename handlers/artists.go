package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"html/template"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	//Only allow the /artist route
	if r.URL.Path != "/artist" {
		http.NotFound(w, r)
		return
	}

	//Get the id from the URL
	id := r.URL.Query().Get("id")

	// Convert it to an integer
	artistID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch all artist
	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Unable to fetch artist", http.StatusInternalServerError)
		return
	}

	//Find the requested artist
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

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
