package main

import "fmt"

// 父结构体
type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHi() {
	fmt.Println("Hi, I'm", p.Name)
}

// 组合 Person = 继承效果
type Student struct {
	Person // 匿名嵌套，类似继承
	School string
}

// // 如果是这种情况，则无法直接调用父结构体的方法
// type Student struct {
// 	Person Person
// 	School string
// }

func main() {
	s := Student{
		Person: Person{Name: "张三", Age: 18},
		School: "清华",
	}

	s.SayHi()           // 直接调用！就像继承来的
	fmt.Println(s.Name) // 直接访问
}
