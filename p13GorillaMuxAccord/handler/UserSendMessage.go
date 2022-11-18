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
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) UserSendMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	var resError model.ErrorRes
	type TempMessage struct {
		ConversationID string `json:"ConversationID" bson:"ConversationID"`
		SenderID string `json:"SenderID" bson:"SenderID"`
		SenderName string `json:"SenderName" bson:"SenderName"`
		Message string `json:"Message" bson:"Message"`
		Reactions []int `json:"Reactions" bson:"Reactions"`
		UserReaction []model.ReactionStruct `json:"UserReaction" bson:"UserReaction"`
		Timestamp string `json:"Timestamp" bson:"Timestamp"`
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
	var MyMsg model.Message
	MyMsg.SenderID = ReqUserData.SenderID
	MyMsg.SenderName = ReqUserData.SenderName
	MyMsg.Message = ReqUserData.Message
	MyMsg.Reactions = ReqUserData.Reactions
	MyMsg.UserReaction = ReqUserData.UserReaction
	MyMsg.Timestamp = ReqUserData.Timestamp



	// var UserDataDB model.Conversation
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"ConversationID", ReqUserData.ConversationID}}
	update := bson.M{"$push": bson.M{
		"Messages": MyMsg}}
	// fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 72 ~ func ~ FrndReqShort : ", FrndReqShort)
	res, err := H.MongoUser.Collection(os.Getenv("MONGO_USER_MSG_COL")).UpdateOne(ctx, filter, update, opts)
    fmt.Println("🚀 ~ file: UserSendMessage.go ~ line 62 ~ func ~ res : ", res)


	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
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
