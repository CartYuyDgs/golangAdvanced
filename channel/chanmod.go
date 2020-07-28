package main

import (
	"fmt"
	"sync"
)

type Ticker struct {
	Name  string
	Price int
}

type Vector struct {
	Data chan Ticker
}

func demo1() {
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

func demo2() {
	v := &Vector{Data: make(chan Ticker, 1)}

	go func(vector *Vector) {
		for {
			if v, ok := <-vector.Data; ok {
				fmt.Println(v.Name, v.Price)
			}
		}
	}(v)

	for {
		v.Data <- Ticker{
			Name:  "vector",
			Price: 20,
		}
	}
}

func main() {
	v := &Vector{Data: make(chan Ticker, 100)}
	//生产
	for j := 0; j < 100; j++ {
		v.Data <- Ticker{
			Name:  "vector",
			Price: j,
		}
	}

	//消费
	wg := new(sync.WaitGroup)
	for k := 0; k < 100; k++ {
		wg.Add(1)
		go func(vector *Vector) {
			for {
				if len(vector.Data) == 0 {
					wg.Done()
					fmt.Println("，没有票了。。。。")
					return
				}
				if v, ok := <-vector.Data; ok {
					fmt.Println(v.Price, v.Name)
				}
			}

		}(v)

	}
	wg.Wait()
}
