package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

// Event struct
type Event struct {
	ID         int    `json:"id,string,omitempty"`
	Name       string `json:"name"`
	Date       string `json:"date"`
	LocationID int    `json:"locationId,string,omitempty"`
}

var eventStatements = map[string]string{
	"GetEventByID": `
		select *
		from events
		where id = $1`,

	"Save": `
		insert into events
		values (default, $1, $2, $3)
		returning id`,
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

// Save (insert) the event
func (e *Event) Save() errors.DatabaseError {
	id := 0
	if err := database.Db.QueryRow(eventStatements["Save"],
		e.Name, e.Date, e.LocationID).Scan(&id); err != nil {

		return &errors.InternalDatabaseError{Message: err.Error()}
	}

	e.ID = id
	return nil
}
