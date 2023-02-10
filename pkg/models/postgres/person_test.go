package postgres

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestPersonsTable(t *testing.T) {
	dataBase := DataBase{}
	err := dataBase.openDb()
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
	dataBase := DataBase{}
	dataBase.openDb()
	dataBase.createPersonsTable()

	if err := dataBase.addPerson("e@mail.com", "pass1234", "John", "Doe", nil, nil); err != nil {
		t.Errorf("John Doe wasn't insert in database: %s", err)
	}
	if err := dataBase.addPerson("em@ail.com", "passwrd123", "Marry", "Sue", []int{3, 8}, nil); err != nil {
		t.Errorf("Mary Sue wasn't insert in database: %s", err)
	}
	if err := dataBase.addPerson("e-m@il.com", "crocodail8", "Ted", "", []int{6}, []int{0}); err != nil {
		t.Errorf("Ted wasn't insert in database: %s", err)
	}

	dataBase.DB.Exec("DROP TABLE persons")
}

func TestDeletePerson(t *testing.T) {
	dataBase := DataBase{}
	dataBase.openDb()
	dataBase.createPersonsTable()

	if err := dataBase.deletePerson("em@ail.com"); err != nil {
		// delete non-exist row is not error
		t.Errorf("Something wrong with deleting non-exist person: %s", err)
	}

	dataBase.addPerson("e@mail.com", "pass1234", "John", "Doe", nil, nil)
	dataBase.addPerson("em@ail.com", "passwrd123", "Marry", "Sue", nil, nil)
	dataBase.addPerson("e-m@il.com", "crocodail8", "Ted", "", nil, nil)

	if err := dataBase.deletePerson("em@ail.com"); err != nil {
		t.Errorf("Mary Sue wasn't deleted: %s", err)
	}
	if err := dataBase.deletePerson("em@ail.com"); err != nil {
		// delete non-exist row is not error
		t.Errorf("Something wrong with deleting non-exist person: %s", err)
	}

	dataBase.DB.Exec("DROP TABLE persons")
}
