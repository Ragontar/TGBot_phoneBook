package crud

import (
	"context"
	"github.com/Ragontar/TGBot_phoneBook/pkg/database"
)

/*

Methods to get data from database

GetRecordByName(userId int64, name string) ([]Person, error) --- returns a slice of Person with records about people
	with Name registered in user's phonebook. Does not return people with the same Name, who are not registered
	in user's phonebook from Global Phonebook table. Returns a nil-value and error in case of failure.

GetAllRecords(userId int64) ([]Person, error) --- returns a slice of Person with all records in user's phonebook.
	Returns a nil-value and error in case of failure.

*/

type Person struct {
	Name   string
	Number string
}

func GetRecordByName(userId int64, name string) ([]Person, error) {
	db := database.GetDB()
	records := make([]Person, 0)

	row := db.QueryRow(context.Background(), "SELECT phonebook FROM users WHERE user_id = $1;", userId)

	pbIDs := make([]int, 0)
	err := row.Scan(&pbIDs)
	if err != nil {
		return nil, err
	}

	sqlString := "SELECT phone_number FROM phonebook WHERE name = $1 AND id = $2;"
	for _, id := range pbIDs {
		row = db.QueryRow(context.Background(), sqlString, name, id)
		var number string
		err := row.Scan(&number)
		if err != nil {
			return nil, err
		}

		records = append(records, Person{name, number})
	}

	return records, nil
}

func GetAllRecords(userId int64) ([]Person, error) {
	records := make([]Person, 0)
	db := database.GetDB()

	row := db.QueryRow(context.Background(), "SELECT phonebook FROM users WHERE user_id = $1;", userId)

	pbIDs := make([]int, 0)
	err := row.Scan(&pbIDs)
	if err != nil {
		return nil, err
	}

	sqlString := "SELECT name, phone_number FROM phonebook WHERE id = $1;"
	for _, id := range pbIDs {
		row = db.QueryRow(context.Background(), sqlString, id)
		var name, number string
		err := row.Scan(&name, &number)
		if err != nil {
			return nil, err
		}

		records = append(records, Person{name, number})
	}
	return records, nil
}
