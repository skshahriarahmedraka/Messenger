package model

import (
	// "time"

	// "gorm.io/datatypes"
	"gorm.io/gorm"
)


type AccordUser struct {

	gorm.Model
	UUID string `gorm:"type:varchar(100);unique_index;not null" json:"UUID" bson:"UUID"` 
	
	Email string `gorm:"type:varchar(100);unique_index;not null" json:"Email"  bson:"Email" `
	Password string `gorm:"type:varchar(100);not null" json:"Password" bson:"Password"`
	UserID string `gorm:"type:varchar(100);not null" json:"UserID" bson:"UserID"`
	UserName string `gorm:"type:varchar(100);not null" json:"UserName" bson:"UserName"` 
	Mobile string `gorm:"type:varchar(50);not null" json:"Mobile" bson:"Mobile"`
	BirthDate string `gorm:"type:varchar(20);not null" json:"BirthDate" bson:"BirthDate"`
	
	Accounttype string `gorm:"type:varchar(20);not null" json:"Accounttype" bson:"Accounttype"`
	ProfileImg string  `gorm:"type:varchar(200);not null" json:"ProfileImg" bson:"ProfileImg"` 
	BannerImg string   `gorm:"type:varchar(200);not null" json:"BannerImg" bson:"BannerImg"`
	Coin float64  ` json:"Coin" bson:"Coin"` 
	FrinedListID []string  ` json:"FriendListID"   bson:"FriendListID" `
	
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






