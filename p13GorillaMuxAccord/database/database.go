package database

import (
	"app/handler"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	// "gorm.io/gorm"
)


func DatabaseInitialization(db1 *mongo.Database,db2 *mongo.Database) handler.DatabaseCollections {
    fmt.Println("🚀 ~ file: database.go ~ line 12 ~ funcDatabaseInitialization ~ db1 : ", db1)
    fmt.Println("🚀 ~ file: database.go ~ line 12 ~ funcDatabaseInitialization ~ db2 : ", db2)
	return handler.DatabaseCollections{MongoUser:db1,MongoUserMsg: db2}
}

