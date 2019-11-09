package main

//引入包不论自定义包还是他人的包都应该使用绝对路径
import (
	"lesson/lesson4/mylib"
	"fmt"
)

func main() {
	//导入包名后，利用包名.方法名来调用
	//方法名应该首字母大写才可以被其他包下的方法调用
	fmt.Println(mylib.Add(1,2))
	//没有返回值的方法不可以被fmt.Println格式化输出
	mylib.Say()
}	