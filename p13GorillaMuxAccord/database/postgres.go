package database

// import (
// 	"app/model"
// 	"fmt"
// 	"log"
// 	"os"
// 	// "time"

// 	// "gorm.io/datatypes"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func PostgresInit() *gorm.DB {
// 	// v:= config.EnvironmentVar()

// 	c := &model.PostgresConfig{
// 		HOST:     os.Getenv("POSTGRES_HOST"),
// 		PORT:     os.Getenv("POSTGRES_PORT"),
// 		PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
// 		USER:     os.Getenv("POSTGRES_USER"),
// 		DBNAME:   os.Getenv("POSTGRES_DBNAME"),
// 		SSLMODE:  os.Getenv("POSTGRES_SSLMODE"),
// 		TIMEZONE: os.Getenv("POSTGRES_TIMEZONE"),
// 	}
// 	db, err := postgresConnect(c)
// 	if err != nil {
// 		fmt.Println("❌ Failed to connect Postgres ")
// 		return nil
// 	} else {
// 		fmt.Println("✨🥰 ~ file: postgres.go ~ line 28 ~ funcPostgresInit ~ db : ", db)

// 	}
// 	return db
// }

// func postgresConnect(config *model.PostgresConfig) (*gorm.DB, error) {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", config.HOST, config.PORT, config.USER, config.PASSWORD, config.DBNAME, config.SSLMODE, config.TIMEZONE)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Println("❌ Failed to connect postgres : ", err)
// 		return db, err
// 	}
// 	db.AutoMigrate(&model.AccordUser{})
// 	// location,_ := time.LoadLocation("Etc/GMT")

// 	// this should give you time in location
// 	// t := time.Now().In(location)
// 	// fmt.Println("🚀 ~ file: postgres.go ~ line 50 ~ funcpostgresConnect ~ t : ", t)

// 	return db, nil
// }
