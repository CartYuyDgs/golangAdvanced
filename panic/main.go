package main

import (
	"fmt"
)

func main() {

	//defer func() {
	//	message := recover()
	//	fmt.Println(message)
	//}()
	defer receiveRecover()
	//panic("i am panic")
	//panic(errors.New("i am error"))
	panic(1)
}

func receiveRecover() {
	message := recover()
	fmt.Println(message)
	switch message.(type) {
	case string:
		fmt.Println("string message", message)
	case error:
		fmt.Println("err: ", message)
	default:
		fmt.Println(message)
	}
}
