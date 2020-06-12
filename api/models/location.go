package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Location struct
type Location struct {
	ID   int     `json:"id,string,omitempty"`
	Lat  float64 `json:"lat,string,omitempty"`
	Long float64 `json:"long,string,omitempty"`
}

var locationStatements = map[string]string{
	"GetLocationByID": `
		select *
		from locations
		where id = $1`,

	"Save": `
		insert into locations
		values (default, $1, $2)
		returning id`,

	"DeleteLocationByID": `
		delete from locations
		where id = $1`,
}

// GetLocationByID gives the Location object which matches the provided id
func GetLocationByID(id int) (Location, errors.DatabaseError) {
	row := database.Db.QueryRow(locationStatements["GetLocationByID"], id)
	var ret Location

	switch err := row.Scan(&ret.ID, &ret.Lat, &ret.Long); err {
	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.IDNotFoundError{TableName: "locations", ID: id}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

// Save - saves the Location object on which you call this method
func (l *Location) Save() errors.DatabaseError {
	id := 0
	if err := database.Db.QueryRow(locationStatements["Save"], l.Lat, l.Long).Scan(&id); err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	l.ID = id
	return nil
}

// DeleteLocationByID deletes the location with id that matches the provided id
func DeleteLocationByID(id int) error {
	result, err := database.Db.Exec(locationStatements["DeleteLocationByID"], id)
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &errors.InternalDatabaseError{Message: err.Error()}
	}
	if rowsAffected == 0 {
		return &errors.IDNotFoundError{TableName: "locations", ID: id}
	}

	return nil
}
