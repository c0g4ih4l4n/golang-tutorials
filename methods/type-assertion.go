package methods

import "fmt"

// PrintDemoTypeAssertion :
// A demo for type assert
func PrintDemoTypeAssertion() {
	var i interface{} = "hi"

	t := i.(T)
	fmt.Printf("Value T: %v", t)
}
