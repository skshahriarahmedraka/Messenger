package handler 


import (


	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type DatabaseCollections struct {
	Mongo *mongo.Database
	Postgres *gorm.DB 
}

