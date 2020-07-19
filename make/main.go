package main

import "fmt"

//slice map chan
//返回引用

func makeSlice() {
	s := make([]string, 3)
	s[0] = "d1"
	s[1] = "s2"
	s[2] = "33"

	fmt.Println(s)
}

func makeMap() {
	m := make(map[string]int, 3)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println(m)
}

func makeChan() chan int {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3
	close(c)
	return c
}

func main() {
	makeSlice()
	makeMap()

	c := makeChan()
	for {
		num, ok := <-c
		if ok {
			fmt.Println(num)
		} else {
			break
		}
	}
}
