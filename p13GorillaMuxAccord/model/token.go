package model

import (

	"github.com/golang-jwt/jwt/v4"
	// "github.com/dgrijalva/jwt-go"
)


type Claims struct {
	UserName string `json:"UserName"`
	UserID string `json:"UserID"`
	Email string `json:"email"`
	UUID string `json:"UUID"`
	jwt.StandardClaims	
}

type RefreshClaims struct {
	UserName string `json:"UserName"`
	UserID string `json:"UserID"`
	Email string `json:"email"`
	UUID string `json:"UUID"`
	jwt.StandardClaims	
}