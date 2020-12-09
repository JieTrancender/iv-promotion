package main

import (
	"runtime"
)

func sum(c3 chan struct{}) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	println(sum)
	c3 <- struct{}{}
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

	c3 := make(chan struct{})
	go sum(c3)

	println("NumGoroutine", runtime.NumGoroutine())
	<-c
	<-c3
	println("NumGoroutine", runtime.NumGoroutine())

	// 通道c1还可以继续读取
	for v := range c1 {
		println(v)
	}
}
