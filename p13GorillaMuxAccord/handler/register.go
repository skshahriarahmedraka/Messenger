package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"fmt"
	"os"
	"time"

	// "context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "gorm.io/datatypes"
	// "time"
)

func (H *DatabaseCollections) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var resError model.ErrorRes

	var user model.AccordUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ user : ", user)
	

	user.Accounttype="normal"
	user.BannerImg=""
	user.Coin=0.0
	user.FrinedListID=[]string{}
	user.ProfileImg=""
	myid:= uuid.New()
	user.UUID=myid.String()

	

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// SEARCH EMAIL
	count, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).CountDocuments(ctx, bson.M{"Email": user.Email})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongodb countDocument email connection error"
		json.NewEncoder(w).Encode(resError)
		return
	}
	if count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "User already registered"
		json.NewEncoder(w).Encode(resError)
		return
	}
	//MOBILE
	count, err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).CountDocuments(ctx, bson.M{"Mobile": user.Mobile})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongodb countDocument mobile connection error"
		json.NewEncoder(w).Encode(resError)
		return

	}
	if count > 0 {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "mobile already in use"})
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "mobile number already registered"
		json.NewEncoder(w).Encode(resError)
		return
	}



	res, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).InsertOne(ctx, user)
	logerror.ERROR("🚀 ~ file: register.go ~ line 102 ~ func ~ err : ", err)

	if err == nil {
		fmt.Println("🚀 ~ file: register.go ~ line 116 ~ func ~ res : ", res)
	}

	mongoRes, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).InsertOne(ctx, user)
	fmt.Println("🚀 ~ file: register.go ~ line 80 ~ func ~ mongoRes : ", mongoRes)

	if err != nil {
    logerror.ERROR("🚀 ~ file: register.go ~ line 130 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
		return
	}
	// expirationTime := time.Now().Add(time.Hour * 1000)
	// myClaim := &model.Claims{
	// 	Username:   user.Username,
	// 	Email:      user.Email,
	// 	Userid:     user.Userid,
	// 	Accouttype: user.Accounttype,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }
	// // LOGIN SUCCESSFUL
	// // token,err := jwt.ParseWithClaims(jwt.SigningMethodHS256,myClaim)
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	// fmt.Println("🚀 ~ file: login.go ~ line 51 ~ func ~ token : ", token)
	// tokenString, err := token.SignedString([]byte(os.Getenv("JWTSECRET")))

	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Println("❌ ~ file: login.go ~ line 52 ~ func ~ StatusInternalServerError : ", err)
	// 	json.NewEncoder(w).Encode(`{status: "internal server Error"}`)
	// 	return

	// }
	// http.SetCookie(w, &http.Cookie{
	// 	Name:     "Auth1",
	// 	Value:    tokenString,
	// 	Expires:  expirationTime,
	// 	HttpOnly: true,

	// 	// SameSite: SameSiteLaxMode,
	// })

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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Auth1",
		Value:   tokenString,
		Expires: expirationTime,
		HttpOnly:true,
		
		// SameSite: SameSiteLaxMode,
	})


	RefreshExpTime := time.Now().Add(time.Minute * 1000)
	RefreshClaims := model.RefreshClaims{
		UserName:   user.UserName,
		Email:      user.Email,
		UserID:     user.UserID,
		UUID : user.UUID,
		// Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: RefreshExpTime.Unix(),
		},
	}


	refreshtokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims).SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		logerror.ERROR("🚀 ~ file: register.go ~ line 214 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(`{status: "StatusInternalServerError"}`)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "RefAuth1",
		Value:   refreshtokenStr,
		Expires: RefreshExpTime,
		HttpOnly:true,
		
		// SameSite: SameSiteLaxMode,
	})

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(`{status: "Register successfull"}`)

}
