package concurrency

import "math/rand"

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

func same(t1 *Tree, t2 *Tree) {
	ch1, ch2 := make(chan int), make(chan int)
	walk(t1, ch1)
	walk(t2, ch2)
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

func PrintResultExercise() {

}
