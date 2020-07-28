package main

import "fmt"

type Ticker struct {
	Name  string
	Price int
}

type Vector struct {
	Data chan Ticker
}

func main() {
	v := &Vector{Data: make(chan Ticker, 20)}

	for i := 0; i < 20; i++ {
		v.Data <- Ticker{
			Name:  "vector",
			Price: i,
		}
	}

	done := make(chan int)

	go func(vector *Vector) {
		for {
			if len(vector.Data) == 0 {
				done <- 1
				fmt.Println("没有票了")
			}
			if v, ok := <-vector.Data; ok {
				fmt.Println(v.Name, v.Price)
			}
		}
	}(v)

	<-done
}
