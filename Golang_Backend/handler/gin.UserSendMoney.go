package handler

import (
	"app/logerror"
	// "app/model"
	"context"
	// "net/http"
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

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func (H *DatabaseCollections) GinUserSendMoney() gin.HandlerFunc {
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
			// type UserInfo struct {
			// 	InsType string `json:"InsType" bson:"InsType"`
			// 	UserUUID string `json:"UserUUID" bson:"UserUUID"`
			// 	FrndUUID string `json:"frndUUID" bson:"frndUUID"`
			// }
			type TempMessage struct {
				SenderUUID string `json:"SenderUUID" bson:"SenderUUID"`
				ReceiverUUID string `json:"ReceiverUUID" bson:"ReceiverUUID"`
				Amount float64 `json:"Amount" bson:"Amount"`
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


			// var resError model.ErrorRes
	
	
	
	
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"UUID", ReqData.SenderUUID}}
	update :=  bson.D{{ "$inc" , bson.D{{ "Coin", -ReqData.Amount}}}}

	UpdateRes, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		
		break

	}
	fmt.Println("🚀 ~ file: UserSendMoney.go ~ line 49 ~ func ~ UpdateRes : ", UpdateRes)


	opts = options.Update().SetUpsert(true)
	filter = bson.D{{"UUID", ReqData.ReceiverUUID}}
	update =  bson.D{{ "$inc" , bson.D{{ "Coin", ReqData.Amount}}}}

	UpdateRes, err = H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
	if err != nil {
		
		break

	}
	fmt.Println("🚀 ~ file: UserSendMoney.go ~ line 49 ~ func ~ UpdateRes : ", UpdateRes)
	
	
		// c.JSON(http.StatusOK, UserDataDB)
        // fmt.Println("🚀 ~ file: UserConversationMsg.go ~ line 57 ~ func ~ UserDataDB.Messages : ", UserDataDB.Messages)

	

			
			//If client message is ping will return pong

			//Response message to client
			// type welcomeMessage struct {
			// 	Message string `json:"Message" bson:"Message"`
			// }
			// var welcome welcomeMessage
			// welcome.Message = "Welcome to  brother " 
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
