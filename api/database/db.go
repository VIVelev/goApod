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

// Prepare database
func Prepare() error {
	queries := []string{
		`create table if not exists authors (
			id serial primary key,
			name varchar not null
		)`,

		`create table if not exists locations (
			id serial primary key,
			lat double precision not null,
			long double precision not null
		)`,

		`create table if not exists articles (
			id serial primary key,
			title varchar not null,
			image_url varchar not null,
			text text not null,
			author_id int references authors(id) not null,
			date date not null,
			event_id int references events(id)
		)`,

		`create table if not exists events (
			id serial primary key,
			name varchar not null,
			date date not null,
			location_id int references locations(id)
		)`,

		`create table if not exists likes (
			id serial primary key,
			author_id int references authors(id) not null,
			article_id int references articles(id) not null
		)`,
	}

	for _, q := range queries {
		if _, err := Db.Exec(q); err != nil {
			return err
		}
	}

	return nil
}
