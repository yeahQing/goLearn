package main

import "fmt"

//定义一个接口
type writer interface {
	Write([]byte) error
}

//Sayer ...
type Sayer interface {
	say()
}

type dog struct{}
type cat struct{}

//いぬ dog实现了Sayer接口
func (d dog) say() {
	fmt.Println("汪汪汪")
}

//ねこ cat实现了Sayer接口
func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	//声明一个接口
	var x Sayer
	a := cat{}
	b := dog{}
	x = a
	x.say()
	x = b
	x.say()
}
