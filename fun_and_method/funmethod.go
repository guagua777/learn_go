package main

import "fmt"

// 结构体
type User struct {
	Name string
}

// 方法（绑定 User）
func (u User) Say() {
	fmt.Println("Hello", u.Name)
}

// 方法（指针，修改结构体）
func (u *User) ChangeName(newName string) {
	u.Name = newName
}

// 普通函数
func Hi(s string) {
	fmt.Println("Hi", s)
}

func main() {
	u := User{Name: "myname"}

	u.Say()            // 调用方法
	u.ChangeName("张三") // 调用方法
	Hi("test")         // 调用函数
}
