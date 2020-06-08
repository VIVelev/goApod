package models

import "github.com/VIVelev/goApod/database"

// Author model
type Author struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetAuthors get all authors
func GetAuthors() ([]Author, error) {
	rows, err := database.Db.Query("select id, name from authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []Author

	for rows.Next() {
		var id uint
		var name string
		if err = rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		authors = append(authors, Author{ID: id, Name: name})
	}

	return authors, nil
}
