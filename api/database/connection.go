package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // postgresql
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

// Connection conn
var Connection *sql.DB

// Connect establishing a connection to the database
func Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	fmt.Println("Successfully connected!")
	Connection = db
	return nil // no error !!!
}
