package main

import (
	"fmt"
)

type Man struct {
	Name string
	flag int
}

func main() {
	m := map[int]int{
		1: 1,
		2: 2,
	}

	m2 := make(map[int]int)
	m2[1] = 1
	m2[2] = 2

	fmt.Println(m, m2)

	mans := []Man{
		{Name: "a", flag: 1},
		{Name: "b", flag: 2},
		{Name: "c", flag: 3},
	}

	mm := make(map[int]Man)
	for _, m := range mans {
		mm[m.flag] = m
		//mm[m.flag]=&m
	}

	fmt.Println(mm)

	for _, m := range mm {
		fmt.Println(m.Name)
	}
}
