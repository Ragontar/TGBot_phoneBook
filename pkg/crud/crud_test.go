package crud

import (
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
	"os"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	setup.Init("/home/reaver/GolandProjects/TGBot_phoneBook/cfg/config.txt")
	err := registerUser(1488)
	if err != nil {
		fmt.Println("Error: registerUser()")
	}
}

func TestAddToGlobalPhonebook(t *testing.T) {
	setup.Init("/home/reaver/GolandProjects/TGBot_phoneBook/cfg/config.txt")
	id, err := addToGlobalPhonebook("vasya", "14881488")
	if err != nil {
		setup.GetLogger().Println("Error adding Vasya with phone 14881488")
		fmt.Println(err)
		os.Exit(40)
	} else {
		setup.GetLogger().Printf("Added Vasya with phone 14881488 and id %v", id)
	}
}
