package subpackage

import "fmt"

func Add(a, b int) int {
	return a + b
}

func init() {
	fmt.Println("subpackage 先执行init")
}
