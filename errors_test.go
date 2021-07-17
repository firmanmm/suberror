package suberror

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseError(t *testing.T) {
	runtimeErr := RuntimeError.New("system crashed")
	if !runtimeErr.TypeOf(RuntimeError) {
		t.Error("runtimeErr is not a subtype of RuntimeError")
	}
	ioErrType := InternalError.Derive()
	ioErr := ioErrType.New("there is an IO error")
	if !ioErr.TypeOf(RuntimeError) {
		t.Error("ioErr is not a subtype of RuntimeError")
	}
	if !ioErr.TypeOf(ioErrType) {
		t.Error("ioErr is not a subtype of IOError")
	}

	netErrType := ioErrType.Derive()
	netErr := netErrType.New("there is an Network error")
	if !netErr.TypeOf(RuntimeError) {
		t.Error("netErr is not a subtype of RuntimeError")
	}
	if !netErr.TypeOf(ioErrType) {
		t.Error("netErr is not a subtype of IOError")
	}
	if !netErr.TypeOf(netErrType) {
		t.Error("netErr is not a subtype of NetworkError")
	}
}

func TestDeriveError(t *testing.T) {
	testcase := make([]ErrorType, 100)
	testcase[0] = RuntimeError.Derive()
	for i := 1; i < len(testcase); i++ {
		testcase[i] = testcase[i-1].Derive()
	}

	for i := 0; i < len(testcase); i++ {
		for j := i; j < len(testcase); j++ {
			if !testcase[j].TypeOf(testcase[i]) {
				t.Fatalf("testcase %d is not a subtype of %d", j, i)
			}
		}
	}
}

func TestDeriveWithCode(t *testing.T) {
	func() {
		defer func() {
			err := recover()
			assert.Nil(t, err)
		}()
		customErrorCodeType := RuntimeError.DeriveWithCode(400)
		assert.Equal(t, ErrorCode(400), customErrorCodeType.GetCode())
	}()
	func() {
		defer func() {
			err := recover()
			assert.Error(t, err.(error))
		}()
		RuntimeError.DeriveWithCode(400)
	}()
}

func TestTryCatchLikeError(t *testing.T) {
	ioErrType := InternalError.Derive()
	netErrType := ioErrType.Derive()
	subIOError := ioErrType.Derive()
	err := subIOError.New("there was an sub IO error")
	var res ErrorType
	switch true {
	case err.TypeOf(netErrType):
		fmt.Println("Im Network error")
	case err.TypeOf(ioErrType):
		fmt.Println("Im IOError error")
	case err.TypeOf(RuntimeError):
		fmt.Println("Im Runtime error")
	default:
		t.Fatal("non matching error") //Something wrong
	}
	if !err.TypeOf(ioErrType) {
		t.Errorf("got %v want %v", res.GetCode(), err.GetCode())
	}
}

func TestSimple(t *testing.T) {
	err := InternalError.New("Internal")
	assert.False(t, err.TypeOf(ClientError))
	assert.True(t, err.TypeOf(RuntimeError))

	err = ClientError.New("Client")
	assert.False(t, err.TypeOf(InternalError))
	assert.True(t, err.TypeOf(RuntimeError))
}
