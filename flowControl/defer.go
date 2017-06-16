package flowControl

import "fmt"

// PrintDefer : defer test function
func PrintDefer() {
	defer fmt.Println("There !!")

	fmt.Println("Hi ")
}
