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
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) GetConversationID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	var resError model.ErrorRes
	type GetConversationIDModel struct {
		ReqUUID string `json:"ReqUUID"`
		FrndUUID string `json:"FrndUUID"`
	}
	type ResModel struct {
		ConversationID string `json:"ConversationID"`
	}
	var result ResModel
	var ReqUserData GetConversationIDModel
	err := json.NewDecoder(r.Body).Decode(&ReqUserData)
	if err != nil {
    logerror.ERROR("🚀 ~ file: UserGetConversationID.go ~ line 25 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "UserGetConversationID error"
		_ = json.NewEncoder(w).Encode(resError)
		return
	}


	var UserDataDB model.AccordUser
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UUID": ReqUserData.ReqUUID}).Decode(&UserDataDB)
	if err != nil {
    logerror.ERROR("🚀 ~ file: UserProfileData.go ~ line 27 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
		return
	} else {

		for _, v := range UserDataDB.FriendList {
			if v.UUID == ReqUserData.FrndUUID {
				result.ConversationID = v.ConversationID
                fmt.Println("🚀 ~ file: UserGetConversationID.go ~ line 55 ~ func ~ result : ", result)
				w.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(w).Encode(result)
				return
			}
		}
		
		// c.JSON(http.StatusOK, UserDataDB)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "no id found UserGetConversationID error"
		_ = json.NewEncoder(w).Encode(resError)
	}



}
