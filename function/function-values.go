package function

import (
	"fmt"
	"math"
)

// PrintFunctionValue :
// desp : use function as a value
// assign function to variable
// use that variable with arguments
func PrintFunctionValue() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
}
