package app

import (
	"encoding/json"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World2")
}
