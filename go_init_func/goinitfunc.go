// package goinitfunc

package main

import (
	"fmt"
	subpackage "init_func/sub_package"
)

// 自动执行init函数
func init() {
	fmt.Println("先执行init")
}

func main() {
	fmt.Println("再执行main")
	var result = subpackage.Add(1, 2)
	fmt.Println(result)
}
