package main

import (
	"html/template"
	"log"
	"net/http"
)


func main() {
	//Serve static files(CSS, images, js)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route
	http.HandleFunc("/", homeHandler)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Home Page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// only allow the home route
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

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
