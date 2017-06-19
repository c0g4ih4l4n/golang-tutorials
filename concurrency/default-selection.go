package concurrency

import (
	"fmt"
	"time"
)

// PrintDefaultSelection :
// 2 channel tick and boom
// if both tick and boom doesn't ready
// call default
func PrintDefaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
