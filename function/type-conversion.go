package function

import (
	"fmt"
	"math"
)

// PrintTypeConversion : Print result of type conversion
func PrintTypeConversion() {
	var x, y = 3, 4
	var z = math.Sqrt(float64(x*x + y*y))
	fmt.Println(z)
}
