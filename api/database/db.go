package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgresql
)

// Db conn
var Db *sql.DB

// Connect establishing a connection to the database
func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	fmt.Println("Successfully connected!")

	Db = db
	return nil
}
