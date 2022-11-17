package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) SendFriendReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes

	var ReqData model.UserFrndReqSend

	err := json.NewDecoder(r.Body).Decode(&ReqData)
	fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 27 ~ func ~ ReqData : ", ReqData)
	// err := c.BindJSON(&ReqData)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserTokenReqList.go ~ line 30 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// logerror.ERROR("❌🔥 error in c.BindJSON(&ReqData) ", err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		// return
	}
	if ReqData.UUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid "})
		// return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{{"UUID", ReqData.UUID}}
	var result model.AccordUser
	err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, filter).Decode(&result)
	fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 52 ~ func ~ result : ", result)
	if err != nil {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
		resError.ErrorRes = "mongodb FindOne(ctx, filter) error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
	}
	var FrndReqShort model.FrndReqShort
	FrndReqShort.UUID = result.UUID
	FrndReqShort.UserID = result.UserID
	FrndReqShort.ProfileImg = result.ProfileImg
	FrndReqShort.UserBio = result.UserBio
	FrndReqShort.UserName = result.UserName
	opts := options.Update().SetUpsert(true)
	filter = bson.D{{"UUID", ReqData.Frnd}}
	update := bson.M{"$push": bson.M{
		"FrndReq": FrndReqShort}}
	fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 72 ~ func ~ FrndReqShort : ", FrndReqShort)
	res, err := H.MongoUser.Collection(os.Getenv("MONGO_FRND_REQ_COL")).UpdateOne(ctx, filter, update, opts)

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

	// update your own pending request list
	filter = bson.D{{"UUID", ReqData.Frnd}}
	// var result model.AccordUser
	err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, filter).Decode(&result)
    fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 93 ~ func ~ result : ", result)
	if err != nil {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
		resError.ErrorRes = "mongodb FindOne(ctx, filter) error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
	}
	// var FrndReqShort model.FrndReqShort
	FrndReqShort.UUID = result.UUID
	FrndReqShort.UserID = result.UserID
	FrndReqShort.ProfileImg = result.ProfileImg
	FrndReqShort.UserBio = result.UserBio
	FrndReqShort.UserName = result.UserName
	opts = options.Update().SetUpsert(true)
	filter = bson.D{{"UUID", ReqData.UUID}}
	update = bson.M{"$push": bson.M{
		"FrndReqPending": FrndReqShort}}
	fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 72 ~ func ~ FrndReqShort : ", FrndReqShort)
	res, err = H.MongoUser.Collection(os.Getenv("MONGO_FRND_REQ_COL")).UpdateOne(ctx, filter, update, opts)
	// fmt.Println("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ results : ", results)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserSendFriendReq.go ~ line 96 ~ func ~ err : ", err)
		resError.ErrorRes = "mongodb  update one error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
		// LogError.LogError("mongodb update one error", err)
		// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb update one error"})
		// return
		
	}
	fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 114 ~ func ~ res : ", res)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&FrndReqShort)
	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)

}
