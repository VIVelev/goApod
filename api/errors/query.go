package errors

import "fmt"

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
