package app

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	router := New()
	//os.Setenv("PORT", "8000")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
