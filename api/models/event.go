package models

import (
	"database/sql"
	"time"

	"github.com/VIVelev/goApod/database"
)

//Event struct
type Event struct {
	ID         int       `json:"id"`
	Date       time.Time `json:"date"`
	LocationID int       `json:"locationId"`
	ArticleID  int       `json:"articleId"`
}

//GetEventByID gives you the event which maches the provided ID
func GetEventByID(ID int) (*Event, error) {
	var event Event
	sqlStatement := `
	SELECT * FROM events
	WHERE id = $1`
	row := database.Db.QueryRow(sqlStatement, ID)
	err := row.Scan(&event.ID, &event.Date, &event.LocationID, &event.ArticleID)
	if err != nil {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	} else {
		return &event, nil
	}
}

//Save saves the event
func (e *Event) Save() error {
	sqlStatement := `
	INSERT INTO events (date, location_id, article_id)
	VALUES ($1, $2, $3)`
	if _, err := database.Db.Exec(sqlStatement, e.Date, e.LocationID, e.ArticleID); err != nil {
		return err
	}
	return nil
}
