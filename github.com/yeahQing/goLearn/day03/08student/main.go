package main

import (
	"fmt"
	"os"
)

type studentManager struct {
	addStudent map[int64]*studentTest
}

type studentTest struct {
	id   int64
	name string
}

func newStudentManager(allStudent map[int64]*studentTest) *studentManager {
	return &studentManager{
		addStudent: allStudent,
	}
}
func newStudentTest(id int64, name string) *studentTest {
	return &studentTest{
		id:   id,
		name: name,
	}
}

func main() {
	allStudentMap := make(map[int64]*studentTest, 48)
	allStudentMap[10086] = newStudentTest(10086, "google")
	var studentM = newStudentManager(allStudentMap)
	for {
		fmt.Print(`
	欢迎来到学生管理系统！
	1.查看所有学生信息
	2.新增学生信息
	3.修改学生信息
	4.删除学生信息
	5.退出
	请选择您需要的功能：
`)

		var tmp int8
		fmt.Scanln(&tmp)
		switch tmp {
		case 1:
			studentM.showStudentTest()
		case 2:
			var (
				id   int64
				name string
			)
			fmt.Print("请输入学生的学号：")
			fmt.Scanln(&id)
			fmt.Print("请输入学生的姓名：")
			fmt.Scanln(&name)
			studentM.addStudentTest(id, name)
		case 3:
			var (
				id   int64
				name string
			)
			fmt.Print("请输入学生的学号：")
			fmt.Scanln(&id)
			fmt.Print("请输入学生想要修改的姓名：")
			fmt.Scanln(&name)
			studentM.changeStudentTest(id, name)
		case 4:
			var (
				id int64
			)
			fmt.Print("请输入删除学生的学号：")
			fmt.Scanln(&id)
			studentM.deleteStudentTest(id)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("无效的输入，请重试")
		}
	}
}

func (s studentManager) showStudentTest() {
	for index, value := range s.addStudent {
		fmt.Println("学号：", index, "  姓名：", value.name)
	}
}

func (s studentManager) addStudentTest(id int64, name string) {
	_, flag := s.addStudent[id]
	if flag {
		fmt.Println("已存在，若想修改，请重新选择3：修改功能")
	} else {
		var student *studentTest
		student = &studentTest{
			id:   id,
			name: name,
		}
		fmt.Println("我要开始添加", student)
		s.addStudent[id] = student
		fmt.Println("添加成功！")
		for index, value := range s.addStudent {
			fmt.Println("学号：", index, "  姓名：", value.name)
		}
	}
}

func (s studentManager) changeStudentTest(id int64, name string) {
	value, flag := s.addStudent[id]
	if !flag {
		fmt.Println("查无此人，请先添加对应学生的信息，请重新选择2：添加功能")
	} else {
		value.name = name
		fmt.Println("修改成功")
		for index, value := range s.addStudent {
			fmt.Println("学号：", index, "  姓名：", value.name)
		}
	}
}

func (s studentManager) deleteStudentTest(id int64) {
	delete(s.addStudent, id)
	fmt.Println("删除成功")
	for index, value := range s.addStudent {
		fmt.Println("学号：", index, "  姓名：", value.name)
	}
}
