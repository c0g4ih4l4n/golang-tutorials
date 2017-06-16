package moreType

import "fmt"

// Vertex : type for test struct
type Vertex struct {
	x, y int
}

// PrintStruct : function init one Vertex and print it
func PrintStruct() {
	a := Vertex{1, 2}

	p := &a
	fmt.Printf("X: %v \nY: %v", p.x, p.y)
}
