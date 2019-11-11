package main

import "fmt"

func ifDemo() {
	score := 0
	if score := 65; score >= 90 {
		//此处的score是if语句内部声明的
		//其只在if、else if、else可见
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	//此处的score是函数内部声明的
	fmt.Println(score)
}

func main() {
	ifDemo() //C 0
}
