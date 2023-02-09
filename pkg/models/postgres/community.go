package postgres

import (
	"database/sql"
	"strconv"
)

type CommunityModel struct {
	DB *sql.DB
}

func (m *CommunityModel) CreateTableCommunity() {
	query := `CREATE TABLE IF NOT EXISTS community (
		id SERIAL PRIMARY KEY,
		name TEXT,
		persons SERIAL[],
		events SERIAL[]
		)`
	statement, err := m.DB.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func (m *CommunityModel) newCommunity(name string, peronsId []int) error {
	list := "'{"
	for _, id := range peronsId {
		list += strconv.Itoa(id) + ", "
	}
	list += "}'"
	stmt := `INSERT INTO community (name, persons) VALUES($1, $2)`
	_, err := m.DB.Exec(stmt, name, list) // list = '{1,2,4,5,11}'
	if err != nil {
		return err
	}
	return nil
}
