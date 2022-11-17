package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserChatFrndList struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	UUID string `json:"UUID" bson:"UUID"`
	FrndChatList []FrndChatShort `json:"FrndChatList" bson:"FrndChatList"`
}

type FrndChatShort struct {
		UUID string `json:"UUID" bson:"UUID"`
		UserID string `json:"UserID" bson:"UserID"`
		UserImg string `json:"UserImg" bson:"UserImg"`
		UserName string `json:"UserName" bson:"UserName"`
		LastMessage string `json:"LastMessage" bson:"LastMessage"`
		LastMessageTime string `json:"LastMessageTime" bson:"LastMessageTime"`
		SilentNotification bool `json:"SilentNotification" bson:"SilentNotification"`
		NumberOfNotification int `json:"NumberOfNotification" bson:"NumberOfNotification"`
		ActiveStatus bool `json:"ActiveStatus" bson:"ActiveStatus"`
		LastActiveTime string `json:"LastActiveTime" bson:"LastActiveTime"`
}