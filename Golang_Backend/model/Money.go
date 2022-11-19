package model

import "github.com/golang-jwt/jwt/v4"


type UserMoney struct {
	UUID  string         `json:"UUID" bson:"UUID"`
	Coin    float64        `json:"Coin" bson:"Coin"`
	ReqList []UserMoneyReq `json:"ReqList" bson:"ReqList"`
}
type UserMoneyReq struct {
	Amount float64 `json:"Amount" bson:"Amount"`
	JWT    string  `json:"JWT" bson:"JWT"`
	Status string  `json:"Status" bson:"Status"` // true = already used token, false = Not used token
}

type JWTModel struct {
	UUID string  `json:"UUID" bson:"UUID"`
	Amount float64 `json:"Amount" bson:"Amount"`
	jwt.StandardClaims
}
