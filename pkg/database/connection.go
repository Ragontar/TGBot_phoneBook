package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Init() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *gorm.DB {
	if db == nil {
		db = Init()
		sleep := time.Duration(1)

		for db == nil {
			sleep *= 2
			fmt.Printf("DB is unavailable. Sleeping %ds....", sleep)
			time.Sleep(sleep * time.Second)
			db = Init()
		}
	}

	return db
}
