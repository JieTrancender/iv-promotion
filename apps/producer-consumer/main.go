package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producer(value int, out chan int) {
	for i := 1; ; i++ {
		out <- i * value
	}
}

func consumer(in <-chan int) {
	for v := range in {
		println(v)
	}
}

func main() {
	ch := make(chan int, 3)

	go producer(5, ch)
	go producer(3, ch)
	go consumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", sig)
}
