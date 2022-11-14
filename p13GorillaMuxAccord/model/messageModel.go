package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// {
//     conversation_id: 12345,
//     time: time,
//     members: ['user1', 'user2'],
//     messages: [
//       {
//          sender: 'user1',
//          message: 'Hello World',
//          timestamp: time
//       },
//       {
//          sender: 'user1',
//          message: 'Hello World',
//          timestamp: time
//       }],
//    total_messages: 2
// }


type Conversation struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	GroupID string `json:"GroupID" bson:"GroupID"`
	ConversationID string `gorm:"type:varchar(100);unique_index;not null" json:"ConversationID" bson:"ConversationID"`
	Members []string `json:"Members" bson:"Members"`
	Messages []Message `json:"Messages" bson:"Messages"`
	TotalMessages uint64 `json:"TotalMessages" bson:"TotalMessages"`
}

type Message struct {
	SenderID string `json:"SenderID" bson:"SenderID"`
	SenderName string `json:"Sender" bson:"Sender"`
	Message string `json:"Message" bson:"Message"`
	Reactions []int `json:"Reactions" bson:"Reactions"`
	UserReaction []ReactionStruct `json:"ReactionID" bson:"ReactionID"`
	Timestamp time.Time `json:"Timestamp" bson:"Timestamp"`

}

type ReactionStruct struct {
	UserID string `json:"UserID" bson:"UserID"`
	ReactionID int `json:"ReactionID" bson:"ReactionID"`
}