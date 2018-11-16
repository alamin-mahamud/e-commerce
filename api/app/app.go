package app

import (
	"log"
	"net/http"
)

// Start the App
func Start() {
	router := New()

	port := 8010 // os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
