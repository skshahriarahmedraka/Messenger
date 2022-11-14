package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"encoding/json"
	// "fmt"

	// "fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) UserFriendReqList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes


	paramsMap := mux.Vars(r)
	ReqUUID := paramsMap["UUID"]

	var UserDataDB model.FrndReqList
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err := H.MongoUser.Collection(os.Getenv("MONGO_FRND_REQ_COL")).FindOne(ctx, bson.M{"UUID": ReqUUID}).Decode(&UserDataDB)

	if err != nil {
    logerror.ERROR("🚀 ~ file: UserFriendReqList.go ~ line 34 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// logerror.ERROR("❌🔥 error in c.BindJSON(&ReqData) ", err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		// return
	}else {
		w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&UserDataDB.FrndReq)
	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)
	}

	
	// var resError model.ErrorRes

	// var ReqData model.UserFrndReqSend

	// err := json.NewDecoder(r.Body).Decode(&ReqData)
    // fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 27 ~ func ~ ReqData : ", ReqData)
	// // err := c.BindJSON(&ReqData)
	// if ReqData.UUID == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	resError.ErrorRes = "http.StatusBadRequest"
	// 	_ = json.NewEncoder(w).Encode(resError)
	// 	return
	// 	// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid "})
	// 	// return
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// filter := bson.D{{"UUID", ReqData.UUID}}
	// var result model.AccordUser
	// err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, filter).Decode(&result)
    // fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 52 ~ func ~ result : ", result)
	// if err != nil {
	// 	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
	// 	resError.ErrorRes = "mongodb FindOne(ctx, filter) error"
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	err = json.NewEncoder(w).Encode(resError)
	// 	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
	// 	return
	// }
	// var FrndReqShort model.FrndReqShort
	// FrndReqShort.UUID = result.UUID
	// FrndReqShort.UserID = result.UserID
	// FrndReqShort.ProfileImg = result.ProfileImg
	// FrndReqShort.UserName = result.UserName
	// opts := options.Update().SetUpsert(true)
	// filter = bson.D{{"UUID", ReqData.Frnd}}
	// update := bson.M{"$push": bson.M{
	// 	"FrndReq": FrndReqShort,
	// },
	// }
	// res, err := H.MongoUser.Collection(os.Getenv("MONGO_FRND_REQ_COL")).UpdateOne(ctx, filter, update, opts)

	// fmt.Println("🚀 ~ file: UserRechargeWallet.go ~ line 72 ~ returnfunc ~ res : ", res)

	// if err != nil {
	// 	logerror.ERROR("🚀 ~ file: UserRechargeWallet.go ~ line 63 ~ func ~ err : ", err)
	// 	resError.ErrorRes = "mongodb  update one error"
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	err = json.NewEncoder(w).Encode(resError)
	// 	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
	// 	return
	// 	// LogError.LogError("mongodb update one error", err)
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb update one error"})
	// 	// return

	// }

	

	// fmt.Println("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ results : ", results)
	

}
