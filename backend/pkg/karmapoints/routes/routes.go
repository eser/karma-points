package routes

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Karma Points!")
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Notes API")
}

func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create Note API")
}
