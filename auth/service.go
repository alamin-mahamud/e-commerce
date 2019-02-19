package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo         Repository
	TokenService Authable
}

func (srv *Service) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		SendJSONErrResponse(w, http.StatusBadRequest, "You must Provide an ID")
		return
	}

	user, err := srv.Repo.Get(id)
	if err != nil {
		SendJSONErrResponse(w, http.StatusNotFound, err.Error())
		return
	}

	SendJSONResponse(w, http.StatusOK, *user)
}

func (srv *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := srv.Repo.GetAll()
	if err != nil {
		SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(w, http.StatusOK, users)
}

func (srv *Service) Auth(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailFromReq, ok := params["email"]
	if !ok {
		SendJSONErrResponse(w, http.StatusBadRequest, "You must Provide an email")
		return
	}

	passwordFromReq, ok := params["password"]
	if !ok {
		SendJSONErrResponse(w, http.StatusBadRequest, "You must Provide a password")
		return
	}

	log.Println("Logging in with:", emailFromReq, passwordFromReq)
	user, err := srv.Repo.GetByEmail(emailFromReq)

	log.Println(user)
	if err != nil {
		SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Compares our given password against the hashed password
	// stored in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordFromReq)); err != nil {
		SendJSONErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := srv.TokenService.Encode(user)
	if err != nil {
		SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	m := map[string]string{
		"token": token,
	}

	SendJSONResponse(w, http.StatusOK, m)
	return
}

func (srv *Service) Create(w http.ResponseWriter, r *http.Request) {

	user := User{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		SendJSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = string(hashedPass)
	if err := srv.Repo.Create(&user); err != nil {
		SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	SendJSONResponse(w, http.StatusCreated, map[string]interface{}{
		"success": true,
		"msg":     "successfully created",
	})
	return
}

func (srv *Service) ValidateToken(w http.ResponseWriter, r *http.Request) {

	token := struct {
		Token string `json:"token"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		SendJSONErrResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Decode Token
	claims, err := srv.TokenService.Decode(token.Token)
	if err != nil {
		SendJSONErrResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println(claims)
	if claims.User.Id == "" {
		SendJSONErrResponse(w, http.StatusInternalServerError, "invalid user")
		return
	}

	validatedResponse := struct {
		Valid bool `json:"valid"`
	}{
		Valid: true,
	}

	SendJSONResponse(w, http.StatusAccepted, validatedResponse)
	return
}
