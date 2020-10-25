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
