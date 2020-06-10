package errors

// DatabaseError - database error
type DatabaseError interface {
	error
	Code() int
}
