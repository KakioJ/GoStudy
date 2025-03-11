package convertcf

import (
	"fmt"
)

type Foot float64
type Meter float64
type Pound float64
type Kilogram float64

func (f Foot) String() string     { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string    { return fmt.Sprintf("%g m", m) }
func (p Pound) String() string    { return fmt.Sprintf("%g lb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

func FToM(f Foot) Meter {
	return Meter(f * 0.3048)
}

func MToF(m Meter) Foot {
	return Foot(m / 0.3048)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.45359237)
}

func KToP(k Kilogram) Pound {
	return Pound(k / 0.45359237)
}
