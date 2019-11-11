package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When`s Saturday?")
	//获取今天是一天中的那一天
	today := time.Now().Weekday()
	fmt.Println(today)

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}