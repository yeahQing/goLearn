package mylib

import "fmt"

//Test struct
type Test struct {
	A int8
	B int8
	C int8
	D int8
}

//Address 地址
type Address struct {
	Province   string
	City       string
	CreateTime string
}

//Email 邮箱
type Email struct {
	Account    string
	CreateTime string
}

//User 用户
type User struct {
	Name    string
	Gender  string
	Address Address
	Email   // 匿名结构体
}

//继承

//Animal 动物结构体
type Animal struct {
	Name string
}

//Move 移动
func (a *Animal) Move() {
	fmt.Printf("%s会动! \n", a.Name)
}

//Dog 狗的结构体
type Dog struct {
	Feet int8
	*Animal
}

//Wang 叫
func (d *Dog) Wang() {
	fmt.Printf("%s会汪汪汪", d.Name)
}
