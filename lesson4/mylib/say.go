package mylib

import (
	"fmt"
)

//声明变量列表
var x, y int = 4, 5

//声明方法并定义
func Say() {
	//调用同包下的方法
	fmt.Println(Add(x,y))
}