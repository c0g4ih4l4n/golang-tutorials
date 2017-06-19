package function

import "fmt"

// fibonacci is a fun<Plug>yankstack_after_pastection that returns
// a function that returns an int.
func fibonacci1() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

// PrintAnotherResultForFiboClosure :
// demo by zxyar
func PrintAnotherResultForFiboClosure() {
	f := fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
