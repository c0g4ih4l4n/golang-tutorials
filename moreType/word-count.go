package moreType

import (
	"fmt"
	"strings"
)

// PrintWordCount :
// desp: print number of word
// given string
// count number of word in that string
func PrintWordCount() {
	var s = "To day is Friday"
	fmt.Println(countWord(s))
}

func countWord(s string) int {
	var words = strings.Fields(s)
	return len(words)
}
