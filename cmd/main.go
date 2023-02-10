package main

import (
	"fmt"
	"html/template"
	"net/http"
	"santa/pkg/postgres"

	_ "github.com/lib/pq"
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

	// 'go run ./cmd' from root dir, thats compile both files (main.go and handlers.go)
	fmt.Println("*** Listen and serve ***")
	http.ListenAndServe("localhost:8080", nil)
}
