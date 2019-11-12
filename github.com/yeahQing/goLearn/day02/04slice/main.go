package main

import "fmt"

var (
	d = []bool{false, true}
)

func main() {
	//声明切片
	var a []string
	var b = []int{}
	var c = []bool{true, false}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)

	//基于数组创造切片
	a1 := [5]int{55, 56, 57, 58, 59}
	b1 := a1[:4]
	fmt.Println(b1)
	fmt.Printf("type of b1: %T\n", b1)
}
