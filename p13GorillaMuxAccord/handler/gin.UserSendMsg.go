package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"os"
	"time"

	// "app/model"
	// "context"
	// "encoding/json"
	"fmt"
	// "net/http"
	// "os"
	// "time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "github.com/gorilla/websocket"
	// "go.mongodb.org/mongo-driver/bson"
)



func (H *DatabaseCollections) GinUserSendMsg() gin.HandlerFunc {
	return func(c *gin.Context) {
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()

		// for{
		// 	// Read message from browser
		// 	type UserInfo struct {
		// 		UUID string `json:"UUID" bson:"UUID"`
		// 	}
		// 	var ReqData UserInfo
		// 	err := ws.ReadJSON(&ReqData)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}

		// }

		for {
			//Read Message from client
			
			type TempMessage struct {
				ConversationID string `json:"ConversationID" bson:"ConversationID"`
				SenderID string `json:"SenderID" bson:"SenderID"`
				SenderName string `json:"SenderName" bson:"SenderName"`
				Message string `json:"Message" bson:"Message"`
				Reactions []int `json:"Reactions" bson:"Reactions"`
				UserReaction []model.ReactionStruct `json:"UserReaction" bson:"UserReaction"`
				Timestamp string `json:"Timestamp" bson:"Timestamp"`
			}
			var ReqData TempMessage

			err := ws.ReadJSON(&ReqData)
			fmt.Println("🚀 ~ file: gin.websocket.go ~ line 56 ~ returnfunc ~ ReqData : ", ReqData)
			// mt, message, err := ws.ReadMessage()
			if err != nil {
				logerror.ERROR("🚀 ~ file: Instantbuy.go ~ line 58 ~ func ~ err : ", err)
				// fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			// var resError model.ErrorRes
	

	var MyMsg model.Message
	MyMsg.SenderID = ReqData.SenderID
	MyMsg.SenderName = ReqData.SenderName
	MyMsg.Message = ReqData.Message
	MyMsg.Reactions = ReqData.Reactions
	MyMsg.UserReaction = ReqData.UserReaction
	MyMsg.Timestamp = ReqData.Timestamp



	// var UserDataDB model.Conversation
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"ConversationID", ReqData.ConversationID}}
	update := bson.M{"$push": bson.M{
		"Messages": MyMsg}}
	// fmt.Println("🚀 ~ file: UserSendFriendReq.go ~ line 72 ~ func ~ FrndReqShort : ", FrndReqShort)
	res, err := H.MongoUser.Collection(os.Getenv("MONGO_USER_MSG_COL")).UpdateOne(ctx, filter, update, opts)
    fmt.Println("🚀 ~ file: UserSendMessage.go ~ line 62 ~ func ~ res : ", res)


	

			 
			err = ws.WriteJSON(ReqData)
			if err != nil {
				fmt.Println("🚀 ~ file: gin.websocket.go ~ line 73 ~ returnfunc ~ err : ", err)
				break
			}
		}

		// for {

		// 	mt, message, err := ws.ReadMessage()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}
		// 	//If client message is ping will return pong
		// 	if string(message) == "ping" {
		// 		message = []byte("pong")
		// 	}
		// 	//Response message to client
		// 	err = ws.WriteMessage(mt, message)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		break
		// 	}
		// }
	}
}

// func (H *DatabaseCollections)UpdateFrndChatShort(ReqUUID string) {
// 	var resError model.ErrorRes
// 	var UserDataDB model.UserChatFrndList
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
// 	defer cancel()
// 	err := H.MongoUser.Collection(os.Getenv("MONGO_FRND_CHATSHORT_COL")).FindOne(ctx, bson.M{"UUID": ReqUUID}).Decode(&UserDataDB)
//     logerror.ERROR("🚀 ~ file: UserFriendReqList.go ~ line 34 ~ func ~ err : ", err)

// }
