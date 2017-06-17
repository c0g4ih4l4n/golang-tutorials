package methods

import "fmt"

// ErrNegativeSqrt :
// Err declare for sqrt function
// that given value is negative number
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot sqrt negative number : %v", float64(e))
}

func sqrt(x float64) (float64, *ErrNegativeSqrt) {
	if x < 0 {
		var e = ErrNegativeSqrt(x)
		return 0, &e
	} else if x == 0 {
		return 0, nil
	}

	// calculate follow Newton method
	result := 1.0
	for i := 0; i < 10; i++ {
		result = result - ((result*result - x) / (2 * result))
	}
	return result, nil
}

// PrintExError :
// Demo for error exercise
// create own error and modify how to display that error
func PrintExError() {
	fmt.Println("****** Exercise Error ******")
	fmt.Println("----------------------------")

	var requestNumber float64
	fmt.Print("Enter Your Number : ")
	fmt.Scanf("%f", &requestNumber)
	result, err := sqrt(requestNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result of sqrt(%v) is %v", requestNumber, result)
	return
}
