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

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections)UserTokenReqList(w http.ResponseWriter, r *http.Request){
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
	//bson.M{"UserID": ReqData.UserID}
	filter := bson.M{"UUID": ReqData.UUID}
	// findOptions := options.Find().SetProjection(bson.D{{"ReqList", 1},{"_id",0},})

	result := []model.UserMoney{}
	cursor, err := H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).Find(ctx, filter)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resError.ErrorRes = "http.StatusBadRequest"
		_ = json.NewEncoder(w).Encode(resError)
		return

		// LogError.LogError("🚀 ~ mogodb FindOne(ctx, filter, findOptions).Decode(&result) : ", err)
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}

	// var results []bson.D
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	var Umoney model.UserMoney
	for _, r := range result {
		Umoney = r
	}
	fmt.Println("🚀 ~ file: userTokenReq.go ~ line 100 ~ returnfunc ~ Umoney.ReqList", Umoney.ReqList)
	// _ = json.NewEncoder(w).Encode(Umoney.ReqList)
	// c.JSON(http.StatusOK, Umoney.ReqList)


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Umoney.ReqList)

}