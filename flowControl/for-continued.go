package flowControl

import "fmt"

// PrintForContinued : Print result for function
func PrintForContinued() {
	sum := 1
	for sum < 1000 {
		sum += sum

	}
	fmt.Println(sum)
}
