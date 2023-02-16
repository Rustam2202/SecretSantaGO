package postgres

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestPersonsTable(t *testing.T) {
	dataBase := DataBase{}
	err := dataBase.OpenDb()
	if err != nil {
		t.Errorf("Databsae could not open. %s", err)
	}
	err = dataBase.createPersonsTable()
	if err != nil {
		t.Errorf("Table 'persons' could not be created. %s", err)
	}
	err = dataBase.dropPersonsTable()
	if err != nil {
		t.Errorf("Table 'persons' could not be droped. %s", err)
	}
}

func TestAddPerson(t *testing.T) {
	db := DataBase{}
	db.OpenDb()
	db.dropPersonsTable()
	db.createPersonsTable()

	if err := db.AddPersonToDB("e@mail.com", "John", "Doe", []byte{'p', 'a', 's', 1, 2, 3}, nil, nil); err != nil {
		t.Errorf("John Doe wasn't insert in database: %s", err)
	}
	if err := db.AddPersonToDB("em@ail.com", "Marry", "Sue", []byte{'p', 'a', 's', 'w', 'o', 'r', 'd', 1, 2, 3, 4}, []int{3, 8}, nil); err != nil {
		t.Errorf("Mary Sue wasn't insert in database: %s", err)
	}
	if err := db.AddPersonToDB("e-m@il.com", "Ted", "", []byte{'c', 'r', 'o', 'c', 'o', 'd', 'a', 'i', 'l', '8'}, []int{6}, []int{0}); err != nil {
		t.Errorf("Ted wasn't insert in database: %s", err)
	}
	if err := db.AddPersonToDB("only@email.net", "", "", []byte{}, []int{}, []int{}); err != nil {
		t.Errorf("Ted wasn't insert in database: %s", err)
	}
	db.dropPersonsTable()
}

func TestDeletePerson(t *testing.T) {
	db := DataBase{}
	db.OpenDb()
	db.dropPersonsTable()
	db.createPersonsTable()

	if err := db.deletePersonFromDB("em@ail.com"); err != nil {
		// delete non-exist row is not error
		t.Errorf("Something wrong with deleting non-exist person: %s", err)
	}

	db.AddPersonToDB("e@mail.com", "John", "Doe", []byte{}, nil, nil)
	db.AddPersonToDB("em@ail.com", "Marry", "Sue", []byte{}, nil, nil)
	db.AddPersonToDB("e-m@il.com", "Ted", "", []byte{}, nil, nil)

	if err := db.deletePersonFromDB("em@ail.com"); err != nil {
		t.Errorf("Mary Sue wasn't deleted: %s", err)
	}
	if err := db.deletePersonFromDB("em@ail.com"); err != nil {
		// delete non-exist row is not error
		t.Errorf("Something wrong with deleting non-exist person: %s", err)
	}
	db.dropPersonsTable()
}

func TestExistPerson(t *testing.T) {
	db := DataBase{}
	db.OpenDb()
	db.createPersonsTable()

	db.AddPersonToDB("e@mail.com", "John", "Doe", []byte{}, nil, nil)
	db.AddPersonToDB("em@ail.com", "Marry", "Sue", []byte{}, nil, nil)
	db.AddPersonToDB("e-m@il.com", "Ted", "", []byte{}, nil, nil)

	if b := db.IsPersonExist("e@mail.com"); b != true {
		t.Errorf("Expected true")
	}
	if b := db.IsPersonExist("non-exist@email.com"); b != false {
		t.Errorf("Expected false")
	}
}
