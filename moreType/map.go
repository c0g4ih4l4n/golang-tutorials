package moreType

import "fmt"

// PrintMap :
// demo Map Type
func PrintMap() {
	m := make(map[string]Vertex)

	m["first"] = Vertex{1, 2}

	var m2 = map[string]Vertex{
		"second": Vertex{2, 3},
		"third":  {3, 4},
	}

	delete(m2, "third")

	fmt.Println(m["first"])
	fmt.Println(m2)

	v, ok := m["third"]
	if !ok {
		fmt.Printf("Not found")
	} else {
		fmt.Printf("%v", v)
	}
}
