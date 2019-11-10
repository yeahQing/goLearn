package mylib

import (
	"math/cmplx"
)

/* 
bool string 
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte rune 
float32 float64 
complex64 complex128
*/

// 变量声明也可以分组成语法糖
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	Z complex128 = cmplx.Sqrt(-5 + 12i)
	//小写开头的变量和方法不能被其他被使用，是protected修饰的
	z complex128 = cmplx.Sqrt(-2 + 10i)
)

func Transform(x, y int) float64 {
	//将两个整型的乘显式转换为float64
	f := float64(x*x + y*y)
	return f
}
