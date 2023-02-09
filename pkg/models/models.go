package models

type Person struct {
	Id        int
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type Event struct {
	Id      int
	Name    string
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
