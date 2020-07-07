package main

import (
	"container/ring"
	"fmt"
)

//环状结构
func main() {
	r := ring.New(3)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		fmt.Println(r.Value)
		r = r.Next()
	}
}
