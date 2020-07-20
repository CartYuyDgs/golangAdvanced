package main

import (
	"fmt"
	"sync"
)

type DataSoure struct {
	connect int
}

//触发一次GC，全部释放
func main() {
	p := new(sync.Pool)

	for i := 0; i < 50; i++ {
		p.Put(DataSoure{connect: i})
		fmt.Println(p.Get())
	}
	fmt.Println(p.Get())
}
