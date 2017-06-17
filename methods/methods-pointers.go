package methods

import "fmt"

// Scale :
// scale value of Vertex var by number f
func (v *Vertex) Scale(f float64) {
	v.X, v.Y = v.X*f, v.Y*f
}

func scale(v *Vertex, f float64) {
	v.X, v.Y = v.X*f, v.Y*f
}

func (v Vertex) testFunc(f float64) {
	fmt.Println("Test Function ")
}

// PrintMethodPointerDemo :
// A demo for this
func PrintMethodPointerDemo() {
	v := Vertex{1, 2}

	p := &v
	fmt.Println(v)
	fmt.Println(v.Abs())
	p.testFunc(2)
}
