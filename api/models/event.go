package models

import "time"

//Event struct
type Event struct {
	ID         int
	Date       time.Time
	LocationID int
}

//GetEvents gives you all the recorded events
func GetEvents() (*[]*Event, error) {
	return nil, nil
}

//GetEventByID gives you the event which maches the provided ID
func GetEventByID(ID int) (*Event, error) {
	return nil, nil
}

func (event *Event) save() error {
	return nil
}
