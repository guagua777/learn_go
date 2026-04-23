package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享变量：多协程会同时读写
var count int
var mu sync.Mutex // 互斥锁：保护共享变量

func worker(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束，计数-1

	// 1. 操作共享变量，加锁防竞争
	mu.Lock()
	count++
	mu.Unlock()

	// 2. 往 channel 发送消息（协程间通信）
	msg := fmt.Sprintf("goroutine-%d 完成任务", id)
	ch <- msg
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3) // 有缓冲通道

	// 启动 3 个协程
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	// 开启单独协程：等待全部任务结束后关闭通道
	go func() {
		wg.Wait()
		close(ch) // 全部干完，关闭通道
	}()

	// 3. select 多路监听：接收通道数据 + 超时控制
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				// 通道关闭，退出循环
				fmt.Println("\n通道已关闭，所有任务结束")
				fmt.Println("最终 count =", count)
				return
			}
			fmt.Println("收到：", msg)

		case <-time.After(2 * time.Second):
			fmt.Println("等待超时，强制退出")
			return
		}
	}
}
