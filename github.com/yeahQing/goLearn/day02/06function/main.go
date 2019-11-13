package main

import (
	"errors"
	"fmt"
	"strings"
)

//type用于定义一个函数类型
//定义一个caclulation是一个函数类型，接收两个int类型的参数，并且返回一个int类型
type calculation func(int, int) int

func intSum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//sliceTest arg1: []int 切片
func sliceTest(y []int) {
	for _, v := range y {
		fmt.Println(v)
	}
}

//calc 返回类型和返回参数都在返回参数列表中显示
func calc(x, y int) (sum int, sub int) { //此处可以写为sum, sub int
	sum = x + y
	sub = x - y
	return
}

//calc1 返回类型在参数列表中声明，返回参数在return 语句中显示
func calc1(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//add
func add(x, y int) int {
	return x + y
}

//sub
func sub(x, y int) int {
	return x - y
}

//函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//闭包
//func 函数名为addr() 返回值是一个函数
//func addr()函数名为一个函数
func addr() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

//makeSuffixFunc
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}

		return name
	}
}

//闭包实例3
//函数名为函数 返回值为两个匿名函数
func calc2(base int) (func(int) int, func(int) int) {
	//第一个匿名函数
	//用add变量接收
	add := func(i int) int {
		base += i
		return base
	}

	//第二个匿名函数
	//用sub变量接收
	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

func main() {
	//声明一个切片
	x := []int{1, 2, 3, 4, 5}
	sum := intSum(10, 12, 13, 14)
	fmt.Println(sum)
	// intSum(x)
	sliceTest(x)
	sum, sub := calc(20, 10)
	fmt.Println(sum, sub)

	fmt.Println(calc1(10, 20))
	var c calculation
	c = add
	fmt.Println(c(10, 20))

	fmt.Println(do("+"))

	//匿名函数
	//用变量接收函数
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	//调用函数
	add(10, 2)

	//自执行匿名函数
	func(x, y int) {
		fmt.Println(x - y)
	}(3, 2)

	//调用闭包
	var f = addr()
	fmt.Println(f(12))
	//12
	fmt.Println(f(10))
	//22

	jpgFuc := makeSuffixFunc(".jpg")
	fmt.Println(jpgFuc("test"))

	testFuc := makeSuffixFunc(".txt")
	fmt.Println(testFuc("test"))

	f1, f2 := calc2(10)       //base会随着内部函数的改变而改变
	fmt.Println(f1(1), f2(2)) //10+1=11 11-2=9

}
