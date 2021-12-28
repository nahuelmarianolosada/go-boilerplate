package models

import "github.com/golang-jwt/jwt"

type Login struct {
	// TODO: Implement Login model
	Username string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}


type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}