package postgres

import (
	"database/sql"
)

type PersonModel struct {
	DB *sql.DB
}

func (m *PersonModel) Insert(email string, password string, firstName string, lastName string) error {
	stmt := `INSERT INTO persons (email, password, firstname, lastname) VALUES(?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, email, password, firstName, lastName)
	if err != nil {
		return err
	}
	return nil
}
