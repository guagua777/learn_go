package main

import "fmt"

// 1. 定义接口
type Animal interface {
	Speak() string
}

// 2. 定义 Dog，实现 Speak()
type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪"
}

// 3. 定义 Cat，实现 Speak()
type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵"
}

// 4. 接口作为参数（多态）
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main1() {
	d := Dog{}
	c := Cat{}

	MakeSound(d) // 汪汪
	MakeSound(c) // 喵喵
}
