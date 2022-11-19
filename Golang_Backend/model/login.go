package model



type LoginModel struct {
	Email string `gorm:"type:varchar(100);not null" json:"Email"`
	Password string `gorm:"type:varchar(100);not null" json:"Password"`
} 




// type PostgresConfig struct {
// 	HOST string 
// 	PORT string 
// 	PASSWORD string 
// 	USER string 
// 	DBNAME string 
// 	SSLMODE string
// 	TIMEZONE string 
// }
