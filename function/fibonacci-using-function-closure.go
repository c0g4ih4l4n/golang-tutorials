package function

import "fmt"

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

// PrintFiboUsingFunctionClosure :
// desp : funny example about fibonacci using function closure
func PrintFiboUsingFunctionClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
