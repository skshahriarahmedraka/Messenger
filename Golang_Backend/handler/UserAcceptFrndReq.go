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
	newConversation.Messages=[]model.Message{}
	newConversation.TotalMessages=0
	// newConversation.Messages =
	// newConversation.TotalMessages =

	// opts := options.FindOne().SetProjection(bson.M{"FriendList":1})
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
		"FriendList": FriendStruct,
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
		"FriendList": FriendStruct,
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
		logerror.ERROR("🚀ReqData.Acceptor ~ file: UserAcceptFrndReq.go ~ line 134 ~ func ~ err Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts) : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}

	
	fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 133 ~ func ~ UpdateRes : ", UpdateRes)
	filter = bson.D{{"UUID", ReqData.Requestor}}
	update = bson.M{"$pull": bson.M{
		"FrndReqPending": bson.M{"UUID": ReqData.Acceptor},
	},
	}

	UpdateRes, err = H.MongoUser.Collection(os.Getenv("MONGO_FRND_REQ_COL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		logerror.ERROR("🚀 ReqData.Requestor ~ file: UserAcceptFrndReq.go ~ line 134 ~ func ~ err Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts) : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	fmt.Println("🚀ReqData.Requestor ~ file: UserAcceptFrndReq.go ~ line 109 ~ func ~ UpdateRes : ", UpdateRes)

	// GET ACCEPTOR DATA
	var AcceptorData model.AccordUser
	err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UUID": ReqData.Acceptor}).Decode(&AcceptorData)
	if err != nil {
		logerror.ERROR("🚀 AcceptorData ~ file: UserAcceptFrndReq.go ~ line 134 ~ func ~ err Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts) : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
    fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 248 ~ func ~ AcceptorData : ", AcceptorData)
	// GET REQUESOR  DATA
	var RequestorData model.AccordUser
	err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UUID": ReqData.Requestor}).Decode(&RequestorData)
	if err != nil {
		logerror.ERROR("🚀 RequestorData ~ file: UserAcceptFrndReq.go ~ line 134 ~ func ~ err Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts) : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "error Collection(os.Getenv(MONGO_FRND_REQ_COL)).UpdateOne(ctx, filter, update, opts)"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
    fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 262 ~ func ~ RequestorData : ", RequestorData)

	
	// create chat short 
	var FrndChatShort model.FrndChatShort
	FrndChatShort.UUID = AcceptorData.UUID
	FrndChatShort.UserID = AcceptorData.UserID
	FrndChatShort.UserName = AcceptorData.UserName
	FrndChatShort.ActiveStatus =false
	FrndChatShort.UserImg=AcceptorData.ProfileImg
	FrndChatShort.LastMessage = ""
	FrndChatShort.LastMessageTime = time.Now().String()
	FrndChatShort.LastActiveTime = time.Now().String()
	FrndChatShort.NumberOfNotification=0
	FrndChatShort.SilentNotification=false
	FrndChatShort.ConversationID=FriendStruct.ConversationID


	opts2 := options.FindOneAndUpdate().SetUpsert(true)

	filter = bson.D{{"UUID", RequestorData.UUID}}
	update = bson.M{"$push": bson.M{
		"FrndChatList": FrndChatShort,
	},
	}

	res2 := H.MongoUser.Collection(os.Getenv("MONGO_FRND_CHATSHORT_COL")).FindOneAndUpdate(ctx, filter, update, opts2)
    fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 219 ~ func ~ res2 : ", res2)
	
	// create chat short 
	// var FrndChatShort model.FrndChatShort
	FrndChatShort.UUID = RequestorData.UUID
	FrndChatShort.UserID = RequestorData.UserID
	FrndChatShort.UserName = RequestorData.UserName
	FrndChatShort.ActiveStatus =false
	FrndChatShort.UserImg=RequestorData.ProfileImg
	FrndChatShort.LastMessage = ""
	FrndChatShort.LastMessageTime = time.Now().String()
	FrndChatShort.LastActiveTime = time.Now().String()
	FrndChatShort.NumberOfNotification=0
	FrndChatShort.SilentNotification=false
	FrndChatShort.ConversationID=FriendStruct.ConversationID


	opts2 = options.FindOneAndUpdate().SetUpsert(true)

	filter = bson.D{{"UUID", AcceptorData.UUID}}
	update = bson.M{"$push": bson.M{
		"FrndChatList": FrndChatShort,
	},
	}

	res2 = H.MongoUser.Collection(os.Getenv("MONGO_FRND_CHATSHORT_COL")).FindOneAndUpdate(ctx, filter, update, opts2)
    fmt.Println("🚀 ~ file: UserAcceptFrndReq.go ~ line 219 ~ func ~ res2 : ", res2)
	

	// opts := options.FindOne().SetProjection(bson.M{"FriendList":1})
	// res, err := H.MongoUser.Collection(os.Getenv("MONGO_USER_MSG_COL")).InsertOne(ctx, newConversation)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&UpdateRes)
	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)


}
