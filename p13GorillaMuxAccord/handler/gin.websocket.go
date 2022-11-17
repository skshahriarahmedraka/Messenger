package handler

import (
	"app/logerror"
	// "app/model"
	// "context"
	// "encoding/json"
	"fmt"
	"net/http"
	// "os"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	// "go.mongodb.org/mongo-driver/bson"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (H *DatabaseCollections) Instantbuy() gin.HandlerFunc {
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
			type UserInfo struct {
				UUID string `json:"UUID" bson:"UUID"`
			}
			var ReqData UserInfo

			err := ws.ReadJSON(&ReqData)
			fmt.Println("🚀 ~ file: gin.websocket.go ~ line 56 ~ returnfunc ~ ReqData : ", ReqData)
			// mt, message, err := ws.ReadMessage()
			if err != nil {
				logerror.ERROR("🚀 ~ file: Instantbuy.go ~ line 58 ~ func ~ err : ", err)
				// fmt.Println(err)
				break
			}
			//If client message is ping will return pong

			//Response message to client
			type welcomeMessage struct {
				Message string `json:"Message" bson:"Message"`
			}
			var welcome welcomeMessage
			welcome.Message = "Welcome to  brother " + ReqData.UUID
			err = ws.WriteJSON(welcome)
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
