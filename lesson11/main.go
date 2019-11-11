package main

import "fmt"

func main() {
	const count = 1 << 2
	for i := 1; i < 1<<2; i <<= 1 {
		defer fmt.Print(i)
	}
	fmt.Print(count + 1)
}
