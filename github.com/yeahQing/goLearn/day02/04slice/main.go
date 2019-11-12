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

	a2 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a2:%v type: %T len:%d cap:%d\n", a2, a2, len(a2), cap(a2))
	b2 := a2[0:3] // 切片
	fmt.Printf("b2:%v type: %T len:%d cap:%d\n", b2, b2, len(b2), cap(b2))
	c2 := b2[0:5] // 切片再切片
	fmt.Printf("c2:%v type: %T len:%d cap:%d\n", c2, c2, len(c2), cap(c2))
	//cap()求切片或数组的容量，容量是从数组的起始位置到其原数组长度的表示
	//比如b2是由
}
