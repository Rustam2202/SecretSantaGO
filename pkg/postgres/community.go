package postgres

import (
	"database/sql"
	"strconv"
)

type CommunityModel struct {
	DB *sql.DB
}

func (db *DataBase) CreateCommunitiesTable() error {
	query := `CREATE TABLE IF NOT EXISTS communities (
		name TEXT,
		persons INTEGER[],
		events INTEGER[]
		)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (db *DataBase) dropCommunitiesTable() error {
	query := `DROP TABLE communities`
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

func (m *CommunityModel) addCommunity(name string, peronsIDs, eventsIDs []int) error {
	pesons := "'{"
	for _, id := range peronsIDs {
		pesons += strconv.Itoa(id) + ","
	}
	pesons += "}'"
	stmt := `INSERT INTO communities (name, persons, events) VALUES($1, $2, $3)`
	_, err := m.DB.Exec(stmt, name, pesons) // list = '{1,2,4,5,11}'
	if err != nil {
		return err
	}
	return nil
}
