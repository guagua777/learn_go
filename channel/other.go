// 这个错误是因为在同一个目录 /home/zhoujun/golang/learn_go/channel 中存在两个不同的包声明：

// channel.go 文件声明了 package main
// other.go 文件声明了 package other
// 在 Go 语言中，同一个目录下的所有文件必须属于同一个包。这是 Go 语言的一个基本规则，目的是保持包的一致性和可维护性。

// 解决方案：
// 修改包声明：将 other.go 文件的包声明改为 package main，使其与 channel.go 保持一致
// 分离文件：如果确实需要不同的包，可以将 other.go 文件移动到一个单独的目录中，例如 /home/zhoujun/golang/learn_go/other
// 建议：
// 对于可执行程序，通常使用 package main
// 对于库代码，使用描述性的包名
// 确保同一个目录下的所有文件使用相同的包声明
// 这样修改后，Go 编译器就能正确识别和处理这些文件了。

// package other

package main

func main1() {

}
