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

func (H *DatabaseCollections) ChatFrndList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	// paramsMap := mux.Vars(r)
	// ReqUUID := paramsMap["UUID"]
	type ReqStruct struct {
		UUID string `json:"UUID" bson:"UUID"`
	}
	var ReqData ReqStruct
	err := json.NewDecoder(r.Body).Decode(&ReqData)

	var UserChatFrndList model.UserChatFrndList
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = H.MongoUser.Collection(os.Getenv("MONGO_FRND_CHATSHORT_COL")).FindOne(ctx, bson.M{"UUID": ReqData.UUID}).Decode(&UserChatFrndList)
	if err != nil {
		logerror.ERROR("🚀 ~ file: UserProfileData.go ~ line 27 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
		return
	} else {

        fmt.Println("🚀 ~ file: UserChatFrndList.go ~ line 38 ~ func ~ UserChatFrndList.FrndChatList : ", UserChatFrndList.FrndChatList)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(UserChatFrndList.FrndChatList)
		return

		// fmt.Println("🚀 ~ file: UserFrndList.go ~ line 28 ~ func ~ UserDataDB : ", UserDataDB)

		// // objID, _ := primitive.ObjectIDFromHex("")
		// // UserDataDB.ID = objID
		// // UserDataDB.Password = ""
		// // c.JSON(http.StatusOK, UserDataDB)

		// ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
		// defer cancel()

		// var UserChatFrndList []model.UserChatFrndList
		// // var temp []model.FriendStruct
		// temp := UserDataDB.FriendList
		// fmt.Println("🚀 ~ file: UserFrndList.go ~ line 48 ~ func ~ UserDataDB.FriendList : ", UserDataDB.FriendList)
		// for _, i := range temp {
		// 	fmt.Println("🚀 ~ file: UserFrndList.go ~ line 49 ~ func ~ i : ", i)
		// 	var FrndShort model.FrndReqShort
		// 	// var UserStruct model.AccordUser
		// 	err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UUID": i.UUID}).Decode(&FrndShort)
        //     fmt.Println("🚀 ~ file: UserFrndList.go ~ line 54 ~ func ~ FrndShort : ", FrndShort)
		// 	// fmt.Println("🚀 ~ file: UserFrndList.go ~ line 50 ~ func ~ UserStruct : ", UserStruct)
		// 	if err != nil {
		// 		logerror.ERROR("🚀 ~ file: UserFrndList.go ~ line 48 ~ func ~ err : ", err)
		// 		w.WriteHeader(http.StatusInternalServerError)
		// 		_ = json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
		// 		return
		// 	}
		// 	FrndReqShort = append(FrndReqShort, FrndShort)
		// }
		// fmt.Println("🚀 ~ file: UserFrndList.go ~ line 56 ~ func ~ FrndReqShort : ", FrndReqShort)
		// w.WriteHeader(http.StatusOK)
		// _ = json.NewEncoder(w).Encode(FrndReqShort)
		// return
	}

}
