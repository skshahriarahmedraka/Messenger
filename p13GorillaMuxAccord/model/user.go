package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// "time"

// "gorm.io/datatypes"


type AccordUser struct {

	ID primitive.ObjectID `json:"_id" bson:"_id"`
	UUID string `gorm:"type:varchar(100);unique_index;not null" json:"UUID" bson:"UUID"` 
	UserID string `gorm:"type:varchar(100);not null" json:"UserID" bson:"UserID"`
	
	Email string `gorm:"type:varchar(100);unique_index;not null" json:"Email"  bson:"Email" `
	Password string `gorm:"type:varchar(100);not null" json:"Password" bson:"Password"`
	UserName string `gorm:"type:varchar(100);not null" json:"UserName" bson:"UserName"` 
	Mobile string `gorm:"type:varchar(50);not null" json:"Mobile" bson:"Mobile"`
	BirthDate string `gorm:"type:varchar(20);not null" json:"BirthDate" bson:"BirthDate"`
	
	ProfileImg string  `gorm:"type:varchar(200);not null" json:"ProfileImg" bson:"ProfileImg"` 
	BannerImg string   `gorm:"type:varchar(200);not null" json:"BannerImg" bson:"BannerImg"`
	
	Coin float64  ` json:"Coin" bson:"Coin"`
	CoinReq []CoinReq `json:"CoinReq" bson:"CoinReq"` 
	Accounttype string `gorm:"type:varchar(20);not null" json:"Accounttype" bson:"Accounttype"`
	TransactionHistory []string `json:"TransactionHistory" bson:"TransactionHistory"`
	
	
	ContactAdminMsg []string `json:"ContactAdminMsg" bson:"ContactAdminMsg"`
	FrinedListID []string  ` json:"FriendListID"   bson:"FriendListID" `
	GroupListID []string  ` json:"GroupListID"   bson:"GroupListID" `

	City    string `json:"City" bson:"City"`
	Address string `json:"Address" bson:"Address"`
	Country string `json:"Country"  bson:"Country"`
	ZipCode string `json:"ZipCode" bson:"ZipCode"`

	WishList           []string `json:"WishList" bson:"WishList"`
	
}

type CoinReq struct {
	Amount float64 `json:"Amount" bson:"Amount"`
	JWT string `json:"JWT" bson:"JWT"`
	Validity bool `json:"Validity" bson:"Validity"`
}

// type UserAuth struct {
//     gorm.Model
//     ProfileID int      `gorm:"not null" json:"profile_id"`
//     Username  string   `gorm:"size:20;unique_index" json:"username"`
//     Email     string   `gorm:"type:varchar(100);unique_index" json:"email"`
//     Password  string   `gorm:"not null" json:"password"`
//     Remember  bool     `gorm:"not null" json:"remeber"`
//     TwoFA     bool     `gorm:"not null" json:"twofa"`
//     Access    int      `gorm:"not null" json:"access"`
//     State     int      `gorm:"not null" json:"state"`
//     LastSeen  string   `gorm:"not null" json:"lastseen"`
// }






