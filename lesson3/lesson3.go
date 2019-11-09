package main

import (
	"fmt"
)

// var 标签用于声明变量列表
var c, java, golang bool

func main() {
	c = true
	var x, y int
	// bool 类型初始值为false
	fmt.Println(c,java,golang)
	//int 类型初始值为0
	fmt.Println(x, y)
}