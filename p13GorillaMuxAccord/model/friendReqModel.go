package model

// import "go.mongodb.org/mongo-driver/bson/primitive"


// type FrndList struct {
// 	ID primitive.ObjectID `json:"_id" bson:"_id"`
// 	UUID string `gorm:"type:varchar(100);unique_index;not null" json:"UUID" bson:"UUID"` 
// 	UserID string `gorm:"type:varchar(100);not null" json:"UserID" bson:"UserID"`
	
// }

type FrndReqList struct {
	UUID string `gorm:"type:varchar(100);unique_index;not null" json:"UUID" bson:"UUID"`
	FrndReq []FrndReqShort `json:"FrndReq" bson:"FrndReq"`
}

type FrndReqShort struct {
	UUID string `json:"UUID" bson:"UUID"`
	UserID string `json:"UserID" bson:"UserID"`
	ProfileImg string `json:"ProfileImg" bson:"ProfileImg"`
	UserBio string `json:"UserBio" bson:"UserBio"`
	UserName string `json:"UserName" bson:"UserName"`
}


type UserFrndReqSend struct {
	UUID string `gorm:"type:varchar(100);unique_index;not null" json:"UUID" bson:"UUID"`
	Frnd string `gorm:"type:varchar(100);not null" json:"Frnd" bson:"Frnd"`
}
