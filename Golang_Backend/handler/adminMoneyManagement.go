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
)

func (H *DatabaseCollections) MoneyManagement(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	var resError model.ErrorRes

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	filter := bson.D{{}}
	cursor, err := H.MongoUser.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).Find(ctx, filter)
	if err != nil {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 25 ~ func ~ err : ", err)
		resError.ErrorRes = "mongodb Find(ctx, filter) error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
	}
	var results []model.UserMoney
	if err = cursor.All(context.TODO(), &results); err != nil {
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 35 ~ iferr=cursor.All ~ err : ", err)
		resError.ErrorRes = "mongodb cursor iteration error"
		w.WriteHeader(http.StatusInternalServerError)
		err = json.NewEncoder(w).Encode(resError)
		logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 30 ~ func ~ err : ", err)
		return
	}

	fmt.Println("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ results : ", results)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&results)
	logerror.ERROR("🚀 ~ file: adminMoneyManagement.go ~ line 45 ~ func ~ err : ", err)
}
