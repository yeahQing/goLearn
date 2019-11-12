package mylib

import (
	"fmt"
	"sort"
)

//TestSlice ...
func TestSlice() {
	slice := make([]string, 5, 10)
	// slice := []string{} //声明一个空的字符串切片
	//一开始创建的切片不为空
	fmt.Println(slice)
	for i := 0; i < 10; i++ {
		slice = append(slice, fmt.Sprintf("%v", i))
		fmt.Printf("len(slice)=%d cap(slice)=%d\n", len(slice), cap(slice))
	}
	fmt.Println(slice)
}

//Sort test
func Sort() {
	var a = [...]int{3, 7, 8, 9, 1}
	// fmt.Println(a[:]) //将数组切片
	//sort包内部实现了内部数据类型的排序
	//升序排列
	sort.Ints(a[:])
	fmt.Println(a)
	//降序排列
	sort.Sort(sort.Reverse(sort.IntSlice(a[:])))
	fmt.Println(a)
}
