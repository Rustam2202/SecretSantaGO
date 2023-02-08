package postgres

import (
	"database/sql"
	"strconv"
)

type EventModel struct {
	DB *sql.DB
}

func (m *PersonModel) CreateTableEvents() {
	query := `CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT,
		persons SERIAL[]
		)`
	statement, err := m.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (m *PersonModel) newEvent(peronsId []int) error {
	list := "'{"
	for _, id := range peronsId {
		list += strconv.Itoa(id) + ", "
	}
	list += "}'"
	stmt := `INSERT INTO events (name, persons) VALUES(?, ?)`
	_, err := m.DB.Exec(stmt, "NewYear Party 2022-2023", list) // list = '{1,2,4,5,11}'
	if err != nil {
		return err
	}
	return nil
}
