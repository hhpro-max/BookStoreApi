package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nimda"
	password = "admin"
	dbname   = "bookstore"
)

var db *sqlx.DB

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	schema := `CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title TEXT,
		author TEXT,
		price REAL
	);`
	db.MustExec(schema)
}
