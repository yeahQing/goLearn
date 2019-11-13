package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		fmt.Printf("before:%v\n", x)
		x++
		fmt.Printf("after:%v\n", x)
	}()
	fmt.Println("调用")
	return x
}

func f2() (x int) {
	defer func() {
		fmt.Printf("before:%v\n", x)
		x++
		fmt.Printf("after:%v\n", x)
	}()
	fmt.Println("调用")
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		fmt.Printf("before:%v\n", x)
		x++
		fmt.Printf("after:%v\n", x)
	}()
	fmt.Println("调用")
	return x
}
func f4() (x int) {
	defer func(x int) {
		fmt.Printf("before:%v\n", x)
		x++
		fmt.Printf("after:%v\n", x)
	}(x)
	fmt.Println("调用")
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
