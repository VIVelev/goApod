package models

import (
	"database/sql"
	"net/http"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Author model
type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var authorStatements = map[string]string{
	"GetAllAuthors": `
		select *
		from authors`,

	"GetAuthorByID": `
		select *
		from authors
		where id = $1`,

	"GetAuthorByName": `
		select *
		from authors
		where name = $1`,

	"Save": `
		insert into authors
		values (default, $1)`,

	"Update": `
		update authors
		set name = $1
		where id = $2`,

	"DeleteAuthorByID": `
		delete from authors
		where id = $1`,
}

// GetAllAuthors get all authors
func GetAllAuthors() ([]Author, errors.DatabaseError) {
	rows, err := database.Db.Query(authorStatements["GetAllAuthors"])
	defer rows.Close()

	if err != nil {
		return nil, &errors.InternalDatabaseError{Message: err.Error()}
	}

	var authors []Author
	for rows.Next() {
		var author Author
		if err = rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, &errors.InternalDatabaseError{Message: err.Error()}
		}

		authors = append(authors, author)
	}

	return authors, nil
}

// GetAuthorByID get single author by id
func GetAuthorByID(id int) (Author, errors.DatabaseError) {
	row := database.Db.QueryRow(authorStatements["GetAuthorByID"], id)
	var ret Author

	switch err := row.Scan(&ret.ID, &ret.Name); err {
	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.IDNotFoundError{TableName: "authors", ID: id}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// GetAuthorByName - get author by name
func GetAuthorByName(name string) (Author, errors.DatabaseError) {
	row := database.Db.QueryRow(authorStatements["GetAuthorByName"], name)
	var ret Author

	switch err := row.Scan(&ret.ID, &ret.Name); err {
	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.GenericError{HTTPCode: http.StatusNotFound, Message: "Invalid name"}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// Save insert new record
func (a *Author) Save() errors.DatabaseError {
	if _, err := database.Db.Exec(authorStatements["Save"], a.Name); err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	return nil
}

// Update updates the author
func (a *Author) Update() errors.DatabaseError {
	result, err := database.Db.Exec(authorStatements["Update"], a.Name, a.ID)
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
func (a *Author) Delete() error {
	return DeleteAuthorByID(a.ID)
}

// DeleteAuthorByID - delete author by id
func DeleteAuthorByID(id int) errors.DatabaseError {
	result, err := database.Db.Exec(authorStatements["DeleteAuthorByID"], id)
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}
	if rowsAffected == 0 {
		return &errors.IDNotFoundError{TableName: "authors", ID: id}
	}

	return nil
}
