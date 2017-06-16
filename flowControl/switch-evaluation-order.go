package flowControl

import (
	"fmt"
	"time"
)

// PrintSwitchEvalutionOrder : answer to day is saturday or not ?
func PrintSwitchEvalutionOrder() {
	fmt.Println("Is today is Saturday ? ")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("The next two day.")
	default:
		fmt.Println("Too far away")

	}
}
