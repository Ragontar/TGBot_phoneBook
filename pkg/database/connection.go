package database

import (
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Init() *gorm.DB {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable" //Переписать

	cfgMap := setup.GetCfgSet().ConfigMap

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfgMap["host"], cfgMap["user"],
		cfgMap["password"], cfgMap["dbname"], cfgMap["port"], cfgMap["sslmode"])

	fmt.Println(dsn)
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
