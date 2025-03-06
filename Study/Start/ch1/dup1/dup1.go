package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    // 创建一个map，用于存储每个字符串出现的次数，键为字符串，值为整数
	counts :=  make(map[string] int)
    // 创建一个新的扫描器，用于从标准输入读取数据
	input := bufio.NewScanner(os.Stdin)
    // 循环读取标准输入的每一行
	for input.Scan() {
        // 将读取的行作为键，值加1，表示该行出现一次
		counts[input.Text()]++
	}

    // 遍历map，打印出现次数大于1的行及其出现次数
	for line, n := range counts {
		if n > 1 {
            // 打印格式为：出现次数\t 行内容
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}