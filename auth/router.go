package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(s *Service) http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	versionedRouter := router.PathPrefix("/" + "v1").Subrouter()
	initAuthRouter(versionedRouter, s)
	return versionedRouter
}

func initAuthRouter(r *mux.Router, s *Service) {
	authRouter := r.PathPrefix("/authentication").Subrouter()
	registerAuthRoutes(authRouter, s)
}

func registerAuthRoutes(r *mux.Router, s *Service) {
	routes := getAuthUserRoutes(s)
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}

// Route - Defines the URL Structure for a typical REST API
// "NAME" => METHOD + PATTERN => HandlerFunc
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// GET
// Auth
// Create
// GETAll
// Validate

func getAuthUserRoutes(authService *Service) Routes {

	return Routes{
		Route{"List", "GET", "", authService.GetAll},
		Route{"Get", "GET", "/{id}", authService.Get},
		Route{"Create", "POST", "", authService.Create},
		Route{"Auth", "POST", "/auth", authService.Auth},
		Route{"Validate", "POST", "/validate", authService.ValidateToken},
	}
}
