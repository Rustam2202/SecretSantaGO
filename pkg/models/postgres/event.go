package postgres

import (
	"database/sql"
	"strconv"
)

type EventModel struct {
	DB *sql.DB
}

type EventInterface interface {
	createTable()
}

func (m *EventModel) openDb() error {
	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	m.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
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

func (m *PersonModel) newEvent(name string, peronsId []int) error {
	list := "'{"
	for _, id := range peronsId {
		list += strconv.Itoa(id) + ", "
	}
	list += "}'"
	stmt := `INSERT INTO events (name, persons) VALUES($1, $2)`
	_, err := m.DB.Exec(stmt, name, list) // list = '{1,2,4,5,11}'
	if err != nil {
		return err
	}
	return nil
}

func (m *PersonModel) addPerson(name string, personId int) {

}
