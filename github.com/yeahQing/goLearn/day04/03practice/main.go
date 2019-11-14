package main

import "fmt"

type student struct {
	id   int64
	name string
}

type studentList map[int64]*student

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

//展示学生列表
func (s studentList) showStudents() {
	for k, v := range s {
		fmt.Println("学号:", k, "姓名:", v.name)
	}
}

//添加学生
func (s studentList) addStudents(id int64, name string) {
	stu, exist := s[id]
	if stu != nil && exist {
		fmt.Println("已经存在了!", stu, exist)
		return
	}
	s[id] = &student{
		id:   id,
		name: name,
	}
}

//修改学生
func (s studentList) modifyStudent(id int64, name string) {
	stu, exist := s[id]
	if stu == nil && !exist {
		fmt.Println("没有该学生!", stu, exist)
		return
	}
	stu.name = name
}

//删除学生
func (s studentList) deleteStudent(id int64) {
	delete(s, id)
	fmt.Println("删除成功!")
}

func showMenu() {
	fmt.Print(`
	欢迎来到学生管理系统！
	1.查看所有学生信息
	2.新增学生信息
	3.修改学生信息
	4.删除学生信息
	5.退出
	请选择您需要的功能：
`)
}

func main() {
	//初始化学生类表
	list := make(studentList, 48)
forEnd:
	for {
		showMenu()
		var option int
		switch fmt.Scanln(&option); option {
		case 1:
			list.showStudents()
		case 2:
			var (
				id   int64
				name string
			)
			fmt.Scanln(&id, &name)
			list.addStudents(id, name)
		case 3:
			var (
				id   int64
				name string
			)
			fmt.Scanln(&id, &name)
			list.modifyStudent(id, name)
		case 4:
			var (
				id int64
			)
			fmt.Scanln(&id)
			list.deleteStudent(id)
		case 5:
			fmt.Println("谢谢使用!")
			break forEnd
		}
	}

}
