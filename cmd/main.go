package main

import (
	"fmt"
	"html/template"
	"net/http"
	"santa/pkg/postgres"

	"github.com/gorilla/context"
	_ "github.com/lib/pq"
)

type Application struct {
	db  *postgres.DataBase
	tpl *template.Template
}

// 'go run ./cmd' from root dir, thats compile both files (main.go and handlers.go)
func main() {
	var app Application
	app.db = &postgres.DataBase{}
	app.db.InitializDatabase()
	app.tpl, _ = template.ParseGlob("ui/templates/*.html")

	http.HandleFunc("/", Auth(app.homeHandler))
	http.HandleFunc("/register", app.registerHandler)
	http.HandleFunc("/registerAuth", app.registerAuthHandler)
	http.HandleFunc("/login", app.loginHandler)
	http.HandleFunc("/loginAuth", app.loginAuthHandler)
	http.HandleFunc("/logout", app.logoutHandler)
	http.HandleFunc("/event", app.newEvent)
	http.HandleFunc("/addPerson", app.addPersonToEventHandler)

	fmt.Println("*** Listen and serve ***")
	http.ListenAndServe("localhost:8080", context.ClearHandler(http.DefaultServeMux))
}
