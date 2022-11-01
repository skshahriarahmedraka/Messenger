package database

import (
	"app/logerror"
	"context"
	"fmt"

	"os"


	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoUserDBConn() *mongo.Database {

	mongoURI := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=" + os.Getenv("MONGO_MAXPOOLSIZE") + "&w=" + os.Getenv("MONGO_W")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
    logerror.ERROR("🚀 ~ file: mongo.go ~ line 24 ~ funcMongoUserDBConn ~ err : ", err)
	return nil 
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
    logerror.ERROR("🚀 ~ file: mongo.go ~ line 30 ~ funcMongoUserDBConn ~ err : ", err)
		return nil 
}
	// defer client.Disconnect(ctx)
    fmt.Println("🚀 ~ file: mongo.go ~ line 34 ~ funcMongoUserDBConn ~ os.Getenv(\"MONGO_USER_DB\") : ", os.Getenv("MONGO_USER_DB"))
	Mydb := client.Database(os.Getenv("MONGO_USER_DB"))
    fmt.Println("⚡😍 ~ file: mongo.go ~ line 35 ~ funcMongoUserDBConn ~ Mydb : ", Mydb)

	// Mycol := Mydb.Collection("stackdb")


	return Mydb

}

func MongoMsgDBConn() *mongo.Database {

	mongoURI := "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=" + os.Getenv("MONGO_MAXPOOLSIZE") + "&w=" + os.Getenv("MONGO_W")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		logerror.ERROR("🚀 ~ file: mongo.go ~ line 46 ~ funcMongoMsgDBConn ~ err : ", err)
		return nil
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
    logerror.ERROR("🚀 ~ file: mongo.go ~ line 54 ~ funcMongoMsgDBConn ~ err : ", err)
	return nil
	}
	// defer client.Disconnect(ctx)
    fmt.Println("🚀 ~ file: mongo.go ~ line 61 ~ funcMongoMsgDBConn ~ os.Getenv(\"MONGO_MSG_DB\") : ", os.Getenv("MONGO_MSG_DB"))
	Mydb := client.Database(os.Getenv("MONGO_MSG_DB"))
    fmt.Println("⚡😍 ~ file: mongo.go ~ line 61 ~ funcMongoMsgDBConn ~ Mydb : ", Mydb)



	return Mydb

}


