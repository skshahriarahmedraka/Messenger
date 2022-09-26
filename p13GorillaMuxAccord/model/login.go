package model



type LoginModel struct {
	Email string `gorm:"type:varchar(100);not null" json:"Email"`
	Password string `gorm:"type:varchar(100);not null" json:"Password"`
} 