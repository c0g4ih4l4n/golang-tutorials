package moreType

import "fmt"

// PrintStructLiteral : Init Struct with init Vertex
func PrintStructLiteral() {
	var (
		a = Vertex{x: 1, y: 2}
		p = &a
	)
	fmt.Printf("Address of p: %v", p)
}
