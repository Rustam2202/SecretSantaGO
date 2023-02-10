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
	dataBase := DataBase{}
	dataBase.OpenDb()
	dataBase.createPersonsTable()

	if err := dataBase.AddPerson("e@mail.com", "John", "Doe", []byte{'p', 'a', 's', 1, 2, 3}, nil, nil); err != nil {
		t.Errorf("John Doe wasn't insert in database: %s", err)
	}
	if err := dataBase.AddPerson("em@ail.com", "Marry", "Sue", []byte{'p', 'a', 's', 'w', 'o', 'r', 'd', 1, 2, 3, 4}, []int{3, 8}, nil); err != nil {
		t.Errorf("Mary Sue wasn't insert in database: %s", err)
	}
	if err := dataBase.AddPerson("e-m@il.com", "Ted", "", []byte{'c', 'r', 'o', 'c', 'o', 'd', 'a', 'i', 'l', '8'}, []int{6}, []int{0}); err != nil {
		t.Errorf("Ted wasn't insert in database: %s", err)
	}

	dataBase.DB.Exec("DROP TABLE persons")
}

func TestDeletePerson(t *testing.T) {
	dataBase := DataBase{}
	dataBase.OpenDb()
	dataBase.createPersonsTable()

	if err := dataBase.deletePerson("em@ail.com"); err != nil {
		// delete non-exist row is not error
		t.Errorf("Something wrong with deleting non-exist person: %s", err)
	}

	dataBase.AddPerson("e@mail.com", "John", "Doe", []byte{}, nil, nil)
	dataBase.AddPerson("em@ail.com", "Marry", "Sue", []byte{}, nil, nil)
	dataBase.AddPerson("e-m@il.com", "Ted", "", []byte{}, nil, nil)

	if err := dataBase.deletePerson("em@ail.com"); err != nil {
		t.Errorf("Mary Sue wasn't deleted: %s", err)
	}
	if err := dataBase.deletePerson("em@ail.com"); err != nil {
		// delete non-exist row is not error
		t.Errorf("Something wrong with deleting non-exist person: %s", err)
	}

	dataBase.DB.Exec("DROP TABLE persons")
}
