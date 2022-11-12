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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) UserProfileUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")

	var resError model.ErrorRes
	var ReqUserData model.AccordUser
	err := json.NewDecoder(r.Body).Decode(&ReqUserData)
	if err != nil {
		logerror.ERROR("🚀 ~ file: userProfileUpdate.go ~ line 21 ~ func ~ err : ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	count, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).CountDocuments(ctx, bson.M{"UserID": ReqUserData.UserID})
	if err != nil {
		logerror.ERROR("🚀 ~ file: userProfileUpdate.go ~ line 30 ~ func ~ err : ", err)
		w.WriteHeader(http.StatusInternalServerError)
		resError.ErrorRes = "mongodb Count Document error"
		_ = json.NewEncoder(w).Encode(resError)
		return

	}
	if count > 0 {
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{"UUID", ReqUserData.UUID}}
		update := bson.D{
			{"$set", bson.D{{"UserName", ReqUserData.UserName}}},
			{"$set", bson.D{{"UserID", ReqUserData.UserID}}},
			{"$set", bson.D{{"Email", ReqUserData.Email}}},
			{"$set", bson.D{{"Mobile", ReqUserData.Mobile}}},
			{"$set", bson.D{{"BirthDate", ReqUserData.BirthDate}}},
			{"$set", bson.D{{"UserBio", ReqUserData.UserBio}}},

			{"$set", bson.D{{"ProfileImg", ReqUserData.ProfileImg}}},
			{"$set", bson.D{{"BannerImg", ReqUserData.BannerImg}}},

			{"$set", bson.D{{"Address", ReqUserData.Address}}},
			{"$set", bson.D{{"City", ReqUserData.City}}},
			{"$set", bson.D{{"Country", ReqUserData.Country}}},
			{"$set", bson.D{{"ZipCode", ReqUserData.ZipCode}}},
		}
		res, err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).UpdateOne(ctx, filter, update, opts)
		if err != nil {
			logerror.ERROR("🚀 ~ file: userProfileUpdate.go ~ line 61 ~ func ~ err : ", err)
			w.WriteHeader(http.StatusInternalServerError)
			resError.ErrorRes = "mongodb update one error"
			_ = json.NewEncoder(w).Encode(resError)
			return
			
			
			} else {
				fmt.Println("🚀 ~ file: userProfileUpdate.go ~ line 59 ~ func ~ res : ", res)
		resError.ErrorRes = "mongodb Count Document error"
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(ReqUserData)
        logerror.ERROR("🚀 ~ file: userProfileUpdate.go ~ line 74 ~ func ~ err : ", err)
		return
			
		}

	} else {
		fmt.Println("🚀 ~ file: userProfileUpdate.go ~ line 88 ~ func ~ count : ", count)
		w.WriteHeader(http.StatusBadRequest)
			resError.ErrorRes = "Count document not > 0"
			_ = json.NewEncoder(w).Encode(resError)
			return
	}

}
