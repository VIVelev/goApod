package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
	"github.com/VIVelev/goApod/errors"
)

//Location model
type Location struct {
	ID   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

var locationStatements = map[string]string{
	"GetLocationByID": `
		SELECT * FROM locations
		WHERE id = $1`,
	"Save": `
		INSERT INTO locations (lat, long)
		VALUES ($1, $2)`,
	"DeleteLocationByID": `
		DELETE FROM locations
		WHERE id = $1`,
}

//GetLocationByID gives the Location object which matches the provided ID
func GetLocationByID(ID int) (*Location, error) {
	var location Location
	row := database.Db.QueryRow(locationStatements["GetLocationByID"], ID)
	err := row.Scan(&location.ID, &location.Lat, &location.Long)
	if err != nil {
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	} else {
		return &location, nil
	}
}

//Save - saves the Location object on which you call this method
func (l *Location) Save() (*Location, error) {
	row := database.Db.QueryRow(locationStatements["Save"], l.Lat, l.Long)
	err := row.Scan(&l.ID, &l.Lat, &l.Long)
	if err != nil {
		return nil, err
	}
	return l, nil
}

//DeleteLocationByID deletes the location with id that matches the provided ID
func DeleteLocationByID(ID int) error {
	res, err := database.Db.Exec(locationStatements["DeleteLocationByID"], ID)
	if err != nil {
		return err
	}
	if _, err := res.RowsAffected(); err != nil {
		return &errors.IDNotFoundError{TableName: "Locations", ID: ID}
	}
	return nil
}
