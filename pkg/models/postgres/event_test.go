package postgres

import (
	"testing"
)

func TestEventsTable(t *testing.T) {
	dataBase := DataBase{}
	err := dataBase.openDb()
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
	dataBase.openDb()
	dataBase.createEventsTable()
	err:=dataBase.addEvent("Test Event", nil)
	if err != nil {
		t.Errorf("'Test Event' was not added. %s", err)
	}
	dataBase.dropEventsTable()
}

func TestAddEvent(t *testing.T) {

}
