package main

import (
	"fmt"
	"net/http"

	"github.com/eser/karma-points/backend/pkg/karmapoints/routes"
)

func main() {
	fmt.Println("Starting server on :8080")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", routes.HomeHandler)
	mux.HandleFunc("GET /api/notes", routes.NotesHandler)

	mux.HandleFunc("POST /api/notes", routes.CreateNoteHandler)

	http.ListenAndServe(":8080", mux)

}
