package handler

import (
	"app/model"
	"fmt"
	"os"
	"time"

	// "log"
	// "context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	// "github.com/dgrijalva/jwt-go"
	// "gorm.io/datatypes"
	// "time"
)

func (H *DatabaseCollections) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var l model.LoginModel
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ l : ", l.Password)

	var user model.AccordUser

	// CHECK EMAIL AND PASSWORD IN DATABASE
	// e:=H.Postgres.Where("email = ?",l.Email).Where("password=?",l.Password).Find(&user).Error
	// e:=H.Postgres.Raw("SELCT * FROM accord_users WHERE email=? AND password=?",l.Email,l.Password).Scan(&user).Error
	e := H.Postgres.Where(map[string]interface{}{"email": l.Email, "password": l.Password}).Find(&user).Error
    fmt.Println("🚀 ~ file: login.go ~ line 39 ~ func ~ user : ", user)
	// fmt.Println("🚀 ~ file: login.go ~ line 31 ~ func ~ error : ", e)

	if e!=nil {
		w.WriteHeader(http.StatusInternalServerError)
       
		fmt.Println("❌ ~ file: login.go ~ line 50 ~ func ~ http.StatusInternalServerError : ", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{status: "StatusInternalServerError"}`)
	}
	if user.ID != 0  {
		//CREATE COOKIE
		expirationTime := time.Now().Add(time.Hour * 1000)
		myClaim := &model.Claims{
			UserName:   user.UserName,
			Email:      user.Email,
			UserID:     user.UserID,
			UUID : user.UUID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		// LOGIN SUCCESSFUL
		// token,err := jwt.ParseWithClaims(jwt.SigningMethodHS256,myClaim)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

        fmt.Println("🚀 ~ file: login.go ~ line 51 ~ func ~ token : ", token)
		tokenString, err := token.SignedString( []byte(os.Getenv("JWT_SECRET")))
        
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
            fmt.Println("❌ ~ file: login.go ~ line 52 ~ func ~ StatusInternalServerError : " ,err)
			json.NewEncoder(w).Encode(`{status: "internal server Error"}`)
			return

		}else {
			http.SetCookie(w, &http.Cookie{
				Name:     "Auth1",
				Value:    tokenString,
				Expires:  expirationTime,
				HttpOnly: true,
	
				// SameSite: SameSiteLaxMode,
			})
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(`{status: "Login Successful"}`)
			return 
		}
	}
	// CHECK IF ONLY EMAIL EXIST 
	H.Postgres.Find(&user, "Email=?", l.Email)
	if user.ID != 0 {
		w.WriteHeader(http.StatusUnauthorized)
        fmt.Println("❌ ~ file: login.go ~ line 71 ~ func ~ StatusUnauthorized : ")

		json.NewEncoder(w).Encode(`{message: "Password Incorrent"}`)

	} else {
		w.WriteHeader(http.StatusUnauthorized)
        fmt.Println("❌ ~ file: login.go ~ line 79 ~ func ~ StatusUnauthorized : ")
		json.NewEncoder(w).Encode(`{status: "Email/Password incorrect"}`)
	}
	// fmt.Println("🚀 ~ file: login.go ~ line 34 ~ func ~ user full data : ", user)
	// if user.ID ==0 {
	// }

	

}
