package methods

import (
	"fmt"
	"math"
)

// Vertex : Type struct to comment
type Vertex struct {
	X, Y float64
}

// Abs : funcion to get Abs of vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// PrintMethodsResult :
// Result to print to main.go
func PrintMethodsResult() {
	a := Vertex{3, 4}
	fmt.Println(a.Abs())
}
