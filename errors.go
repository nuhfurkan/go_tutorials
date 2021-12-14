package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	z := float64(x)
	if z < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	
	// to set the number of significant figures change the variable sigFig
	var sigFig float64 = 4
	for z*z-x > math.Pow(10, -sigFig) {
		z -= (z*z - x) / (2 * z)
	}
	
	return math.Round(z*math.Pow(10, sigFig)) / math.Pow(10, sigFig), nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
