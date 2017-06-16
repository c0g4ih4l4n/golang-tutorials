package packages

import (
	"fmt"
	"math/rand"
)

// PrintPackages : Print Random number
func PrintPackages() {
	fmt.Printf("Generate random number : %v", rand.Intn(10))
}
