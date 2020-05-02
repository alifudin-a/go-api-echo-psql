package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "postgres"
	dbname   = "books_database"
	sslmode  = "disable"
)

// OpenDB : open database connection
func OpenDB() *sql.DB {
	psqlInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", user, password, dbname, sslmode)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	} else {
		log.Println(" Database Successfully Connected!")
	}

	return db
}
