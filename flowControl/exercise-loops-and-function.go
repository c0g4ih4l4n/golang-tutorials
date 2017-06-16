package flowControl

import "fmt"

func sqrtNewtonMethod(x int) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z = z - (z*z-float64(x))/(2*z)
	}
	return z
}

// PrintExerciseResult : Print result of sqrt calculate by Newton method
func PrintExerciseResult() {
	fmt.Printf("Sqrt of 2 is : %v", sqrtNewtonMethod(2))
}
