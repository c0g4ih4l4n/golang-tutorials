package methods

import "fmt"

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type int")
	case string:
		fmt.Println("Type String")
	default:
		fmt.Println("Not Type i know")
	}
}

// PrintTypeSwitchDemo :
// demo type switch
// use interface in param
func PrintTypeSwitchDemo() {
	do(1)
	do("hello")
}
