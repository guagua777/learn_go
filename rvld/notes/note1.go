package main

import "fmt"

func main() {
	// 类型转换：字符串转byte切片
	elfMagic := []byte("\177ELF")
	fmt.Println("ELF魔数（byte切片）：", elfMagic) // 输出：[127 69 76 70]

	// 切片字面量：手动指定byte值
	elfMagicLiteral := []byte{127, 69, 76, 70}
	fmt.Println("ELF魔数（字面量）：", elfMagicLiteral) // 输出：[127 69 76 70]

	// 两者等价
	// fmt.Sprintf("%v", elfMagic)
	// fmt.Sprintf("%v", elfMagicLiteral)
	fmt.Println("是否相等：", fmt.Sprintf("%v", elfMagic) == fmt.Sprintf("%v", elfMagicLiteral)) // 输出：true
}
