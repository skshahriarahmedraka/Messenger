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

	// "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) UserSendMoney(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	var resError model.ErrorRes
	type TempMessage struct {
		SenderUUID string `json:"SenderUUID" bson:"SenderUUID"`
		ReceiverUUID string `json:"ReceiverUUID" bson:"ReceiverUUID"`
		Amount float64 `json:"Amount" bson:"Amount"`
	}
	
	var ReqUserData TempMessage
	err := json.NewDecoder(r.Body).Decode(&ReqUserData)
	if err != nil {
    logerror.ERROR("🚀 ~ file: UserGetConversationID.go ~ line 25 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "UserGetConversationID error"
		_ = json.NewEncoder(w).Encode(resError)
		return
	}
	
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"UUID", ReqUserData.SenderUUID}}
	update :=  bson.D{{ "$inc" , bson.D{{ "Coin", -ReqUserData.Amount}}}}

	UpdateRes, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logerror.ERROR("🚀 ReqUserData.SenderUUID error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = " error ReqUserData.SenderUUID "
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	fmt.Println("🚀 ~ file: UserSendMoney.go ~ line 49 ~ func ~ UpdateRes : ", UpdateRes)


	opts = options.Update().SetUpsert(true)
	filter = bson.D{{"UUID", ReqUserData.ReceiverUUID}}
	update =  bson.D{{ "$inc" , bson.D{{ "Coin", ReqUserData.Amount}}}}

	UpdateRes, err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logerror.ERROR("🚀 ReqUserData.ReceiverUUID error ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error ReqUserData.ReceiverUUID"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	fmt.Println("🚀 ~ file: UserSendMoney.go ~ line 49 ~ func ~ UpdateRes : ", UpdateRes)
	
	
		// c.JSON(http.StatusOK, UserDataDB)
        // fmt.Println("🚀 ~ file: UserConversationMsg.go ~ line 57 ~ func ~ UserDataDB.Messages : ", UserDataDB.Messages)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(ReqUserData)
		return
	


	// err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UUID": ReqUserData.ReqUUID}).Decode(&UserDataDB)
	// if err != nil {
    // logerror.ERROR("🚀 ~ file: UserProfileData.go ~ line 27 ~ func ~ err : ", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	_ = json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
	// 	return
	// } else {

	// 	for _, v := range UserDataDB.FriendList {
	// 		if v.UUID == ReqUserData.FrndUUID {
	// 			result.ConversationID = v.ConversationID
    //             fmt.Println("🚀 ~ file: UserGetConversationID.go ~ line 55 ~ func ~ result : ", result)
	// 			w.WriteHeader(http.StatusOK)
	// 			_ = json.NewEncoder(w).Encode(result)
	// 			return
	// 		}
	// 	}
		
	// 	// c.JSON(http.StatusOK, UserDataDB)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resError.ErrorRes = "no id found UserGetConversationID error"
	// 	_ = json.NewEncoder(w).Encode(resError)
	// }



}
