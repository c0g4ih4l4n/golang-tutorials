package methods

import "fmt"

// Abser : defined a set of method signatures
type Abser interface {
	Abs() float64
}

// PrintInterfaceDemo :
// printdemo interface
func PrintInterfaceDemo() {
	var a Abser

	var f MyFloat
	//v := Vertex{1, 2}

	a = f
	//a = &v

	//fmt.Println(f.Abs())
	fmt.Println(a.Abs())
}
