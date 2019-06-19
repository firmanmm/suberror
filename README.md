# Suberror
A Tree like error handling for golang because i need better Error handling in golang. Give you support to tree like error handling that you can find in C++, C# and Java programming language.

## Usage
```
//Derive from RuntimeError
package main

import(
    "github.com/firmanmm/suberror"
)

func main() {
    randError := suberror.RuntimeError.Derive() //Derive from RuntimeError
    err := randError.New("An error")
    log.Println(err.Error()) //Print "An error"
    if err.TypeOf(suberror.RuntimeError) { //match since we derive from RuntimeError
        log.Println("Yay i'am a runtime error")
    }
    if err.TypeOf(randError) { //match since its a randError
        log.Println("Yay i'am a rand error")
    }
    reRandError := randError.Derive()
    newErr := reRandError.New("a new error")
    if newErr.TypeOf(suberror.RuntimeError) { //match since reRandError derive from randError which derive from RuntimeError
        ...
    }
}
```
## Example
Example of "Try Catch" error handling in golang:
```
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