package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"net/http"
	"text/template"
)

// Home Page handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// only allow the home route
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "Unable to fetch artists", http.StatusInternalServerError)
		return
	}

	pageData := models.PageData{
		Title:   "Groupie Tracker",
		Artists: artists,
	}

	//parsing the html template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	//rendering template
	err = tmpl.Execute(w, pageData)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
