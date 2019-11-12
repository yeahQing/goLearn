package main

import "fmt"
import "lesson/github.com/yeahQing/goLearn/day02/04slice/mylib"

var (
	d = []bool{false, true}
)

func appendTest() {
	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}
}

func main() {
	//声明切片
	var a []string
	var b = []int{}
	var c = []bool{true, false}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(a == nil)

	//基于数组创造切片
	a1 := [5]int{55, 56, 57, 58, 59}
	b1 := a1[:4]
	fmt.Println(b1)
	fmt.Printf("type of b1: %T\n", b1)

	a2 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a2:%v type: %T len:%d cap:%d\n", a2, a2, len(a2), cap(a2))
	b2 := a2[2:6] // 切片
	fmt.Printf("b2:%v type: %T len:%d cap:%d\n", b2, b2, len(b2), cap(b2))
	c2 := b2[2:4] // 切片再切片
	fmt.Printf("c2:%v type: %T len:%d cap:%d\n", c2, c2, len(c2), cap(c2))
	//cap()求切片或数组的容量，容量是从数组的起始位置到其原数组长度的表示
	//比如b2是a2的切片，那么b2如果从下标为1开始切，则容量为5,如果从下标为0切，则容量为6
	//c2如果再切割，则其起始位置应该为b2的起始位置，其容量应该是从切b2切片的位置开始到b2的长度

	//go语言中内置的make函数:make([]T, size, cap)
	a3 := make([]int, 2, 10)
	fmt.Println(a3)
	fmt.Println(len(a3))
	fmt.Println(cap(a3))

	s := []int{1, 3, 5}
	//判断切片为空: len(s) == 0 ? true : false
	//for range遍历切片
	for index, value := range s {
		fmt.Println(index, value)
	}
	// append的用法
	appendTest()

	//追加多个切片
	var citySlice []string
	citySlice = append(citySlice, "北京")
	citySlice = append(citySlice, "上海", "广州", "深圳")
	slice1 := []string{"成都", "重庆"}
	citySlice = append(citySlice, slice1...)
	fmt.Println(citySlice)

	//copy() 方法 copy(destSlice, srcSlice [] T) arg1: 目标切片 arg2: 源切片
	srcSlice := []int{1, 2, 3, 4, 5}
	destSlice := make([]int, 5, 5)
	copy(destSlice, srcSlice) //复制切片，复制的切片会开辟一块新的内存空间
	fmt.Println(srcSlice)
	destSlice[0] = 1000
	fmt.Println(destSlice)

	//删除切片中的某一个元素删除下标为destSlice[2] = 3的元素
	destSlice = append(destSlice[:2], destSlice[3:]...)
	fmt.Println(destSlice)
	mylib.TestSlice()
	mylib.Sort()
}
