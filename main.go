package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title  string
	Artist []Artist
}

type Artist struct {
	ID           int
	Name         string
	Image        string
	Members      []string
	CreationDate int
	FirstAlbum   int
}

func main() {
	http.HandleFunc("/", homeHandler)

	log.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//parsing the html template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Groupie Tracker",
		Artist: []Artist{
			{
				Name:         "Queen",
				Image:        "queen.jpg",
				ID:           1,
				CreationDate: 1970,
			},

			{
				Name:         "Coldplay",
				Image:        "coldplay.jpg",
				ID:           2,
				CreationDate: 1997,
			},
		},
	}

	//rendering template
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
