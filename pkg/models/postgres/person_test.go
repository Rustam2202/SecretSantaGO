package postgres

import (
	"testing"

	_ "github.com/lib/pq"
)

func Test1(t *testing.T) {
	model := PersonModel{}
	err := model.openDb()
	if err != nil {
		t.Error(err)
	}

	model.CreateTablePersons()

	if err = model.Insert("e@mail.com", "pass1234", "John", "Doe"); err != nil {
		t.Errorf("John Doe wasn't insert in database: %s", err)
	}

	if err = model.Insert("em@ail.com", "passwrd123", "Marry", "Sue"); err != nil {
		t.Errorf("Mary Sue wasn't insert in database: %s", err)
	}
	if err = model.Insert("e-m@il.com", "crocodail8", "Ted", ""); err != nil {
		t.Errorf("Ted wasn't insert in database: %s", err)
	}

	model.DB.Exec("DROP TABLE persons")
}

func Test2(t *testing.T){
	model:=PersonModel{}
	model.createTable()
	model.Insert("e@mail.com", "pass1234", "John", "Doe")
}
