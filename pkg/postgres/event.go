package postgres

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
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

func (db *DataBase) addPersonToEvent(eventName string, personId int) error {
	stmt := `SELECT persons FROM events WHERE name = $1`
	row := db.DB.QueryRow(stmt, eventName)

	var personsId []sql.NullInt32 // can't simply to Scan in []int
	if err := row.Scan(pq.Array(&personsId)); err != nil {
		return err
	}
	var ids []int
	for _, id := range personsId {
		if int(id.Int32) == personId {
			return errors.New("person already exist in event")
		}
		ids = append(ids, int(id.Int32))
	}
	ids = append(ids, personId)

	stmt = `UPDATE events SET persons = $2 WHERE name = $1`
	if _, err := db.DB.Exec(stmt, eventName, makeSQLArray(ids)); err != nil {
		return err
	}
	return nil
}
