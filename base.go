package suberror

import (
	"fmt"
	"runtime/debug"
)

//RuntimeError a global error that is categorized as runtime error
var RuntimeError = newBaseErrorType()

//ClientError represent any error due to client input
var ClientError = RuntimeError.Derive()

//InternalError represent any error that happen because of server problem
//Comes pre setup with a stack trace generated
var InternalError ErrorType

//IOError represent any IO related error
var IOError ErrorType

//NetworkError represent any network related error
var NetworkError ErrorType

func init() {
	InternalError = RuntimeError.Derive()
	InternalError.SetPreNewError(func(err Error) {
		stackTrace := debug.Stack()
		message := fmt.Sprintf("Error : %s\n%s", err.Error(), string(stackTrace))
		err.setMessage(message)
	})
	IOError = InternalError.Derive()
	NetworkError = IOError.Derive()
}
