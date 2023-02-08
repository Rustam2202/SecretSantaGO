package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"sqlite-golang/pkg/models/postgres"

	_ "github.com/lib/pq"
	//"golang.org/x/crypto/bcrypt"
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

var tplExt *template.Template
var DBextern *sql.DB
var Persons *postgres.PersonModel

func main() {
	openDB()
	//AddPerson(database, Person{"Jhon", "Doe"})
	//AddGiftedInfo(database, GiftedInfo{1, 2, 2011, "True Gift"})

	tplExt, _ = template.ParseGlob("ui/templates/*.html")

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerAuth", registerAuthHandler)
	fmt.Println("*** Listen and serve ***")
	http.ListenAndServe("localhost:8080", nil)
}

func openDB() {
	connStr := "user=postgres password=postgres dbname=SecretSantaDB sslmode=disable"
	var err error
	DBextern, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplExt.ExecuteTemplate(w, "index.html", nil)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** registerHandler running ***")
	tplExt.ExecuteTemplate(w, "register.html", nil)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** registerAuthHandler running ***")
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	Persons.Insert(email, password, firstName, lastName)
	fmt.Println(email, password, firstName, lastName)
	//	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}
