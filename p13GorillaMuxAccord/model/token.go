package model

import (

	"github.com/golang-jwt/jwt/v4"
	// "github.com/dgrijalva/jwt-go"
)


type Claims struct {
	Username string `json:"username"`
	Userid string `json:"userid"`
	Email string `json:"email"`
	Accouttype string `json:"accounttype"`
	jwt.StandardClaims	
}