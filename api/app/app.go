package app

import (
	"log"
	"net/http"
	"os"
)

// Start the App
func Start() {
	router := New()

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
