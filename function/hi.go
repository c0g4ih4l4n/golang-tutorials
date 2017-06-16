package function

import "fmt"

func add(a, b int) int {
	return a + b
}

// PrintSum : Print Sum of 1 and 2
func PrintSum() {
	fmt.Printf("Sum 1 and 2 : %v", add(1, 2))
}
