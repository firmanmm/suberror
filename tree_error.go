package suberror

import (
	"fmt"
	"strconv"
)

var codeCounter ErrorCode

//ErrorCode define the error code
//Ex: 1234
type ErrorCode int

//ErrorType represent error type
type ErrorType interface {
	TypeOf(err ErrorType) bool
	New(message string) Error
	Newf(message string, args ...interface{}) Error
	GetCode() ErrorCode
	Derive() ErrorType
	getParent() ErrorType
}

//BaseErrorType for handling further error
type BaseErrorType struct {
	parent ErrorType
	code   ErrorCode
	family map[ErrorCode]ErrorType
}

//TypeOf check wheter current error is subtype of [err]
//return true if valid
func (t *BaseErrorType) TypeOf(err ErrorType) bool {
	_, ok := t.family[err.GetCode()]
	return ok
}

//New instance of Error with this given type
func (t *BaseErrorType) New(message string) Error {
	err := new(BaseError)
	err.errType = t
	err.message = message
	return err
}

//Newf instance of formatted error
func (t *BaseErrorType) Newf(format string, args ...interface{}) Error {
	return t.New(fmt.Sprintf(format, args...))
}

//Derive a new BaseErrorType from this error type
func (t *BaseErrorType) Derive() ErrorType {
	errType := newBaseErrorType()
	errType.parent = t
	for i, v := range t.family {
		errType.family[i] = v
	}
	errType.family[t.code] = t
	return errType
}

//String represent struct as string of error code
func (t *BaseErrorType) String() string {
	return strconv.Itoa(int(t.code))
}

//GetCode return current error code
func (t *BaseErrorType) GetCode() ErrorCode {
	return t.code
}

func (t *BaseErrorType) getParent() ErrorType {
	return t.parent
}

func newBaseErrorType() *BaseErrorType {
	errType := new(BaseErrorType)
	codeCounter++
	errType.code = codeCounter
	errType.family = make(map[ErrorCode]ErrorType)
	errType.family[codeCounter] = errType
	return errType
}
