package models

import "time"

type Person struct {
	Id        int
	Email     string
	Passwor   string
	FirstName string
	LastName  string
}

type Event struct {
	Id      int
	Date    time.Time
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
