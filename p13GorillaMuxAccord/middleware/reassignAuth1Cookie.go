package middleware

import (
	"app/model"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


func ReAssignAuth1Cookie (user model.AccordUser) *http.Cookie{
	expirationTime := time.Now().Add(time.Minute * 100)

	claims := &model.Claims{
				UserName:   user.UserName,
		Email:      user.Email,
		UserID:     user.UserID,
		UUID : user.UUID,
		// Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil 
	}

	return &http.Cookie{
		Name:    "Auth1",
		Value:   tokenString,
		Expires: expirationTime,
		HttpOnly:true,
		
		// SameSite: SameSiteLaxMode,
	}

}