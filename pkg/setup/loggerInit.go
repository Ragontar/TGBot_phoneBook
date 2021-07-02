package setup

import (
	"log"
	"os"
)

var logger *log.Logger

func initLogger() *log.Logger {
	f, err := os.OpenFile("records.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logger := log.New(f, "", log.Flags())
	return logger
}

func GetLogger() *log.Logger {
	if logger == nil {
		return initLogger()
	}
	return logger
}
