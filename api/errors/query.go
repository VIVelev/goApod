package errors

import (
	"fmt"
	"net/http"
)

// IDNotFoundError - used when getting a row by id
type IDNotFoundError struct {
	TableName string // name of the SQL table
	ID        int    // the index that couldn't be found
}

func (e *IDNotFoundError) Error() string {
	return fmt.Sprintf(`
		IDNotFoundError: ID '%d' could not be found in table \"%s\"`,
		e.ID, e.TableName)
}

// Code - implements code
func (e *IDNotFoundError) Code() int {
	return http.StatusInternalServerError
}
