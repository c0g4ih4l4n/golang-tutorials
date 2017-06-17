package methods

import (
	"fmt"
	"io"
)

// MyReader : My own reader to has method Read
type MyReader struct{}

// Read : of type MyReader to override method Read (interface io.Reader)
func (r MyReader) Read(b []byte) (n int, e error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

// PrintResultExReader : Print result of exercise
func PrintResultExReader() {
	var r MyReader

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		if err == io.EOF {
			fmt.Println("End !!")
			break
		}
		fmt.Printf("%s", b[:n])
	}

}
