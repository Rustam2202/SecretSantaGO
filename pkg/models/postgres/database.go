package postgres

import "database/sql"

type DataBase struct {
	DB *sql.DB
}

func (db *DataBase) OpenDb() error {
	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	db.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) InitializDatabase() {
	db.OpenDb()
	db.createPersonsTable()
	db.createEventsTable()
	db.createGiftsTable()
	db.CreateCommunitiesTable()
}
