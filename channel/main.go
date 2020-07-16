package main

import "fmt"

func f1() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	e := <-ch
	fmt.Println(e)
}

func f2() {
	ch := make(chan int, 2)
	go func() {
		for i := 10; i > 0; i-- {
			fmt.Println("send num .", i)
			ch <- i
		}
		fmt.Println("send end")
		close(ch)
	}()

	for {
		elem, ok := <-ch
		if !ok {
			fmt.Println("receive over")
			break
		}

		fmt.Println("receive num :", elem)
	}

	fmt.Println("End")
}

func main() {
	f2()
}
