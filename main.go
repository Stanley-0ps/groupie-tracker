package main

import (
	"html/template"
	"log"
	"net/http"
)

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

	//rendering template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}
}
