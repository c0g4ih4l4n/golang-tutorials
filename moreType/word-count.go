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
	var s = "To day is si Friday"
	fmt.Println(s)
	fmt.Println(countWord(s))
}

func countWord(s string) map[string]int {

	listWord := make(map[string]int)

	for _, w := range strings.Fields(s) {
		listWord[strings.ToLower(w)]++
	}

	return listWord
}
