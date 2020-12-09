package main

import (
	"fmt"
)

const (
	// MaxGoroutineNum max goroutine num
	MaxGoroutineNum = 10
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
	workers := MaxGoroutineNum

	taskChan := make(chan task, 10)

	resultChan := make(chan int, 10)

	done := make(chan struct{}, 10)

	// 初始化task的goroutine，计算100个自然数之和
	go InitTask(taskChan, resultChan, 100)

	// distribute tasks to MaxGoroutineNum workers
	DistributeTask(taskChan, workers, done)

	go CloseResult(done, resultChan, workers)

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
func DistributeTask(taskChan <-chan task, workers int, done chan struct{}) {
	for i := 0; i < workers; i++ {
		go ProcessTask(taskChan, done)
	}
}

// ProcessTask 处理具体工作，并将结果发送到结果通道
func ProcessTask(taskchan <-chan task, done chan struct{}) {
	for tsk := range taskchan {
		tsk.Do()
	}
	done <- struct{}{}
}

// CloseResult 获取每个goroutine处理完任务的通知，并关闭结果通道
func CloseResult(done chan struct{}, resultChan chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}

	close(done)
	close(resultChan)
}

// ProcessResult 读取结果通道，汇总结果
func ProcessResult(resultChan chan int) int {
	sum := 0
	for v := range resultChan {
		sum += v
	}

	return sum
}
