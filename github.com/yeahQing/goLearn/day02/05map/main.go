package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

/* map[KeyType](ValueType) */
func main() {
	//map为引用类型，需要初始化
	scoreMap := make(map[string]int, 8)
	//给map添加键值
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100

	fmt.Println(scoreMap)
	//根据键取值
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of scoreMap: %T \n", scoreMap)

	userInfo := map[string]string{
		"username": "张三丰",
		"password": "123456",
	}

	fmt.Println(userInfo)

	//判断某个map中是否有某个键
	v, ok := userInfo["username"]

	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("没有姓名!")
	}

	rand.Seed(time.Now().UnixNano())
	//初始化Map容量为200 cap = 200
	var scoreMap1 = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap1[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap1 {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap1[key])
	}

	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	//初始化一个map集合,键为string类型,值为[]string切片类型
	var mapSlice1 = make(map[string][]string, 3) // cap = 3
	fmt.Println(mapSlice1)                       //输出mapSlice1
	fmt.Println("After init")

	key := "中国"

	v1, ok := mapSlice1[key]
	//如果不存在键为"中国"的字段
	if !ok {
		v1 = make([]string, 0, 2) //创建一个切片，容量为2
	}
	//向切片中填充值
	v1 = append(v1, "北京", "上海")
	mapSlice1[key] = v1 //给键为中国的字段赋值
	fmt.Println(mapSlice1)

	//要统计的句子
	sentence := "how do you do"
	words := strings.Split(sentence, " ")
	fmt.Println(words) //输出切分后的字符串数组
	//初始化一个map结果集,cap=3,因为句子拆分后只有3种键
	result := make(map[string]int, 3)
	//遍历字符串数组,让map中对应的键++
	for _, v := range words {
		result[v]++
	}
	//输出结果集
	fmt.Println(result)
}
