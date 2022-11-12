package handler

import (
	"app/logerror"
	"app/model"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "gorm.io/gorm"
)

type DatabaseCollections struct {
	MongoUser *mongo.Database
	MongoUserMsg *mongo.Database
	// Postgres *gorm.DB 
}


func (H *DatabaseCollections)UserIDtoUUID(UserID string) string {
	var UserDataDB model.AccordUser
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err := H.MongoUser.Collection(os.Getenv("MONGO_USERCOL")).FindOne(ctx, bson.M{"UserID": UserID}).Decode(&UserDataDB)
    logerror.ERROR("🚀 ~ file: handler.go ~ line 25 ~ func ~ err : ", err)
	if err != nil {
		
		return UserDataDB.UUID
	} else {
		return ""
	}
} 