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

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) UserTokenReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes

	type ReqStruct struct {
		UUID   string  `json:"UserID" bson:"UserID"`
		Amount float64 `json:"Amount" bson:"Amount"`
	}
	var ReqData ReqStruct
	err := json.NewDecoder(r.Body).Decode(&ReqData)
	// err := c.BindJSON(&ReqData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return
		
	}
	if ReqData.UUID == "" || ReqData.Amount == 0.0 {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "bad request userid or  amount is empty"
		_ = json.NewEncoder(w).Encode(resError)
		return
	}
	fmt.Println("🚀 ~ file: userTokenReq.go ~ line 36 ~ returnfunc ~ ReqData : ", ReqData)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	count, err := H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).CountDocuments(ctx, bson.M{"UUID": ReqData.UUID})
    fmt.Println("🚀 ~ file: UserTokenReq.go ~ line 46 ~ func ~ count : ", count)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserTokenReq.go ~ line 49 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "CountDocuments(ctx, bson.M{UserID: ReqData.UserID})"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	if count > 0 {

		var SendJwt model.UserMoneyReq

		SendJwt.JWT = "Your Request is Pending for Admin Acceptance"
		SendJwt.Amount = ReqData.Amount
		SendJwt.Status = "pending"

		opts := options.FindOneAndUpdate().SetUpsert(true)
		filter := bson.D{{"UUID", ReqData.UUID}}
		update := bson.M{"$push": bson.M{
			"ReqList": SendJwt,
		},
		}
	
		var UpdatedDocument model.UserMoney
		err = H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).FindOneAndUpdate(ctx, filter, update, opts).Decode(&UpdatedDocument)
		if err != nil {
			logerror.ERROR("🚀 ~ file: UserTokenReq.go ~ line 80 ~ func ~ err : ", err)
			w.WriteHeader(http.StatusInternalServerError)
			resError.ErrorRes = "error FindOneAndUpdate(ctx, filter, update, opts).Decode(&UpdatedDocument)"
			_ = json.NewEncoder(w).Encode(resError)
			return
		}
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(SendJwt)
		logerror.ERROR("🚀 ~ file: userProfileUpdate.go ~ line 74 ~ func ~ err : ", err)
		return

	} else {
		logerror.ERROR("🚀 ~ file: UserTokenReq.go ~ line 80 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongo count error"
		_ = json.NewEncoder(w).Encode(resError)
		return
	}
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(`{status: "User public Profile successfull"}`)

}
