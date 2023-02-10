package postgres

import (
	"database/sql"
)

type PersonModel struct {
	DB *sql.DB
}

func (db *DataBase) createPersonsTable() error {
	query := `CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		firstname TEXT NOT NULL,
		lastname TEXT,
		events INTEGER[],
		communities INTEGER[]
		)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) dropPersonsTable() error {
	query := `DROP TABLE persons`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) addPerson(email, password, firstName, lastName string, events, communities []int) error {
	stmt := `INSERT INTO persons (email, password, firstname, lastname, events, communities) VALUES($1, $2, $3, $4, $5, $6)`
	_, err := db.DB.Exec(stmt, email, password, firstName, lastName, makeSQLArray(events), makeSQLArray(communities))
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) deletePerson(email string) error {
	stmt := `DELETE FROM  persons WHERE email = $1`
	_, err := db.DB.Exec(stmt, email)
	if err != nil {
		return err
	}
	return nil
}
