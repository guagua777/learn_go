package main

import (
	"fmt"
	"strings"
)

// 在这个代码中，var t interface {...} 和 type t interface {...} 是两种完全不同的语法结构，它们的作用和用法也不同：

// 1. 当前代码的含义

// go
// var t interface {
//     talk() string
// }
// 这是变量声明，表示：

// 声明了一个名为 t 的变量
// 变量类型是一个匿名接口，该接口要求实现 talk() string 方法
// 后续可以将任何实现了该接口的类型值赋给这个变量（如 martian{} 和 laser(3)）
// 2. 尝试修改后的含义

// go
// type t interface {
//     talk() string
// }
// 这是类型定义，表示：

// 定义了一个名为 t 的新接口类型
// 该接口要求实现 talk() string 方法
// 但这只是定义了一个类型，并没有创建任何变量
// 3. 为什么不能直接替换
// 如果改为类型定义，代码会出现以下问题：

// 第38行 t = martian{} 会报错，因为 t 现在是类型名，不是变量名
// 第41行 t = laser(3) 同样会报错
// 第39行和42行的 t.talk() 也会报错，因为 t 是类型而不是变量
// 4. 正确的修改方式
// 如果要使用类型定义，需要：

// 先定义接口类型
// 再声明该类型的变量
// 例如：

// go
// // 定义接口类型
// type Talker interface {
//     talk() string
// }

// // 声明变量
// var t Talker
// 这样修改后，代码就能正常工作了。

// 总结
// var t interface {...} 是声明变量，而 type t interface {...} 是定义类型，它们的语法和用途完全不同，不能直接替换。

// var t interface{}

// 定义变量，变量为接口类型
var t interface {
	talk() string
}

// type t interface {
// 	talk() string
// }

// 定义struct
type martian struct{}

// 接口实现
// 定义struct的方法
// 同时实现interface的方法
func (m martian) talk() string {
	return "nack nack"
}

// 定义自定义类型（全新类型）
// 这不是定义别名 ******
type laser int

// 给laser类型，实现接口
func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
	// 赋值给t变量
	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}
