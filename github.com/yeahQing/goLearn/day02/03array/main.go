package main

import "fmt"

func main() {
	//定义一个长度为3的int类型的数组
	//初始化为0
	// var a [3]int
	// var b [3]int
	//使用指定的值初始化
	var testArray = [3]int{1, 2}
	var a = testArray[0]
	fmt.Println(a)
	//让编译器根据初始值的个数自行推断数组的长度
	var cityArray = [...]string{"青岛", "济南", "武汉"}
	fmt.Printf("type of cityArray:%T\n", cityArray)
	fmt.Println(cityArray)

	//通过索引值的方式初始化数组
	indexArray := [...]int{1: 1, 3: 5}
	fmt.Println(indexArray)
	fmt.Printf("type of indexArrat: %T \n", indexArray)

	//遍历数组
	for i := 0; i < len(indexArray); i++ {
		fmt.Println(indexArray[i])
	}

	//for range 遍历
	for index, value := range indexArray {
		fmt.Println(index, value)
	}

	//二维数组
	dArray := [...][2]string{
		{"北京", "上海"},
		{"重庆", "成都"},
		{"广州", "深圳"},
	}

	fmt.Println(dArray)

	var arrayTest = [...]int{1, 3, 5, 7, 8}
	sum := 0
	for index, value := range arrayTest {
		sum += value
		for i := index + 1; i < len(arrayTest); i++ {
			if arrayTest[i]+value == 8 {
				fmt.Printf("(%v,%v)\n", index, i)
			}
		}
	}
	fmt.Println(sum)
}
