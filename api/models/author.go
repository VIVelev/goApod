package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
)

// Author model
type Author struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetAuthors get all authors
func GetAuthors() (*[]*Author, error) {
	rows, err := database.Db.Query("select id, name from authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*Author

	for rows.Next() {
		var author Author
		if err = rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}

		authors = append(authors, &author)
	}

	return &authors, nil
}

// GetAuthor get single author by id
func GetAuthor(id uint) (*Author, error) {
	row := database.Db.QueryRow("select id, name from authors where id=$1", id)
	var author Author
	switch err := row.Scan(&author.ID, &author.Name); err {
	case nil:
		return &author, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}
