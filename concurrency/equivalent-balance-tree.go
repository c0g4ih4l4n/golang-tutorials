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

func walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func walker(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		walk(t, ch)
		close(ch)
	}()
	return ch
}

func same(t1 *Tree, t2 *Tree) bool {
	ch1, ch2 := walker(t1), walker(t2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || (!ok1 && ok2) || (ok1 && !ok2) {
			return false
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
	var t1, t2 *Tree
	t1 = New(10, 2)
	t2 = New(10, 2)

	fmt.Println(same(t1, t2))
}
