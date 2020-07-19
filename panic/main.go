package main

import "fmt"

func main() {

	defer func() {
		message := recover()
		fmt.Println(message)
	}()
	panic("i am panic")
}
