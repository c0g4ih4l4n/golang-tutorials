package moreType

import "fmt"

// PrintRange : Use to demo range feature
func PrintRange() {
	var s = [6]int{1, 2, 3, 4, 5, 6}
	for i := range s {
		fmt.Println(i)
	}

	for _, value := range s {
		fmt.Println(value)
	}
}
