package function

import "fmt"

// Swap : Swap two strings
func Swap(x, y string) (string, string) {
	return y, x
}

// PrintSwap : print result of swap function
func PrintSwap() {
	input1, input2 := "Hi", "there"
	result1, result2 := Swap(input1, input2)
	fmt.Printf("\nStrings before swap : %v and %v", input1, input2)
	fmt.Printf("\nResult after swap : %v and %v", result1, result2)
}
