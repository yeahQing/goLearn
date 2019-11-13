package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {

	v := 1
	ptr := &v //ptr用来接收v的地址
	fmt.Println(ptr)
	//ptr值存放的是v的地址
	fmt.Printf("type of ptr : %T\n", ptr)
	//ptr也有自己的地址
	fmt.Println(&ptr)
	//根据ptr的值地址取址寻值
	//此处寻找的是ptr中指针指向的地址,即v的地址
	fmt.Println(*ptr)

	x := 10
	//此处传递的只是x的复制
	modify1(x)
	fmt.Println(x)
	//ptr *int = &x
	//此处传递的是内存地址
	modify2(&x)
	fmt.Println(x)

	//声明一个指针类型的变量
	var a *int
	a = new(int) //初始化之后分配内存空间,返回地址
	*a = 10      //取a的地址赋值
	fmt.Println(a)
	fmt.Println(*a)

}
