package mylib

import "fmt"

//Student 实体类
type Student struct {
	name string
	age  int
}

//NewStudent 构造函数
func NewStudent(name string, age int) *Student {
	return &Student{
		name: name,
		age:  age,
	}
}

//Learn 学生学习方法
func (s *Student) Learn() {
	fmt.Printf("%s的任务是学习\n", s.name)
}

//SetName 学生姓名setter
func (s *Student) SetName(name string) {
	s.name = name
}

//GetName 学生姓名getter
/* 虽然这里可以不使用指针类型的接收者Student，但是为了和其他方法一致，也要使用*Student */
func (s *Student) GetName() string {
	return s.name
}
