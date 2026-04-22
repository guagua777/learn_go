package main

import (
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {
	// 方法中定义变量，变量为接口类型
	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
}
