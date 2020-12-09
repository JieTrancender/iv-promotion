package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func generateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}

		// 收到通知后关闭通道ch
		close(ch)
	}()

	return ch
}

func main() {
	// 设置随机数种子
	rand.Seed(time.Now().Unix())

	done := make(chan struct{})
	ch := generateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 发送通知，告诉生产者停止生产
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// 此时生产者已经退出了
	println("NumGoroutine", runtime.NumGoroutine())
}
