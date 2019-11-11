package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//if前可以有简短的语句声明
	//v if else内部的变量通过x,n计算得出
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrt(x float64) float64 {

	z := float64(1)
	count := 10//声明一个变量10
	for i := 0; i < count; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}

	return z

}

func main(){
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),//每一个方法后都需要跟逗号
	)

	fmt.Println(Sqrt(2))
}