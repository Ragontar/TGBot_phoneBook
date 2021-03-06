package database

/*

Тут реализованы методы подключения DB PostgreSQL к проекту. Драйвер: pgx.

func Init() *pgx.Conn  ---   инициализирует подключение. Параметры dsn-строки читаются из структуры с конфигом
							 (см. setup/cfgSetup.go) Также, возвращает базу данных.
func GetDB() *pgx.Conn ---   Лучше использовать этот метод для получения DB. В нем же реализована инициализация
							 подключения.

PS Перед инициализацией подключение конфигурационная структура должна быть определена и корректно заполнена!!!
Иначе паника и габелла.

*/

import (
	"context"
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"github.com/jackc/pgx/v4"
	"time"
)

var db *pgx.Conn

func Init() *pgx.Conn {
	cfgMap := setup.GetCfgSet().ConfigMap

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfgMap["host"], cfgMap["user"],
		cfgMap["password"], cfgMap["dbname"], cfgMap["port"], cfgMap["sslmode"])

	fmt.Println(dsn)

	//db, err := sql.Open("pq", dsn)
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	return db
}

func GetDB() *pgx.Conn {
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
