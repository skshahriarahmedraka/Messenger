package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) UserFrndData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	paramsMap := mux.Vars(r)
	ReqUUID := paramsMap["UUID"]

	var UserDataDB model.AccordUser
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UserID": ReqUUID}).Decode(&UserDataDB)
	if err != nil {
    logerror.ERROR("🚀 ~ file: UserProfileData.go ~ line 27 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(`{status: " mongodb Connection error"}`)
		return
	} else {
		objID, _ := primitive.ObjectIDFromHex("")
		UserDataDB.ID = objID
		UserDataDB.Password = ""
		// c.JSON(http.StatusOK, UserDataDB)
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(UserDataDB)
		return
	}



}
