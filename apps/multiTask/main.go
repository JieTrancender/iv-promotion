package main

import (
	"fmt"
	"sync"
)

type task struct {
	begin  int
	end    int
	result chan<- int
}

func (t *task) Do() {
	sum := 0
	for i := t.begin; i <= t.end; i++ {
		sum += i
	}

	t.result <- sum
}

func main() {
	taskChan := make(chan task, 10)

	resultChan := make(chan int, 10)

	wait := &sync.WaitGroup{}

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskChan, resultChan, 100)

	// 每个task启动一个goroutine进行处理
	go DistributeTask(taskChan, wait, resultChan)

	// 通过结果通道获取结果并汇总
	sum := ProcessResult(resultChan)

	fmt.Println("sum = ", sum)
}

// InitTask 构建task并写入task通道
// chan<- 表示只能写入
// <-chan 表示只能读取
func InitTask(taskChan chan<- task, r chan int, p int) {
	qu := p / 10
	mod := p % 10
	high := qu * 10
	for j := 0; j < qu; j++ {
		b := 10*j + 1
		e := 10 * (j + 1)
		tsk := task{
			begin:  b,
			end:    e,
			result: r,
		}
		taskChan <- tsk
	}

	if mod != 0 {
		tsk := task{
			begin:  high + 1,
			end:    p,
			result: r,
		}
		taskChan <- tsk
	}

	close(taskChan)
}

// DistributeTask 读取每个task chan，每个task启动一个worker goroutine进行处理
// 并等待每个task运行完，关闭结果通道
func DistributeTask(taskChan <-chan task, wait *sync.WaitGroup, resultChan chan int) {
	for v := range taskChan {
		wait.Add(1)
		go ProcessTask(v, wait)
	}
	wait.Wait()
	close(resultChan)
}

// ProcessTask 处理具体工作，并将结果发送到结果通道
func ProcessTask(t task, wait *sync.WaitGroup) {
	t.Do()
	wait.Done()
}

// ProcessResult 读取结果通道，汇总结果
func ProcessResult(resultChan chan int) int {
	sum := 0
	for v := range resultChan {
		sum += v
	}

	return sum
}
