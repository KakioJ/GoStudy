package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// 记录程序开始时间
	start := time.Now()
	// 创建一个字符串类型的通道，用于接收fetch函数返回的结果
	ch := make(chan string)
	args := os.Args[1:]
	args = append(args, args...)
	// 遍历命令行参数中的URL，为每个URL启动一个goroutine进行数据抓取
	for _, url := range args {
		go fetch(url, ch)
	}

	// 遍历命令行参数中的URL，从通道中接收每个URL的抓取结果并打印
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	// 打印程序运行的总时间，保留两位小数
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch 函数用于从指定的URL获取数据，并通过通道返回结果
func fetch(url string, ch chan<- string) {
	// 记录开始时间
	start := time.Now()
	// 发送HTTP GET请求
	resp, err := http.Get(url)
	if err != nil {
		// 如果请求出错，将错误信息发送到通道并返回
		ch <- fmt.Sprint(err)
		return
	}
	out, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: creating output file %v\n", err)
		os.Exit(1)
	}

	nbytes, err := io.Copy(out, resp.Body)
	// 关闭响应体
	resp.Body.Close()
	if err != nil {
		// 如果读取响应体出错，将错误信息发送到通道
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
	}

	// 计算请求耗时（秒）
	secs := time.Since(start).Seconds()
	// 将请求耗时、响应体字节数和URL发送到通道
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
