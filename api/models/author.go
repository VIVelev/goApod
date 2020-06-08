package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
)

// Author model
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetAuthors get all authors
func GetAuthors() (*[]*Author, error) {
	statement := `
		select id, name
		from authors`
	rows, err := database.Db.Query(statement)
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
func GetAuthor(id int) (*Author, error) {
	statement := `
		select id, name
		from authors
		where id = $1`
	row := database.Db.QueryRow(statement, id)
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

// Save insert new record
func (a *Author) Save() error {
	statement := `
		insert into authors (id, name)
		values (default, $1)`
	if _, err := database.Db.Exec(statement, a.Name); err != nil {
		return err
	}
	return nil
}

// Update updates the author
func (a *Author) Update() (bool, error) {
	statement := `
		update authors
		set name = $1
		where id = $2`
	result, err := database.Db.Exec(statement, a.Name, a.ID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

// Delete author
func (a *Author) Delete() (bool, error) {
	statement := `
		delete from authors
		where id = $1`
	result, err := database.Db.Exec(statement, a.ID)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
