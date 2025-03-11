package main

import (
	"bufio"
	"fmt"
	"kakio/ch2/convertcf"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 获取命令行参数，从第二个参数开始（即忽略程序名）
	sm := os.Args[1:]
	// 检查获取的参数列表长度是否为0
	if len(sm) == 0 {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.")
			return
		}
		inputString := strings.Split(strings.TrimSpace(input), " ")
		fmt.Println(inputString)
		for i, s := range inputString {
			if s == " " || s == "\n" {
				continue
			}
			t, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf:%v\n", err)
				os.Exit(1)
			}
			switch {
			case i == 0:
				f := convertcf.Foot(t)
				fmt.Printf("%s = %s\n", f.String(), convertcf.FToM(f).String())
			case i == 1:
				m := convertcf.Meter(t)
				fmt.Printf("%s = %s\n", m.String(), convertcf.MToF(m).String())
			case i == 2:
				p := convertcf.Pound(t)
				fmt.Printf("%s = %s\n", p.String(), convertcf.PToK(p).String())
			case i == 3:
				k := convertcf.Kilogram(t)
				// p := convertcf.Pound(t)
				fmt.Printf("%s = %s\n", k.String(), convertcf.KToP(k).String())
			}
		}
	} else {
		for i, s := range sm {
			if s == " " || s == "\n" {
				continue
			}
			t, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf:%v\n", err)
				os.Exit(1)
			}
			switch {
			case i == 0:
				f := convertcf.Foot(t)
				// m := convertcf.Meter(t)
				fmt.Printf("%s = %s\n", f.String(), convertcf.FToM(f).String())
			case i == 1:
				// f := convertcf.Foot(t)
				m := convertcf.Meter(t)
				fmt.Printf("%s = %s\n", m.String(), convertcf.MToF(m).String())
			case i == 2:
				// k := convertcf.Kilogram(t)
				p := convertcf.Pound(t)
				fmt.Printf("%s = %s\n", p.String(), convertcf.PToK(p).String())
			case i == 3:
				k := convertcf.Kilogram(t)
				// p := convertcf.Pound(t)
				fmt.Printf("%s = %s\n", k.String(), convertcf.KToP(k).String())
			}
		}
	}
}
