package models

import (
	"database/sql"

	"github.com/VIVelev/goApod/database"
)

//Location model
type Location struct {
	ID   int     `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

//GetLocationByID gives the Location object which matches the provided ID
func GetLocationByID(ID int) (*Location, error) {
	var location Location
	sqlStatement := `
	SELECT * FROM locations
	WHERE id = $1`
	row := database.Db.QueryRow(sqlStatement, ID)
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
func (l *Location) Save() error {
	sqlStatement := `
	INSERT INTO locations (lat, long)
	VALUES ($1, $2)`
	if _, err := database.Db.Exec(sqlStatement, l.Lat, l.Long); err != nil {
		return err
	}
	return nil
}

//DeleteLocationByID deletes the location with id that matches the provided ID
func DeleteLocationByID(ID int) error { // TO DO: VIVelev's custom error
	sqlStatement := `
	DELETE FROM locations
	WHERE id = $1;`
	res, err := database.Db.Exec(sqlStatement, ID)
	if err != nil {
		return err // database error
	}
	if _, err := res.RowsAffected(); err != nil {
		return err // Invalid id error
	}
	return nil
}
