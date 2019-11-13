package main

import "fmt"

//NewInt 自定义类型
type NewInt int

//注意: NewInt是一个新的类型,其具有int的特性

//MyInt 类型别名
// type TypeAlias = Type
// type byte = uint8
// type rune = int32
type MyInt = int

type person struct {
	name, city string //同种类型可以声明在同一行
	// city string
	age int8
}

//自定义构造函数
func newPerson(name, city string, age int8) *person {
	return &person{ //new(person)
		name: name,
		city: city,
		age:  age,
	}
}

//此为方法而非函数
//函数不属于任何类型
//方法属于特定的类型 这里只有person的实例可以调用Dream()方法
//接收者 接收者类型 接收者形参命名应该为类型首字母小写
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言! \n", p.name)
}

func main() {
	var a NewInt
	var b MyInt
	//自定义类型和类型别名的区别: 一个是一个新的类型，另一个只是给某一个类型起了一个别称
	fmt.Printf("type of a : %T \n", a)
	fmt.Printf("type of b : %T \n", b)

	var p1 person
	p1.name = "张三"
	p1.city = "上海"
	p1.age = 22
	fmt.Println(p1)
	//格式化输出
	fmt.Printf("p1=%#v\n", p1)

	//匿名结构体
	var user struct {
		Name string
		age  int
	}

	user.Name = "小明"
	user.age = 20
	fmt.Printf("user = %#v \n", user)

	//结构体指针
	var p2 = new(person)
	p2.name = "小红"
	p2.city = "深圳"
	p2.age = 12
	fmt.Printf("p2=%#v \n", p2)
	p2.Dream()
}
