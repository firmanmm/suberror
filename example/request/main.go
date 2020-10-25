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
