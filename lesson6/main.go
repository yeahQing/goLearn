package main

import (
	"fmt"
)

func main() {
	sum := 0
	//初始化语句 第一个;之前的可选
	//后置语句 第二个;之后的可选
	for i:= 0;i <= 100;i++ {
			sum += i
	}
	fmt.Println(sum)

	//while
	for sum < 100 {
		sum += 100
	}

	fmt.Println(sum)
}