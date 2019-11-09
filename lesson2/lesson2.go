package main

import (
	"fmt"
	"math"
)

/*

	计算圆的面积

*/

func calculate(r float64) float64 {
	return r * r * math.Pi
}

func main() {
	fmt.Println(calculate(2.0))
}