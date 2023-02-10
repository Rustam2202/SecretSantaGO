package postgres

import (
	"testing"
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
	dataBase.dropEventsTable()
}
