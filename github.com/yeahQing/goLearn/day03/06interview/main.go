package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {
	//map集合
	m := make(map[string]*student)
	//stus是一个切片
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
		{name: "小王八", age: 4500},
	}

	for i, stu := range stus {
		fmt.Println(i, &stu)
		m[stu.name] = &stu
		fmt.Printf("%p\n", &stu)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
}
