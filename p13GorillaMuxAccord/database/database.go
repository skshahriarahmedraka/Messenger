package database

import (
	"app/handler"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)


func DatabaseInitialization(db2 *gorm.DB,db1 *mongo.Database) handler.DatabaseCollections {
    fmt.Println("🚀 ~ file: database.go ~ line 15 ~ funcDatabaseInitialization ~ mongodb Database : ", db1)
    fmt.Println("🚀 ~ file: database.go ~ line 15 ~ funcDatabaseInitialization ~ postgresql Database : ", db2)
	return handler.DatabaseCollections{Mongo:db1,Postgres:db2}
}

