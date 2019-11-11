package main

import (
	"fmt"
	"math"
)

func foo(a1, a2 int) (x, y int) {
	x = a1
	y = a2
	return
}

func main() {
	fmt.Println("Hello World")
	math.Sqrt(2)
	//类型可以推断就不需要加类型了
	//var s1 string = "sss"
	var s1 = "sss"
	fmt.Println(s1)
	//简短变量声明只能在函数中使用
	s2 := "sads"
	fmt.Println(s2)
	x, _ := foo(1, 2)
	fmt.Println(x)
	_, y := foo(3, 2)
	fmt.Println(y)
}
