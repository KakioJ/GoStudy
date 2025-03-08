package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// func CToF(c Celsius) Fahrenheit {
// 	return Fahrenheit(c*9/5 + 32)
// }

// func FToC(f Fahrenheit) Celsius {
// 	return Celsius((f - 32) * 5 / 9)
// }

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }


func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// func main() {
// 	c := FToC(212.0)
// 	fmt.Println(c.String())
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%s\n", c)
// 	fmt.Printf("%g\n", c)
// 	fmt.Println(c == BoilingC, c)
// 	f := CToF(BoilingC)
// 	fmt.Println(f == Fahrenheit(BoilingC), f)
// 	// fmt.Println(f == BoilingC, f)
// 	fmt.Println(CToF(AbsoluteZeroC), " ", CToF(FreezingC), " ", CToF(BoilingC))
// }
