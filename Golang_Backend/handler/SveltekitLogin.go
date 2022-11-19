package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"

	//"os"
	"time"

	// "log"
	// "context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	// "github.com/dgrijalva/jwt-go"
	// "gorm.io/datatypes"
	// "time"
)

func (H *DatabaseCollections) SveltekitLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	var user model.LoginModel
	var dbUser model.AccordUser
	// err := json.NewDecoder(r.Body).Decode(&user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ l : ", user.Password)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ user : ", user)

	var resError model.ErrorRes

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// SEARCH EMAIL
	// count, err := H.Mongo.Collection("usercol").CountDocuments(ctx, bson.M{"Email": user.Email})
	err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"Email": user.Email}).Decode(&dbUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongodb FindOne email connection error"
		_ = json.NewEncoder(w).Encode(resError)
		return
	}

	// if err!=nil{
	// 	response.WriteHeader(http.StatusInternalServerError)
	// 	response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	// 	return
	//  }
	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		logerror.ERROR("🚀 ~ file: SveltekitLogin.go ~ line 71 ~ func ~ passErr : ", passErr)
		resError.ErrorRes = " bcrypt error : " + passErr.Error()
		_ = json.NewEncoder(w).Encode(resError)
		return
	}
	//  jwtToken, err := GenerateJWT()
	//  if err != nil{
	//  response.WriteHeader(http.StatusInternalServerError)
	//  response.Write([]byte(`{"message":"`+err.Error()+`"}`))
	//  return
	//  }
	//  response.Write([]byte(`{"token":"`+jwtToken+`"}`))
	// if user.ID != 0  {
	//CREATE COOKIE
	expirationTime := time.Now().Add(time.Hour * 10000)
	myClaim := &model.Claims{
		UserName: dbUser.UserName,
		Email:    dbUser.Email,
		UserID:   dbUser.UserID,
		UUID:     dbUser.UUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// }
	// LOGIN SUCCESSFUL

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(`{status: "StatusInternalServerError"}`)
		return
	}
	////////////////////////////////
	var SendJwt struct {
		JWT string
	}
	SendJwt.JWT = tokenString
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&SendJwt)

}

// CHECK IF ONLY EMAIL EXIST
//H.Postgres.Find(&user, "Email=?", l.Email)
//if user.ID != 0 {
//	w.WriteHeader(http.StatusUnauthorized)
//    fmt.Println("❌ ~ file: login.go ~ line 71 ~ func ~ StatusUnauthorized : ")
//
//	json.NewEncoder(w).Encode(`{message: "Password Incorrent"}`)
//
//} else {
//	w.WriteHeader(http.StatusUnauthorized)
//    fmt.Println("❌ ~ file: login.go ~ line 79 ~ func ~ StatusUnauthorized : ")
//	json.NewEncoder(w).Encode(`{status: "Email/Password incorrect"}`)
//}
// fmt.Println("🚀 ~ file: login.go ~ line 34 ~ func ~ user full data : ", user)
// if user.ID ==0 {
// }
