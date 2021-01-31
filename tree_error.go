package suberror

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

var (
	codeCounter ErrorCode
	//Maintain synchronization between call
	codeCounterMutex *sync.Mutex = &sync.Mutex{}
	//Store used code to provide custom error code capabilities
	codeMap map[ErrorCode]bool = make(map[ErrorCode]bool)
)

//ErrorCode define the error code
//Ex: 1234
type ErrorCode uint

type PreErrorTypeFunc func(err Error)

//ErrorType represent error type
type ErrorType interface {
	TypeOf(err ErrorType) bool
	New(message string) Error
	Newf(message string, args ...interface{}) Error
	GetCode() ErrorCode
	Derive() ErrorType
	DeriveWithCode(code uint) ErrorType
	getParent() ErrorType
	SetPreNewError(preFunc PreErrorTypeFunc)
}

//BaseErrorType for handling further error
type BaseErrorType struct {
	parent     ErrorType
	code       ErrorCode
	preNewFunc PreErrorTypeFunc
	family     map[ErrorCode]ErrorType
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
	if t.preNewFunc != nil {
		t.preNewFunc(err)
	}
	return err
}

//Newf instance of formatted error
func (t *BaseErrorType) Newf(format string, args ...interface{}) Error {
	return t.New(fmt.Sprintf(format, args...))
}

//Derive a new BaseErrorType from this error type
func (t *BaseErrorType) Derive() ErrorType {
	errType := newBaseErrorType()
	pickedCode := ErrorCode(0)
	for pickedCode == 0 {
		//Safeguard against improper use
		func() {
			codeCounterMutex.Lock()
			defer codeCounterMutex.Unlock()
			codeCounter++
			if _, ok := codeMap[codeCounter]; !ok {
				codeMap[codeCounter] = true
				pickedCode = codeCounter
			}
		}()
	}
	errType.code = pickedCode
	errType.preNewFunc = t.preNewFunc
	errType.parent = t
	for i, v := range t.family {
		errType.family[i] = v
	}
	errType.family[t.code] = t
	errType.family[pickedCode] = errType
	return errType
}

//Derive a new BaseErrorType from this error type
func (t *BaseErrorType) DeriveWithCode(code uint) ErrorType {
	errType := newBaseErrorType()
	pickedCode := ErrorCode(code)
	codeCounterMutex.Lock()
	defer codeCounterMutex.Unlock()
	if _, ok := codeMap[pickedCode]; ok {
		panic(errors.New("Code already used"))
	}
	codeMap[pickedCode] = true
	errType.code = pickedCode
	errType.preNewFunc = t.preNewFunc
	errType.parent = t
	for i, v := range t.family {
		errType.family[i] = v
	}
	errType.family[t.code] = t
	errType.family[pickedCode] = errType
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

//SetPreNewError will be executed before the created error returned
func (t *BaseErrorType) SetPreNewError(preFunc PreErrorTypeFunc) {
	t.preNewFunc = preFunc
}

func newBaseErrorType() *BaseErrorType {
	errType := new(BaseErrorType)
	errType.family = make(map[ErrorCode]ErrorType)
	errType.family[codeCounter] = errType
	return errType
}
