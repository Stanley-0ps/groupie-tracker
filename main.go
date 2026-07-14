package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	log.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Groupie Tracker"))
}
