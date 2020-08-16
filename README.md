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
Try to run `example/main.go` to see the result

```go
package main

import (
	"fmt"

	"github.com/firmanmm/suberror"
)

var (
	RandomError         = suberror.RuntimeError.Derive()
	BranchOfRandomError = RandomError.Derive()
	DerivedRandomError  = RandomError.Derive()
)

func main() {
	runtimeError := suberror.RuntimeError.New("A runtime Error")
	randomError := RandomError.New("A random error")
	branchError := BranchOfRandomError.New("A brnach error")
	derivedErr := DerivedRandomError.New("a derived error")

	fmt.Println("Start of RuntimeError")
	printError(runtimeError)
	fmt.Println("Start of RandomError")
	printError(randomError)
	fmt.Println("Start of BranchOfRandomError")
	printError(branchError)
	fmt.Println("Start of DerivedRandomError")
	printError(derivedErr)
}

func printError(err suberror.Error) {
	if err.TypeOf(suberror.RuntimeError) {
		fmt.Printf("Yay i'am a RuntimeError error %s\n", err.Error())
	}
	if err.TypeOf(RandomError) {
		fmt.Printf("Yay i'am a RandomError error %s\n", err.Error())
	}
	if err.TypeOf(BranchOfRandomError) {
		fmt.Printf("Yay i'am a BranchOfRandomError error %s\n", err.Error())
	}
	if err.TypeOf(DerivedRandomError) {
		fmt.Printf("Yay i'am a DerivedRandomError error %s\n", err.Error())
	}
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