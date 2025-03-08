package main

import (
	"fmt"
	"kakio/ch2/tempconv"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC)) // "212°F"
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
}
