package flowControl

import "fmt"

// PrintIf : Print result of example for if statement
func PrintIf() {
	var check = true

	if check {
		fmt.Printf("The result of check is %v", check)
	} else {
		fmt.Printf("Wrong")
	}
}
