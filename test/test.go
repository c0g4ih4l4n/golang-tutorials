package test

import "fmt"

// PrintTest :
func PrintTest() {
	var name string
	var age int

	fmt.Print("Enter your name : ")
	fmt.Scanf("%s", &name)

	fmt.Print("Enter your age : ")
	fmt.Scanf("%d", &age)

	fmt.Printf("Your name is %s and you're %d years old.", name, age)
}
