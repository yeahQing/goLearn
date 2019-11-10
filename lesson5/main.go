package main

import "fmt"
import "lesson/lesson5/mylib"

func typeof(v interface{}) string {
	return fmt.Sprintf("%T",v)
}

func main() {
	//简洁赋值语句
	i := 1
	//如果变量类型可以被推断，那么可以不用加类型，直接使用:=
	fmt.Println(i)
	//引入另一个包下的布尔类型
	fmt.Println(mylib.ToBe)
	//引入另一个包下的复数
	fmt.Println(mylib.Z)
	//引入另一个包下的整数最大值
	fmt.Println(mylib.MaxInt)
	//调用变量转换方法
	fmt.Println("After Transform type is:"+typeof(mylib.Transform(1,2)))
	v := 0 // 修改这里！
	fmt.Printf("v is of type %T\n", v)
}
