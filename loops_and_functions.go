package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(x)
	
	// to set the number of significant figures change the variable sigFig
	var sigFig float64 = 8
	
	for z*z - x > math.Pow(10, -sigFig) {
		z -= (z*z - x) / (2*z)
	}
	return math.Round(z*math.Pow(10, sigFig))/ math.Pow(10, sigFig)
}

func main() {
	fmt.Println(Sqrt(8))
}
