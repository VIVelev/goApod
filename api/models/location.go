package models

//Location model
type Location struct {
	ID   int
	Lat  float64
	Long float64
}

//GetLocation gives the Location object which matches the provided ID
func GetLocation(ID int) (*Location, error) {
	return nil, nil
}

//Save - saves the Location object on which you call this method
func (*Location) Save() error {
	return nil
}
