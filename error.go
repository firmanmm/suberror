package suberror

//Error represent an error message
type Error struct {
	message string
	errType *ErrorType
}

//TypeOf perform check if this error is part of other ErrorType
func (e *Error) TypeOf(err *ErrorType) bool {
	return e.errType.TypeOf(err)
}

//Error return representative error message
func (e *Error) Error() string {
	return e.message
}
