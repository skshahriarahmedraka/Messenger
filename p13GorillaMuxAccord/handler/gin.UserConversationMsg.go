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
	// "github.com/gorilla/websocket"
	// "go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections) GinUserConversationMsg() gin.HandlerFunc {
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
			}
			var ReqUserData TempMessage

			err := ws.ReadJSON(&ReqUserData)
			fmt.Println("🚀 ~ file: gin.websocket.go ~ line 56 ~ returnfunc ~ ReqData : ", ReqUserData)
			// mt, message, err := ws.ReadMessage()
			if err != nil {
				logerror.ERROR("🚀 ~ file: Instantbuy.go ~ line 58 ~ func ~ err : ", err)
				// fmt.Println(err)
				break
			}
			//If client message is ping will return pong

			// var resError model.ErrorRes

			// var ReqUserData TempMessage

			var UserDataDB model.Conversation

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
			defer cancel()
			err = H.MongoUser.Collection(os.Getenv("MONGO_USER_MSG_COL")).FindOne(ctx, bson.M{"ConversationID": ReqUserData.ConversationID}).Decode(&UserDataDB)
			if err != nil {
				logerror.ERROR("🚀 ~ file: gin.UserConversationMsg.go ~ line 78 ~ returnfunc ~ err : ", err)
				return
			}
			//Response message to client

			err = ws.WriteJSON(UserDataDB.Messages)
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
