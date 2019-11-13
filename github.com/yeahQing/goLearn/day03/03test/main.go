package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	//map集合[string]:int 字符串:整型
	//初始化一个集合,cap = users切片的长度
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	//遍历切片
	for _, value := range users {
		//首先将用户名转化为小写
		name := strings.ToLower(value)
		//分别计算
		distribution[value] += 1 * strings.Count(name, "e")
		distribution[value] += 2 * strings.Count(name, "i")
		distribution[value] += 3 * strings.Count(name, "o")
		distribution[value] += 4 * strings.Count(name, "u")
	}
	//遍历map集合
	for k, v := range distribution {
		fmt.Println("姓名: ", k, "获得硬币数: ", v)
		coins -= v
	}

	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下: ", left)
}
