package main

import (
	"fmt"
	"html/template"
	"net/http"
	"santa/pkg/models/postgres"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type Application struct {
	db  *postgres.DataBase
	tpl *template.Template
}

func main() {
	var app Application
	app.db = &postgres.DataBase{}
	app.db.InitializDatabase()
	app.tpl, _ = template.ParseGlob("ui/templates/*.html")

	http.HandleFunc("/", app.homeHandler)
	http.HandleFunc("/register", app.registerHandler)
	http.HandleFunc("/registerAuth", app.registerAuthHandler)

	fmt.Println("*** Listen and serve ***")
	http.ListenAndServe("localhost:8080", nil)
}

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
	//fmt.Println(email, password, firstName, lastName)
	hashPassw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	app.db.AddPerson(email, firstName, lastName, hashPassw, nil, nil)
	// w.Write([]byte("You been registrated"))
	app.tpl.ExecuteTemplate(w, "index.html", nil)
}
