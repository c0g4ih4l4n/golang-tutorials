package function

import "fmt"

func namedResult(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// PrintNamedResult : print named result
func PrintNamedResult() {
	result1, result2 := namedResult(18)
	fmt.Printf("Result is : %v and %v", result1, result2)

}
