package handler

import (
	"app/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func (H *DatabaseCollections) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	// an example API handler
	fmt.Println("Home has called ...")
	cookie, err := r.Cookie("Auth1")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
	}

	tokenStr := cookie.Value
	fmt.Println("🚀 ~ file: home.go ~ line 28 ~ func ~ tokenStr : ", tokenStr)

	claims := model.Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
	fmt.Println("🚀 ~ file: home.go ~ line 34 ~ func ~ err : ", err)
	if err != nil {
		if err != jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("🚀 ~ file: home.go ~ line 60 ~ func ~ tkn.Valid : ", tkn.Valid)
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
	}

	fmt.Println("✨Token is valid Welcome home")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(`{status: "Welcome home"}`)

}
