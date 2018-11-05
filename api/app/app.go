package app

import (
	"log"
	"net/http"
)

func Start() {
	router := New()
	log.Fatal(http.ListenAndServe(":8000", router))
}
