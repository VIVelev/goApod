package models

import (
	"database/sql"
	"time"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

//Event struct
type Event struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	LocationID int       `json:"locationId"`
	ArticleID  int       `json:"articleId"`
}

var eventStatements = map[string]string{
	"GetEventByID": `
		SELECT * FROM events
		WHERE id = $1`,
	"Save": `
		INSERT INTO events (date, location_id, article_id)
		VALUES ($1, $2, $3)`,
}

//GetEventByID gives you the event which maches the provided ID
func GetEventByID(ID int) (*Event, error) {
	var event Event
	row := database.Db.QueryRow(eventStatements["GetEventByID"], ID)

	switch err := row.Scan(&event.ID, &event.Date,
		&event.LocationID, &event.ArticleID); err {

	case nil:
		return &event, nil
	case sql.ErrNoRows:
		return nil, &errors.IDNotFoundError{TableName: "events", ID: ID}
	default:
		return nil, err
	}
}

//Save saves the event
func (e *Event) Save() (*Event, error) {
	row := database.Db.QueryRow(
		eventStatements["Save"], e.Date, e.LocationID, e.ArticleID)

	err := row.Scan(&e.ID, &e.Date, &e.LocationID, &e.ArticleID)
	if err != nil {
		return nil, err
	}
	return e, nil
}
