package suberror

import (
	"strconv"
)

var codeCounter ErrorCode

//ErrorCode define the error code
//Ex: 1234
type ErrorCode int

//ErrorType for handling further error
type ErrorType struct {
	parent *ErrorType
	code   ErrorCode
}

//TypeOf check wheter current error is subtype of [err]
//return true if valid
func (t *ErrorType) TypeOf(err *ErrorType) bool {
	iter := t
	for iter != nil {
		if iter.code == err.code {
			return true
		}
		iter = iter.parent
	}
	return false
}

//New instance of Error with this given type
func (t *ErrorType) New(message string) *Error {
	err := new(Error)
	err.errType = t
	err.message = message
	return err
}

//Derive a new ErrorType from this error type
func (t *ErrorType) Derive() *ErrorType {
	errType := newErrorType()
	errType.parent = t
	return errType
}

//String represent struct as string of error code
func (t *ErrorType) String() string {
	return strconv.Itoa(int(t.code))
}

func newErrorType() *ErrorType {
	errType := new(ErrorType)
	codeCounter++
	errType.code = codeCounter
	return errType
}
