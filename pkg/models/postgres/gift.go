package postgres

func (db *DataBase) createGiftsTable() error {
	query := `CREATE TABLE IF NOT EXISTS gifts (
		personfrom INTEGER,
		personto INTEGER,
		gift TEXT,
		event INTEGER
		)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (db *DataBase) dropGiftsTable() error {
	query := `DROP TABLE gifts`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (db *DataBase) addGift(persFromId, persToId, eventId int, gift string) error {
	stmt := `INSERT INTO gifts (personfrom, personto, gift, event) VALUES($1, $2, $3, $4)`
	_, err := db.DB.Exec(stmt, persFromId, persToId, gift, eventId)
	if err != nil {
		return err
	}
	return nil
}
