package models

import "time"

//Event struct
type Event struct {
	ID         int
	Date       time.Time
	LocationID int
}

//GetEvents gives you all the recorded events
func GetEvents() (*[]*Events, error) {
	return nil, nil
}

//GetEventByID gives you the event which maches the provided ID
func GetEventByID(ID int) (*Event, error) {
	return nil, nil
}

func (event *Event) save() error {
	sqlStatement := `
		INSERT INTO users (date, email, first_name, last_name)
		VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, 30, "jon@calhoun.io", "Jonathan", "Calhoun")
	if err != nil {
		return err
	}
	return nil
}
