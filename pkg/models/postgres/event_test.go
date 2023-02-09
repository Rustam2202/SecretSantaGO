package postgres

import (
	"testing"
)

func TestNewEvent(t *testing.T) {

	model := Models{}
	model.openDb()

	model.PersonModel.Insert("e@mail.com", "pass1234", "John", "Doe")
	model.PersonModel.Insert("em@ail.com", "passwrd123", "Marry", "Sue")
	model.PersonModel.Insert("e-m@il.com", "crocodail8", "Ted", "")
}
