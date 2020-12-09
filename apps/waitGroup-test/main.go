package main

import (
	"net/http"
	"sync"
)

var (
	wg sync.WaitGroup

	urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.qq.com/",
		"http://www.keyboard-man.com/",
	}
)

func requestURL(url string) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err == nil {
		println("request success", url, resp.Status)
	} else {
		println("request fail", url, err)
	}
}

func main() {
	for _, url := range urls {
		// 每一个url启动一个gorountine
		wg.Add(1)

		go requestURL(url)
	}

	// 等待所有的请求结束
	wg.Wait()
}
