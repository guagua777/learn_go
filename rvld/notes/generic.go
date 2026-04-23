package main

import "fmt"

// 约束：允许 int/float64
type Number interface {
	int | float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main1() {
	fmt.Println(Max(10, 20))     // T=int
	fmt.Println(Max(3.14, 2.71)) // T=float64
}
