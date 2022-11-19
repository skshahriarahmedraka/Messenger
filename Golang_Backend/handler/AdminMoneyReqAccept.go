package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections) AdminMoneyReqAccept(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes

	type ReqDataStruct struct {
		UUID   string  `json:"UserID" bson:"UserID"`
		Amount float64 `json:"Amount" bson:"Amount"`
	}
	var ReqData ReqDataStruct
	err := json.NewDecoder(r.Body).Decode(&ReqData)

	// err := c.BindJSON(&ReqData)
	if err != nil {
		logerror.ERROR("🚀 ~ file: AdminMoneyReqAccept.go ~ line 30 ~ func ~ err : ", err)
		resError.ErrorRes = "bind json error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
		// logerror.ERROR("bind json error", err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "binding error"})
		// return
	}
	if ReqData.UUID == "" || ReqData.Amount == 0.0 {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
		resError.ErrorRes = "UUID or Amount is empty"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: AdminMoneyReqAccept.go ~ line 47 ~ func ~ err : ", err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid or  amount is empty"})
		// return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	count, err := H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).CountDocuments(ctx, bson.M{"UUID": ReqData.UUID})
	fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 55 ~ func ~ err : ", err)
	if err != nil {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
		resError.ErrorRes = "mongodb Count Document error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}
	if count > 0 {

		// var SendJwt model.UserMoneyReq

		// SendJwt.JWT = "Your Request is Pending for Admin Acceptance"
		// SendJwt.Amount = ReqData.Amount
		// SendJwt.Status = false

		//update := bson.D{{"$push", ReqComment.Comment}}

		expirationTime := time.Now().Add(time.Hour * 100000)
		myClaim := &model.JWTModel{
			UUID:   ReqData.UUID,
			Amount: ReqData.Amount,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_MONEY")))
		logerror.ERROR("🚀 ~ file: AdminMoneyReqAccept.go ~ line 89 ~ func ~ err : ", err)

		if err != nil {
			resError.ErrorRes = "mongodb Count Document error"
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resError)
			logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
			return
			// logerror.ERROR("🚀 ~ file: AdminMoneyReqAccept.go ~ line 92 ~ func ~ err : ", err)
			// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			// 	return
		}
		//inserting token in db
		// opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.D{{"UUID", ReqData.UUID}}
		// update := bson.D{{"$set", bson.D{{"ReqList.$.JWT",  tokenString}}}}
		//update2 := bson.D{{
		//	"$addToSet", bson.D{{"Comment", bson.D{{
		//		"$each": []model.CommentModel{ReqComment.Comment}}}}}}}
		var resData model.UserMoney

		err = H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOne(ctx, filter).Decode(&resData)
		logerror.ERROR("🚀 ~ file: AdminMoneyReqAccept.go ~ line 110 ~ func ~ err : ", err)
		if err != nil {
			resError.ErrorRes = "mongodb Count Document error"
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resError)
			logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
			return
		}

		//var ReqListData []model.UserMoneyReq
		//ReqListData = resData.ReqList
		for i, v := range resData.ReqList {
			if v.Amount == ReqData.Amount && v.Status == "pending" {
				resData.ReqList[i].JWT = tokenString
				resData.ReqList[i].Status = "accepted"
				break
			}
		}
		res, err := H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).ReplaceOne(ctx, filter, resData)
		fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 130 ~ func ~ res : ", res)
		if err != nil {
			resError.ErrorRes = "mongodb Count Document error"
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resError)
			logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)

			// logerror.ERROR("mongodb find one and update error", err)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(res)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)

		// c.JSON(http.StatusOK, res)

	} else {
		logerror.ERROR("❌🔥 User is not found ", err)
		resError.ErrorRes = "mongodb Count Document error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)

		// c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		return

	}

}
