package main

import (
	"context"
	"fmt"
	"time"
)

type otherContext struct {
	context.Context
}

type keyStruct struct {
	key string
}

func main() {
	// 使用context.Background()创建一个WithCancel类型的上下文
	ctxa, cancel := context.WithCancel(context.Background())

	// work模拟运行并检测前端的退出通知
	go work(ctxa, "work1")

	// 使用WithDeadline包装前面的上下文对象ctxa
	tm := time.Now().Add(3 * time.Second)
	ctxb, cancel2 := context.WithDeadline(ctxa, tm)

	go work(ctxb, "work2")

	// 使用WithValue包装前面的上下文对象ctxb
	oc := otherContext{ctxb}
	ctxc := context.WithValue(oc, keyStruct{"key"}, "andes, pss from main ")

	go workWithValue(ctxc, "work3")

	// 故意sleep 10s，让work2/work3超时退出
	time.Sleep(10 * time.Second)

	// 显式调用work1的cancel方法通知其退出
	cancel()
	cancel2()

	// 等待work1打印退出信息
	time.Sleep(5 * time.Second)
	fmt.Println("main stop")
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			fmt.Printf("%s is running\n", name)
			time.Sleep(1 * time.Second)
		}
	}
}

// 等待前端的退出通知，并试图获取context传递的数据
func workWithValue(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s get msg to cancel\n", name)
			return
		default:
			value := ctx.Value(keyStruct{"key"}).(string)
			fmt.Printf("%s is running, value = %v\n", name, value)
			time.Sleep(1 * time.Second)
		}
	}
}
