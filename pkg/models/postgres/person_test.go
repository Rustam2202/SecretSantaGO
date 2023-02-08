package postgres

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func Test1(t *testing.T) {
	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	model := PersonModel{}
	model.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	model.CreateTablePersons()

	if err = model.Insert("e@mail.com", "pass1234", "John", "Doe"); err != nil {
		t.Errorf("John Doe wasn't insert in database: %s", err)
	}
	model.Insert("em@ail.com", "passwrd123", "Marry", "Sue")
	model.Insert("e-m@il.com", "crocodail8", "Ted", "")

	model.DB.Exec("DROP TABLE persons")
}
