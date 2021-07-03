package crud

import (
	"context"
	"github.com/Ragontar/TGBot_phoneBook/pkg/database"
)

/*

Methods to update existing data in database

*/

func DeleteRecordByID(userID int64, recordID int) error {
	db := database.GetDB()
	phonebook := make([]int, 0)

	row := db.QueryRow(context.Background(), "SELECT phonebook FROM users WHERE user_id = $1;", userID)
	err := row.Scan(&phonebook)
	if err != nil {
		return err
	}

	for i, id := range phonebook {
		if id == recordID {
			phonebook = append(phonebook[:i], phonebook[i+1:]...)
			break
		}
	}

	_, err = db.Exec(context.Background(), "UPDATE users SET phonebook = $1 WHERE user_id = $2;", phonebook, userID)
	if err != nil {
		return err
	}

	return nil
}
