package main

import (
	"fmt"
	"net/http"
	"santa/pkg/models"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("cookie-store"))

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.tpl.ExecuteTemplate(w, "index.html", nil)
}

func (app *Application) loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginHandler running*****")
	app.tpl.ExecuteTemplate(w, "login.html", nil)
}

func (app *Application) loginAuthHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("*** loginAuthHandler running ***")
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		return
	}

	var hashPassw []byte
	stmt := `SELECT password FROM persons WHERE email = $1`
	row := app.db.DB.QueryRow(stmt, email)
	err := row.Scan(&hashPassw)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(hashPassw, []byte(password))
	if err == nil {
		session, _ := store.Get(r, "session")
		session.Values["email"] = email
		session.Save(r, w)
		app.tpl.ExecuteTemplate(w, "index.html", "Log In")
		return
	}
	app.tpl.ExecuteTemplate(w, "login.html", "Incorrect email or password")
}

func (app *Application) registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** registerHandler running ***")
	app.tpl.ExecuteTemplate(w, "register.html", nil)
}

func (app *Application) registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*** registerAuthHandler running ***")
	r.ParseForm()
	email := r.FormValue("email")
	if app.db.IsPersonExist(email) {
		app.tpl.ExecuteTemplate(w, "register.html", "This email already exist")
		return
	}
	password := r.FormValue("password")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	hashPassw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	app.db.AddPersonToDB(email, firstName, lastName, hashPassw, nil, nil)
	session, _ := store.Get(r, "session")
	session.Values["email"] = email
	session.Save(r, w)
	app.tpl.ExecuteTemplate(w, "index.html", "Log In")
}

func (app *Application) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "email")
	session.Save(r, w)
	app.tpl.ExecuteTemplate(w, "login.html", "Logged Out")
}

func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		_, ok := session.Values["email"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		HandlerFunc.ServeHTTP(w, r)
	}
}

var event models.Event

func (app *Application) newEvent(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	event.Name = r.FormValue("name")
//	app.tpl.ExecuteTemplate(w, "event.html", nil)
}

func (app *Application) addPersonToEventHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	event.Persons = append(event.Persons, models.Person{Email: r.FormValue("person")})
	fmt.Println("Person added")
}
