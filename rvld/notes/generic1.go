package main

import "fmt"

// 混合约束：值类型集合 + 方法约束
type NumStringer interface {
	~int | ~float64
	String() string
}

// 泛型函数，接收满足混合约束的类型
func Print[T NumStringer](v T) {
	fmt.Println(v.String())
}

// 给 int 绑定方法，满足接口
type MyInt int

func (m MyInt) String() string {
	return fmt.Sprintf("整数：%d", m)
}

// 给 float64 绑定方法
type MyFloat float64

func (m MyFloat) String() string {
	return fmt.Sprintf("小数：%.2f", m)
}

func main2() {
	Print(MyInt(100))
	Print(MyFloat(3.14))
}
