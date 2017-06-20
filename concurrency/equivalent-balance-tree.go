package concurrency

import (
	"fmt"
	"math/rand"
)

// Tree :
// basic tree
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// quit to stop goroutine
// but it's doesn't reallyy need
// because we dont loop in channel
func walk(t *Tree, ch, quit chan int) {
	if t == nil {
		return
	}

	walk(t.Left, ch, quit)
	select {
	case ch <- t.Value:
	//val success fully sent
	case <-quit:
		return
	}
	walk(t.Right, ch, quit)
}

func walker(t *Tree, quit chan int) <-chan int {
	ch := make(chan int)
	go func() {
		walk(t, ch, quit)
		close(ch)
	}()
	return ch
}

func same(t1 *Tree, t2 *Tree) bool {
	quit := make(chan int)
	ch1, ch2 := walker(t1, quit), walker(t2, quit)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || !ok1 && ok2 || ok1 && !ok2 {
			quit <- 0
			return false
		}
		if !ok1 && !ok2 {
			return true
		}
	}

	return true
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

// PrintResultExercise :
// result exercise
func PrintResultExercise() {
	fmt.Print("tree.New(100, 1) == tree.New(100, 1): ")
	if same(New(100, 1), New(100, 1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

	fmt.Print("tree.New(10, 1) != tree.New(20, 1): ")
	if !same(New(10, 1), New(20, 1)) {
		fmt.Println("PASSED")
	} else {
		fmt.Println("FAILED")
	}

}
