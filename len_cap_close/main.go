package main

import "fmt"

func main() {
	c := closeChan()
	fmt.Println(<-c)
}

func closeChan() chan int {
	c := make(chan int, 1)
	defer close(c)
	c <- 1
	return c
}

func lenAndCap() {
	s := make([]string, 1, 5)
	s[0] = "11"
	fmt.Println(s, len(s), cap(s))
	s = append(s, "aa")
	fmt.Println(s, len(s), cap(s))
	s = append(s, "aa1")
	fmt.Println(s, len(s), cap(s))
	s = append(s, "aa2")
	fmt.Println(s, len(s), cap(s))
	s = append(s, "aa3")
	fmt.Println(s, len(s), cap(s))
	s = append(s, "aa4")
	fmt.Println(s, len(s), cap(s))
}
