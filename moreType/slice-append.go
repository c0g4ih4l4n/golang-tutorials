package moreType

import "fmt"

// PrintSliceAppend :
// Function create an empty slice
// add item to it an print it out
func PrintSliceAppend() {

	var s []int

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Println(s)

}
