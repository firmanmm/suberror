package suberror

import (
	"testing"
)

func TestBaseError(t *testing.T) {
	runtimeErr := RuntimeError.New("system crashed")
	expected := "system crashed"
	if runtimeErr.Error() != expected {
		t.Errorf("got %v want %v", runtimeErr, expected)
	}
	if !runtimeErr.TypeOf(RuntimeError) {
		t.Error("runtimeErr is not a subtype of RuntimeError")
	}

	ioErr := IOError.New("there is an IO error")
	expected = "there is an IO error"
	if ioErr.Error() != expected {
		t.Errorf("got %v want %v", ioErr, expected)
	}
	if !ioErr.TypeOf(RuntimeError) {
		t.Error("ioErr is not a subtype of RuntimeError")
	}
	if !ioErr.TypeOf(IOError) {
		t.Error("ioErr is not a subtype of IOError")
	}

	netErr := NetworkError.New("there is an Network error")
	expected = "there is an Network error"
	if netErr.Error() != expected {
		t.Errorf("got %v want %v", ioErr, expected)
	}
	if !netErr.TypeOf(RuntimeError) {
		t.Error("netErr is not a subtype of RuntimeError")
	}
	if !netErr.TypeOf(IOError) {
		t.Error("netErr is not a subtype of IOError")
	}
	if !netErr.TypeOf(NetworkError) {
		t.Error("netErr is not a subtype of NetworkError")
	}
}

func TestDeriveError(t *testing.T) {
	testcase := make([]*ErrorType, 100)
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

func TestTryCatchLikeError(t *testing.T) {
	subIOError := IOError.Derive()
	err := subIOError.New("there was an sub IO error")
	var res *ErrorType
	switch true {
	case err.TypeOf(NetworkError):
		res = NetworkError //Should skip
	case err.TypeOf(IOError):
		res = IOError //Should get
	case err.TypeOf(RuntimeError):
		res = RuntimeError //Should not get
	default:
		t.Fatal("non matching error") //Something wrong
	}
	if !err.TypeOf(IOError) {
		t.Errorf("got %v want %v", res, err.errType)
	}
}

func BenchmarkDeriveError(b *testing.B) {
	testcase := make([]*ErrorType, b.N)
	testcase[0] = RuntimeError.Derive()
	for i := 1; i < len(testcase); i++ {
		testcase[i] = testcase[i-1].Derive()
	}

	for i := 0; i < len(testcase); i++ {
		if !testcase[i].TypeOf(testcase[0]) {
			b.Fatalf("testcase %d is not a subtype of %d", 0, i)
		}

	}
}
