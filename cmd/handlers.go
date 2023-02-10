package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.tpl.ExecuteTemplate(w, "index.html", nil)
}

func (app *Application) registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** registerHandler running ***")
	app.tpl.ExecuteTemplate(w, "register.html", nil)
}

func (app *Application) registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** registerAuthHandler running ***")
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	hashPassw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	app.db.AddPerson(email, firstName, lastName, hashPassw, nil, nil)
	app.tpl.ExecuteTemplate(w, "index.html", nil)
}