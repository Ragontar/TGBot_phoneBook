package crud

/*

Methods, implementing addition of new data to DB.

func registerUser(userId int64) error   ---   adds a new user by telegram userId. Returns an error otherwise.

func addToGlobalPhonebook(name string, number string) (int, error)   ---   adds a record to the global phonebook,
	returns id of the new record or -1 and non-nil error in case of failure

*/

import (
	"context"
	"fmt"
	"github.com/Ragontar/TGBot_phoneBook/pkg/database"
	"github.com/Ragontar/TGBot_phoneBook/pkg/setup"
)

func registerUser(userId int64) error {
	db := database.GetDB()
	arr := make([]int, 0)
	sqlString := fmt.Sprintf("INSERT INTO users (user_id, phonebook) VALUES ($1, $2);")

	logger := setup.GetLogger()
	logger.Printf("Executed query: %s", sqlString)
	logger.Printf("VALUES: %v, %v", userId, arr)

	_, err := db.Exec(context.Background(), sqlString, userId, arr)
	if err != nil {
		setup.GetLogger().Printf("Failed to register user %d: %v", userId, err)
		return err
	}

	return nil
}

func addToGlobalPhonebook(name string, number string) (int, error) {
	db := database.GetDB()
	logger := setup.GetLogger()
	sqlString := fmt.Sprintf("INSERT INTO phonebook (name, phone_number) VALUES ($1, $2) RETURNING id;")

	row := db.QueryRow(context.Background(), sqlString, name, number)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}

	logger.Printf("Executed query: %s", sqlString)
	logger.Printf("VALUES: %s, %s", name, number)
	logger.Printf("ID: %v", id)

	return id, err
}
