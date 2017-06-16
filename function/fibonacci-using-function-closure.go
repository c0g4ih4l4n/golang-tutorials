package function

import "fmt"

var x1, x2 = 0, 1

func fibonacci() func() int {
	return func() int {
		x1, x2 = x2, x1+x2
		return x2
	}
}

// PrintFiboUsingFunctionClosure :
// desp : funny example about fibonacci using function closure
func PrintFiboUsingFunctionClosure() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			fmt.Println(0)
		case 1:
			fmt.Println(1)
		default:
			fmt.Println(f())
		}
	}
}
