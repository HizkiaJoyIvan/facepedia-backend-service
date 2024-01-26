package main

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
)

var db *sql.DB

func connectDB() {
	connStr := "user=<your_postgres_user> password=<your_postgres_password> dbname=<your_postgres_dbname> sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes() *mux.Router {
	r := mux.NewRouter()
}
