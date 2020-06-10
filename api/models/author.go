package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Author model
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetAuthors get all authors
func GetAuthors() (*[]*Author, errors.DatabaseError) {
	statement := `
		select id, name
		from authors`
	rows, err := database.Db.Query(statement)
	if err != nil {
		return nil, &errors.InternalDatabaseError{Message: err.Error()}
	}
	defer rows.Close()

	var authors []*Author
	for rows.Next() {
		var author Author
		if err = rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, &errors.InternalDatabaseError{Message: err.Error()}
		}
		authors = append(authors, &author)
	}

	return &authors, nil
}

// GetAuthor get single author by id
func GetAuthor(id int) (*Author, errors.DatabaseError) {
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
		return nil, &errors.IDNotFoundError{TableName: "authors", ID: id}
	default:
		return nil, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// Save insert new record
func (a *Author) Save() errors.DatabaseError {
	statement := `
		insert into authors (id, name)
		values (default, $1)`
	if _, err := database.Db.Exec(statement, a.Name); err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}
	return nil
}

// Update updates the author
func (a *Author) Update() errors.DatabaseError {
	statement := `
		update authors
		set name = $1
		where id = $2`
	result, err := database.Db.Exec(statement, a.Name, a.ID)
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}
	if rowsAffected == 0 {
		return &errors.IDNotFoundError{TableName: "authors", ID: a.ID}
	}
	return nil
}

// Delete author
func (a *Author) Delete() errors.DatabaseError {
	statement := `
		delete from authors
		where id = $1`
	result, err := database.Db.Exec(statement, a.ID)
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}
	if rowsAffected == 0 {
		return &errors.IDNotFoundError{TableName: "authors", ID: a.ID}
	}
	return nil
}
