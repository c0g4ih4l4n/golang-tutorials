package methods

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name is %v, Age is %v", p.Name, p.Age)
}

func PrintDemoStringer() {
	alice := Person{"Alice", 9}

	fmt.Println(alice)
}
