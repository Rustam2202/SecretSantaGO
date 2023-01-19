package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const (
	PersonsTable  string = "persons"
	GiftInfoTable string = "giftinfo"
)

func CreatePersonsTable(db *sql.DB) {
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY,
		firstname TEXT,
		lastname TEXT)`,
		PersonsTable)
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
		GiftInfoTable)
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
	query := fmt.Sprintf(`INSERT INTO %s (
		from_person, to_person, year, gift) 
		VALUES (?, ?, ?, ?)`, GiftInfoTable)
	statement, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	statement.Exec(info.from_person, info.to_person, info.year, info.gift)
}

func MakeEvent(db *sql.DB) {
	
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

type Event struct {
	year     uint
	persons  []Person
	database *sql.DB
}

func main() {
	database, _ := sql.Open("sqlite3", "./party.db")
	CreatePersonsTable(database)
	CreateGiftedTable(database)
	AddPerson(database, Person{"Jhon", "Doe"})
	AddGiftedInfo(database, GiftedInfo{1, 2, 2011, "True Gift"})
}
