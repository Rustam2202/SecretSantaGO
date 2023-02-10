package postgres

import (
	"database/sql"
)

type EventModel struct {
	DB *sql.DB
}

func (db *DataBase) createEventsTable() error {
	query := `CREATE TABLE IF NOT EXISTS events (
		name TEXT,
		persons INTEGER[]
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

func (db *DataBase) dropEventsTable() error {
	query := `DROP TABLE events`
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

func (db *DataBase) addEvent(eventName string, peronIDs []int) error {
	stmt := `INSERT INTO events (name, persons) VALUES($1, $2)`
	_, err := db.DB.Exec(stmt, eventName, makeSQLArray(peronIDs))
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) addPersonToEvent(eventName string, personId int) {

}


