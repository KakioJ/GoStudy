package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// 解析命令行参数
	flag.Parse()
	// 将命令行参数（除程序名外的参数）用指定的分隔符连接成一个字符串
	// flag.Args() 返回命令行参数的切片
	// *sep 是一个指向字符串的指针，表示分隔符
	fmt.Print(strings.Join(flag.Args(), *sep))
	// 如果 n 标志为 false，则输出一个换行符
	// *n 是一个指向布尔值的指针，表示是否需要换行
	if !*n {
		fmt.Println("")
	}

}
