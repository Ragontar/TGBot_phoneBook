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

func TestGetRecordByName(t *testing.T) {
	setup.Init("/home/reaver/GolandProjects/TGBot_phoneBook/cfg/config.txt")

	records, err := GetRecordByName(100, "vasya")
	if err != nil {
		setup.GetLogger().Println("Error selecting Vasyas with ids 1, 3, 6, 8")
		fmt.Println(err)
		os.Exit(40)
	} else {
		setup.GetLogger().Printf("Selected Vasyas with ids 1, 3, 6, 8 %v", records)
	}
	fmt.Println(records)
}

func TestGetAllRecords(t *testing.T) {
	setup.Init("/home/reaver/GolandProjects/TGBot_phoneBook/cfg/config.txt")

	records, err := GetAllRecords(200)
	if err != nil {
		setup.GetLogger().Println("Error getting all records from userId = 100:")
		fmt.Println(err)
		os.Exit(40)
	} else {
		setup.GetLogger().Printf("Got all records of userId = 100, %v", records)
	}
	fmt.Println(records)
}

func TestDeleteRecordByID(t *testing.T) {
	setup.Init("/home/reaver/GolandProjects/TGBot_phoneBook/cfg/config.txt")

	err := DeleteRecordByID(100, 3)
	if err != nil {
		setup.GetLogger().Println("Error deleting phonebook record from userId = 100:")
		fmt.Println(err)
		os.Exit(40)
	} else {
		setup.GetLogger().Printf("Deleted phonebook record from userId = 100")
	}

}
