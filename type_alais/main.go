package main

import "fmt"

// 给 int 起个别名叫 myInt
type myInt = int

func main() {
	var a myInt = 10
	var b int = 20

	// 可以直接运算！因为它们是同一个类型
	fmt.Println(a + b) // 正常运行：30
}
