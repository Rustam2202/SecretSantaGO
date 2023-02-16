package postgres

import (
	"database/sql"
	"santa/pkg/models"
)

type PersonModel struct {
	DB *sql.DB
}

func (db *DataBase) createPersonsTable() error {
	query := `CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password BYTEA,
		firstname TEXT,
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

func (db *DataBase) AddPersonToDB(email, firstName, lastName string, password []byte,
	events, communities []int) error {
	stmt := `INSERT INTO persons (email, password, firstname, lastname, events, communities) 
				VALUES($1, $2, $3, $4, $5, $6)`
	_, err := db.DB.Exec(stmt, email, password, firstName, lastName, makeSQLArray(events),
		makeSQLArray(communities))
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) deletePersonFromDB(email string) error {
	stmt := `DELETE FROM  persons WHERE email = $1`
	_, err := db.DB.Exec(stmt, email)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) getPerson(email string) models.Person {
	var person models.Person
	stmt := `SELECT * FROM persons WHERE email = $1`
	_, err := db.DB.Query(stmt, email)
	if err != nil {
		return models.Person{}
	}
	return person
}

func (db *DataBase) IsPersonExist(email string) bool {
	stmt := `SELECT email FROM persons WHERE email = $1`
	row := db.DB.QueryRow(stmt, email)
	var em string
	err := row.Scan(&em)
	return err != sql.ErrNoRows
}
