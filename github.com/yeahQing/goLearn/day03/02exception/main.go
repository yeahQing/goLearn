package main

import "fmt"

func a() {
	fmt.Println("A")
}

func b() {
	//抛出异常
	//defer 匿名函数
	x := 1
	y := 2
	defer func(x int) {
		err := recover()
		if err != nil {
			//修复异常
			fmt.Println(x) // x = 1因为程序执行到这里的时候x的值已经赋予了
			fmt.Println(y) // y = 20
			fmt.Println("recover in B")
		}
	}(x)
	x = 10
	y = 20
	//异常抛出
	panic("there has been a error...")
	// x = 10 unable reached因为panic引发了程序崩溃
}

func c() {
	fmt.Println("C")
}

func main() {
	a()
	b()
	c()
}
