package postgres

import (
	"database/sql"
)

type PersonModel struct {
	DB *sql.DB
}

type PersonInteface interface {
	createTable()
}

func (db *DataBase) createPersonsTable() error {
	query := `CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		firstname TEXT NOT NULL,
		lastname TEXT
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

func (db *DataBase) addPerson(email string, password string, firstName string, lastName string) error {
	stmt := `INSERT INTO persons (email, password, firstname, lastname) VALUES($1, $2, $3, $4)`
	_, err := db.DB.Exec(stmt, email, password, firstName, lastName)
	if err != nil {
		return err
	}
	return nil
}

/*
func (p PersonModel) createTable() {
	query := `CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		firstname TEXT NOT NULL,
		lastname TEXT
		)`
	statement, err := p.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (m *PersonModel) openDb() error {
	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	m.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}

func (m *PersonModel) CreateTablePersons() {
	query := `CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		firstname TEXT NOT NULL,
		lastname TEXT
		)`
	statement, err := m.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (m *PersonModel) Insert(email string, password string, firstName string, lastName string) error {
	stmt := `INSERT INTO persons (email, password, firstname, lastname) VALUES($1, $2, $3, $4)`
	_, err := m.DB.Exec(stmt, email, password, firstName, lastName)
	if err != nil {
		return err
	}
	return nil
}

func (m *PersonModel) GetById() {

}

func (m *PersonModel) Edit() {

}

func (m *PersonModel) Delete() {

}
*/
