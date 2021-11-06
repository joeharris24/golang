// Maps are just like dictionaries (I think)

package main

import "fmt"

func main() {

	m := make(map[string]int)
	fmt.Println(m)
	// > map[]

	m["test"] = 102
	m["test2"] = 201
	fmt.Println(m)
	// map[test:102 test2:201]

	v1 := m["test"]
	fmt.Println(v1)
	// 102

	delete(m, "test")
	fmt.Println(m)
	// > map[test2:201]

	_, prs := m["test2"]
	fmt.Println(prs)
	// m["test2"] = value, present
	// So _, prs = blank, present (bool)
	// > true

	n := map[string]int{"A": 1, "B": 2}
	fmt.Println(n)
	// > map[A:1 B:2]

}
