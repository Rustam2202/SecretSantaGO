package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


func CreatePersonsTable(db *sql.DB) {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		firstname TEXT,
		lastname TEXT)`,
		"persons")
	statement, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func CreateGiftedTable(db *sql.DB) {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		from_person INTEGER PRIMARY KEY,
		to_person INTEGER,
		year INTEGER,
		gift TEXT)`,
		"giftinfo")
	statement, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec()
}

func AddPerson(db *sql.DB, person Person) {
	statement, err := db.Prepare(`INSERT INTO persons (
		firstname, lastname) 
		VALUES (?, ?)`)
	if err != nil {
		panic(err)
	}
	statement.Exec(person.firstName, person.lastName)
}

func AddGiftedInfo(db *sql.DB, info GiftedInfo) {
	statement, err := db.Prepare(`INSERT INTO persons (
		from_person, to_person, year, gift) 
		VALUES (?, ?, ?, ?)`)
	if err != nil {
		panic(err)
	}
	statement.Exec(info.from_person, info.to_person, info.year, info.gift)
}

type Person struct {
	//ID        uint
	firstName string
	lastName  string
}

type GiftedInfo struct {
	from_person uint
	to_person   uint
	year        uint
	gift        string
}

func main() {
	database, _ := sql.Open("sqlite3", "./party.db")
	CreatePersonsTable(database)
	CreateGiftedTable(database)
	AddPerson(database, Person{"Jhon", "Doe"})
}
