package moreType

import "fmt"

// PrintSlice : Init one slice and print and array create from that slice
func PrintSlice() {
	primes := [6]int{1, 2, 3, 4, 5, 6}

	var s = primes[:0]
	printSliceInfo(s)

	s = primes[:4]
	printSliceInfo(s)

	s = primes[2:]
	printSliceInfo(s)

	s = primes[1:]
	printSliceInfo(s)
}

func printSliceInfo(s []int) {
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
