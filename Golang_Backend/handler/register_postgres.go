package handler

// import (
// 	"app/model"
// 	"fmt"
// 	"os"
// 	"time"

// 	// "context"
// 	"encoding/json"
// 	"net/http"

// 	"github.com/golang-jwt/jwt/v4"
// 	// "gorm.io/datatypes"
// 	// "time"
// )

// func (H *DatabaseCollections) Register_postgres(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	w.Header().Add("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Credentials", "true")


// 	var l model.AccordUser
// 	err := json.NewDecoder(r.Body).Decode(&l)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ l : ", l)

// 	// ctx,cancel:=context.WithTimeout(context.Background(),time.Second*20)

// 	// defer cancel()

// 	// u := model.AccordUser{
// 	// Email:       "raka@gmail.com",
// 	// Password:    "whatsupraka",
// 	// Userid:      "skraka",
// 	// Username:    "sk shahriar ahmed raka",
// 	// Mobile:      "018388810189",
// 	// Birthdate:   datatypes.Date(),
// 	// Accounttype: "premium",
// 	// }
// 	var user model.AccordUser

// 	H.Postgres.Find(&user, "Email=?", l.Email)
// 	if user.Email == l.Email {
// 		fmt.Println("🚀 ~ file: register.go ~ line 42 ~ func ~ http.StatusOK")
// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(`{status: "User already Exist"}`)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	err = H.Postgres.Select("Email", "Password", "Userid", "Username", "Mobile", "Birthdate", "Accounttype").Create(&l).Error

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
//         fmt.Println("❌ ~ file: register.go ~ line 52 ~ func ~ StatusInternalServerError  ")
// 		json.NewEncoder(w).Encode(`{status: "StatusInternalServerError"}`)
// 		return
// 	}
// 	expirationTime := time.Now().Add(time.Hour * 1000)
// 		myClaim := &model.Claims{
// 			UserName:   user.UserName,
// 			Email:      user.Email,
// 			// Userid:     user.Userid,
// 			// Accouttype: user.Accounttype,
// 			StandardClaims: jwt.StandardClaims{
// 				ExpiresAt: expirationTime.Unix(),
// 			},
// 		}
// 		// LOGIN SUCCESSFUL
// 		// token,err := jwt.ParseWithClaims(jwt.SigningMethodHS256,myClaim)
// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

//         fmt.Println("🚀 ~ file: login.go ~ line 51 ~ func ~ token : ", token)
// 		tokenString, err := token.SignedString( []byte(os.Getenv("JWTSECRET")))
        
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
//             fmt.Println("❌ ~ file: login.go ~ line 52 ~ func ~ StatusInternalServerError : " ,err)
// 			json.NewEncoder(w).Encode(`{status: "internal server Error"}`)
// 			return

// 		}else {
// 			http.SetCookie(w, &http.Cookie{
// 				Name:     "Auth1",
// 				Value:    tokenString,
// 				Expires:  expirationTime,
// 				HttpOnly: true,
	
// 				// SameSite: SameSiteLaxMode,
// 			})
// 		}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(`{status: "Register successfull"}`)

// }
