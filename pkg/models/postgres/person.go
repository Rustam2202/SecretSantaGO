package postgres

import (
	"database/sql"
)

type PersonModel struct {
	DB *sql.DB
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
	stmt := `INSERT INTO persons (email, password, firstname, lastname) VALUES(?, ?, ?, ?)`
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
