package main

import (
	"groupie-tracker/handlers"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", handlers.HomeHandler)

	log.Println("Server running on port http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
