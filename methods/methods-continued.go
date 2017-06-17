package methods

import "fmt"

// MyFloat : type declare in my own
// just basic type
type MyFloat float64

// Abs : methods on MyFloat
// return abs of variable
func (v MyFloat) Abs() float64 {
	if v < 0 {
		return float64(-v)
	}
	return float64(v)
}

// PrintMethodContinued :
// Print demo of method continued
func PrintMethodContinued() {
	var x MyFloat
	x = 4
	fmt.Println(x.Abs())
}
