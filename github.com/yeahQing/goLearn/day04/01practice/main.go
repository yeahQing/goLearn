package main

import (
	"encoding/json"
	"fmt"

	m "lesson/github.com/yeahQing/goLearn/day04/01practice/mylib"
)

//声明结构体
type student struct {
	name    string
	age     int
	float32 //匿名字段
}

func (s student) speak() {
	fmt.Printf("%s别说话!\n", s.name)
}

func main() {
	// var stu student
	// stu := &student{}

	// stu := student{
	// 	"小红",
	// 	12,
	// }

	stu := &student{
		"小红",
		12,
		12.4,
	}

	// stu.name = "小红"
	stu.speak()
	fmt.Println(stu)
	n := m.Test{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
	}
	fmt.Printf("n.a %p \n", &n.A)
	fmt.Printf("n.b %p \n", &n.B)
	fmt.Printf("n.c %p \n", &n.C)
	fmt.Printf("n.d %p \n", &n.D)

	s := m.NewStudent("拜山头", 12)
	s.Learn()
	s.SetName("小明1")
	fmt.Println(s.GetName())

	user1 := m.User{
		Name:   "晔清",
		Gender: "男",
		Address: m.Address{
			Province: "山东",
			City:     "威海",
		},
		Email: m.Email{
			Account: "100222",
		},
	}

	user1.Address.CreateTime = "2019-02-12"
	user1.Email.CreateTime = "2019-12-12"
	fmt.Println(user1)

	d1 := &m.Dog{
		Feet: 4,
		Animal: &m.Animal{
			Name: "哈士奇",
		},
	}
	//Dog继承了Animal的move方法和姓名属性
	d1.Wang()
	d1.Move()

	//首先创立一个部门
	department := &m.Department{
		Title:    "鬼畜部",
		Teachers: make([]*m.Teacher, 0, 200),
	}

	for i := 0; i < 10; i++ {
		teacher := &m.Teacher{
			Name:   fmt.Sprintf("teacher%02d", i),
			Gender: "男",
			ID:     i,
		}
		//向部门中添加老师
		department.Teachers = append(department.Teachers, teacher)
	}
	//json序列化
	data, err := json.Marshal(department)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	//输出json序列化之后的结果
	fmt.Printf("json: %s\n", data)
}
