package errors

import "net/http"

// InternalDatabaseError - used when getting a row by id
type InternalDatabaseError struct {
	Message string
}

func (e *InternalDatabaseError) Error() string {
	return "Internal database error"
}

// Code - implements code
func (e *InternalDatabaseError) Code() int {
	return http.StatusInternalServerError
}
