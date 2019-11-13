package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	//执行了这个函数，只是结果是最后才输出
	// x = 1, y = x + y = 1 + 2 = 3
	// AA = 4
	defer calc("AA", x, calc("A", x, y))
	x = 10
	//22
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
