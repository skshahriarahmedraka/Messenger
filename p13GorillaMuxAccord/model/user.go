package model

import (
	// "time"

	// "gorm.io/datatypes"
	"gorm.io/gorm"
)


type AccordUser struct {

	gorm.Model
	Email string `gorm:"type:varchar(100);unique_index;not null" json:"Email"`
	Password string `gorm:"type:varchar(100);not null" json:"Password"`
	Userid string `gorm:"type:varchar(100);not null" json:"Userid"`
	Username string `gorm:"type:varchar(100);not null" json:"Username"` 
	Mobile string `gorm:"type:varchar(50);not null" json:"Mobile"`
	Birthdate string `gorm:"type:varchar(20);not null" json:"BirthDate"`
	Accounttype string `gorm:"type:varchar(20);not null" json:"Accounttype"`
	
	
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






