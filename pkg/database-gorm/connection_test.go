package database_gorm

import (
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	os.Chdir("..")
	os.Chdir("..")

	setup.Init("./cfg/config.txt")
	Init()
}

func TestGetDB(t *testing.T) {
	os.Chdir("..")
	os.Chdir("..")

	setup.Init("./cfg/config.txt")
	GetDB()
}
