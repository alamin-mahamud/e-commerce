package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	userRepo := &UserRepository{db}
	tokenService := &TokenService{userRepo}
	authService := &Service{userRepo, tokenService}
	r := NewRouter(authService)

	http.Handle("/", r)
	http.ListenAndServe(":40000", nil)
}
