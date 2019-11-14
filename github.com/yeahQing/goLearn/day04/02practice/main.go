package main

import (
	"fmt"
	"io"
)

var (
	data []byte
	err  error
)

//student 学生实体结构体
type student struct {
	id    int
	name  string
	age   int
	score float32
}

//showStudents
func showStudents() {

}

//addStudents
func addStudents(stu *student) {

}

//ReadFrom ...
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	//初始化切片
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func main() {
	// data, err = ReadFrom(os.Stdin, 11)
	// data, err = ReadFrom(strings.NewReader("from string"), 12)
	// fmt.Println(data, err)
	// var ch string
	// fmt.Scanln(&ch)
	// fmt.Println(ch)
	//初始化student
	stu := &student{
		id:    1,
		name:  "张明",
		age:   12,
		score: 60.5,
	}

	fmt.Printf("%#v\n", stu)
	//fmt.Printf("%+v",stu) 输出结构体时同时输出属性名
	//%#v会输出具体输出的变量的类型
	//空接口类型
	// var i interface{} = 23
	// fmt.Printf("%v\n", i)
	// var j uint = 2
	// fmt.Printf("%d", j)
	var (
		s string
		i int
	)
	fmt.Printf("%T %T", s, i)
	fmt.Sscanf(" 1234567 ", "%5s%d", &s, &i)
	fmt.Println(s, i)
}
