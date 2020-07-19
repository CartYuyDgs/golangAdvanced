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

func main() {
	NewMap()
}
