package moreType

import (
	"fmt"
	"strings"
)

// PrintSliceOfSlice : print Caro 3 * 3
func PrintSliceOfSlice() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "O"
		}
	}
	board[0][0] = "X"
	board[1][1] = "X"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

}
