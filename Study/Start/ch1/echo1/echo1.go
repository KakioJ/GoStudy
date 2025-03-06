package main

import (
	"fmt"
	"os"
)

func main() {
	var s, seep string
	for i := 1; i < len(os.Args); i++ {
		s += seep + os.Args[i]
		seep = " "
	}
	fmt.Println(s)
}
