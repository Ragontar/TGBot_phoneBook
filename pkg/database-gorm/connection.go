package database_gorm

/*

!!! НЕ ИСПОЛЬЗУЕТСЯ В ДАННЫЙ МОМЕНТ !!!

Тут реализованы методы подключения PostgreSQL к проекту с помощью GORM. Драйвер: pgx.

func Init() *gorm.DB   ---   инициализирует подключение. Параметры dsn-строки читаются из структуры с конфигом
							 (см. setup/cfgSetup.go) Также, возвращает базу данных.
func GetDB() *gorm.DB  ---   Лучше использовать этот метод для получения DB. В нем же реализована инициализация
							 подключения.

PS Перед инициализацией подключение конфигурационная структура должна быть определена и корректно заполнена!!!
Иначе паника и габелла.

*/

import (
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func Init() *gorm.DB {
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
