package main

import "fmt"

func basename1(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}

func main() {
	fmt.Println(basename1("a/b/c.go"))
	fmt.Println(basename1("c.d.go"))
	fmt.Println(basename1("abc"))
}
