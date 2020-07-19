package main

import (
	"fmt"
	"reflect"
)

//返回指针
//清空

func NewMap() {
	m := new(map[int]string)
	m1 := make(map[int]string, 2)
	fmt.Println(reflect.TypeOf(m), reflect.TypeOf(m1))
}

func deleteElement() {
	m := make(map[int]string, 3)
	m[1] = "ss"
	m[0] = "ss1"
	m[2] = "ss2"
	fmt.Println(m)
	delete(m, 1)
	fmt.Println(m)
}

func copyElement() {
	mslice := make([]string, 3)
	mslice[0] = "aa"
	mslice[1] = "aa1"
	mslice1 := make([]string, 3)
	copy(mslice1, mslice)
	fmt.Println(mslice, mslice1)
}

func appendElement() {
	mslice := make([]string, 3)
	fmt.Println(mslice, len(mslice), cap(mslice))
	mslice[0] = "aa"
	fmt.Println(mslice, len(mslice), cap(mslice))
	mslice = append(mslice, "aaa")

	fmt.Println(mslice, len(mslice), cap(mslice))
}

func main() {
	//NewMap()
	//appendElement()
	//copyElement()
	deleteElement()
}
