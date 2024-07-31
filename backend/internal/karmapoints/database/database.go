package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect(connectionString string) {
	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection established")
}

func GetDB() *sql.DB {
	return db
}
