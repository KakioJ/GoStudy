package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var status int64

func main() {
	// 遍历命令行参数中的所有URL
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			fmt.Println(url, "具备前缀http://或https://")
		} else {
			fmt.Println(url, "不具备具备前缀http://或https://")
			url = "http://" + url
		}
		// 发送HTTP GET请求获取URL的内容
		resp, err := http.Get(url)
		if err != nil {
			// 如果请求出错，打印错误信息到标准错误输出并退出程序
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		// 读取连接的状态
		signal := resp.Status
		fmt.Println("访问网站", url, "的连接状态是：", signal)
		// 读取响应体内容
		status, err = io.Copy(os.Stdout, resp.Body) // 将响应体内容拷贝到标准输出
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println()
		// fmt.Printf("%s", )
	}
}
