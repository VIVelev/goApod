package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Event struct
type Event struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	LocationID int    `json:"locationId"`
}

var eventStatements = map[string]string{
	"GetEventByID": `
		select *
		from events
		where id = $1`,

	"Save": `
		insert into events
		values (default, $1, $2, $3)`,
}

// GetEventByID gives you the event which maches the provided ID
func GetEventByID(id int) (Event, errors.DatabaseError) {
	row := database.Db.QueryRow(eventStatements["GetEventByID"], id)
	var ret Event

	switch err := row.Scan(&ret.ID, &ret.Name,
		&ret.Date, &ret.LocationID); err {

	case nil:
		return ret, nil
	case sql.ErrNoRows:
		return ret, &errors.IDNotFoundError{TableName: "events", ID: id}
	default:
		return ret, &errors.InternalDatabaseError{Message: err.Error()}
	}
}

//Save saves the event
func (e *Event) Save() errors.DatabaseError {
	if _, err := database.Db.Exec(eventStatements["Save"],
		e.Name, e.Date, e.LocationID); err != nil {

		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	return nil
}
