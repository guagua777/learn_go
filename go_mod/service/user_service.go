package service

// Go 用的是就近原则：编译某个 .go 文件时，Go 工具链会从该文件所在目录开始，向上逐级查找 go.mod，找到的第一个就是它所属的模块。

// 具体到你的例子：

// go_web/service/user_service.go 向上找，先遇到 go_web/go.mod（module 名是 go_web）
// 所以 import "go_web/common" 就在 go_web/ 目录下找 common 包
// 不会继续往上找 learn_go/go.mod
// 这两个 go.mod 定义了两个独立的模块：

// learn_go/go.mod → 模块 learn_go
// learn_go/go_web/go.mod → 模块 go_web
// 它们互相独立，各自管理自己的依赖。如果 learn_go 模块想用 go_web 模块的代码，需要在 learn_go/go.mod 里显式 require go_web，或者用 replace 指令指向本地路径。

import (
	"go_web/common"
)

func Calc() int {
	return common.Add(1, 2)
}
