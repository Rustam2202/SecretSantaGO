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

	if err := dataBase.addPerson("e@mail.com", "pass1234", "John", "Doe"); err != nil {
		t.Errorf("John Doe wasn't insert in database: %s", err)
	}

	if err := dataBase.addPerson("em@ail.com", "passwrd123", "Marry", "Sue"); err != nil {
		t.Errorf("Mary Sue wasn't insert in database: %s", err)
	}
	if err := dataBase.addPerson("e-m@il.com", "crocodail8", "Ted", ""); err != nil {
		t.Errorf("Ted wasn't insert in database: %s", err)
	}

	dataBase.DB.Exec("DROP TABLE persons")
}

func Test2(t *testing.T) {

}
