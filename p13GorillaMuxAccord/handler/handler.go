package handler 


import (


	"go.mongodb.org/mongo-driver/mongo"
	// "gorm.io/gorm"
)

type DatabaseCollections struct {
	MongoUser *mongo.Database
	MongoUserMsg *mongo.Database
	// Postgres *gorm.DB 
}

