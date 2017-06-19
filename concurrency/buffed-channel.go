package concurrency

import "fmt"

// PrintBufferedChannel : a demo for buffered channel
// and example about deadlock when
// buffered channel full or empty
func PrintBufferedChannel() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	ch <- 1
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
