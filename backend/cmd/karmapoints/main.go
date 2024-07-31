package main

import (
	"fmt"
	"net/http"

	"github.com/eser/karma-points/backend/pkg/karmapoints/routes"
)

func main() {
	fmt.Println("Starting server on :8080")
	http.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":8080", nil)
}
