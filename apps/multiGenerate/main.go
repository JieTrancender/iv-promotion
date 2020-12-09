package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateIntA() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

func generateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()

	return ch
}

func generateInt() chan int {
	ch := make(chan int, 20)
	go func() {
		for {
			select {
			case ch <- <-generateIntA():
			case ch <- <-generateIntB():
			}
		}
	}()

	return ch
}

func main() {
	rand.Seed(time.Now().Unix())

	ch := generateInt()

	for i := 0; i < 10000; i++ {
		fmt.Println(<-ch)
	}
}
