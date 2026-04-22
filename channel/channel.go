package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 定义通道
	var ch chan int // 初始值是 nil

	// 2. 启动循环
	go func() {
		for {
			fmt.Println("循环中...")
			select {
			case ch <- 1: // 这里如果 ch 是 nil，这行代码会一直阻塞！
				fmt.Println("发送成功")
			case <-time.After(time.Second * 1):
				fmt.Println("超时...")
			}
		}
	}()

	// 3. 主协程休眠2秒，此时 ch 还是 nil，循环里一直在阻塞
	time.Sleep(5 * time.Second)

	fmt.Println("准备好通道了！")
	// 4. 关键一步：给通道初始化（从 nil 变成 非 nil）
	ch = make(chan int)

	// 现在循环里的 case ch <- 1 不再阻塞，会正常执行发送操作
	select {
	case <-ch:
		fmt.Println("收到数据")
	}
}
