package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=apigo-sample port=9000 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error: ", err)
	}
	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDb() *gorm.DB {
	return db
}
