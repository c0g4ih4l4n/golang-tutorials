package concurrency

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// PrintGoRoutine :
// Demo for GoRoutin
func PrintGoRoutine() {
	go say("Hi")
	say("There !!")
}
