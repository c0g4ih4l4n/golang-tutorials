package methods

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (n int, err error) {
	return rot13.r.Read(b)
}

// PrintResultRot13Reader :
// Print demo result
func PrintResultRot13Reader() {
	s := strings.NewReader("This is sample text for exercise Rot13Reader")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
