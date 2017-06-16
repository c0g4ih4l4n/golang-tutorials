package flowControl

import (
	"fmt"
	"time"
)

// PrintSwitchNoCondition : Print Switch No Condition = switch true
func PrintSwitchNoCondition() {
	fmt.Println("Get up ? ")
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("Morning men !")
	case time.Now().Hour() < 18:
		fmt.Println("Afternoon men !")
	default:
		fmt.Println("Nightmare !")
	}

}
