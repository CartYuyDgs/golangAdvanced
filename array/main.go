package main

import "fmt"

func main() {
	a1 := []string{"a", "b", "c"}
	//a2 := []string(nil)
	a2 := make([]string, 20)

	a3 := append(a2, "c", "d", "e", "f")
	fmt.Println(a3, len(a3))
	//地址
	fmt.Printf("%p\n", &a1)

	//a4是a[1:2]的引用
	a4 := a1[1:2]
	fmt.Println(a4)
	a1[1] = "ccc"
	fmt.Println(a4)

	a1 = []string(nil)
	fmt.Println(a1, a4)
}
