package main

import (
	"fmt"
	"time"
)

type query struct {
	sql chan string

	result chan string
}

func execQuery(q query) {
	go func() {
		sql := <-q.sql

		// 访问数据库

		q.result <- "result from " + sql
	}()
}

func main() {
	q := query{make(chan string, 1), make(chan string, 1)}

	go execQuery(q)

	// 发送参数
	q.sql <- "select * from table"

	// 做其他事情，通过sleep描述
	time.Sleep(1 * time.Second)

	// 获取结果
	fmt.Println(<-q.result)
}
