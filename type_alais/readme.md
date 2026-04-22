二、和「自定义类型」的区别（最容易混淆）
很多人分不清这两个：
1. 类型别名（= 同一个类型）
go
运行
type myInt = int
2. 自定义类型（全新类型）
go
运行
type myInt int
区别演示
go
运行
// 自定义类型（新类型）
type A int

// 类型别名（同一个类型）
type B = int

func main() {
    var a A = 10
    var b B = 10
    var c int = 10

    // a + c ❌ 报错！A 和 int 是不同类型
    // b + c ✅ 正常！B 就是 int
}



type A = B → 别名，A 和 B 完全一样
type A B → 新类型，A 和 B 不能混用
别名只是外号，不创建新类型