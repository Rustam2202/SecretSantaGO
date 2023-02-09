package postgres

import "database/sql"

type DataBase struct {
	DB *sql.DB
}

type Databaser interface {
	createTable()
	insert()
	add()
}

type Models struct {
	DB          *sql.DB
	PersonModel PersonModel
	EventModel  EventModel
	pi          PersonInteface
	ei          EventInterface
}

func (model *Models) openDb() error {
	model.pi.creatTable()


	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	model.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}
