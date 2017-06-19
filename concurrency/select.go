package concurrency

import "fmt"

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

// PrintSelect :
// Demo for select statement
// run when one case is ready
// run random when multiple case is ready
func PrintSelect() {
	c, quit := make(chan int), make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacciSelect(c, quit)
}
