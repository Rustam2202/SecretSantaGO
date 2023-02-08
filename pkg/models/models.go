package models


type Person struct {
	Id        int
	Email     string
	Passwor   string
	FirstName string
	LastName  string
}

type Event struct {
	Id      int
	name    string
	Persons []Person
}

type Gifts struct {
	PersonFrom Person
	PersonTo   Person
	Gift       string
	Event      Event
}

type Community struct {
	Persons []Person
	Events  []Event
}
