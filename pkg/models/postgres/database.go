package postgres

import "database/sql"

type DataBase struct {
	DB *sql.DB
}

func (db *DataBase) openDb() error {
	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	db.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	return nil
}
