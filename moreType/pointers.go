package moreType

import "fmt"

// PrintPointer : assign pointer to variable and print value that pointer point to
func PrintPointer() {
	i, j := 1, 2

	var p, q *int

	p = &i
	q = &j
	var r **int
	r = &q
	fmt.Println(*p)
	fmt.Printf("\nAddress of variable q is %v\n And value of variable r is %v", q, *r)
	fmt.Printf("Type of variable r is %T", r)
}
