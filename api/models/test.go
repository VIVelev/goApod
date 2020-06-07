package models

// Test struct
type Test struct {
	Info string
}

// GetInfo return
func (T *Test) GetInfo() string {
	return T.Info
}
