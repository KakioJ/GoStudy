package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 创建一个map，用于存储每行文本出现的次数，键为文本行，值为出现次数
	counts := make(map[string]map[string]int)
	// 获取命令行参数中的文件名列表，跳过第一个参数（程序名）
	files := os.Args[1:]
	// 如果没有提供文件名，则从标准输入读取
	if len(files) == 0 {
		countlines("stdin", os.Stdin, counts)
	} else {
		// 遍历每个文件名
		for _, arg := range files {
			// 打开文件
			f, err := os.Open(arg)
			// 如果打开文件出错，打印错误信息到标准错误输出，并继续处理下一个文件
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 调用countlines函数统计文件中的行出现次数
			countlines(arg, f, counts)
			// 关闭文件
			f.Close()
		}
	}

	// 遍历counts map，打印出现次数大于1的行及其出现次数
	for filename, count := range counts {
		for line, n := range count {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}
	}
}

func countlines(filename string, f *os.File, counts map[string]map[string]int) {
	// 为每个文件名创建一个新的map，用于存储该文件中每行文本的出现次数
	if counts[filename] == nil {
		counts[filename] = make(map[string]int)
	}
	// 创建一个新的Scanner，用于读取文件内容
	input := bufio.NewScanner(f)
	// 循环读取文件的每一行
	for input.Scan() {
		// 将当前行作为键，在counts映射中对应的值加1
		counts[filename][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
