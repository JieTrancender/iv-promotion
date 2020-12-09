package main

import (
	"runtime"
	"time"
)

func sum() {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	time.Sleep(1 * time.Second)
}

func main() {
	println("GOMAXPROCS", runtime.GOMAXPROCS(0))

	c := make(chan struct{})
	c1 := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)

		c <- struct{}{}
	}(c, c1)

	go func(i chan struct{}) {
		sum := 0
		for i := 0; i < 10000; i++ {
			sum += i
		}
		println(sum)

		c <- struct{}{}
	}(c)

	println("NumGoroutine", runtime.NumGoroutine())
	<-c
	println("NumGoroutine", runtime.NumGoroutine())

	// 通道c1还可以继续读取
	for v := range c1 {
		println(v)
	}
}
