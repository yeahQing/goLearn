package main

import "fmt"

func testSwitch() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func testSwitch1() {
	age := 25
	switch {
	case age < 18:
		fmt.Println("未成年")
	case age >= 18 && age < 30:
		fmt.Println("青年人")
	case age >= 30:
		fmt.Println("中年人")
	default:
		fmt.Println("不小了!")
	}
}

func switchDemo() {
	switch s := "a"; s {
	case "a":
		fmt.Println("a")
		fallthrough //相当于不写break
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("...")

	}
}

func main() {
	testSwitch()
	testSwitch1()
	switchDemo()
}
