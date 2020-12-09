package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func generateIntA(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}

		close(ch)
	}()

	return ch
}

func generateIntB(done chan struct{}) chan int {
	ch := make(chan int, 5)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}

		close(ch)
	}()

	return ch
}

func generateInt(done chan struct{}) chan int {
	ch := make(chan int)
	seed := make(chan struct{})
	go func() {
		randA := generateIntA(seed)
		randB := generateIntB(seed)
	Lable:
		for {
			select {
			case ch <- <-randA:
			case ch <- <-randB:
			case <-done:
				seed <- struct{}{}
				seed <- struct{}{}
				break Lable
			}
		}

		close(ch)
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().Unix())

	done := make(chan struct{})
	ch := generateInt(done)

	time.Sleep(1 * time.Second)
	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	// 通知生成器停止生产
	done <- struct{}{}

	time.Sleep(2 * time.Second)
	fmt.Println("stop generate", runtime.NumGoroutine())
}
