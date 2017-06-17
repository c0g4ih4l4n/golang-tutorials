package methods

import "fmt"

// I : test interface
type I interface {
	M()
}

// T : test struct
type T struct {
	S string
}

// M : test function M
func (t *T) M() {
	if t == nil {
		fmt.Println("Nil Value")
		return
	}
	fmt.Println(t.S)
}

// F : test strcut
type F float64

// M : test function
func (f *F) M() {
	fmt.Println(f)
}

// PrintInterfaceValueDemo : test function
func PrintInterfaceValueDemo() {
	var i I

	t := T{"hi"}
	f := F(1)

	i = &t
	describe(i)
	i.M()

	i = &f
	describe(i)
	i.M()

	var test *F
	i = test
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("value: %v, type: %T\n", i, i)
}
