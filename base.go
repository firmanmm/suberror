package suberror

import (
	"fmt"
	"runtime/debug"
)

var (
	//RuntimeError a global error that is categorized as runtime error
	RuntimeError = newBaseErrorType()

	//ClientError represent any error due to client input
	ClientError = RuntimeError.Derive()

	//InternalError represent any error that happen because of server problem
	//Comes pre setup with a stack trace generated
	InternalError = RuntimeError.Derive()
)

func init() {
	InternalError.SetPreNewError(func(err Error) {
		stackTrace := debug.Stack()
		message := fmt.Sprintf("Error : %s\n%s", err.Error(), string(stackTrace))
		err.setMessage(message)
	})
}
