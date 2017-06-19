package concurrency

import "fmt"

func sum(arr []int, c chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	c <- sum
}

// PrintChannel :
// demo using 2 goroutine to multithread
// sum an array
func PrintChannel() {
	var arr []int
	var n int
	var c chan int

	arr = []int{1, 2, 3, 4, 5, 6}
	c = make(chan int)
	n = len(arr)
	go sum(arr[:(n/2)], c)
	go sum(arr[(n/2):], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
