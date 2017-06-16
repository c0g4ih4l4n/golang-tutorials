package flowControl

import "fmt"

// PrintDeferMulti : Multi defer
func PrintDeferMulti() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
