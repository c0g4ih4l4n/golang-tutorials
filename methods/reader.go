package methods

import (
	"fmt"
	"io"
	"strings"
)

// PrintReader :
// Demo for interface io.Reader
// has Read method
func PrintReader() {
	fmt.Println("**Reader example**")
	fmt.Println("------------------")
	r := strings.NewReader("This is a sample text")

	b := make([]byte, 6)

	for {
		n, err := r.Read(b)

		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("\nHave read : %q\n", b[:n])

		if err == io.EOF {
			fmt.Println("End!!")
			break
		}
	}

}
