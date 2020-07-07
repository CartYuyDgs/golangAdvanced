package main

import (
	"container/list"
	"fmt"
)

type Man struct {
	name string
}

func main() {
	l := list.New()

	l1 := l.PushBack("a")

	l.InsertBefore("b", l1)

	l.InsertAfter("c", l1)
	l.PushBack(1234)
	l.PushBack(Man{"abc"})

	fmt.Println(l, l.Len())

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		man, ok := e.Value.(Man)
		if ok {
			fmt.Println(man)
		}
	}

	for e := l.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case int, int8, int16, int32, int64:
			fmt.Println("int: ", e.Value)
		case Man:
			fmt.Println("Man: ", e.Value)
		case string:
			fmt.Println("string: ", e.Value)
		default:

		}
	}
}
