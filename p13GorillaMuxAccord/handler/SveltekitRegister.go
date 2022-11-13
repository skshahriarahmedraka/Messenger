package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	// "context"
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "gorm.io/datatypes"
	// "time"
)

func (H *DatabaseCollections) SveltekitRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	var resError model.ErrorRes

	var user model.AccordUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	user.ID = primitive.NewObjectID()
	user.UUID =uuid.New().String()
	user.ProfileImg = ""
	user.BannerImg = ""
	user.Coin = 0.0
	// user.CoinReq= []model.CoinReq{}
	if user.Email == os.Getenv("ADMIN_EMAIL") {
		user.Accounttype = os.Getenv("ADMIN_ACCOUNT_TYPE")
	} else {
		user.Accounttype = "normal"
	}
	user.TransactionHistory= []string{}
	

	user.FrinedList = []model.FriendStruct{}
	user.GroupList = []model.GroupStruct{}
	user.City=""
	user.Address=""
	user.Country=""
	user.ZipCode=""
	user.UserBio=""
	// x := time.Now().In(location)
	dt,_:= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.AccountCreatedTime =primitive.NewDateTimeFromTime(dt)
    fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 60 ~ func ~ user.AccountCreatedTime : ", user.AccountCreatedTime)
    // fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 60 ~ func ~ user.AccountCreatedTime : ", user.AccountCreatedTime)
	// location,_ := time.LoadLocation("Europe/Rome")
    // x:= user.AccountCreatedTime.In(location)
    // fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 62 ~ func ~ x : ", x)
	fmt.Println("🚀 ~ file: login.go ~ line 44 ~ func ~ user : ", user)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	
	// SEARCH EMAIL
	count, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).CountDocuments(ctx, bson.M{"Email": user.Email})
	if err != nil {
    logerror.ERROR("🚀 ~ file: SveltekitRegister.go ~ line 56 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongodb countDocument email connection error"
		_ = json.NewEncoder(w).Encode(resError)
		return
	}
	if count > 0 {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "User already registered"
        fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 65 ~ func ~ resError.ErrorRes  : ", resError.ErrorRes )
		_ = json.NewEncoder(w).Encode(resError)
		return
	}
	//MOBILE
	count, err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).CountDocuments(ctx, bson.M{"Mobile": user.Mobile})
	if err != nil {
    logerror.ERROR("🚀 ~ file: SveltekitRegister.go ~ line 72 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongodb countDocument mobile connection error"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	if count > 0 {
		// c.JSON(http.StatusBadRequest, gin.H{"error": "mobile already in use"})
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "mobile number already registered"
        fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 83 ~ func ~ resError.ErrorRes : ", resError.ErrorRes)
		_ = json.NewEncoder(w).Encode(resError)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	if err != nil {
		log.Println(err)
	}
	//return string(hash)

	res, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).InsertOne(ctx, user)
	logerror.ERROR("🚀 ~ file: register.go ~ line 102 ~ func ~ err : ", err)

	if err == nil {
		fmt.Println("🚀 ~ file: register.go ~ line 116 ~ func ~ res : ", res)
	}

	// mongoRes, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).InsertOne(ctx, user)
	// fmt.Println("🚀 ~ file: register.go ~ line 80 ~ func ~ mongoRes : ", mongoRes)

	if err != nil {
		logerror.ERROR("🚀 ~ file: register.go ~ line 130 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
		return
	}


	var uMoney model.UserMoney
	uMoney.UUID = user.UUID
	uMoney.Coin = 0.0
	uMoney.ReqList = []model.UserMoneyReq{}
	res, err = H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).InsertOne(ctx, uMoney)
	fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 115 ~ returnfunc ~ res : ", res)
	if err != nil {
		logerror.ERROR("🚀 ~ file: SveltekitRegister.go ~ line 141 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "InsertOne(ctx, uMoney) error"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
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

	expirationTime := time.Now().Add(time.Hour * 10000)

	claims := &model.Claims{
		UserName: user.UserName,
		Email:    user.Email,
		UserID:   user.UserID,
		UUID:     user.UUID,
		Accounttype : user.Accounttype,
		// Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

	/////////////////

	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "Auth1",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// 	HttpOnly:true,

	// 	// SameSite: SameSiteLaxMode,
	// })

	// RefreshExpTime := time.Now().Add(time.Minute * 1000)
	// RefreshClaims := model.RefreshClaims{
	// 	UserName:   user.UserName,
	// 	Email:      user.Email,
	// 	UserID:     user.UserID,
	// 	UUID : user.UUID,
	// 	// Username: credentials.Username,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: RefreshExpTime.Unix(),
	// 	},
	// }

	// refreshtokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshClaims).SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	// if err != nil {
	// 	logerror.ERROR("🚀 ~ file: register.go ~ line 214 ~ func ~ err : ", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	json.NewEncoder(w).Encode(`{status: "StatusInternalServerError"}`)
	// 	return
	// }
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "RefAuth1",
	// 	Value:   refreshtokenStr,
	// 	Expires: RefreshExpTime,
	// 	HttpOnly:true,

	// 	// SameSite: SameSiteLaxMode,
	// })

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(`{status: "Register successfull"}`)

}
