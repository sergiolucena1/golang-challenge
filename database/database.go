package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB(){
	str := "host=localhost port=5432 user=postgres dbname=GolangChallenge sslmode=disable password=postgres"

	database , err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil{
		log.Fatal("error", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB{
	return db
}
