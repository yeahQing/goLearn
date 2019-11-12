package main

import "fmt"

func gotoDemo1() {
	//小驼峰命名
	var breakFlag bool
	//不使用goto 语句
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	//goto语句跳到这里
breakTag:
	fmt.Println("结束for循环")
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

func mpt() {
loop:
	for i := 1; i < 10; i++ {
		// for j := 1; j <= i; j++ {
		// 	fmt.Printf("%v*%v=%v\t", j, i, i*j)
		// }
		for j := 1; j < 10; j++ {
			if i == j {
				fmt.Printf("%d*%d=%d\n", j, i, i*j)
				continue loop
			}
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
	}
}

func main() {
	// gotoDemo1()
	// gotoDemo2()
	// continueDemo()
	mpt()
}
