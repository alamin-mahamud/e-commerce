package main

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte(os.Getenv("JWT_SECRET"))
)

type CustomClaims struct {
	User *User
	jwt.StandardClaims
}

type Authable interface {
	Encode(user *User) (string, error)
	Decode(token string) (*CustomClaims, error)
}

type TokenService struct {
	repository Repository
}

// Encode Token
func (srv *TokenService) Encode(user *User) (string, error) {
	// Create the claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 86400,
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token and return
	return token.SignedString(key)
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {
	// Parse the Token
	tokenType, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	}

	return nil, err
}
