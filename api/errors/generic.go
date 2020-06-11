package errors

// GenericError - used when getting error
type GenericError struct {
	Message  string
	HTTPCode int
}

func (e *GenericError) Error() string {
	return e.Message
}

// Code - implements code
func (e *GenericError) Code() int {
	return e.HTTPCode
}
