# Suberror
A Tree like error handling for golang because i need better Error handling in golang. Give you support to tree like error handling that you can find in C++, C# and Java programming language.

## Installation
```bash
go get github.com/firmanmm/suberror
```

## Performance
```
goos: windows
goarch: amd64
pkg: github.com/firmanmm/suberror
BenchmarkDeriveErrorDepth10-8                 	  131271	      8894 ns/op	    4881 B/op	      37 allocs/op
BenchmarkDeriveErrorDepth25-8                 	   22488	     53296 ns/op	   31423 B/op	     120 allocs/op
BenchmarkDeriveErrorDepth50-8                 	    6004	    202846 ns/op	  124388 B/op	     309 allocs/op
BenchmarkDeriveErrorDepth100-8                	    1484	    819404 ns/op	  527456 B/op	     839 allocs/op
BenchmarkCheckTypeErrorDepth100TargetRoot-8   	66721897	        19.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkCheckTypeErrorDepth100Target50-8     	66850876	        19.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkCheckTypeErrorDepth100TargetLeaf-8   	80351941	        14.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/firmanmm/suberror	9.789s
```
As you can see here that `suberror` provides a pretty good error handling. Also deriving from error is not really that bad. Checking from `Leaf` to `Root` tooks only `19.9ns` per operation which is pretty amazing. However, you can also see that deriving error into 100 depth (or inheritance in OOP) tooks `819404ns` which is good enough i think. Especially that it should only happen once at program initialization.

## Usage
Try to expore [example](example) to for more

```go
package main

import (
	"fmt"
	"log"

	"github.com/firmanmm/suberror"
)

var (
	AccessError      = suberror.ClientError.Derive()
	TransactionError = suberror.ClientError.Derive()
	PaymentError     = TransactionError.Derive()
)

func main() {
	err := doBusinessStuff(1, 5)
	if err != nil {
		handleError(err)
		//No Access : Not worth my time
	}
	err = doBusinessStuff(5, 10)
	if err != nil {
		handleError(err)
		//2020/10/25 14:47:34 Insufficient Fund
	}
	err = doBusinessStuff(101, 10)
	if err != nil {
		handleError(err)
		//2020/10/25 14:47:34 Balance Too high
	}
	err = doBusinessStuff(35, 11)
	if err != nil {
		handleError(err)
		/*
			Error : End of Transaction error
			goroutine 1 [running]:
			runtime/debug.Stack(0xc00000e2c0, 0xc000092030, 0x5904c0)
					c:/go/src/runtime/debug/stack.go:24 +0xa4
			github.com/firmanmm/suberror.init.0.func1(0x4fee60, 0xc000004560)
					D:/Projects/Rendoru/OpenSource/suberror/base.go:27 +0x2d
			github.com/firmanmm/suberror.(*BaseErrorType).New(0xc00006c3f0, 0x4e5aed, 0x18, 0x4bb040, 0xc000046210)
					D:/Projects/Rendoru/OpenSource/suberror/tree_error.go:48 +0xb2
			main.doBusinessStuff(0x23, 0xb, 0x4fee60, 0xc000004540)
					D:/Projects/Rendoru/OpenSource/suberror/example/request/main.go:62 +0xc9
			main.main()
					D:/Projects/Rendoru/OpenSource/suberror/example/request/main.go:29 +0xb1
		*/
	}
}

func handleError(err suberror.Error) {
	switch true {
	case err.TypeOf(AccessError):
		fmt.Println("No Access : " + err.Error())
	case err.TypeOf(TransactionError):
		//In this case we want to log instead of print the error
		log.Println(err.Error())
	default:
		fmt.Println(err.Error())
	}
}

func doBusinessStuff(money, cost int) suberror.Error {
	//Assume that we can only process business deal when there is atleast 10 coin involved
	if cost < 10 {
		return AccessError.New("Not worth my time")
	}
	//We need to make sure if they have enough coin or not
	if money < cost {
		return TransactionError.New("Insufficient Fund")
	}

	//Let's also assume that we can't accept payment if their balance is higher than 10 times the cost
	if money > cost*10 {
		return PaymentError.New("Balance Too high")
	}
	//Assume the transaction is buggy
	return suberror.InternalError.New("End of Transaction error")
}

```
## Tricks
`Note :` Deriving inside function is a really bad idea, you won't see it done in `Java` as it's not a good idea to do it since it's usually performed at compile time.


Example of "Try Catch" error handling in golang:
```go
subIOError := IOError.Derive()
err := subIOError.New("there was an sub IO error")
switch true {
case err.TypeOf(NetworkError):
    res = NetworkError //Should skip
case err.TypeOf(IOError):
    res = IOError //Should get
case err.TypeOf(RuntimeError):
    res = RuntimeError //Should not get
default:
    log.Fatal("non matching error") //Something wrong
}
```
For more example please see errors_test.go