package flowControl

import "fmt"

// PrintFor : Print Result for file for.go
func PrintFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
