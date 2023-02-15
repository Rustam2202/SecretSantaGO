package models

type Person struct {
	Id        int
	Email     string
	Password  string
	FirstName string
	LastName  string
//	Community []int
}

type Event struct {
	Id      int
	Name    string
	Persons []int
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
