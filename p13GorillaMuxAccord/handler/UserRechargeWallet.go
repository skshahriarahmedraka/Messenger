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

	// "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) UserRechargeWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes
	type ReqStruct struct {
		UUID string `json:"UUID" bson:"UUID"`
		JWT  string `json:"JWT" bson:"JWT"`
	}

	var ReqUserData ReqStruct
	err := json.NewDecoder(r.Body).Decode(&ReqUserData)
	// err := c.BindJSON(&ReqUserData)
	if err != nil {
		logerror.ERROR("bind json error", err)
		resError.ErrorRes = "bind json error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("❌🔥 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "binding error"})
		// return
	}
	fmt.Println("❌🔥 ~ file: UserRechargeWallet.go ~ line 29 ~ func ~ ReqUserData : ", ReqUserData)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).CountDocuments(ctx, bson.M{"UUID": ReqUserData.UUID})
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserRechargeWallet.go ~ line 43 ~ func ~ err : ", err)
		resError.ErrorRes = "Count Document error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("❌🔥 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb Count Document error"})
		// return
	}
	if count > 0 {
		filter := bson.D{{"UUID", ReqUserData.UUID}}
		// update := bson.D{{"$set", bson.D{{"ReqList.$.JWT",  tokenString}}}}
		//update2 := bson.D{{
		//	"$addToSet", bson.D{{"Comment", bson.D{{
		//		"$each": []model.CommentModel{ReqComment.Comment}}}}}}}
		var resData model.UserMoney

		err = H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOne(ctx, filter).Decode(&resData)
		fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 90 ~ returnfunc ~ resData : ", resData)
		if err != nil {
			logerror.ERROR("❌🔥 ~ file: UserRechargeWallet.go ~ line 63 ~ func ~ err : ", err)
			resError.ErrorRes = "mongodb find one  error"
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resError)
			logerror.ERROR("❌🔥 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
			return
			// LogError.LogError("mongodb find one and update error", err)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
			// return
		}

		//var ReqListData []model.UserMoneyReq
		//ReqListData = resData.ReqList
		for i, v := range resData.ReqList {
			if v.JWT == ReqUserData.JWT && v.Status == "accepted" {
				resData.Coin += v.Amount
				resData.ReqList[i].Status = "used"

				opts := options.Update().SetUpsert(true)
				filter := bson.D{{"UUID", ReqUserData.UUID}}
				update := bson.D{
					{"$inc", bson.D{{"Coin", v.Amount}}},
				}
				res, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
				fmt.Println("🚀 ~ file: UserRechargeWallet.go ~ line 72 ~ returnfunc ~ res : ", res)
				if err != nil {
					logerror.ERROR("🚀 ~ file: UserRechargeWallet.go ~ line 63 ~ func ~ err : ", err)
					resError.ErrorRes = "mongodb  update one error"
					w.WriteHeader(http.StatusInternalServerError)
					err = json.NewEncoder(w).Encode(resError)
					logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
					return
					// LogError.LogError("mongodb update one error", err)
					// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb update one error"})
					// return

				}
				break
			}
		}
		res, err := H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).ReplaceOne(ctx, filter, resData)
		fmt.Println("🚀 ~ file: AdminMoneyReqAccept.go ~ line 107 ~ returnfunc ~ res : ", res)
		if err != nil {
			logerror.ERROR("🚀 ~ file: UserRechargeWallet.go ~ line 63 ~ func ~ err : ", err)
			resError.ErrorRes = "mongodb find one and update error"
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resError)
			logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
			return
			// LogError.LogError("mongodb find one and update error", err)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
			// return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(res)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)
		// c.JSON(http.StatusOK, res)

	} else {
		logerror.ERROR("❌🔥 User is not found ", err)
		resError.ErrorRes = "mongodb count not > 0 error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": "User is not found"})
		// return

	}
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(`{status: "User public Profile successfull"}`)

}
