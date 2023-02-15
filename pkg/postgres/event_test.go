package postgres

import (
	"database/sql"
	"santa/pkg/models"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestEventsTable(t *testing.T) {
	dataBase := DataBase{}
	err := dataBase.OpenDb()
	if err != nil {
		t.Errorf("Databsae could not open. %s", err)
	}
	err = dataBase.createEventsTable()
	if err != nil {
		t.Errorf("Table 'events' could not be created. %s", err)
	}
	err = dataBase.dropEventsTable()
	if err != nil {
		t.Errorf("Table 'events' could not be droped. %s", err)
	}
}

func TestAddEmptyEvent(t *testing.T) {
	dataBase := DataBase{}
	dataBase.OpenDb()
	dataBase.createEventsTable()
	err := dataBase.addEvent("Test Event", nil)
	if err != nil {
		t.Errorf("Empty event was not added. %s", err)
	}

	dataBase.dropEventsTable()
}

func TestAddEventWithOnePerson(t *testing.T) {
	dataBase := DataBase{}
	dataBase.OpenDb()
	dataBase.createEventsTable()
	err := dataBase.addEvent("Test Event", []int{11})
	if err != nil {
		t.Errorf("Event with one person was not added. %s", err)
	}
	dataBase.dropEventsTable()
}

func TestAddEvent(t *testing.T) {
	dataBase := DataBase{}
	dataBase.OpenDb()
	dataBase.createEventsTable()
	err := dataBase.addEvent("Test Event", []int{1, 2, 4, 7, 11})
	if err != nil {
		t.Errorf("Event with persons was not added. %s", err)
	}
	table := checkEventsTable(dataBase.DB, "Test Event")
	assert.Equal(t, []int{1, 2, 4, 7, 11}, table.Persons)
	dataBase.dropEventsTable()
}

func TestAddPersonToEvent(t *testing.T) {
	db := DataBase{}
	db.OpenDb()
	db.createEventsTable()
	eventName := "Test Event"
	db.addEvent(eventName, nil)
	if err := db.addPersonToEvent(eventName, 5); err != nil {
		t.Error(err)
	}
	if err := db.addPersonToEvent(eventName, 11); err != nil {
		t.Error(err)
	}
	if err := db.addPersonToEvent(eventName, 5); err == nil {
		t.Error("Person already exist, expected error.", err)
	}
	table := checkEventsTable(db.DB, eventName)
	assert.Equal(t, eventName, table.Name)
	assert.Equal(t, []int{5, 11}, table.Persons)
	db.dropEventsTable()
}

func checkEventsTable(db *sql.DB, eventName string) models.Event {
	stmt := `SELECT * FROM events WHERE name = $1`
	row := db.QueryRow(stmt, eventName)
	var name string
	var personsId []sql.NullInt32
	row.Scan(&name, pq.Array(&personsId))
	var ids []int
	for _, id := range personsId {
		ids = append(ids, int(id.Int32))
	}
	return models.Event{Name: name, Persons: ids}
}
