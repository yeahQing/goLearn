package main

import "fmt"

var (
	s1   = "sad"
	var1 = 1
	var2 int
	var3 int64
	u1   uint16
)

const (
	statusOk = 200
	notFound = 404
)

const (
	n1 = 100
	n2
	n3
)

const (
	a1 = iota // iota从0开始
	a2
	a3
	a4
	//iota出现则默认值为0
)

func main() {
	fmt.Println(a2)
}
