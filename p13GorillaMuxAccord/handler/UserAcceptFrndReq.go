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

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) UserAcceptFrndReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes

	type ReqStruct struct {
		Acceptor  string `json:"Acceptor" bson:"Acceptor"`
		Requestor string `json:"Requestor" bson:"Requestor"`
	}
	var ReqData ReqStruct

	err := json.NewDecoder(r.Body).Decode(&ReqData)
	fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 31 ~ func ~ ReqData : ", ReqData)
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
	if ReqData.Acceptor == "" || ReqData.Requestor == "" {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid "})
		// return
	}

	// CREATE CONVERSATION COLLECTION
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// filter := bson.D{{"UUID", ReqData.Acceptor}}
	var newConversation model.Conversation
	newConversation.ID = primitive.NewObjectID()
	newConversation.GroupID = uuid.New().String()
	newConversation.ConversationID = uuid.New().String()
	newConversation.Members = append(newConversation.Members, ReqData.Acceptor, ReqData.Requestor)
	// newConversation.Messages =
	// newConversation.TotalMessages =

	// opts := options.FindOne().SetProjection(bson.M{"FrinedList":1})
	res, err := H.MongoUser.Collection(os.Getenv("MONGO_USER_MSG_COL")).InsertOne(ctx, newConversation)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserAcceptFrndReq.go ~ line 69 ~ func ~ err Collection(os.Getenv(MONGO_USER_MSG_COL)).InsertOne(ctx,newConversation): ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error ; Collection(os.Getenv(MONGO_USER_MSG_COL)).InsertOne(ctx,newConversation)"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// logerror.ERROR("❌🔥 error in c.BindJSON(&ReqData) ", err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		// return
	}
	fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 57 ~ func ~ res : ", res)

	// INSERT FRIEND IN USER DATA
	var FriendStruct model.FriendStruct
	FriendStruct.UUID = ReqData.Requestor
	FriendStruct.ConversationID = newConversation.ConversationID

	filter := bson.D{{"UUID", ReqData.Acceptor}}
	update := bson.M{"$push": bson.M{
		"FrinedList": FriendStruct,
	},
	}
	opts := options.Update().SetUpsert(true)
	UpdateRes, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserAcceptFrndReq.go ~ line 94 ~ func ~ err Collection(os.Getenv(MONGO_USERCOL)).UpdateOne(ctx, filter, update, opts): ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error Collection(os.Getenv(MONGO_USERCOL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 83 ~ func ~ UpdateRes : ", UpdateRes)

	// INSERT FRIEND IN USER DATA
	// var FriendStruct model.FriendStruct
	FriendStruct.UUID = ReqData.Acceptor
	FriendStruct.ConversationID = newConversation.ConversationID

	filter = bson.D{{"UUID", ReqData.Requestor}}
	update = bson.M{"$push": bson.M{
		"FrinedList": FriendStruct,
	},
	}
	opts = options.Update().SetUpsert(true)
	UpdateRes, err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserAcceptFrndReq.go ~ line 118 ~ func ~ err :.Collection(os.Getenv(MONGO_USERCOL)).UpdateOne(ctx, filter, update, opts) ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error .Collection(os.Getenv(MONGO_USERCOL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 98 ~ func ~ UpdateRes : ", UpdateRes)

	// PULL OUT THE FRIEND REQUEST

	filter = bson.D{{"UUID", ReqData.Acceptor}}
	update = bson.M{"$pull": bson.M{
		"FrndReq": bson.M{"UUID": ReqData.Requestor},
	},
	}

	UpdateRes, err = H.MongoUser.Collection(os.Getenv("MONGO_FRND_REQ_COL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
    logerror.ERROR("🚀 ~ file: UserAcceptFrndReq.go ~ line 134 ~ func ~ err Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts) : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 109 ~ func ~ UpdateRes : ", UpdateRes)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&UpdateRes)
	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)
	// ////////////////////////////////////////////////////////////////////////////
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// filter := bson.D{{"UUID", ReqData.Acceptor}}
	// var result []model.FriendStruct
	// opts := options.FindOne().SetProjection(bson.M{"FrinedList":1})
	// err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, filter,opts).Decode(&result)
	// fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 57 ~ func ~ result : ", result)

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
	// FrndReqShort.UserBio = result.UserBio
	// FrndReqShort.UserName = result.UserName
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
	// w.WriteHeader(http.StatusOK)
	// err = json.NewEncoder(w).Encode(&FrndReqShort)
	// logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)

}
