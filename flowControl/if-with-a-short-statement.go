package flowControl

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// PrintPow : Print Result of check Pow Check Function
func PrintPow() {
	fmt.Println(
		pow(3, 2, 10),
		pow(4, 2, 10),
	)
}
