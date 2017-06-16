package moreType

import "fmt"

// PrintMakingSlice : Making Slice
func PrintMakingSlice() {
	slice := make([]int, 5)

	slice[0] = 0
	slice[1] = 1

	fmt.Println(slice)
}
