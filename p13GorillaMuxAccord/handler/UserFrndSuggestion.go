package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"encoding/json"
	// "fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)


func(H *DatabaseCollections) FrndSuggestion(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Add("Content-Type","application/json")
	var resError model.ErrorRes

	type ReqStruct struct {
		UUID string `json:"UUID" bson:"UUID"`
	}
	var ReqData ReqStruct

	err := json.NewDecoder(r.Body).Decode(&ReqData)
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
	if ReqData.UUID == "" {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return
		// c.JSON(http.StatusBadRequest, gin.H{"error": "bad request userid "})
		// return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{{}}
	cursor, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).Find(ctx, filter)
	if err != nil {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
		resError.ErrorRes = "mongodb Find(ctx, filter) error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
	}

	var results []model.UserFrndSuggStruct
	for cursor.Next(context.TODO()) {
		var res model.UserFrndSuggStruct
		var result model.AccordUser
		if err := cursor.Decode(&result); err != nil {
            logerror.ERROR("🚀 ~ file: UserFrndSuggestion.go ~ line 62 ~ iferr:=cursor.Decode ~ err : ", err)
			// log.Fatal(err)
		}
		if result.UUID != ReqData.UUID && NotInFrndList(result.FrinedList,result.UUID) {
			
		// 	H.MongoUser.Collection("MONGO_USERCOL").CountDocuments(ctx, bson.M{ })
		// logerror.ERROR("🚀 ~ file: AdminMoneyReqAccept.go ~ line 110 ~ func ~ err : ", err)
		// if err != nil {
		// 	resError.ErrorRes = "mongodb Count Document error"
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	err = json.NewEncoder(w).Encode(resError)
		// 	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb find one and update error"})
		// 	return
		// }

			res.UUID = result.UUID
			res.UserName = result.UserName
			res.UserID = result.UserID
			res.ProfileImg= result.ProfileImg
			res.UserBio= result.UserBio
			results=append(results, res)
		}
	}
	if err := cursor.Err(); err != nil {
        logerror.ERROR("🚀 ~ file: UserFrndSuggestion.go ~ line 76 ~ iferr:=cursor.Err ~ err : ", err)
		resError.ErrorRes = "mongo cursor.Err() error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
        logerror.ERROR("🚀 ~ file: UserFrndSuggestion.go ~ line 80 ~ iferr:=cursor.Err ~ err : ", err)
		return
	}


	// if err = cursor.All(context.TODO(), &results); err != nil {
	// 	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 35 ~ iferr=cursor.All ~ err : ", err)
	// 	resError.ErrorRes = "mongodb cursor iteration error"
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	err = json.NewEncoder(w).Encode(resError)
	// 	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
	// 	return
	// }

	// fmt.Println("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ results : ", results)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&results)
	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)




}


func  NotInFrndList(frndList []model.FriendStruct,UUID string) bool {
	for _,frnd := range frndList{
		if frnd.UUID == UUID{
			return false
		}
	}
	return true
}

