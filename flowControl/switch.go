package flowControl

import (
	"fmt"
	"runtime"
)

// PrintSwitch : Print os using switch
func PrintSwitch() {
	fmt.Println("Go run on : ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux ")
	default:
		fmt.Printf("%s", os)
	}
}
